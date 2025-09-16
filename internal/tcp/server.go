package tcp

import (
	"crypto/tls"
	"digital_signage_api/internal/db"
	"digital_signage_api/internal/dto"
	"digital_signage_api/internal/models"
	"digital_signage_api/internal/utils"
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"time"
)

var clientCount int
var mu sync.Mutex

func StartTCPServer() {
	certFile := "/etc/certs/fullchain.pem"
	keyFile  := "/etc/certs/privkey.pem"

	cer, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		panic(fmt.Sprintf("failed to load cert: %v", err))
	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}}

	listener, err := tls.Listen("tcp", ":9000", config)
	if err != nil {
		panic(fmt.Sprintf("failed to start TLS listener: %v", err))
	}
	fmt.Println("🎧 TLS TCP server listening on :9000")

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
	payload := dto.PlaylistPayloadDTO{
		ScheduleID: schedule.ScheduleID,
		PlaylistID: playlist.PlaylistID,
		Name:       playlist.Name,
		Contents:   []dto.TCPContentResponseDTO{},
	}

	for _, pc := range playlist.PlaylistContent {
		if pc.Content != nil {
			payload.Contents = append(payload.Contents, dto.TCPContentResponseDTO{
				ContentID: pc.ContentID,
				Title:     pc.Content.Title,
				URL:       utils.BuildContentURL(pc.Content.Title),
				Order:     pc.Order,
			})
		}
	}

	pretty, _ := json.MarshalIndent(payload, "", "  ")
	fmt.Println("Sending (server log):", string(pretty))

	data, _ := json.Marshal(payload)
	conn.Write(append(data, '\n'))
}
