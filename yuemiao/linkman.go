package yuemiao

import (
	"encoding/json"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type LinkManResp struct {
	Code string `json:"code"`
	Data []struct {
		ID           int    `json:"id"`
		UserID       int    `json:"userId"`
		Name         string `json:"name"`
		IDCardNo     string `json:"idCardNo"`
		Birthday     string `json:"birthday"`
		Sex          int    `json:"sex"`
		RegionCode   string `json:"regionCode"`
		Address      string `json:"address"`
		IsDefault    int    `json:"isDefault"`
		RelationType int    `json:"relationType"`
		CreateTime   string `json:"createTime"`
		ModifyTime   string `json:"modifyTime"`
		Yn           int    `json:"yn"`
	} `json:"data"`
	Ok bool `json:"ok"`
}

func (y YueMiao) LinkMan() (int, error) {
	res := LinkManResp{}

	resp, err := y.Requests.Get(LinkManUrl, nil, nil)
	if err != nil {
		y.logger.Error("get province failed", zap.Error(err))
		return 0, err
	}
	if err := json.Unmarshal(resp, &res); err != nil {
		y.logger.Error("json unmarshal failed", zap.Error(err), zap.Any("body", resp))
	}
	if len(res.Data) == 0 {
		return 0, errors.New("请添加联系人")
	}

	return res.Data[2].ID, nil
}
