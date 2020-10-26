package yuemiao

import (
	"encoding/json"

	"go.uber.org/zap"
)

type ValidateCodeResp struct {
	Code string `json:"code"`
	Data string `json:"data"`
	Ok   bool   `json:"ok"`
}

func (y YueMiao) ValidateCode() (string, error) {
	res := ValidateCodeResp{}

	resp, err := y.Requests.Get(ValidateCodeUrl, nil, nil)
	if err != nil {
		y.logger.Error("get validate code result failed", zap.Error(err))
		return "", err
	}
	if err := json.Unmarshal(resp, &res); err != nil {
		y.logger.Error("json unmarshal failed", zap.Error(err), zap.Any("body", resp))
	}

	return res.Data, nil
}
