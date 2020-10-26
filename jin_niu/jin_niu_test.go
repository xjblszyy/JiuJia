package jin_niu

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"yuemiao/config"
)

const (
	verbose        = true
	cookie         = "Hm_lpvt_24e9753272ef5f4ff11548de77419e85=1593672164; Hm_lvt_24e9753272ef5f4ff11548de77419e85=1593389057,1593506477,1593582599,1593672164; COOKIE_JSESSIONID_2153_1=1593672162597-514A0A9FF0B0E21E2FAF34"
	hisID          = "2153"
	platformSource = "1"
	platformID     = "2153"
)

func TestJinNiu_DateScheduleList(t *testing.T) {
	cfg := config.JinNiuConfig{
		Verbose:        verbose,
		Cookie:         cookie,
		HisID:          hisID,
		PlatformSource: platformSource,
		PlatformID:     platformID,
	}

	jinNiu := NewJinNiu(zap.L(), cfg)
	res, err := jinNiu.DateScheduleList()
	assert.NoError(t, err)
	t.Log(t, res)
}

func TestJinNiu_DoctorList(t *testing.T) {
	cfg := config.JinNiuConfig{
		Verbose:        verbose,
		Cookie:         cookie,
		HisID:          hisID,
		PlatformSource: platformSource,
		PlatformID:     platformID,
	}

	jinNiu := NewJinNiu(zap.L(), cfg)
	res, err := jinNiu.DoctorList("2020-07-03")
	assert.NoError(t, err)
	t.Log(t, res)
}

func TestJinNiu_DoctorTimeList(t *testing.T) {
	cfg := config.JinNiuConfig{
		Verbose:        verbose,
		Cookie:         cookie,
		HisID:          hisID,
		PlatformSource: platformSource,
		PlatformID:     platformID,
	}

	jinNiu := NewJinNiu(zap.L(), cfg)
	res, err := jinNiu.DoctorTimeList("2020-07-03", "5255")
	assert.NoError(t, err)
	t.Log(t, res)
}

func TestJinNiu_RegisterConfirmList(t *testing.T) {
	cfg := config.JinNiuConfig{
		Verbose:        verbose,
		Cookie:         cookie,
		HisID:          hisID,
		PlatformSource: platformSource,
		PlatformID:     platformID,
	}
	doctorId := "5255"
	scheduleId := "95065"
	scheduleDate := "2020-07-03"
	visitBeginTime := "14:00:00"
	visitEndTime := "14:10:00"
	jinNiu := NewJinNiu(zap.L(), cfg)
	res, err := jinNiu.RegisterConfirmList(doctorId, scheduleId, scheduleDate, visitBeginTime, visitEndTime)
	assert.NoError(t, err)
	t.Log(t, res)
}

func TestJinNiu_GeneratorOrder(t *testing.T) {
	cfg := config.JinNiuConfig{
		Verbose:        verbose,
		Cookie:         cookie,
		HisID:          hisID,
		PlatformSource: platformSource,
		PlatformID:     platformID,
	}
	doctorId := "5255"
	scheduleId := "95065"
	scheduleDate := "2020-07-02"
	visitBeginTime := "14:00:00"
	visitEndTime := "14:10:00"
	patientId := "2939582631628701708"

	jinNiu := NewJinNiu(zap.L(), cfg)
	res, err := jinNiu.GeneratorOrder(doctorId, scheduleId, scheduleDate, visitBeginTime, visitEndTime, patientId)
	assert.NoError(t, err)
	t.Log(t, res)
}
