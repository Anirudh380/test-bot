package feed

import "github.com/gorilla/websocket"

type FeedRepository interface {
	GetMarketFeedUrl(string) (string, error)
	ConnectWebSocket(string, []string) (*websocket.Conn, error)
}

