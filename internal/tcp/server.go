package tcp

import (
	"digital_signage_api/internal/db"
	"digital_signage_api/internal/models"
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"time"
)

type PlaylistPayload struct {
	ScheduleID uint              `json:"schedule_id"`
	PlaylistID uint              `json:"playlist_id"`
	Name       string            `json:"name"`
	Contents   []ContentResponse `json:"contents"`
}

type ContentResponse struct {
	ContentID uint   `json:"content_id"`
	Title     string `json:"title"`
	URL       string `json:"url"`
	Order     int    `json:"order"`
}



var clientCount int
var mu sync.Mutex

func StartTCPServer() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	fmt.Println("🎧 TCP server listening on :9000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("❌ accept error:", err)
			continue
		}

		mu.Lock()
		clientCount++
		mu.Unlock()
		fmt.Println("Client connected. Total:", clientCount)

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	mu.Lock()
	clientCount++
	mu.Unlock()
	fmt.Println("Client connected. Total:", clientCount)

	// kirim sekali saat connect
	sendActivePlaylist(conn)

	// biarin koneksi tetap terbuka, tunggu client yang nutup
	buf := make([]byte, 1)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			break // client putus → keluar loop
		}
	}

	mu.Lock()
	clientCount--
	mu.Unlock()
	conn.Close()
	fmt.Println("Client disconnected. Total:", clientCount)
}


func sendActivePlaylist(conn net.Conn) {
	now := time.Now().Unix() // epoch second

	var schedule models.Schedule
	err := db.DB.Preload("Playlist.PlaylistContent.Content").
		Where("start_time <= ? AND end_time > ?", now, now).
		First(&schedule).Error
	if err != nil {
		msg := `{"message":"no active playlist"}`
		fmt.Println("Sending (server log):", msg)
		conn.Write([]byte(msg + "\n"))
		return
	}

	playlist := schedule.Playlist
	payload := PlaylistPayload{
		ScheduleID: schedule.ScheduleID,
		PlaylistID: playlist.PlaylistID,
		Name:       playlist.Name,
		Contents:   []ContentResponse{},
	}

	for _, pc := range playlist.PlaylistContent {
		if pc.Content != nil {
			payload.Contents = append(payload.Contents, ContentResponse{
				ContentID: pc.ContentID,
				Title:     pc.Content.Title,
				URL:       pc.Content.FileURL,
				Order:     pc.Order,
			})
		}
	}

	pretty, _ := json.MarshalIndent(payload, "", "  ")
	fmt.Println("Sending (server log):", string(pretty))

	data, _ := json.Marshal(payload)
	conn.Write(append(data, '\n'))
}
