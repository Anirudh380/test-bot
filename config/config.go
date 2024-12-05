package config

const UPSTOX_CLIENT_ID = ""
const UPSTOX_CLIENT_SECRET = ""
const UPSTOX_REDIRECT_URI = "https://www.google.com"
const MARKET_DATA_FEED_URL = "https://api.upstox.com/v2/feed/market-data-feed/authorize"

type Config struct {
	UpstoxClientId     string
	UpstoxClientSecret string
	UpstoxRedirectUri  string
	MarketDataFeedUrl  string
	TokenRefreshTimes  []string
}

func NewConfig() (*Config, error) {
	var config Config
	config.UpstoxClientId = UPSTOX_CLIENT_ID
	config.UpstoxClientSecret = UPSTOX_CLIENT_SECRET
	config.UpstoxRedirectUri = UPSTOX_REDIRECT_URI
	config.MarketDataFeedUrl = MARKET_DATA_FEED_URL
	config.TokenRefreshTimes = []string{"15:30"}

	return &config, nil
}
