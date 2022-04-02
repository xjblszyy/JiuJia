package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEccHs(t *testing.T) {
	var (
		memberID  = "12372032"
		seckillID = "1085"
		st        = "1630902134216"
	)

	res := EccHs(seckillID, st, memberID)
	assert.Equal(t, res, "7a6e7b94684fa3b50ebf59bc3a76de40")
}
