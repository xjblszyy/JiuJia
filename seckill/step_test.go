package seckill

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"yuemiao/config"
)

const (
	verbose        = true
	tk             = "c2aa36791d9249d6924af7575cb9760d_e1cde20dac4de1c1fca4671c1f3833f8"
	province       = "四川省"
	city           = "成都市"
	vaccines       = "1"
	departmentName = "成都市锦江区莲新社区卫生服务中心"
)

func TestAllSteps_GetProvinceCode(t *testing.T) {
	cfg := config.YueMiaoConfig{
		Verbose:        verbose,
		TK:             tk,
		Province:       province,
		City:           city,
		Vaccines:       vaccines,
		DepartmentName: departmentName,
	}

	s := NewAllSteps(zap.L(), cfg)
	res, err := s.GetProvinceCode()
	assert.NoError(t, err)
	assert.Equal(t, "51", res)
}

func TestAllSteps_GetCityCode(t *testing.T) {
	cfg := config.YueMiaoConfig{
		Verbose:        verbose,
		TK:             tk,
		Province:       province,
		City:           city,
		Vaccines:       vaccines,
		DepartmentName: departmentName,
	}
	s := NewAllSteps(zap.L(), cfg)
	res, err := s.GetCityCode()
	assert.NoError(t, err)
	assert.Equal(t, "5101", res)
}

func TestAllSteps_GetAllDepartments(t *testing.T) {
	cfg := config.YueMiaoConfig{
		Verbose:        verbose,
		TK:             tk,
		Province:       province,
		City:           city,
		Vaccines:       vaccines,
		DepartmentName: departmentName,
	}
	s := NewAllSteps(zap.L(), cfg)
	res, err := s.GetAllDepartments()
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Data.Rows)
}

func TestAllSteps_FetchDepartment(t *testing.T) {
	cfg := config.YueMiaoConfig{
		Verbose:        verbose,
		TK:             tk,
		Province:       province,
		City:           city,
		Vaccines:       vaccines,
		DepartmentName: departmentName,
	}
	s := NewAllSteps(zap.L(), cfg)
	res, err := s.FetchDepartmentID()
	assert.NoError(t, err)
	assert.Equal(t, res, []int{6962})
}

func TestAllSteps_FetchDepartmentInfo(t *testing.T) {
	cfg := config.YueMiaoConfig{
		Verbose:        verbose,
		TK:             tk,
		Province:       province,
		City:           city,
		Vaccines:       vaccines,
		DepartmentName: departmentName,
	}
	s := NewAllSteps(zap.L(), cfg)
	res, err := s.FetchDepartmentInfo(6962)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Data)
}

func TestAllSteps_GetWorkDay(t *testing.T) {
	cfg := config.YueMiaoConfig{
		Verbose:        verbose,
		TK:             tk,
		Province:       province,
		City:           city,
		Vaccines:       vaccines,
		DepartmentName: departmentName,
	}
	s := NewAllSteps(zap.L(), cfg)

	info, err := s.FetchDepartmentInfo(6962)
	assert.NoError(t, err)

	res, err := s.GetWorkDay(info)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Data)
}

func TestAllSteps_GetWorkTime(t *testing.T) {
	cfg := config.YueMiaoConfig{
		Verbose:        verbose,
		TK:             tk,
		Province:       province,
		City:           city,
		Vaccines:       vaccines,
		DepartmentName: departmentName,
	}
	s := NewAllSteps(zap.L(), cfg)

	info, err := s.FetchDepartmentInfo(6962)
	assert.NoError(t, err)

	days, err := s.GetWorkDay(info)
	assert.NoError(t, err)
	assert.NotEmpty(t, days.Data.DateList)
	day := days.Data.DateList[0]

	res, err := s.GetWorkTime(info, day)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Data.Times.Data)
}

func TestAllSteps_GetAllCitiesCode(t *testing.T) {
	cfg := config.YueMiaoConfig{
		Verbose:        verbose,
		TK:             tk,
		Province:       province,
		City:           city,
		Vaccines:       vaccines,
		DepartmentName: departmentName,
	}
	s := NewAllSteps(zap.L(), cfg)
	res, err := s.GetAllCitiesCode()
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}
