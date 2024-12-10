package replay

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
	tpr "test-bot/proto"
	"time"

	"google.golang.org/protobuf/proto"
)

func Replay() {
	file, err := os.Open("messages.log")

	if err != nil {
		fmt.Errorf("failed to open log file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	startTime := time.Now()
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " | ")

		// Parse the elapsed time and message
		if len(parts) != 2 {
			continue
		}

		// Parse the elapsed time
		elapsedStr := parts[0]
		elapsedTime, err := time.ParseDuration(elapsedStr)
		if err != nil {
			log.Println("Error parsing elapsed time:", err)
			continue
		}

		// Parse the message in hexadecimal format
		messageHex := parts[1][9:] // Skip "Message: "
		messageBytes, err := hex.DecodeString(messageHex)
		if err != nil {
			log.Println("Error decoding message:", err)
			continue
		}

		var protoMessage tpr.FeedResponse
		err = proto.Unmarshal(messageBytes, &protoMessage)
		if err != nil {
			log.Println("Error unmarshaling Protobuf message:", err)
			continue
		}
		fmt.Println(protoMessage)
		// Wait until the appropriate time for replay
		time.Sleep(elapsedTime - time.Since(startTime))

	}
}
