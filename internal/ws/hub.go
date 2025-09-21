package ws

import (
	"digital_signage_api/internal/db"
	"digital_signage_api/internal/dto"
	"digital_signage_api/internal/models"
	"digital_signage_api/internal/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// DeviceConn membungkus websocket.Conn dengan mutex agar aman untuk concurrent write
type DeviceConn struct {
	Conn *websocket.Conn
	Mu   sync.Mutex
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // TODO: perketat origin
	},
}

// Global maps
var (
	deviceConns = make(map[uint]*DeviceConn)      // DeviceID -> DeviceConn
	airportDevices = make(map[uint]map[uint]struct{}) // AirportID -> set of DeviceIDs
	lastScheduleSent = make(map[uint]uint)             // DeviceID -> ScheduleID
	airportTimezones = make(map[uint]string)            // AirportID -> IANA timezone
	mu sync.Mutex
)

// helper write aman
func safeWrite(dc *DeviceConn, messageType int, data []byte) error {
	dc.Mu.Lock()
	defer dc.Mu.Unlock()
	return dc.Conn.WriteMessage(messageType, data)
}

func safeWriteControl(dc *DeviceConn, messageType int, data []byte, deadline time.Time) error {
	dc.Mu.Lock()
	defer dc.Mu.Unlock()
	return dc.Conn.WriteControl(messageType, data, deadline)
}

func HandleDeviceConnection(w http.ResponseWriter, r *http.Request, device models.Device) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("❌ upgrade error:", err)
		return
	}
	dc := &DeviceConn{Conn: conn}

	// simpan device
	mu.Lock()
	deviceConns[device.DeviceID] = dc
	if airportDevices[device.AirportID] == nil {
		airportDevices[device.AirportID] = make(map[uint]struct{})
	}
	airportDevices[device.AirportID][device.DeviceID] = struct{}{}
	airportTimezones[device.AirportID] = mapToIANATz(device.Airport.Timezone)
	mu.Unlock()

	fmt.Printf("✅ device %d connected (airport %d)\n", device.DeviceID, device.AirportID)

	// kirim active schedule saat connect
	SendActiveSchedule(device)

	// set deadline awal
	conn.SetReadDeadline(time.Now().Add(60 * time.Second))

	// kalau terima Pong → perpanjang deadline
	conn.SetPongHandler(func(appData string) error {
		fmt.Printf("📨 received pong from device %d\n", device.DeviceID)
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	// goroutine ping loop
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		for {
			<-ticker.C
			mu.Lock()
			_, ok := deviceConns[device.DeviceID]
			mu.Unlock()
			if !ok {
				return // device sudah disconnect
			}
			if err := safeWriteControl(dc, websocket.PingMessage, nil, time.Now().Add(5*time.Second)); err != nil {
				fmt.Printf("⚠️ failed to send ping to device %d: %v\n", device.DeviceID, err)
				return
			}
			fmt.Printf("📡 sent ping to device %d\n", device.DeviceID)
		}
	}()

	// listen sampai disconnect
	for {
		mt, r, err := conn.NextReader()
		if err != nil {
			fmt.Printf("❌ device %d read error: %v\n", device.DeviceID, err)
			break
		}
		// reset deadline setiap pesan apapun
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))

		if mt == websocket.TextMessage {
			var sb strings.Builder
			_, _ = io.Copy(&sb, r)
			msg := sb.String()
			if strings.TrimSpace(msg) != "" {
				fmt.Printf("💬 text message from device %d: %s\n", device.DeviceID, msg)
			}
		}
	}

	// cleanup
	mu.Lock()
	delete(deviceConns, device.DeviceID)
	if set, ok := airportDevices[device.AirportID]; ok {
		delete(set, device.DeviceID)
		if len(set) == 0 {
			delete(airportDevices, device.AirportID)
		}
	}
	delete(lastScheduleSent, device.DeviceID)
	mu.Unlock()
	conn.Close()
	fmt.Printf("❌ device %d disconnected\n", device.DeviceID)
}

// Kirim schedule hanya ke satu device (dipanggil saat connect)
func SendActiveSchedule(device models.Device) {
	mu.Lock()
	dc, ok := deviceConns[device.DeviceID]
	mu.Unlock()
	if !ok {
		fmt.Printf("⚠️ device %d not connected\n", device.DeviceID)
		return
	}

	mu.Lock()
	locName := airportTimezones[device.AirportID]
	mu.Unlock()
	
	loc, _ := time.LoadLocation(locName)
	now := time.Now().In(loc)
	nowEpochMs := now.UnixMilli()
	nowTime := now.Format("15:04:05")

	// Ambil semua schedule milik airport device
	var schedules []models.Schedule
	err := db.DB.
		Preload("Playlist.PlaylistContent.Content").
		Where("airport_id = ? AND start_date <= ? AND end_date >= ?", device.AirportID, nowEpochMs, nowEpochMs).
		Find(&schedules).Error
	if err != nil || len(schedules) == 0 {
		msg := `{"message":"no active schedule"}`
		safeWrite(dc, websocket.TextMessage, []byte(msg))
		return
	}

	active := pickActiveSchedule(schedules, nowTime)
	sendScheduleToDevice(device.DeviceID, active)
}

