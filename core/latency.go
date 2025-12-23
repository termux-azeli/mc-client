package core

import (
	"encoding/binary"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const PROTO_PING = 0xFF

func HandlePing(ws *websocket.Conn) {
	go func() {
		for range time.Tick(1 * time.Second) {
			buf := make([]byte, 9)
			buf[0] = PROTO_PING
			binary.BigEndian.PutUint64(buf[1:], uint64(time.Now().UnixNano()))
			ws.WriteMessage(websocket.BinaryMessage, buf)
		}
	}()

	go func() {
		for {
			_, data, err := ws.ReadMessage()
			if err != nil {
				return
			}
			if data[0] == PROTO_PING {
				ts := int64(binary.BigEndian.Uint64(data[1:]))
				log.Println("[RTT]", time.Since(time.Unix(0, ts)))
			}
		}
	}()
}
