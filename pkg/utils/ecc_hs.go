package utils

import (
	"crypto/md5"
	"encoding/hex"

	"jiujia/pkg/constants"
)

func EccHs(seckillId, st, memberID string) string {
	salt := constants.Salt
	h1 := md5.New()
	h1.Write([]byte(seckillId + memberID + st))
	data1 := hex.EncodeToString(h1.Sum(nil))

	h2 := md5.New()
	h2.Write([]byte(data1 + salt))
	return hex.EncodeToString(h2.Sum(nil))
}
