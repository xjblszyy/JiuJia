package seckill

const (
	// 登陆地址
	LoginUrl = "https://wx.healthych.com/passport/wx/login.do"
	// 城市地址
	AllCitiesUrl = "https://wx.healthych.com/base/region/childRegions.do"
	// 门诊地址
	DepartmentUrl = "https://wx.healthych.com/base/department/getDepartments.do"
	// 门诊详情
	DepartmentInfoUrl = "https://wx.healthych.com/base/departmentVaccine/item.do"
	// 当前门诊可用日期查询地址
	WorkDayUrl = "https://wx.healthych.com/order/subscribe/workDays.do"
	// 当前门诊可用日期下的时间查询地址
	WorkTimeUrl = "https://wx.healthych.com/order/subscribe/departmentWorkTimes2.do"
	// 是否有秒杀信息
	HasSeckillUrl = "https://miaomiao.scmttec.com/seckill/seckill/list.do"
	// 获取所有城市列表信息
	CitiesCodeUrl = "https://miaomiao.scmttec.com/base/region/childRegions.do"

	UserAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 6_1_3 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Mobile/10B329 micromessenger/5.0.1"
	// 二价
	CustomIdErJia = "1"
	// 四价
	CustomIdSiJia = "2"
	// 九价
	CustomIdJiuJia = "3"
)
