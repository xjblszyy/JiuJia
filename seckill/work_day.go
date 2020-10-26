package seckill

import (
	"encoding/json"
	"strconv"

	"go.uber.org/zap"
)

type WorkDayResp struct {
	Code string `json:"code"`
	Data struct {
		DateList      []string `json:"dateList"`
		SubscribeDays int      `json:"subscribeDays"`
	} `json:"data"`
	NotOk bool `json:"notOk"`
	Ok    bool `json:"ok"`
}

// 获取这个门诊的可预约日期
func (s AllSteps) GetWorkDay(info DepartmentInfo) (WorkDayResp, error) {
	res := WorkDayResp{}

	body := make(map[string]string)
	body["depaCode"] = info.Data.DepartmentCode
	body["linkmanId"] = "2772838"
	body["vaccCode"] = info.Data.VaccineCode
	body["vaccIndex"] = "1"
	body["departmentVaccineId"] = strconv.Itoa(info.Data.ID)

	resp, err := s.Requests.Get(WorkDayUrl, body, nil)
	if err != nil {
		s.logger.Error("get work day failed", zap.Error(err))
		return res, err
	}
	err = json.Unmarshal(resp, &res)
	if err != nil {
		s.logger.Error("json unmarshal failed", zap.Error(err))
		return res, err
	}
	return res, nil
}