// pilih schedule aktif dari daftar
func pickActiveSchedule(schedules []models.Schedule, nowTime string) *models.Schedule {
	// 1. Cari urgent
	for _, sch := range schedules {
		if sch.IsUrgent && isTimeInRange(nowTime, sch.StartTime, sch.EndTime) {
			return &sch
		}
	}
	// 2. Cari normal
	for _, sch := range schedules {
		if isTimeInRange(nowTime, sch.StartTime, sch.EndTime) {
			return &sch
		}
	}
	return nil
}

// kirim schedule ke device tertentu
func sendScheduleToDevice(deviceID uint, active *models.Schedule) {
	mu.Lock()
	dc, ok := deviceConns[deviceID]
	mu.Unlock()
	if !ok {
		return
	}

	if active == nil {
		msg := `{"message":"no active schedule"}`
		safeWrite(dc, websocket.TextMessage, []byte(msg))
		return
	}

	// cek apakah sama dengan last
	mu.Lock()
	lastSent := lastScheduleSent[deviceID]
	mu.Unlock()
	if lastSent == active.ScheduleID {
		return
	}

	// build payload
	playlist := active.Playlist
	payload := dto.ActiveScheduleRes{
		ScheduleID: active.ScheduleID,
		IsUrgent:   active.IsUrgent,
		StartDate:  active.StartDate,
		EndDate:    active.EndDate,
		StartTime:  active.StartTime,
		EndTime:    active.EndTime,
		PlaylistID: playlist.PlaylistID,
		Name:       playlist.Name,
	}
	for _, pc := range playlist.PlaylistContent {
		if pc.Content != nil {
			payload.Contents = append(payload.Contents, struct {
				ContentID uint   `json:"content_id"`
				Title     string `json:"title"`
				URL       string `json:"url"`
				Order     int    `json:"order"`
			}{
				ContentID: pc.ContentID,
				Title:     pc.Content.Title,
				URL:       utils.BuildContentURL(pc.Content.Title),
				Order:     pc.Order,
			})
		}
	}

	data, _ := json.Marshal(payload)
	fmt.Println("Sending to device:", deviceID, string(data))
	safeWrite(dc, websocket.TextMessage, data)

	mu.Lock()
	lastScheduleSent[deviceID] = active.ScheduleID
	mu.Unlock()
}

// cek apakah nowTime ada di range StartTime–EndTime
func isTimeInRange(nowStr, startStr, endStr string) bool {
	layout := "15:04:05"

	nowT, err1 := time.Parse(layout, strings.TrimSpace(nowStr))
	startT, err2 := time.Parse(layout, strings.TrimSpace(startStr))
	endT, err3 := time.Parse(layout, strings.TrimSpace(endStr))

	if err1 != nil || err2 != nil || err3 != nil {
		return false
	}

	if startT.Before(endT) {
		// normal range (ex: 08:00–17:00)
		return !nowT.Before(startT) && !nowT.After(endT)
	} else {
		// lewat tengah malam (ex: 21:00–03:00)
		return !nowT.Before(startT) || !nowT.After(endT)
	}
}

// RunScheduler jalan di background setiap 1 menit
func RunScheduler() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		<-ticker.C

		// snapshot airports + devices
		mu.Lock()
		airports := make(map[uint][]uint) // AirportID -> []DeviceID
		for airportID, set := range airportDevices {
			for deviceID := range set {
				airports[airportID] = append(airports[airportID], deviceID)
			}
		}
		mu.Unlock()

		// proses per airport
		for airportID, deviceIDs := range airports {
			mu.Lock()
			locName := airportTimezones[airportID]
			mu.Unlock()

			loc, _ := time.LoadLocation(locName)
			now := time.Now().In(loc)
			nowEpochMs := now.UnixMilli()
			nowTime := now.Format("15:04:05")

			// query sekali saja untuk airport ini
			var schedules []models.Schedule
			err := db.DB.
				Preload("Playlist.PlaylistContent.Content").
				Where("airport_id = ? AND start_date <= ? AND end_date >= ?", airportID, nowEpochMs, nowEpochMs).
				Find(&schedules).Error
			if err != nil {
				fmt.Printf("⚠️ gagal ambil schedule airport %d: %v\n", airportID, err)
				continue
			}

			active := pickActiveSchedule(schedules, nowTime)
			for _, deviceID := range deviceIDs {
				sendScheduleToDevice(deviceID, active)
			}
		}
	}
}

func mapToIANATz(code string) string {
   switch code {
   	case "WIB":
      	return "Asia/Jakarta"
    	case "WITA":
      	return "Asia/Makassar"
    	case "WIT":
      	return "Asia/Jayapura"
    	default:
      	return "Asia/Jakarta"
   }
}
