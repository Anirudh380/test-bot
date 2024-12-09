package feed

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test-bot/constants"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func (m *feedRepositoryImpl) GetMarketFeedUrl(bearerToken string) (wsUrl string, err error) {
	req, err := http.NewRequest(http.MethodGet, constants.MARKET_DATA_FEED_URL, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+bearerToken)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	fmt.Println(res.StatusCode)

	if res.StatusCode != http.StatusOK {
		fmt.Println("Failed to authorize")
		return
	}

	response := make(map[string]any)
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		fmt.Println(err)
		return
	}

	return response["data"].(map[string]any)["authorizedRedirectUri"].(string), nil
}

func (m *feedRepositoryImpl) ConnectWebSocket(wsUrl string, instrumentKeys []string) (ws *websocket.Conn, err error) {
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
