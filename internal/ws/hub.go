package ws

import (
	"digital_signage_api/internal/db"
	"digital_signage_api/internal/dto"
	"digital_signage_api/internal/models"
	"digital_signage_api/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
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
	nowTime := now.Format("15:04:05") // jam:menit:detik

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

	// Build payload ActiveScheduleRes
	playlist := active.Playlist
	payload := dto.ActiveScheduleRes{
		ScheduleID: active.ScheduleID,
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
}

// Helper: cek apakah nowTime ada di range StartTime–EndTime
func isTimeInRange(now, start, end string) bool {
	layout := "15:04:05"
	nowT, _ := time.Parse(layout, now)
	startT, _ := time.Parse(layout, start)
	endT, _ := time.Parse(layout, end)

	if startT.Before(endT) {
		// normal range (ex: 08:00–17:00)
		return !nowT.Before(startT) && !nowT.After(endT)
	} else {
		// lewat tengah malam (ex: 21:00–03:00)
		return !nowT.Before(startT) || !nowT.After(endT)
	}
}