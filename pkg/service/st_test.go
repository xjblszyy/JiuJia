package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"jiujia/config"
	"jiujia/pkg/http"
)

func TestST(t *testing.T) {
	cfg := config.Config{
		TK:         "",
		MemberID:   "",
		IDCard:     "",
		RegionCode: "",
		Cookie:     "",
	}
	var (
		vaccineID = "1085"
	)
	config.C = &cfg
	http.InitClient("", true)
	logger := zap.L()
	s := New(cfg, logger)
	st, err := s.GetST(vaccineID)
	assert.NoError(t, err)
	assert.NotEmpty(t, st)
}
