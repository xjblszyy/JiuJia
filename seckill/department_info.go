package seckill

import (
	"encoding/json"
	"strconv"

	"go.uber.org/zap"
)

type DepartmentInfo struct {
	Code string `json:"code"`
	Data struct {
		ID               int           `json:"id"`
		DepartmentCode   string        `json:"departmentCode"`
		VaccineCode      string        `json:"vaccineCode"`
		DepartmentName   string        `json:"departmentName"`
		Describtion      string        `json:"describtion"`
		InstructionsUrls []interface{} `json:"instructionsUrls"`
		IsArriveVaccine  int           `json:"isArriveVaccine"`
		Name             string        `json:"name"`
		Prompt           string        `json:"prompt"`
		Subscribed       int           `json:"subscribed"`
		Total            int           `json:"total"`
		Urls             []string      `json:"urls"`
		Items            []struct {
			ID             int    `json:"id"`
			VaccineCode    string `json:"vaccineCode"`
			FactoryName    string `json:"factoryName"`
			Specifications string `json:"specifications"`
			Name           string `json:"name"`
			Price          int    `json:"price"`
		} `json:"items"`
	} `json:"data"`
	Ok bool `json:"ok"`
}

// 获取某一个门诊详情数据
func (s AllSteps) FetchDepartmentInfo(id int) (DepartmentInfo, error) {
	res := DepartmentInfo{}

	body := make(map[string]string)
	body["id"] = strconv.Itoa(id)
	body["isShowDescribtion"] = "true"
	resp, err := s.Requests.Get(DepartmentInfoUrl, body, nil)
	if err != nil {
		s.logger.Error("get all department failed", zap.Error(err))
		return res, err
	}
	err = json.Unmarshal(resp, &res)
	if err != nil {
		s.logger.Error("json unmarshal failed", zap.Error(err))
		return res, err
	}
	return res, nil
}
