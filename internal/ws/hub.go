package ws

import (
	"digital_signage_api/internal/db"
	"digital_signage_api/internal/dto"
	"digital_signage_api/internal/models"
	"digital_signage_api/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // TODO: perketat origin
	},
}

var deviceConns = make(map[uint]*websocket.Conn) // DeviceID -> Conn
var mu sync.Mutex

// cache terakhir schedule yang dikirim -> biar tidak spam kirim schedule sama
var lastScheduleSent = make(map[uint]uint) // DeviceID -> ScheduleID

func HandleDeviceConnection(w http.ResponseWriter, r *http.Request, device models.Device) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("❌ upgrade error:", err)
		return
	}

	mu.Lock()
	deviceConns[device.DeviceID] = conn
	mu.Unlock()

	fmt.Printf("✅ device %d connected (airport %d)\n", device.DeviceID, device.AirportID)

	// kirim active schedule saat connect
	SendActiveSchedule(device)

	// listen until disconnect
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}

	mu.Lock()
	delete(deviceConns, device.DeviceID)
	delete(lastScheduleSent, device.DeviceID)
	mu.Unlock()
	conn.Close()
	fmt.Printf("❌ device %d disconnected\n", device.DeviceID)
}

func SendActiveSchedule(device models.Device) {
	mu.Lock()
	conn, ok := deviceConns[device.DeviceID]
	mu.Unlock()

	if !ok {
		fmt.Printf("⚠️ device %d not connected\n", device.DeviceID)
		return
	}

	now := time.Now()
	nowEpochMs := now.UnixMilli()
	nowTime := now.Format("15:04:05") // current time string

	// Ambil semua schedule milik airport device
	var schedules []models.Schedule
	err := db.DB.
		Preload("Playlist.PlaylistContent.Content").
		Where("airport_id = ? AND start_date <= ? AND end_date >= ?", device.AirportID, nowEpochMs, nowEpochMs).
		Find(&schedules).Error
	if err != nil || len(schedules) == 0 {
		msg := `{"message":"no active schedule"}`
		conn.WriteMessage(websocket.TextMessage, []byte(msg))
		return
	}

	var active *models.Schedule

	// 1. Cari urgent schedule aktif
	for _, sch := range schedules {
		if sch.IsUrgent && isTimeInRange(nowTime, sch.StartTime, sch.EndTime) {
			active = &sch
			break
		}
	}

	// 2. Kalau tidak ada urgent → cari schedule normal yang aktif
	if active == nil {
		for _, sch := range schedules {
			if isTimeInRange(nowTime, sch.StartTime, sch.EndTime) {
				active = &sch
				break
			}
		}
	}

	if active == nil {
		msg := `{"message":"no active schedule"}`
		conn.WriteMessage(websocket.TextMessage, []byte(msg))
		return
	}

	// cek apakah schedule sama dengan sebelumnya
	mu.Lock()
	lastSent := lastScheduleSent[device.DeviceID]
	mu.Unlock()
	if lastSent == active.ScheduleID {
		// tidak kirim ulang kalau belum berubah
		return
	}

	// Build payload ActiveScheduleRes
	playlist := active.Playlist
	payload := dto.ActiveScheduleRes{
		ScheduleID: active.ScheduleID,
		IsUrgent:   active.IsUrgent,
		StartDate:  active.StartDate,
		EndDate:    active.EndDate,
		StartTime:  active.StartTime, // langsung string "HH:MM:SS"
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
	fmt.Println("Sending to device:", string(data))
	conn.WriteMessage(websocket.TextMessage, data)

	mu.Lock()
	lastScheduleSent[device.DeviceID] = active.ScheduleID
	mu.Unlock()
}

// Helper: cek apakah nowTime ada di range StartTime–EndTime
func isTimeInRange(nowStr, startStr, endStr string) bool {
	layout := "15:04:05"

	nowT, err1 := time.Parse(layout, strings.TrimSpace(nowStr))
	startT, err2 := time.Parse(layout, strings.TrimSpace(startStr))
	endT, err3 := time.Parse(layout, strings.TrimSpace(endStr))

	if err1 != nil {
		fmt.Printf("[DEBUG isTimeInRange] gagal parse nowStr=%q: %v\n", nowStr, err1)
	}
	if err2 != nil {
		fmt.Printf("[DEBUG isTimeInRange] gagal parse startStr=%q: %v\n", startStr, err2)
	}
	if err3 != nil {
		fmt.Printf("[DEBUG isTimeInRange] gagal parse endStr=%q: %v\n", endStr, err3)
	}

	if err1 != nil || err2 != nil || err3 != nil {
		return false
	}

	fmt.Printf("[DEBUG isTimeInRange] now=%s, start=%s, end=%s\n",
		nowT.Format(layout), startT.Format(layout), endT.Format(layout))

	if startT.Before(endT) {
		// normal range (ex: 08:00–17:00)
		inRange := !nowT.Before(startT) && !nowT.After(endT)
		fmt.Printf("[DEBUG isTimeInRange] normal range → %v\n", inRange)
		return inRange
	} else {
		// lewat tengah malam (ex: 21:00–03:00)
		inRange := !nowT.Before(startT) || !nowT.After(endT)
		fmt.Printf("[DEBUG isTimeInRange] cross-midnight range → %v\n", inRange)
		return inRange
	}
}


// RunScheduler jalan di background setiap 1 menit
func RunScheduler() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		<-ticker.C

		mu.Lock()
		// copy biar aman dipakai di loop
		conns := make(map[uint]*websocket.Conn)
		for id, conn := range deviceConns {
			conns[id] = conn
		}
		mu.Unlock()

		for deviceID := range conns {
			// ambil device dari DB biar tahu airport_id
			var device models.Device
			if err := db.DB.Preload("Airport").First(&device, deviceID).Error; err != nil {
				fmt.Println("⚠️ gagal ambil device:", err)
				continue
			}

			// kirim schedule aktif
			SendActiveSchedule(device)
		}
	}
}
