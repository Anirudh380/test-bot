package app

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sync"
	"test-bot/feed"
	"test-bot/login"
	tpr "test-bot/proto"

	"google.golang.org/protobuf/proto"
)

func getInstrumentKeys(limit int) []string {
	file, err := os.Open("complete.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return []string{}
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return []string{}
	}
	res := []string{}
	for i, r := range records {
		if i == 0 {
			continue
		}
		res = append(res, r[0])
	}
	return res[:min(limit, len(res))]
}

type App interface {
	Start()
}

func NewTestBotApp(login login.LoginRepository, feed feed.FeedRepository) App {
	iKeys := getInstrumentKeys(40)
	return &impl{login: login, instrumentKeys: &iKeys, feed: feed}
}

type impl struct {
	login          login.LoginRepository
	instrumentKeys *[]string
	feed           feed.FeedRepository
}

func (i *impl) Start() {

	var wg sync.WaitGroup
	wg.Add(1)

	// Start login process in a goroutine
	go i.login.NewLogin()

	// Fetch token and market feed URL
	token := i.login.GetToken()
	wsUrl, err := i.feed.GetMarketFeedUrl(token)
	if err != nil {
		fmt.Println("Error fetching market feed URL:", err)
		return
	}

	// Connect to WebSocket
	ws, err := i.feed.ConnectWebSocket(wsUrl, *i.instrumentKeys)
	if err != nil {
		fmt.Println("Error connecting to WebSocket:", err)
		return
	}

	defer ws.Close()

	// Read messages from WebSocket
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}

		// Decode the Protobuf message
		var protoMessage tpr.FeedResponse
		err = proto.Unmarshal(message, &protoMessage)
		if err != nil {
			log.Println("Error unmarshaling Protobuf message:", err)
			continue
		}

		fmt.Println(protoMessage)
	}

	wg.Wait()
}

// func (i *impl) Start() {
// 	// Open a file to log messages
// 	file, err := os.Create("messages.log")
// 	if err != nil {
// 		log.Fatalf("Failed to create log file: %v", err)
// 	}
// 	defer file.Close()

// 	// Initialize time tracking
// 	startTime := time.Now()

// 	var wg sync.WaitGroup
// 	wg.Add(1)

// 	// Start login process in a goroutine
// 	go i.login.NewLogin()

// 	// Fetch token and market feed URL
// 	token := i.login.GetToken()
// 	wsUrl, err := i.feed.GetMarketFeedUrl(token)
// 	if err != nil {
// 		fmt.Println("Error fetching market feed URL:", err)
// 		return
// 	}

// 	// Connect to WebSocket
// 	ws, err := i.feed.ConnectWebSocket(wsUrl, *i.instrumentKeys)
// 	if err != nil {
// 		fmt.Println("Error connecting to WebSocket:", err)
// 		return
// 	}

// 	defer ws.Close()

// 	// Read messages from WebSocket
// 	for {
// 		_, message, err := ws.ReadMessage()
// 		if err != nil {
// 			fmt.Println("Error reading message:", err)
// 			return
// 		}

// 		// Log the message with elapsed time
// 		elapsed := time.Since(startTime)
// 		logLine := fmt.Sprintf("%s | Message: %x\n", elapsed.String(), message)

// 		// Write to the file
// 		_, writeErr := file.WriteString(logLine)
// 		if writeErr != nil {
// 			fmt.Println("Error writing to log file:", writeErr)
// 		}

// 		// Decode the Protobuf message
// 		var protoMessage tpr.FeedResponse
// 		err = proto.Unmarshal(message, &protoMessage)
// 		if err != nil {
// 			log.Println("Error unmarshaling Protobuf message:", err)
// 			continue
// 		}

// 		// Print the decoded Protobuf message
// 		// fmt.Println(protoMessage)
// 	}

// 	wg.Wait()
// }
