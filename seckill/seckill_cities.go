// 查看有哪些可以秒杀疫苗的城市信息
package seckill

import (
	"encoding/json"

	"go.uber.org/zap"
)

func (s AllSteps) GetSeckillCities() ([]string, error) {
	var res []string

	cityCodes, err := s.GetAllCitiesCode()
	if err != nil {
		s.logger.Error("获取所有城市编码有误", zap.Error(err))
		return res, err
	}
	for code, city := range cityCodes {
		hasSeckill, err := s.hasSeckill(code)
		if err != nil {
			s.logger.Error("判断是否有秒杀城市失败", zap.Error(err), zap.Any("code", code), zap.Any("city", city))
			continue
		}
		if hasSeckill {
			res = append(res, city)
		}
	}
	return res, nil
}

type HasSeckill struct {
	Code  string        `json:"code"`
	Data  []interface{} `json:"data"`
	Ok    bool          `json:"ok"`
	NotOk bool          `json:"notOk"`
}

func (s AllSteps) hasSeckill(code string) (bool, error) {
	res := HasSeckill{}
	param := map[string]string{
		"regionCode": code,
		"offset":     "0",
		"limit":      "10",
	}
	resp, err := s.Requests.Get(HasSeckillUrl, param, nil)
	if err != nil {
		s.logger.Error("get province failed", zap.Error(err))
		return false, err
	}
	if err := json.Unmarshal(resp, &res); err != nil {
		s.logger.Error("json unmarshal failed", zap.Error(err), zap.Any("body", resp))
		return false, err
	}
	if len(res.Data) != 0 {
		return true, nil
	}
	return false, nil
}
