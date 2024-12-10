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
}

// protoc --proto_path=. --go_out=. --go_opt=Mmarket_feed_p.proto=github.com/Anirudh380/test-bot market_feed.proto
