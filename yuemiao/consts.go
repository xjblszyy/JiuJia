package yuemiao

const (
	UserAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 6_1_3 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Mobile/10B329 micromessenger/5.0.1"

	// 门诊列表
	DepartmentsUrl = "https://wx.healthych.com/seckill/department/pageList.do"
	// 门诊详情
	DetailVoUrl = "https://wx.healthych.com/seckill/vaccine/detailVo.do"
	// 订阅
	SubscribeUrl = "https://wx.healthych.com/seckill/vaccine/subscribe.do"
	// 联系人
	LinkManUrl = "https://wx.healthych.com/seckill/linkman/findByUserId.do"
	// 验证码
	ValidateCodeUrl = "https://wx.healthych.com/seckill/validateCode/vcode.do"

	// 验证码保存的文件名
	VcodeFileName      = "vcode.dat"
	ParseVcodeFileName = "parsed_vcode.dat"

	// 改版后的下单地址
	V2SubscribeUrl = "https://miaomiao.scmttec.com/seckill/seckill/subscribe.do"
)
