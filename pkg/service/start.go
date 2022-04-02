package service

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"go.uber.org/zap"
	"jiujia/pkg/constants"
)

func (s Service) Start() error {
	// 先获取抢购的疫苗的开始时间
	vaccines, err := s.GetVaccineList(s.cfg.RegionCode)
	if err != nil {
		s.logger.Error("get vaccine list failed", zap.Error(err))
		return errors.New("获取抢购疫苗失败！")
	}
	startTime := ""
	vaccineID := ""
	for i := 0; i < len(vaccines.Data); i++ {
		if strconv.Itoa(vaccines.Data[i].ID) == s.cfg.VaccineID {
			startTime = vaccines.Data[i].StartTime
			vaccineID = strconv.Itoa(vaccines.Data[i].ID)
			break
		}
	}
	if startTime == "" {
		return errors.New("抢购的疫苗id(配置文件中的vaccine_id项)配置错误！")
	}
	startDateTime, err := time.ParseInLocation(constants.Layout, startTime, time.Local)
	if err != nil {
		s.logger.Error("parse start time failed", zap.Error(err), zap.String("startTime", startTime))
		return errors.New("解析时间错误！")
	}
	now := time.Now()
	if now.Add(5 * time.Second).Before(startDateTime) {
		s.logger.Info("还未到获取st时间，等待中......")
		time.Sleep(startDateTime.Sub(now.Add(5 * time.Second)))
	}

	s.logger.Info("到达获取st时间！")

	// 循环获取st，直到成功为止
	var (
		st string
	)
	for {
		st, err = s.GetST(s.cfg.VaccineID)
		if err != nil {
			s.logger.Error("get st failed", zap.Error(err))
		}
		break
	}

	now = time.Now()
	if now.Add(500 * time.Millisecond).Before(startDateTime) {
		s.logger.Info("获取st成功，但是还未到抢购时间，等待中......")
		time.Sleep(startDateTime.Sub(now.Add(500 * time.Millisecond)))
	}

	s.logger.Info("到达抢购时间，正在抢购！")
	s.start(vaccineID, st)
	return nil
}

func (s Service) start(seckillID, st string) {
	wg := sync.WaitGroup{}
	wg.Add(s.cfg.Total)
	success := false
	for i := 0; i < s.cfg.Total; i++ {
		s.logger.Info(fmt.Sprintf("当前第%d个协程正在抢购！", i+1))
		go func(seckillID, st string) {
			err := s.SecKill(seckillID, s.cfg.MemberID, s.cfg.IDCard, st)
			if err != nil {
				s.logger.Info(fmt.Sprintf("当前第%d个协程正在抢购失败！", i+1), zap.Error(err))
			} else {
				success = true
			}

		}(seckillID, st)
		s.logger.Info(fmt.Sprintf("正在休息%d毫秒，等待下一个协程抢购", s.cfg.Step))
		time.Sleep(time.Duration(s.cfg.Step) * time.Millisecond)
		wg.Done()
	}
	wg.Wait()

	if success {
		s.logger.Info("抢购成功，请在小程序中查看！")
	} else {
		s.logger.Info("所有协程都抢购失败，再接再厉！")
	}
}
