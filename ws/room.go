package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// Client menyimpan koneksi WebSocket dan ID user
type Client struct {
	Conn *websocket.Conn
	ID   string
}

// Struktur fleksibel untuk broadcast
type Broadcast struct {
	Channel string          `json:"channel"`
	IDUser  string          `json:"id-user"`
	Message json.RawMessage `json:"message"` // isi bebas, fleksibel
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var (
	channels = make(map[string][]Client)
	mu       sync.Mutex
)

// /ws/room?channel=abc&id=nadia
func RoomHandler(w http.ResponseWriter, r *http.Request) {
	channel := r.URL.Query().Get("channel")
	id := r.URL.Query().Get("id")

	if channel == "" || id == "" {
		http.Error(w, "Parameter 'channel' dan 'id' wajib diisi", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "WebSocket upgrade gagal", http.StatusInternalServerError)
		return
	}

	client := Client{Conn: conn, ID: id}

	// Tambahkan ke channel
	mu.Lock()
	channels[channel] = append(channels[channel], client)
	mu.Unlock()

	fmt.Printf("✅ %s bergabung ke channel [%s]\n", id, channel)

	defer func() {
		// Hapus client saat disconnect
		mu.Lock()
		newList := []Client{}
		for _, c := range channels[channel] {
			if c.Conn != conn {
				newList = append(newList, c)
			}
		}
		channels[channel] = newList
		mu.Unlock()
		conn.Close()
		fmt.Printf("❌ %s keluar dari channel [%s]\n", id, channel)
	}()

	// Loop pesan masuk dan broadcast
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break // koneksi putus
		}

		// Pastikan pesan JSON valid
		var test map[string]interface{}
		if err := json.Unmarshal(msg, &test); err != nil {
			fmt.Println("⚠️ Pesan dari", id, "bukan JSON valid:", string(msg))
			continue
		}

		b := Broadcast{
			Channel: channel,
			IDUser:  id,
			Message: json.RawMessage(msg), // ✅ bungkus sebagai RawMessage
		}

		data, err := json.Marshal(b)
		if err != nil {
			fmt.Println("❌ Gagal marshal broadcast:", err)
			continue
		}

		// Kirim ke semua client (termasuk pengirim)
		mu.Lock()
		for _, c := range channels[channel] {
			c.Conn.WriteMessage(websocket.TextMessage, data)
		}
		mu.Unlock()
	}

}
