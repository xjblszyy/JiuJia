package service

import (
	"go.uber.org/zap"
	"jiujia/pkg/http"
	"jiujia/pkg/utils"
)

// GetVaccineList 获取疫苗列表
func (s Service) GetVaccineList(regionCode string) (VaccineList, error) {
	url := s.baseUrl + "/seckill/seckill/list.do"
	params := map[string]string{
		"offset":     "0",
		"limit":      "100",
		"regionCode": regionCode, // 4位，例如成都：5101
	}
	headers := utils.CommonHeader()
	result := VaccineList{}
	resp, err := http.Get(url, params, headers, &result)
	if err != nil {
		s.logger.Error("get failed", zap.Error(err))
		return result, err
	}
	s.logger.Debug("get vaccine list success", zap.String("data", resp.String()))

	return result, nil
}
