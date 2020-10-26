package vcode

import (
	"go.uber.org/zap"

	"yuemiao/config"
)

type VCode struct {
	Identify
	logger *zap.Logger
	cfg    config.VCodeConfig
}

func NewVCode(logger *zap.Logger, cfg config.VCodeConfig) VCode {
	headers := make(map[string]string)
	headers["appKey"] = cfg.AppKey
	headers["appCode"] = cfg.AppCode
	headers["appSecret"] = cfg.AppSecret

	identify := NewIdentify(logger, false, headers)
	return VCode{
		Identify: identify,
		logger:   logger,
		cfg:      cfg,
	}
}
