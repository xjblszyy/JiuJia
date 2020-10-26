package jin_niu

import (
	"encoding/json"

	"go.uber.org/zap"
)

type RegisterConfirmList struct {
	Code int `json:"code"`
	Data struct {
		DeptName    string `json:"deptName"`
		DoctorLevel string `json:"doctorLevel"`
		DoctorName  string `json:"doctorName"`
		DoctorTitle string `json:"doctorTitle"`
		HisName     string `json:"hisName"`
		LeftBindNum int    `json:"leftBindNum"`
		PatientList []struct {
			BindStatus      int    `json:"bindStatus"`
			HealthCardFlag  int    `json:"healthCardFlag"`
			IDNo            string `json:"idNo"`
			IDType          int    `json:"idType"`
			IDTypeName      string `json:"idTypeName"`
			IsDefault       int    `json:"isDefault"`
			PatCardNo       string `json:"patCardNo"`
			PatCardNoEncry  string `json:"patCardNoEncry"`
			PatCardType     int    `json:"patCardType"`
			PatCardTypeName string `json:"patCardTypeName"`
			PatHisNo        string `json:"patHisNo"`
			PatInNo         string `json:"patInNo"`
			PatientFullIDNo string `json:"patientFullIdNo"`
			PatientID       string `json:"patientId"`
			PatientMobile   string `json:"patientMobile"`
			PatientName     string `json:"patientName"`
			PatientSex      string `json:"patientSex"`
			PatientType     int    `json:"patientType"`
			RelationName    string `json:"relationName"`
			RelationType    int    `json:"relationType"`
		} `json:"patientList"`
		RegisterType     string `json:"registerType"`
		RegisterTypeName string `json:"registerTypeName"`
		ScheduleDate     string `json:"scheduleDate"`
		TotalFee         int    `json:"totalFee"`
		VisitBeginTime   string `json:"visitBeginTime"`
		VisitEndTime     string `json:"visitEndTime"`
		VisitWeekName    string `json:"visitWeekName"`
	} `json:"data"`
}

func (j JinNiu) RegisterConfirmList(doctorId, scheduleId, scheduleDate, visitBeginTime, visitEndTime string) (RegisterConfirmList, error) {
	res := RegisterConfirmList{}

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
	}

	resp, err := j.Post(RegisterConfirmUrl, param, body)
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
