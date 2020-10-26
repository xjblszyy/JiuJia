package jin_niu

import (
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
)

type ScheduleDoctorList struct {
	Code int `json:"code"`
	Data struct {
		DoctorList   DoctorList `json:"doctorList"`
		VisitNum     int        `json:"visitNum"`
		ScheduleDate string     `json:"scheduleDate"`
		RegisterNum  int        `json:"registerNum"`
	} `json:"data"`
}

type DoctorList []struct {
	CanSubscribe    int    `json:"canSubscribe"`
	DeptName        string `json:"deptName"`
	DoctorID        string `json:"doctorId"`
	DoctorImg       string `json:"doctorImg"`
	DoctorLevel     int    `json:"doctorLevel"`
	DoctorName      string `json:"doctorName"`
	DoctorRemark    string `json:"doctorRemark"`
	DoctorSex       string `json:"doctorSex"`
	DoctorSkill     string `json:"doctorSkill"`
	DoctorTitle     string `json:"doctorTitle"`
	LeftSource      int    `json:"leftSource"`
	RegisterFee     int    `json:"registerFee"`
	ScheduleDate    int64  `json:"scheduleDate"`
	SortNo          int    `json:"sortNo"`
	Status          int    `json:"status"`
	SubscribeStatus int    `json:"subscribeStatus"`
	TotalSource     int    `json:"totalSource"`
}

func (j JinNiu) DoctorList(date string) (DoctorList, error) {
	res := ScheduleDoctorList{}

	param := map[string]string{
		"hisId":          j.cfg.HisID,
		"platformSource": j.cfg.PlatformSource,
		"platformId":     j.cfg.PlatformID,
		"_route":         fmt.Sprintf("h%s", j.cfg.HisID),
	}
	body := map[string]string{
		"deptId":       DeptID,
		"scheduleDate": date,
	}

	resp, err := j.Post(ScheduleDoctorListUrl, param, body)
	if err != nil {
		j.logger.Error("post failed", zap.Error(err))
		return nil, err
	}

	if err := json.Unmarshal(resp, &res); err != nil {
		j.logger.Error("json unmarshal failed", zap.Error(err), zap.Any("body", resp))
		return nil, err
	}
	return res.Data.DoctorList, nil
}

func (j JinNiu) JiuJiaDoctorList(list DoctorList) DoctorList {
	res := DoctorList{}
	for i := 0; i < len(list); i++ {
		// 仅限乙肝疫苗加强针
		// 九价疫苗预约
		if list[i].DoctorName == "九价疫苗预约" && list[i].LeftSource > 0 {
			res = append(res, list[i])
		}
	}
	return res
}
