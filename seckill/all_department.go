package seckill

import (
	"encoding/json"

	"go.uber.org/zap"
)

type Departments struct {
	Code string     `json:"code"`
	Data Department `json:"data"`
	Ok   bool       `json:"ok"`
}

type Department struct {
	Offset       int   `json:"offset"`
	End          int   `json:"end"`
	Total        int   `json:"total"`
	Limit        int   `json:"limit"`
	PageNumber   int   `json:"pageNumber"`
	PageListSize int   `json:"pageListSize"`
	PageNumList  []int `json:"pageNumList"`
	Rows         []struct {
		Code          string        `json:"code"`
		Name          string        `json:"name"`
		ImgURL        string        `json:"imgUrl"`
		RegionCode    string        `json:"regionCode"`
		Address       string        `json:"address"`
		Tel           string        `json:"tel"`
		IsOpen        int           `json:"isOpen"`
		Latitude      float64       `json:"latitude"`
		Longitude     float64       `json:"longitude"`
		WorktimeDesc  string        `json:"worktimeDesc"`
		Distance      float64       `json:"distance"`
		VaccineCode   string        `json:"vaccineCode"`
		VaccineName   string        `json:"vaccineName"`
		Total         int           `json:"total"`
		IsSeckill     int           `json:"isSeckill"`
		Price         int           `json:"price"`
		IsHiddenPrice int           `json:"isHiddenPrice,omitempty"`
		DepaCodes     []interface{} `json:"depaCodes"`
		Vaccines      []interface{} `json:"vaccines"`
		DepaVaccID    int           `json:"depaVaccId"`
	} `json:"rows"`
	Pages int `json:"pages"`
}

// 获取所有门诊
func (s AllSteps) GetAllDepartments() (Departments, error) {
	res := Departments{}
	code, err := s.GetCityCode()
	if err != nil {
		s.logger.Error("get city code failed", zap.Error(err))
		return res, err
	}
	body := make(map[string]string)
	body["offset"] = "0"
	body["limit"] = "100"
	body["regionCode"] = code
	body["isOpen"] = "1"
	body["sortType"] = "1"
	body["customId"] = s.cfg.Vaccines
	resp, err := s.Requests.Post(DepartmentUrl, nil, body)
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
