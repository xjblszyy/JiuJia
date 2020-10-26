package yuemiao

import (
	"testing"

	"go.uber.org/zap"

	"yuemiao/config"
)

const (
	v2verbose   = true
	v2tk        = ""
	v2seckillId = ""
	v2linkmanId = ""
	v2idCardNo  = ""
)

func TestYueMiao_V2(t *testing.T) {
	cfg := config.YueMiaoConfig{
		Verbose:   v2verbose,
		TK:        v2tk,
		SeckillId: v2seckillId,
		LinkmanId: v2linkmanId,
		IdCardNo:  v2idCardNo,
	}

	s := NewYueMiao(zap.L(), cfg)
	s.V2()
}
