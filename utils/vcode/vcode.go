package vcode

import (
	"encoding/json"
	"errors"

	"github.com/go-resty/resty/v2"
	"github.com/prometheus/common/log"
	"go.uber.org/zap"
)

type CodeType string

const (
	// 不定长汉字、英文、数字、符号、空格组合
	VCodeCN CodeType = "cn"
	// 加减乘除计算结果
	VCodeJS CodeType = "js"
	// 动态图验证码
	VCodeGIF CodeType = "gif"
)

type Identify struct {
	Verbose bool
	logger  *zap.Logger
	headers map[string]string
}

func NewIdentify(logger *zap.Logger, verbose bool, headers map[string]string) Identify {
	return Identify{
		Verbose: verbose,
		logger:  logger,
		headers: headers,
	}
}

type ResBody struct {
	ErrCode int    `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
	Msg     string `json:"msg"`
	VType   string `json:"v_type"`
	VCode   string `json:"v_code"`
}

func (r Identify) VCodeResult(base64Img string, vCodeType CodeType) (string, error) {
	var (
		err  error
		resp *resty.Response
		res  string
	)

	body := make(map[string]string)
	body["v_pic"] = base64Img
	body["v_type"] = string(vCodeType)

	req := resty.New().SetDebug(r.Verbose).R().SetHeaders(r.headers)
	resp, err = req.SetFormData(body).Post(identifyUrl)

	if err != nil {
		return res, err
	}

	respData := ResBody{}
	err = json.Unmarshal(resp.Body(), &respData)
	if err != nil {
		r.logger.Error("json unmarshal failed", zap.Error(err), zap.Any("resp.Body", resp.Body()))
		return res, err
	}

	if respData.ErrCode != 0 || respData.VCode == "" {
		return "", errors.New(respData.Msg + respData.ErrMsg)
	}

	log.Info("验证码解析结果: ", respData.VCode)
	return respData.VCode, nil
}
