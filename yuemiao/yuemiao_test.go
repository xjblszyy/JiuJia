package yuemiao

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"yuemiao/config"
)

const (
	verbose = true
	tk      = ""
	// province       = "四川省"
	// city           = "成都市"
	// vaccines       = "1"
	departmentName = "成都市成华区妇幼保健院"
)

func TestYueMiao_Departments(t *testing.T) {
	cfg := config.YueMiaoConfig{
		Verbose:        verbose,
		TK:             tk,
		DepartmentName: departmentName,
	}

	s := NewYueMiao(zap.L(), cfg)
	res, err := s.Departments()
	assert.NoError(t, err)
	t.Log(res)
}

func TestYueMiao_UsableDepartments(t *testing.T) {
	cfg := config.YueMiaoConfig{
		Verbose:        verbose,
		TK:             tk,
		DepartmentName: departmentName,
	}

	s := NewYueMiao(zap.L(), cfg)
	res, err := s.Departments()
	assert.NoError(t, err)

	id, err := s.UsableDepartments(res)
	assert.NoError(t, err)
	t.Log(id)
}

func TestYueMiao_DetailVo(t *testing.T) {
	cfg := config.YueMiaoConfig{
		Verbose:        verbose,
		TK:             tk,
		DepartmentName: departmentName,
	}

	s := NewYueMiao(zap.L(), cfg)
	id := 6178
	res, err := s.DetailVo(id)
	assert.NoError(t, err)
	t.Log(res)
}

func TestYueMiao_LinkMan(t *testing.T) {
	cfg := config.YueMiaoConfig{
		Verbose:        verbose,
		TK:             tk,
		DepartmentName: departmentName,
	}

	s := NewYueMiao(zap.L(), cfg)
	res, err := s.LinkMan()
	assert.NoError(t, err)
	t.Log(res)
}

func TestYueMiao_Subscribe(t *testing.T) {
	cfg := config.YueMiaoConfig{
		Verbose:        verbose,
		TK:             tk,
		DepartmentName: departmentName,
	}

	d := `{"code":"0000","data":{"id":3452,"name":"九价HPV疫苗（进口）","vaccineCode":"8803","total":1,"prompt":"","startMilliscond":1595210400000,"hospitalName":"福州新莲花医院","ageStart":16,"ageEnd":26,"now":1595210405673,"workTimeStart":"08:30","workTimeEnd":"11:00","packingImgUrl":"[\"https://adultvacc-1253522668.file.myqcloud.com/thematic%20pic/%E4%B9%9D%E4%BB%B71_1585135836062.jpg\",\"https://adultvacc-1253522668.file.myqcloud.com/thematic%20pic/%E4%B9%9D%E4%BB%B72_1585135836131.jpg\",\"https://adultvacc-1253522668.file.myqcloud.com/thematic%20pic/%E4%B9%9D%E4%BB%B73_1585135836184.jpg\",\"https://adultvacc-1253522668.file.myqcloud.com/thematic%20pic/%E4%B9%9D%E4%BB%B74_1585135836239.jpg\"]","specifications":"0.5mL/支","factoryName":"默沙东集团","isSubscribeAll":0,"isSeckill":true,"days":[{"day":"20200729","total":2},{"day":"20200724","total":3},{"day":"20200731","total":2},{"day":"20200725","total":3},{"day":"20200722","total":5}],"time":1595210405694},"ok":true}`
	detail := DetailVoResp{}
	_ = json.Unmarshal([]byte(d), &detail)

	s := NewYueMiao(zap.L(), cfg)

	id, err := s.LinkMan()
	assert.NoError(t, err)
	s.linkMan = strconv.Itoa(id)

	_, err = s.Subscribe(&detail, "2020-07-22")
	assert.NoError(t, err)
}

func TestYueMiao_ValidateCode(t *testing.T) {
	cfg := config.YueMiaoConfig{
		Verbose:        verbose,
		TK:             tk,
		DepartmentName: departmentName,
	}

	s := NewYueMiao(zap.L(), cfg)

	data, err := s.ValidateCode()
	assert.NoError(t, err)
	assert.NotEmpty(t, data)
}

func TestYueMiao_AllVCode(t *testing.T) {
	cfg := config.YueMiaoConfig{
		Verbose:        verbose,
		TK:             tk,
		DepartmentName: departmentName,
	}

	var writeArr []string
	vcodes := make(map[string]int)
	s := NewYueMiao(zap.L(), cfg)
	redundentTimes := 0
	for redundentTimes < 5 {
		data, err := s.ValidateCode()
		if err != nil {
			fmt.Println(err)
			return
		}

		if vcodes[data] != 1 {
			vcodes[data] = 1
			writeArr = append(writeArr, data)
		} else {
			redundentTimes++
		}
	}
	writeBytes, err := json.Marshal(writeArr)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	writeToFile(writeBytes)
}

// func writeToFile(writeString []byte) {
// 	var filename = "vcode.dat"
//
// 	err := ioutil.WriteFile(filename, writeString, 0666)
// 	if err != nil {
// 		fmt.Println(err)
// 		panic(err)
// 	}
// }

func TestYueMiao_SaveAllVCode(t *testing.T) {
	file, err := ioutil.ReadFile(VcodeFileName)
	assert.NoError(t, err)
	var imgs []string
	err = json.Unmarshal(file, &imgs)
	assert.NoError(t, err)

	for i := 0; i < len(imgs); i++ {

		img, err := base64.StdEncoding.DecodeString(imgs[i])
		assert.NoError(t, err)

		h := md5.New()
		h.Write([]byte(img))
		err = ioutil.WriteFile(fmt.Sprintf("./images/%s.jpg", hex.EncodeToString(h.Sum(nil))), img, 0666)
		assert.NoError(t, err)
	}
}

func TestYueMiao_SaveAllVCode2(t *testing.T) {
	file, err := ioutil.ReadFile(VcodeFileName)
	assert.NoError(t, err)
	var base64Images []string
	err = json.Unmarshal(file, &base64Images)
	assert.NoError(t, err)

	md5Map := make(map[string]string)
	base64Map := make(map[string]string)
	path, _ := ioutil.ReadDir("./images")
	for _, f := range path {
		n := strings.Split(f.Name(), ".")
		if len(n) != 3 {
			t.Log("名字有误: " + f.Name())
		}
		md5Map[n[0]] = n[1]
	}

	for i := 0; i < len(base64Images); i++ {

		img, err := base64.StdEncoding.DecodeString(base64Images[i])
		assert.NoError(t, err)

		h := md5.New()
		h.Write([]byte(img))
		key := hex.EncodeToString(h.Sum(nil))
		v, ok := md5Map[key]
		assert.True(t, ok)
		base64Map[base64Images[i]] = v
	}

	filename := ParseVcodeFileName
	writeString, _ := json.Marshal(base64Map)
	err = ioutil.WriteFile(filename, writeString, 0666)
	assert.NoError(t, err)
}

func TestTimeFormat(t *testing.T) {
	day := "20200723"
	d, err := time.Parse("20060102", day)
	assert.NoError(t, err)
	t.Log(d.Format("2006-01-02"))
}
