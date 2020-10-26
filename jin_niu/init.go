package jin_niu

import (
	"go.uber.org/zap"

	"yuemiao/config"
)

type JinNiu struct {
	Requests
	logger *zap.Logger
	cfg    config.JinNiuConfig
}

func NewJinNiu(logger *zap.Logger, cfg config.JinNiuConfig) JinNiu {
	headers := make(map[string]string)
	headers["User-Agent"] = UserAgent
	headers["cookie"] = cfg.Cookie
	headers["Content-Type"] = "application/x-www-form-urlencoded; charset=UTF-8"

	req := NewRequests(logger, cfg.Verbose, headers)
	return JinNiu{
		Requests: req,
		logger:   logger,
		cfg:      cfg,
	}
}
