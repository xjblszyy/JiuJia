package yuemiao

import (
	"go.uber.org/zap"

	"yuemiao/config"
)

type YueMiao struct {
	Requests
	logger  *zap.Logger
	cfg     config.YueMiaoConfig
	vcode   string
	linkMan string
}

func NewYueMiao(logger *zap.Logger, cfg config.YueMiaoConfig) YueMiao {
	headers := make(map[string]string)
	headers["User-Agent"] = UserAgent
	headers["tk"] = cfg.TK
	// todo 是否还需要加其他请求头？！！

	req := NewRequests(logger, cfg.Verbose, headers)
	return YueMiao{
		Requests: req,
		logger:   logger,
		cfg:      cfg,
	}
}
