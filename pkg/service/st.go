package service

import (
	"strconv"

	"go.uber.org/zap"
	"jiujia/pkg/http"
	"jiujia/pkg/utils"
)

// GetST 获取加密参数st
func (s Service) GetST(vaccineID string) (string, error) {
	url := s.baseUrl + "/seckill/seckill/checkstock2.do"
	params := map[string]string{
		"id": vaccineID,
	}
	headers := utils.CommonHeader()
	result := ST{}
	_, err := http.Get(url, params, headers, &result)
	if err != nil {
		s.logger.Error("get failed", zap.Error(err))
		return "", err
	}
	return strconv.FormatInt(result.Data.St, 10), nil
}
