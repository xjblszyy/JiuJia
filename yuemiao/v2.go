package yuemiao

import (
	"encoding/json"
	"errors"
	"time"

	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

type V2Subscribe struct {
	Code  string `json:"code"`
	Msg   string `json:"msg"`
	Ok    bool   `json:"ok"`
	NotOk bool   `json:"notOk"`
}

func (y YueMiao) V2() {
	startTime, _ := time.Parse("2006-01-02 15:04:05", y.cfg.StartTime)
	for {
		now, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
		if startTime.Before(now) {
			break
		} else {
			time.Sleep(time.Millisecond * 100)
			y.logger.Info("time sleep 0.1s")
		}
	}
	y.logger.Info("start order")

	// 限流器速率自己调
	limiter := rate.NewLimiter(rate.Every(100*time.Millisecond), 10)
	for {
		if limiter.Allow() {
			go y.subscribe()
		}
	}

}

func (y YueMiao) subscribe() error {
	res := V2Subscribe{}
	param := map[string]string{
		"seckillId":    y.cfg.SeckillId,
		"linkmanId":    y.cfg.LinkmanId,
		"idCardNo":     y.cfg.IdCardNo,
		"vaccineIndex": "1",
	}

	resp, err := y.Requests.Get(V2SubscribeUrl, param, nil)
	if err != nil {
		y.logger.Error("版本2订阅失败", zap.Error(err))
		return err
	}

	y.logger.Info("resp", zap.Any("body", resp))

	if err := json.Unmarshal(resp, &res); err != nil {
		y.logger.Error("json unmarshal failed", zap.Error(err), zap.Any("body", resp))
		return err
	}
	if res.Ok == false {
		return errors.New(res.Msg)
	}
	return nil
}
