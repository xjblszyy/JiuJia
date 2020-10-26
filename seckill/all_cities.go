package seckill

import (
	"encoding/json"
	"errors"
	"fmt"

	"go.uber.org/zap"
)

type Location struct {
	Code string `json:"code"`
	Data []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"data"`
	Ok bool `json:"ok"`
}

func (s AllSteps) GetProvinceCode() (string, error) {
	// if config.C.JiuJia.Province == "四川省"{
	// 	return "51", nil
	// }

	res := Location{}
	resp, err := s.Requests.Get(AllCitiesUrl, nil, nil)
	if err != nil {
		s.logger.Error("get province failed", zap.Error(err))
		return "", err
	}
	if err := json.Unmarshal(resp, &res); err != nil {
		s.logger.Error("json unmarshal failed", zap.Error(err), zap.Any("body", resp))
	}
	for i := 0; i < len(res.Data); i++ {
		if res.Data[i].Name == s.cfg.Province {
			return res.Data[i].Value, nil
		}
	}
	return "", errors.New("省份不存在")
}

func (s AllSteps) GetCityCode() (string, error) {
	var err error
	var provinceCode string
	//
	// if config.C.JiuJia.City != "成都市" {
	// 	return "5101", nil
	// }

	provinceCode, err = s.GetProvinceCode()
	if err != nil {
		s.logger.Error("get province failed", zap.Error(err))
		return "", err
	}

	res := Location{}
	resp, err := s.Requests.Get(AllCitiesUrl, map[string]string{"parentCode": provinceCode}, nil)
	if err != nil {
		s.logger.Error("get all cities failed", zap.Error(err))
		return "", err
	}

	if err := json.Unmarshal(resp, &res); err != nil {
		s.logger.Error("json unmarshal failed", zap.Error(err), zap.Any("body", resp))
	}
	for i := 0; i < len(res.Data); i++ {
		if res.Data[i].Name == s.cfg.City {
			return res.Data[i].Value, nil
		}
	}
	return "", errors.New("城市不存在")
}

type AllCitiesCode struct {
	Code string `json:"code"`
	Data []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"data"`
	Ok bool `json:"ok"`
}

// 获取所有的城市的编码
func (s AllSteps) GetAllCitiesCode() (map[string]string, error) {

	province := AllCitiesCode{}
	city := AllCitiesCode{}
	r := make(map[string]string)

	resp, err := s.Requests.Get(CitiesCodeUrl, nil, nil)
	if err != nil {
		s.logger.Error("get province failed", zap.Error(err))
		return nil, err
	}
	if err := json.Unmarshal(resp, &province); err != nil {
		s.logger.Error("json unmarshal failed", zap.Error(err), zap.Any("body", resp))
		return nil, err
	}

	for i := 0; i < len(province.Data); i++ {
		resp2, err := s.Requests.Get(CitiesCodeUrl, map[string]string{"parentCode": province.Data[i].Value}, nil)
		if err != nil {
			s.logger.Error("get city failed", zap.Error(err))
			return nil, err
		}
		if err := json.Unmarshal(resp2, &city); err != nil {
			s.logger.Error("json unmarshal failed", zap.Error(err), zap.Any("body", resp2))
			return nil, err
		}
		for j := 0; j < len(city.Data); j++ {
			r[city.Data[j].Value] = fmt.Sprintf("%s-%s", province.Data[i].Name, city.Data[j].Name)
		}
	}

	return r, nil
}
