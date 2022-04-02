package service

import (
	"go.uber.org/zap"
	"jiujia/pkg/http"
	"jiujia/pkg/utils"
)

func (s Service) GetMember() (Member, error) {
	url := s.baseUrl + "/seckill/linkman/findByUserId.do"
	headers := utils.CommonHeader()
	result := Member{}
	_, err := http.Get(url, nil, headers, &result)
	if err != nil {
		s.logger.Error("get failed", zap.Error(err))
		return result, err
	}
	return result, nil
}
