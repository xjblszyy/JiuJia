package yuemiao

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"go.uber.org/zap"

	"yuemiao/config"
	"yuemiao/utils/vcode"
)

func (y YueMiao) Together() error {

	// 0:先解析本地的保存的验证码

	// 1:获取门诊列表
	departments, err := y.Departments()
	if err != nil {
		y.logger.Error("获取门诊列表失败", zap.Error(err))
		return err
	}

	// 2:获取需要订阅的门诊信息
	departmentID, err := y.UsableDepartments(departments)
	if err != nil {
		y.logger.Error("获取可用门诊失败", zap.Error(err))
		return err
	}

	// 3:获取联系人(注意联系人为第一个，其他人无效)
	linkManID, err := y.LinkMan()
	if err != nil {
		y.logger.Error("获取联系人失败", zap.Error(err))
		return err
	}
	y.linkMan = strconv.Itoa(linkManID)

	// 4:获取并解析验证码
	code, err := y.getAndParseCode()
	if err != nil {
		y.logger.Error("获取并解析验证码失败", zap.Error(err))
		return err
	}
	y.vcode = code

	startTime, _ := time.Parse("2006-01-02 15:04:05", y.cfg.StartTime)
	for {
		now, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
		if startTime.Before(now) {
			break
		} else {
			time.Sleep(time.Millisecond * 500)
			y.logger.Info("time sleep 0.5s")
		}
	}
	y.logger.Info("start order")

	// 5:门诊详情
	detail, err := y.DetailVo(departmentID)
	for err != nil {
		detail, err = y.DetailVo(departmentID)
	}

	// 6:订阅
	date := y.suitableDay(&detail, true)
	for {
		errCode, err := y.Subscribe(&detail, date)

		if err != nil {
			y.logger.Error("订阅失败", zap.Error(err))
			code, err = y.getAndParseCode()
			if err != nil {
				y.logger.Error("获取并解析验证码失败", zap.Error(err))
				return err
			}
			y.vcode = code
			continue
		}

		if errCode == "0000" {
			y.logger.Info("订阅成功！")
			break
		}

		if errCode == "9999" {
			y.logger.Info("订阅人数已满，更换日期！")
			date = y.suitableDay(&detail, false)
		}
	}

	return nil
}

// 获取验证码并解析结果
func (y YueMiao) getAndParseCode() (string, error) {

	parsedCode := map[string]string{}
	pwd, _ := os.Getwd()
	file, err := ioutil.ReadFile(pwd + "/yuemiao/" + ParseVcodeFileName)
	if err != nil {
		y.logger.Error("读本地文件失败", zap.Error(err))
		return "", err
	}
	err = json.Unmarshal(file, &parsedCode)
	if err != nil {
		y.logger.Error("json.Unmarshal失败", zap.Error(err))
		return "", err
	}

	// 4:获取验证码
	s := NewYueMiao(zap.L(), config.C.YueMiao)
	codeData, err := s.ValidateCode()
	if err != nil {
		y.logger.Error("获取验证码失败", zap.Error(err))
		return "", err
	}

	var (
		code string
		ok   bool
	)
	code, ok = parsedCode[codeData]
	if !ok {
		y.logger.Info("本地没有保存验证码")
		// 5:解析验证码
		vc := vcode.NewVCode(zap.L(), config.C.VCode)
		code, err = vc.VCodeResult(codeData, vcode.VCodeJS)
		if err != nil {
			y.logger.Error("解析验证码失败", zap.Error(err))
			return "", err
		}
	}

	return code, nil
}

// // 解析本地验证码
// func (y YueMiao) parseLocalCode() (map[string]string, error){
// 	res := map[string]string{}
// 	file, err := ioutil.ReadFile(ParseVcodeFileName)
// 	if err != nil{
// 		y.logger.Error("读本地文件失败", zap.Error(err))
// 		return nil, err
// 	}
// 	err = json.Unmarshal(file, &res)
// 	if err != nil{
// 		y.logger.Error("json.Unmarshal失败", zap.Error(err))
// 		return nil, err
// 	}
// 	return res, nil
// }
