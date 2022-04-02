package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"jiujia/config"
	"jiujia/pkg/http"
)

func TestSecKill(t *testing.T) {
	cfg := config.Config{
		TK:         "",
		MemberID:   "",
		IDCard:     "",
		RegionCode: "",
		Cookie:     "",
	}
	var (
		vaccineID = ""
		linkmanID = cfg.MemberID
		idCard    = cfg.IDCard
		st        = ""
	)
	config.C = &cfg
	http.InitClient("", true)
	logger := zap.L()
	s := New(cfg, logger)
	err := s.SecKill(vaccineID, linkmanID, idCard, st)
	assert.NoError(t, err)
}
