package service

import (
	"go.uber.org/zap"
	"jiujia/pkg/http"
	"jiujia/pkg/utils"
)

// SecKill 获取秒杀资格
func (s Service) SecKill(seckillID, linkmanID, idCard, st string) error {
	url := s.baseUrl + "/seckill/seckill/subscribe.do"
	params := map[string]string{
		"seckillId":    seckillID,
		"vaccineIndex": "1",
		"linkmanId":    linkmanID,
		"idCardNo":     idCard,
	}
	headers := utils.CommonHeader()
	headers["ecc-hs"] = utils.EccHs(seckillID, st, s.cfg.MemberID)
	resp, err := http.Get(url, params, headers, nil)
	if err != nil {
		s.logger.Error("get failed", zap.Error(err))
		return err
	}
	s.logger.Info("seckill success", zap.String("data", resp.String()))
	return nil
}
