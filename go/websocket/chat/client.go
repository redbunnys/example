package chat

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"go-websocket/Common"
	"log"
	"time"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

const (
	// 允許向對等方寫入訊息的時間。
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	//允許從對等方讀取下一條 pong 訊息的時間
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	//以此期間向對等方發送 ping。 必須小於 pongWait。
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	//允許從對等方發送的最大訊息大小。
	maxMessageSize = 512
)

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
	Send chan []byte
	Name string
}

type ClientMsg struct {
	Name string `json:"name"`
	Msg  string `json:"msg"`
}

func (cm *ClientMsg) SerializationJson() []byte {
	byte, err := json.Marshal(cm)
	if err != nil {
		return nil
	}

	return byte
}
func (cm *ClientMsg) DeserializationJson(msg interface{}) {
	er := json.Unmarshal([]byte(msg.(string)), &cm)
	if er != nil {
		log.Println(er)
	}
}

func (c *Client) RandomName() {
	c.Name = Common.GetFullName()
}
func (c *Client) SetName(name string) {
	c.Name = name
}

func (c *Client) Read() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	// 設置讀取限制
	c.Conn.SetReadLimit(maxMessageSize)
	// 設置讀取超時
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	// 設置 pong 訊息處理函式
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	// 無限循環讀取訊息
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			// 如果是正常關閉，則不處
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		// 將訊息發送到廣播頻道
		msg := &ClientMsg{
			Name: c.Name,
			Msg:  string(message),
		}
		//message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		// 將訊息發送到廣播頻道
		c.Hub.Broadcast <- msg.SerializationJson()
	}

}

func (c *Client) Write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	// 無限循環
	for {
		select {
		case message, ok := <-c.Send:
			// 設置寫入超時
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// 關閉連接
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			// 寫入訊息
			w.Write(message)

			// 將排隊的緩衝區中的訊息寫入連接
			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.Send)
			}
			// 寫入訊息
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			// 設置寫入超時
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			// 寫入 ping 訊息
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
