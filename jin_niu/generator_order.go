package jin_niu

import (
	"encoding/json"

	"go.uber.org/zap"
)

type GeneratorOrder struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (j JinNiu) GeneratorOrder(doctorId, scheduleId, scheduleDate, visitBeginTime, visitEndTime, patientId string) (GeneratorOrder, error) {
	res := GeneratorOrder{}

	param := map[string]string{
		"hisId":          j.cfg.HisID,
		"platformSource": j.cfg.PlatformSource,
		"platformId":     j.cfg.PlatformID,
	}
	body := map[string]string{
		"deptId":         DeptID,
		"scheduleId":     scheduleId,
		"doctorId":       doctorId,
		"scheduleDate":   scheduleDate,
		"visitPeriod":    "",
		"visitBeginTime": visitBeginTime,
		"visitEndTime":   visitEndTime,
		"patientId":      patientId,
	}

	resp, err := j.Post(GeneratorOrderUrl, param, body)
	if err != nil {
		j.logger.Error("post failed", zap.Error(err))
		return res, err
	}

	if err := json.Unmarshal(resp, &res); err != nil {
		j.logger.Error("json unmarshal failed", zap.Error(err), zap.Any("body", resp))
		return res, err
	}
	return res, nil
}
