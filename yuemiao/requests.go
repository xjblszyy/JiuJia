package yuemiao

import (
	"encoding/json"
	"errors"

	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

type Requests struct {
	Verbose bool
	Tk      string
	logger  *zap.Logger
	headers map[string]string
}

func NewRequests(logger *zap.Logger, verbose bool, headers map[string]string) Requests {
	return Requests{
		Verbose: verbose,
		logger:  logger,
		headers: headers,
	}
}

type RespData struct {
	Code string      `json:"code,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Ok   bool        `json:"ok,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func (r Requests) do(url, method string, param map[string]string, body map[string]string) ([]byte, error) {
	var (
		err  error
		resp *resty.Response
		res  []byte
	)

	req := resty.New().SetDebug(r.Verbose).R().SetHeaders(r.headers)
	switch method {
	case resty.MethodGet:
		resp, err = req.SetQueryParams(param).SetFormData(body).Get(url)
	case resty.MethodPost:
		resp, err = req.SetQueryParams(param).SetFormData(body).Post(url)
	default:
		return res, errors.New("method not support")
	}

	if err != nil {
		return res, err
	}
	respData := RespData{}
	err = json.Unmarshal(resp.Body(), &respData)
	if err != nil {
		r.logger.Error("json unmarshal failed", zap.Error(err), zap.Any("resp.Body", resp.Body()))
		return res, err
	}
	if respData.Code != "0000" || respData.Ok == false {
		return res, errors.New(respData.Msg)
	}

	return resp.Body(), nil
}

func (r Requests) Get(url string, param, body map[string]string) ([]byte, error) {
	resp, err := r.do(url, "GET", param, body)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (r Requests) Post(url string, param, body map[string]string) ([]byte, error) {
	resp, err := r.do(url, "POST", param, body)
	if err != nil {
		return nil, err
	}
	return resp, err
}
