package service

import (
	"go.uber.org/zap"
	"jiujia/pkg/http"
	"jiujia/pkg/utils"
)

func (s Service) GetArea(parentCode string) (Area, error) {
	url := s.baseUrl + "/base/region/childRegions.do"
	headers := utils.CommonHeader()
	param := map[string]string{
		"parentCode": parentCode,
	}
	result := Area{}
	_, err := http.Get(url, param, headers, &result)
	if err != nil {
		s.logger.Error("get failed", zap.Error(err))
		return result, err
	}
	return result, nil
}
