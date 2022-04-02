# 抢购九价学习工具，`知识无罪`，`请勿贩卖`，`仅供学习`
## 使用平台仅限微信小程序`秒苗`

## 使用方法1
### step1
```
需要去秒苗小程序抓包(如果不会可以百度看一下怎么抓包的)
获取一下数据
tk(几乎该小程序的任何接口的请求头中都有此参数)
member_id(/seckill/linkman/findByUserId.do接口中的id的值)
id_card(/seckill/linkman/findByUserId.do接口中的idCardNo的值)
member_name(/seckill/linkman/findByUserId.do接口中的name的值)
region_code(/base/region/childRegions.do?parentCode=51接口中的value的值,注意是精确到市，不是到省，这个值只有4位，默认成都 5101)
cookie(几乎该小程序的任何接口的请求头中都有此参数)
vaccine_id(/seckill/seckill/list.do接口中的id的值)
```
### step2
```
将step1中的那些参数都填写到config.yaml中(里面的默认值尽量别去改它，除非你知道它是干什么的！)
```

### step3
```
运行 make build 将项目编译成可执行文件，如果你已经编译过了，或者已经下载好了已经编译后的文件，则这个过程可以省略(怕你不知道怎么玩的，啰嗦两句！)
运行命令 jiujia start-without-ui --config=/path/to/your/config.yaml (注意/path/to/your/config.yaml替换成你的配置文件的地址)
```

## 使用方法2
```
直接运行带ui版本，可以从release下载(别着急，马上就做出来)
```

## 说明
```
Makefile和Dockerfile我还没测试，估计会有问题，我后面再改！
```


