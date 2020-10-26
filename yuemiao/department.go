package yuemiao

import (
	"encoding/json"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type DepartmentsResp struct {
	Code string `json:"code"`
	Data struct {
		Offset       int   `json:"offset"`
		End          int   `json:"end"`
		Total        int   `json:"total"`
		Limit        int   `json:"limit"`
		PageNumber   int   `json:"pageNumber"`
		PageListSize int   `json:"pageListSize"`
		PageNumList  []int `json:"pageNumList"`
		Rows         []struct {
			Code         string        `json:"code"`
			Name         string        `json:"name"`
			ImgURL       string        `json:"imgUrl"`
			Address      string        `json:"address"`
			WorktimeDesc string        `json:"worktimeDesc"`
			Total        int           `json:"total"`
			IsSeckill    int           `json:"isSeckill"`
			DepaCodes    []interface{} `json:"depaCodes"`
			Vaccines     []struct {
				Code         string `json:"code"`
				Name         string `json:"name"`
				ID           int    `json:"id"`
				SubDateStart string `json:"subDateStart"`
				IsSeckill    int    `json:"isSeckill"`
			} `json:"vaccines"`
		} `json:"rows"`
		Pages int `json:"pages"`
	} `json:"data"`
	Ok bool `json:"ok"`
}

func (y YueMiao) Departments() (DepartmentsResp, error) {
	res := DepartmentsResp{}
	// ?vaccineCode=8803&cityName=&offset=0&limit=10&name=&regionCode=5101&isSeckill=1
	param := map[string]string{
		"vaccineCode": "8803", // 九价疫苗都是8803
		"offset":      "0",
		"limit":       "10",
		"regionCode":  "5101", // 5101代表成都，这里可以修改的 todo吧
		"isSeckill":   "1",
		"cityName":    "",
	}
	resp, err := y.Requests.Get(DepartmentsUrl, param, nil)
	if err != nil {
		y.logger.Error("get province failed", zap.Error(err))
		return res, err
	}
	if err := json.Unmarshal(resp, &res); err != nil {
		y.logger.Error("json unmarshal failed", zap.Error(err), zap.Any("body", resp))
	}

	return res, nil
}

func (y YueMiao) UsableDepartments(d DepartmentsResp) (int, error) {
	for i := 0; i < len(d.Data.Rows); i++ {
		if d.Data.Rows[i].Name == y.cfg.DepartmentName {
			if len(d.Data.Rows[i].Vaccines) != 0 {
				return d.Data.Rows[i].Vaccines[0].ID, nil
			}
		}
	}
	return 0, errors.New("没有找到要秒杀的门诊，请确认配置中的门诊名称是否正确")
}
