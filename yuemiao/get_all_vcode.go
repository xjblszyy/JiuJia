package yuemiao

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"go.uber.org/zap"
)

func (y YueMiao) GetAllVCode() error {

	var writeArr []string
	vcodes := make(map[string]int)
	redundentTimes := 0
	for redundentTimes < 5 {
		data, err := y.ValidateCode()
		if err != nil {
			y.logger.Error("获取验证码失败", zap.Error(err))
			return err
		}

		if vcodes[data] != 1 {
			vcodes[data] = 1
			writeArr = append(writeArr, data)
		} else {
			redundentTimes++
		}
	}
	writeBytes, err := json.Marshal(writeArr)
	if err != nil {
		y.logger.Error("json marshal failed", zap.Error(err))
		return err
	}
	if err := writeToFile(writeBytes); err != nil {
		y.logger.Error("写入文件失败", zap.Error(err))
		return err
	}
	return nil
}

func writeToFile(writeString []byte) error {
	path, _ := os.Getwd()
	var filename = path + "/yuemiao/" + VcodeFileName

	err := ioutil.WriteFile(filename, writeString, 0666)
	return err
}
