package service

import (
	"go.uber.org/zap"
	"jiujia/config"
)

type Service struct {
	cfg     config.Config
	baseUrl string
	logger  *zap.Logger
}

func New(cfg config.Config, logger *zap.Logger) Service {
	return Service{
		cfg:     cfg,
		baseUrl: "https://miaomiao.scmttec.com",
		logger:  logger,
	}
}
