package utils

import (
	"strings"

	"jiujia/config"
)

func CommonHeader() map[string]string {
	tk := config.C.TK
	cookie := config.C.Cookie
	cookieArray := strings.Split(strings.ReplaceAll(cookie, " ", ""), ";")
	return map[string]string{
		"User-Agent": "Mozilla/5.0 (Linux; Android 5.1.1; SM-N960F Build/JLS36C; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/74.0.3729.136 Mobile Safari/537.36 MMWEBID/1042 MicroMessenger/7.0.15.1680(0x27000F34) Process/appbrand0 WeChat/arm32 NetType/WIFI Language/zh_CN ABI/arm32",
		"Referer":    "https://servicewechat.com/wxff8cad2e9bf18719/2/page-frame.html",
		"tk":         tk,
		"Accept":     "application/json, text/plain, */*",
		"Host":       "miaomiao.scmttec.com",
		"Cookie":     strings.Join(cookieArray, "; "),
	}
}
