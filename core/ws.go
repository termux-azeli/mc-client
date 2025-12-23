package core

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type WSClient struct {
	Conn       *websocket.Conn
	WriteQueue chan []byte
}

func NewWSClient(conn *websocket.Conn) *WSClient {
	client := &WSClient{
		Conn:       conn,
		WriteQueue: make(chan []byte, 512),
	}
	go client.writer()
	return client
}

// writer: handle micro-batching <2ms
func (c *WSClient) writer() {
	if !ENABLE_BATCHING {
		for msg := range c.WriteQueue {
			c.Conn.WriteMessage(websocket.BinaryMessage, msg)
		}
		return
	}

	buf := make([]byte, 0, BATCH_MAX_SIZE)
	ticker := time.NewTicker(BATCH_WINDOW)
	defer ticker.Stop()

	flush := func() {
		if len(buf) > 0 {
			err := c.Conn.WriteMessage(websocket.BinaryMessage, buf)
			if err != nil {
				log.Println("WS write error:", err)
			}
			buf = buf[:0]
		}
	}

	for {
		select {
		case msg := <-c.WriteQueue:
			if len(buf)+len(msg) > BATCH_MAX_SIZE {
				flush()
			}
			buf = append(buf, msg...)
		case <-ticker.C:
			flush()
		}
	}
}

// SendMessage: push ke write queue
func (c *WSClient) SendMessage(data []byte) {
	c.WriteQueue <- data
}

// ReadLoop: optional helper to read messages
func (c *WSClient) ReadLoop(onMessage func([]byte)) {
	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("WS read error:", err)
			return
		}
		onMessage(msg)
	}
}
