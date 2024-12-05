package feed

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test-bot/constants"
)

func GetMarketFeedUrl(bearerToken string) (wsUrl string, err error) {
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
