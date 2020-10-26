package yuemiao

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"go.uber.org/zap"

	"yuemiao/config"
	"yuemiao/utils/vcode"
)

func (y YueMiao) ParseAllVCode() error {
	path, _ := os.Getwd()
	file, err := ioutil.ReadFile(path + "/yuemiao/" + VcodeFileName)
	if err != nil {
		y.logger.Error("读取文件失败", zap.Error(err))
		return err
	}
	var imgs []string
	if err := json.Unmarshal(file, &imgs); err != nil {
		y.logger.Error("json unmarshal failed", zap.Error(err))
		return err
	}

	res := make(map[string]string)
	v := vcode.NewVCode(y.logger, config.C.VCode)
	for i := 0; i < len(imgs); i++ {
		code, err := v.VCodeResult(imgs[i], vcode.VCodeJS)
		if err != nil {
			y.logger.Error("解析验证码出错", zap.Error(err), zap.Any("img", imgs[i]))

		}
		info := fmt.Sprintf("第%d张图片的解析结果是:%s", i+1, code)
		y.logger.Info(info)
		res[imgs[i]] = code
	}

	writeData, err := json.Marshal(res)
	if err != nil {
		y.logger.Error("json marshal failed", zap.Error(err), zap.Any("json", res))
		return err
	}

	var filename = path + "/yuemiao/" + ParseVcodeFileName
	if err := ioutil.WriteFile(filename, writeData, 0666); err != nil {
		y.logger.Error("写入文件失败", zap.Error(err))
		return err
	}
	return nil
}
