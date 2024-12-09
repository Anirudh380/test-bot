package login

import (
	"fmt"
	"sync"
	"test-bot/config"
	"test-bot/logger"
	"time"
)

type lockedToken struct {
	mu    sync.Mutex
	token string
}

func (lt *lockedToken) GetToken() string {
	delay := 100 * time.Millisecond
	for {
		if lt.token == "" {
			time.Sleep(delay)
		} else {
			break
		}
		if delay > 1*time.Minute {
			delay *= 2
		}
	}
	return lt.token
}

func (lt *lockedToken) SetToken(newToken string) {
	lt.mu.Lock()
	defer lt.mu.Unlock()
	lt.token = newToken
}

type loginRepositoryImpl struct {
	conf  *config.Config
	Token lockedToken
}

func NewLoginRepository(conf *config.Config, log *logger.Logger) (LoginRepository, error) {
	l := &loginRepositoryImpl{
		conf:  conf,
		Token: lockedToken{},
	}

	l.NewLogin()
	go func() {
		for {
			now := time.Now()
			var nextTick time.Time

			// Find the next closest tick time
			for _, t := range conf.TokenRefreshTimes {
				hour, min := parseTime(t)
				todayTick := time.Date(now.Year(), now.Month(), now.Day(), hour, min, 0, 0, now.Location())

				if todayTick.After(now) {
					nextTick = todayTick
					break
				}
			}

			// If no future tick today, use the first tick time tomorrow
			if nextTick.IsZero() {
				hour, min := parseTime(conf.TokenRefreshTimes[0])
				nextTick = time.Date(now.Year(), now.Month(), now.Day()+1, hour, min, 0, 0, now.Location())
			}

			// Wait until the next tick
			fmt.Println("Next tick scheduled at:", nextTick)
			time.Sleep(time.Until(nextTick))

			// Perform your action
			// RefreshTokenCh <- 1
			l.NewLogin()
			fmt.Println("Tick at:", time.Now())
		}
	}()

	return l, nil
}

func parseTime(t string) (int, int) {
	parts := make([]int, 2)
	fmt.Sscanf(t, "%d:%d", &parts[0], &parts[1])
	return parts[0], parts[1]
}
