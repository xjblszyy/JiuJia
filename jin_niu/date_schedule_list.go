package jin_niu

import (
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
)

type DateSchedule struct {
	Code int `json:"code"`
	Data struct {
		ScheduleList ScheduleList `json:"scheduleList"`
	} `json:"data"`
}

type ScheduleList []struct {
	DeptID       string `json:"deptId"`
	HisID        int    `json:"hisId"`
	MonthDay     string `json:"monthDay"`
	ScheduleDate string `json:"scheduleDate"`
	Selected     bool   `json:"selected"`
	Status       string `json:"status"`
	WeekDate     string `json:"weekDate"`
}

func (j JinNiu) DateScheduleList() (ScheduleList, error) {
	res := DateSchedule{}
	param := map[string]string{
		"hisId":          j.cfg.HisID,
		"platformSource": j.cfg.PlatformSource,
		"platformId":     j.cfg.PlatformID,
		"_route":         fmt.Sprintf("h%s", j.cfg.HisID),
	}
	body := map[string]string{
		"deptId": DeptID,
	}
	resp, err := j.Post(ScheduleListUrl, param, body)
	if err != nil {
		j.logger.Error("post failed", zap.Error(err))
		return nil, err
	}

	if err := json.Unmarshal(resp, &res); err != nil {
		j.logger.Error("json unmarshal failed", zap.Error(err), zap.Any("body", resp))
		return nil, err
	}
	return res.Data.ScheduleList, nil
}

func (j JinNiu) UsableDateSchedule(list ScheduleList) ScheduleList {
	res := make(ScheduleList, 0, len(list))
	for i := 0; i < len(list); i++ {
		if list[i].WeekDate == "日" && list[i].Status == "1" {
			res = append(res, list[i])
		}
		if list[i].Status == "1" {
			res = append(res, list[i])
		}
	}
	return res
}
