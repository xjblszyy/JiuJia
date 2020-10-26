package seckill

import (
	"encoding/json"
	"strconv"

	"go.uber.org/zap"
)

type WorkTimeResp struct {
	Code string `json:"code"`
	Data struct {
		Times struct {
			Code string `json:"code"`
			Data []struct {
				CreateTime string `json:"createTime"`
				DepaCode   string `json:"depaCode"`
				EndTime    string `json:"endTime"`
				ID         int    `json:"id"`
				MaxSub     int    `json:"maxSub"`
				ModifyTime string `json:"modifyTime,omitempty"`
				StartTime  string `json:"startTime"`
				WorkdayID  int    `json:"workdayId"`
				Yn         int    `json:"yn"`
				TIndex     int    `json:"tIndex,omitempty"`
			} `json:"data"`
			NotOk bool `json:"notOk"`
			Ok    bool `json:"ok"`
		} `json:"times"`
		Now int64 `json:"now"`
	} `json:"data"`
	NotOk bool `json:"notOk"`
	Ok    bool `json:"ok"`
}

// 获取这个门诊的可预约日期下的时间
func (s AllSteps) GetWorkTime(info DepartmentInfo, time string) (WorkTimeResp, error) {
	res := WorkTimeResp{}

	body := make(map[string]string)
	body["depaCode"] = info.Data.DepartmentCode
	body["linkmanId"] = "2772838"
	body["vaccCode"] = info.Data.VaccineCode
	body["vaccIndex"] = "1"
	body["departmentVaccineId"] = strconv.Itoa(info.Data.ID)
	body["subsribeDate"] = time

	resp, err := s.Requests.Get(WorkTimeUrl, body, nil)
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
