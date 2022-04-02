package http

import (
	"encoding/json"
	"errors"

	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"jiujia/pkg/constants"
)

func Get(url string, param, headers map[string]string, result interface{}) (gjson.Result, error) {
	response, err := client.R().SetQueryParams(param).SetHeaders(headers).Get(url)
	if err != nil {
		zap.L().Error("get request failed", zap.Error(err))
		return gjson.Result{}, err
	}
	respBody := string(response.Body())
	if gjson.Get(respBody, "code").String() != constants.SuccessCode {
		return gjson.Result{}, errors.New(gjson.Get(respBody, "msg").String())
	}
	if result != nil {
		if err := json.Unmarshal(response.Body(), &result); err != nil {
			zap.L().Error("json unmarshal failed", zap.Error(err))
			return gjson.Result{}, err
		}
	}
	return gjson.Get(respBody, "data"), nil
}
