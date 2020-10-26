package jin_niu

import (
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
)

type DoctorTimeList struct {
	Code int `json:"code"`
	Data struct {
		DoctorID     string   `json:"doctorId"`
		HisID        int      `json:"hisId"`
		ScheduleDate string   `json:"scheduleDate"`
		DeptID       string   `json:"deptId"`
		ItemList     ItemList `json:"itemList"`
	} `json:"data"`
}

type ItemList []struct {
	LeftSource     int    `json:"leftSource"`
	RegisterFee    int    `json:"registerFee"`
	ScheduleID     string `json:"scheduleId"`
	Status         int    `json:"status"`
	VisitBeginTime string `json:"visitBeginTime"`
	VisitEndTime   string `json:"visitEndTime"`
}

func (j JinNiu) DoctorTimeList(date, doctorId string) (ItemList, error) {
	res := DoctorTimeList{}

	param := map[string]string{
		"hisId":          j.cfg.HisID,
		"platformSource": j.cfg.PlatformSource,
		"platformId":     j.cfg.PlatformID,
		"_route":         fmt.Sprintf("h%s", j.cfg.HisID),
	}
	body := map[string]string{
		"deptId":       DeptID,
		"scheduleDate": date,
		"doctorId":     doctorId,
	}

	resp, err := j.Post(ScheduleTimeListUrl, param, body)
	if err != nil {
		j.logger.Error("post failed", zap.Error(err))
		return nil, err
	}

	if err := json.Unmarshal(resp, &res); err != nil {
		j.logger.Error("json unmarshal failed", zap.Error(err), zap.Any("body", resp))
		return nil, err
	}
	return res.Data.ItemList, nil
}

func (j JinNiu) UsableDoctorTimeList(list ItemList) ItemList {
	res := make(ItemList, 0, len(list))
	for i := 0; i < len(list); i++ {
		if list[i].LeftSource > 0 {
			res = append(res, list[i])
		}
	}
	return res
}
