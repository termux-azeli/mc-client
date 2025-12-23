package core

import (
	"encoding/binary"
	"log"
	"net"
	"net/url"

	"github.com/gorilla/websocket"
)

const PROTO_BEDROCK = 0x01

func RunBedrock(cfg Config) error {
	u := url.URL{
		Scheme: "wss",
		Host:   cfg.Host,
		Path:   cfg.Path,
	}
	ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	udp, _ := net.ListenUDP("udp", nil)
	defer udp.Close()

	myID := uint16(1)

	HandlePing(ws)

	// UDP -> WS
	go func() {
		buf := make([]byte, 2048)
		for {
			n, _, _ := udp.ReadFromUDP(buf)
			packet := make([]byte, 3+n)
			packet[0] = PROTO_BEDROCK
			binary.BigEndian.PutUint16(packet[1:3], myID)
			copy(packet[3:], buf[:n])
			ws.WriteMessage(websocket.BinaryMessage, packet)
		}
	}()

	// WS -> UDP
	for {
		_, data, _ := ws.ReadMessage()
		if data[0] == PROTO_BEDROCK {
			udp.WriteToUDP(data[3:], &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 19132})
		}
	}
}
