package main

import (
	"log"
	"test-bot/wire"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			log.Println("[Recovery from panic]", err)
		}
	}()

	app, err := wire.InitializeApp()
	if err != nil {
		log.Fatalln("Couldn't start the app. Error ->", err)
		return
	}

	app.Start()

	// bearerToken := login.GetToken()

	// res := getInstrumentKeys()
	// start := time.Now()
	// var wg sync.WaitGroup
	// wg.Add(1)

	// // wsUrl, err := feed.GetMarketFeedUrl(bearerToken)
	// // if err != nil {
	// // 	fmt.Println("Error getting WebSocket URL:", err)
	// // 	return
	// // }

	// // ws, err := feed.ConnectWebSocket(wsUrl, res[1:2])
	// // if err != nil {
	// // 	fmt.Println("Error connecting to WebSocket:", err)
	// // }
	// // for {
	// // 	_, message, err := ws.ReadMessage()
	// // 	if err != nil {
	// // 		fmt.Println("Error reading message:", err)
	// // 		return
	// // 	}
	// // 	fmt.Println("Received message:", string(message))

	// // 	var protoMessage tpr.FeedResponse
	// // 	err = proto.Unmarshal(message, &protoMessage)
	// // 	if err != nil {
	// // 		log.Println("Error unmarshaling Protobuf message:", err)
	// // 		continue
	// // 	}
	// // }

	// for i := 0; i < len(res); i += 90 {
	// 	fmt.Println(i)
	// 	j := i
	// 	go func(k int) {
	// 		wsUrl, err := feed.GetMarketFeedUrl(bearerToken)
	// 		if err != nil {
	// 			fmt.Println("Error getting WebSocket URL:", err)
	// 			return
	// 		}
	// 		keys := []string{}
	// 		limit := k + 90
	// 		for ; k < min(len(res)-1, limit); k++ {
	// 			keys = append(keys, res[j])
	// 		}
	// 		_, err = feed.ConnectWebSocket(wsUrl, keys)
	// 		if err != nil {
	// 			fmt.Println("Error connecting to WebSocket:", err)
	// 		}
	// 		// for {
	// 		// 	_, message, err := ws.ReadMessage()
	// 		// 	if err != nil {
	// 		// 		fmt.Println("Error reading message:", err)
	// 		// 		return
	// 		// 	}
	// 		// 	fmt.Println("Received message:", string(message))

	// 		// 	var protoMessage tpr.FeedResponse
	// 		// 	err = proto.Unmarshal(message, &protoMessage)
	// 		// 	if err != nil {
	// 		// 		log.Println("Error unmarshaling Protobuf message:", err)
	// 		// 		continue
	// 		// 	}
	// 		// }
	// 	}(j)
	// }

	// elapsed := time.Since(start)
	// fmt.Printf("Time taken: %s\n", elapsed)

	// wg.Wait()

	// // wsUrl, err := feed.GetMarketFeedUrl(bearerToken)
	// // if err != nil {
	// // 	fmt.Println("Error getting WebSocket URL:", err)
	// // 	return
	// // }
	// // fmt.Println(wsUrl)
	// // ws, err := feed.ConnectWebSocket(wsUrl, []string{"NSE_INDEX|Nifty Bank"})
	// // if err != nil {
	// // 	fmt.Println("Error connecting to WebSocket:", err)
	// // 	return
	// // }
	// // defer ws.Close()

	// // fmt.Println("Connected to WebSocket")

	// // // Read messages from WebSocket
	// // for {
	// // 	_, message, err := ws.ReadMessage()
	// // 	if err != nil {
	// // 		fmt.Println("Error reading message:", err)
	// // 		return
	// // 	}
	// // 	fmt.Println("Received message:", string(message))

	// // 	var protoMessage tpr.FeedResponse
	// // 	err = proto.Unmarshal(message, &protoMessage)
	// // 	if err != nil {
	// // 		log.Println("Error unmarshaling Protobuf message:", err)
	// // 		continue
	// // 	}

	// // 	// Process the message
	// // 	// fmt.Printf("Received message: %+v\n", protoMessage)

	// // 	jsonBytes, err := protojson.Marshal(&protoMessage)
	// // 	if err != nil {
	// // 		log.Println("Error converting Protobuf to JSON:", err)
	// // 		return
	// // 	}

	// // 	var prettyJSON map[string]interface{}
	// // 	if err := json.Unmarshal(jsonBytes, &prettyJSON); err != nil {
	// // 		log.Println("Error formatting JSON:", err)
	// // 		return
	// // 	}

	// // 	prettyFormattedJSON, _ := json.MarshalIndent(prettyJSON, "", "  ")
	// // 	log.Println("JSON Representation:")
	// // 	log.Println(string(prettyFormattedJSON))
	// // }

}

// protoc --proto_path=. --go_out=. --go_opt=Mmarket_feed_p.proto=github.com/Anirudh380/test-bot market_feed.proto
