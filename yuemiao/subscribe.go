package yuemiao

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"strconv"
	"time"

	"go.uber.org/zap"
)

type SubscribeResp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Ok   bool   `json:"ok"`
}

func (y YueMiao) Subscribe(detail *DetailVoResp, subDate string) (string, error) {
	res := SubscribeResp{}
	param := map[string]string{
		"departmentVaccineId": strconv.Itoa(detail.Data.ID),
		"vaccineIndex":        "1",
		"linkmanId":           y.linkMan,
		"subscribeDate":       subDate,
		"sign":                y.sign(detail.Data.Time),
		"vcode":               y.vcode,
	}
	resp, err := y.Requests.Get(SubscribeUrl, param, nil)
	if err != nil {
		y.logger.Error("get province failed", zap.Error(err))
		return "", err
	}
	if err := json.Unmarshal(resp, &res); err != nil {
		y.logger.Error("json unmarshal failed", zap.Error(err), zap.Any("body", resp))
	}

	return res.Code, nil
}

// 生成签名
func (y YueMiao) sign(time int64) string {
	str := strconv.FormatInt(time, 10) + "fuckhacker10000times"
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 选择适合抢的日期，原则：剩余数量最大的
func (y YueMiao) suitableDay(detail *DetailVoResp, firstTime bool) string {
	days := detail.Data.Days

	if len(days) < 1 {
		panic("已全部预约满")
	}

	var day string
	idx := 0
	if firstTime && len(days) != 1 {
		for i := 1; i < len(days); i++ {
			if days[idx].Total <= days[i].Total {
				idx = i
			}
		}
		day = days[idx].Day
	} else {
		rand.Seed(time.Now().UnixNano())
		idx = rand.Intn(len(days) - 1)
		day = days[idx].Day
	}
	detail.Data.Days = append(days[:idx], days[idx+1:]...)

	formatDay, err := time.Parse("20060102", day)
	if err != nil {
		y.logger.Error("解析时间有误")
		return ""
	}
	y.logger.Info("日期：" + formatDay.String())
	return formatDay.Format("2006-01-02")
}
