package jin_niu

import (
	"os"
	"time"

	"go.uber.org/zap"
)

var success chan bool

func (j JinNiu) Together() {
	// 1:获取可用日期
	dateScheduleList, err := j.DateScheduleList()
	if err != nil {
		j.logger.Error("获取可用日期有误", zap.Error(err))
		return
	}
	list := j.UsableDateSchedule(dateScheduleList)

	go j.loop1(list)
	if ok := <-success; ok {
		os.Exit(1)
	}
}

func (j JinNiu) loop1(list ScheduleList) {

	// 获取可用日期下的医生时间信息
	for i := 0; i < len(list); i++ {
		doctorList, err := j.DoctorList(list[i].ScheduleDate)
		var scheduleDate string
		scheduleDate = list[i].ScheduleDate
		if err != nil {
			j.logger.Error("获取可用日期下的医生时间信息", zap.Error(err))
			return
		}

		dList := j.JiuJiaDoctorList(doctorList)
		go j.loop2(scheduleDate, dList)
	}
}

func (j JinNiu) loop2(scheduleDate string, dList DoctorList) {
	for k := 0; k < len(dList); k++ {
		var tList ItemList
		var doctorID string

		// 获取可用日期下的医生信息
		timeList, err := j.DoctorTimeList(scheduleDate, dList[k].DoctorID)
		doctorID = dList[k].DoctorID
		if err != nil {
			j.logger.Error("获取可用日期下的医生时间信息", zap.Error(err))
		}
		tList = j.UsableDoctorTimeList(timeList)
		go j.loop3(tList, doctorID, scheduleDate)

	}
}

func (j JinNiu) loop3(tList ItemList, doctorID string, scheduleDate string) {
	startTime, _ := time.Parse("2006-01-02 15:04:05", j.cfg.StartTime)
	for {
		now, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
		if startTime.Before(now) {
			break
		} else {
			time.Sleep(time.Millisecond * 500)
			j.logger.Info("time sleep 0.5s")
		}
	}
	j.logger.Info("start order")
	for {
		for i := 0; i < len(tList); i++ {
			registerConfirmList, err := j.RegisterConfirmList(doctorID, tList[i].ScheduleID, scheduleDate, tList[i].VisitBeginTime, tList[i].VisitEndTime)
			if err != nil {
				j.logger.Error("RegisterConfirmList", zap.Error(err))
				return
			}

			for k := 0; k < len(registerConfirmList.Data.PatientList); k++ {
				res, err := j.GeneratorOrder(doctorID, tList[i].ScheduleID, scheduleDate, tList[i].VisitBeginTime, tList[i].VisitEndTime, registerConfirmList.Data.PatientList[k].PatientID)
				if err != nil {
					j.logger.Error("GeneratorOrder", zap.Error(err))
					return
				}
				if res.Code == 0 {
					j.logger.Info("抢购成功")
					success <- true
					break
				}
			}
		}
		j.logger.Info("start order but time sleep 0.5s")
		time.Sleep(time.Millisecond * 500)
	}
}
