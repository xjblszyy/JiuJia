package seckill

import (
	"go.uber.org/zap"

	"yuemiao/config"
)

type AllSteps struct {
	Requests
	logger *zap.Logger
	cfg    config.YueMiaoConfig
}

func NewAllSteps(logger *zap.Logger, cfg config.YueMiaoConfig) AllSteps {
	headers := make(map[string]string)
	headers["User-Agent"] = UserAgent
	headers["tk"] = cfg.TK

	req := NewRequests(logger, cfg.Verbose, headers)
	return AllSteps{
		Requests: req,
		logger:   logger,
		cfg:      cfg,
	}
}
