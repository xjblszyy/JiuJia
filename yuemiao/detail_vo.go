package yuemiao

import (
	"encoding/json"
	"strconv"

	"go.uber.org/zap"
)

type DetailVoResp struct {
	Code string `json:"code"`
	Data struct {
		ID              int    `json:"id"`
		Name            string `json:"name"`
		VaccineCode     string `json:"vaccineCode"`
		Total           int    `json:"total"`
		Prompt          string `json:"prompt"`
		StartMilliscond int64  `json:"startMilliscond"`
		HospitalName    string `json:"hospitalName"`
		AgeStart        int    `json:"ageStart"`
		AgeEnd          int    `json:"ageEnd"`
		Now             int64  `json:"now"`
		WorkTimeStart   string `json:"workTimeStart"`
		WorkTimeEnd     string `json:"workTimeEnd"`
		PackingImgURL   string `json:"packingImgUrl"`
		Specifications  string `json:"specifications"`
		FactoryName     string `json:"factoryName"`
		IsSubscribeAll  int    `json:"isSubscribeAll"`
		IsSeckill       bool   `json:"isSeckill"`
		Days            []struct {
			Day   string `json:"day"`
			Total int    `json:"total"`
		} `json:"days"`
		Time int64 `json:"time"`
	} `json:"data"`
	Ok bool `json:"ok"`
}

func (y YueMiao) DetailVo(id int) (DetailVoResp, error) {
	res := DetailVoResp{}
	param := map[string]string{
		"id": strconv.Itoa(id),
	}
	resp, err := y.Requests.Get(DetailVoUrl, param, nil)
	if err != nil {
		y.logger.Error("get province failed", zap.Error(err))
		return res, err
	}
	if err := json.Unmarshal(resp, &res); err != nil {
		y.logger.Error("json unmarshal failed", zap.Error(err), zap.Any("body", resp))
	}

	return res, nil
}
