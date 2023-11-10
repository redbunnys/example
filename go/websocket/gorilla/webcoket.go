package gorilla

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func GorillaWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			panic(err)
		}
		m := string(msg)
		if find := strings.Contains(m, "close"); find {
			conn.Close()
			break
		}
		wmsg := fmt.Sprintf("receive msg: %s \n", m)

		conn.WriteMessage(websocket.TextMessage, []byte(wmsg))
	}

}
