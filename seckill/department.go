package seckill

import (
	"strings"

	"go.uber.org/zap"
)

// 获取某一个门诊数据ID
func (s AllSteps) FetchDepartmentID() ([]int, error) {
	needDepartment := strings.Split(s.cfg.DepartmentName, ",")
	res := make([]int, 0, len(needDepartment))

	departments, err := s.GetAllDepartments()
	if err != nil {
		s.logger.Error("获取门诊列表失败", zap.Error(err))
		return res, nil
	}
	dps := departments.Data.Rows
	for i := 0; i < len(dps); i++ {
		for j := 0; j < len(needDepartment); j++ {
			if dps[i].Name == needDepartment[j] {
				res = append(res, dps[i].DepaVaccID)
			}
		}
	}
	return res, nil
}
