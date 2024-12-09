package feed

import (
	"test-bot/config"
	"test-bot/logger"
)

type feedRepositoryImpl struct {
	conf *config.Config
}

func NewFeedRepository(conf *config.Config, log *logger.Logger) FeedRepository {
	return &feedRepositoryImpl{conf: conf}
}
