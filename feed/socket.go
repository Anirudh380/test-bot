package feed

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func ConnectWebSocket(wsUrl string, instrumentKeys []string) (ws *websocket.Conn, err error) {
	ws, _, err = websocket.DefaultDialer.Dial(wsUrl, nil)
	if err != nil {
		return nil, err
	}

	go func() {
		time.Sleep(1 * time.Second)
		data := map[string]any{
			"guid":   uuid.New().String(),
			"method": "sub",
			"data": map[string]any{
				"mode":           "ltpc",
				"instrumentKeys": instrumentKeys,
			},
		}
		buff, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error marshalling subscription message:", err)
			return
		}
		if err := ws.WriteMessage(websocket.BinaryMessage, buff); err != nil {
			fmt.Println("Error sending subscription message:", err)
		}
		fmt.Println("Sent subscription message")
	}()

	return ws, nil
}
