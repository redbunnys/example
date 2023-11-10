package chat

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

type Hub struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			// 註冊
			h.Clients[client] = true

			log.Println("register client: " + client.Name)
		case client := <-h.Unregister:
			// 取消註冊
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}

			log.Println("unregister client: " + client.Name)

		case message := <-h.Broadcast:
			// 廣播
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
			log.Println("broadcast message: " + string(message))
		}
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // 允許跨域
	},
}

func ChatWs(w http.ResponseWriter, r *http.Request, hub *Hub) {
	if r.Method != "GET" {
		// 請求方法不是 GET
		http.Error(w, "Method not allowed", 405)
		return
	}
	anonymouseBool := true
	var err error
	name := r.URL.Query().Get("name")
	anonymouse := r.URL.Query().Get("anonymouse")
	anonymouseBool, err = strconv.ParseBool(anonymouse)
	if err != nil {
		anonymouseBool = true
	}
	// 升級 HTTP 請求為 WebSocket 連接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	// 建立 Client
	client := &Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	if anonymouseBool {
		client.RandomName()
	} else {
		client.SetName(name)
	}
	log.Println("client name: " + client.Name)
	// 註冊 Client
	client.Hub.Register <- client
	// 啟動寫入
	go client.Write()
	// 啟動讀取
	go client.Read()

}
