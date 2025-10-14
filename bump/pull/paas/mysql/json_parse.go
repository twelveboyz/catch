package mysql

import (
	"github.com/tidwall/gjson"
)

func ResourceInfo(body string, resourceName []string) []Mysql {

	var sliceCloudPort []Mysql
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body|@pretty")

	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		//fmt.Println(key) //打印循环次数

		//循环过滤resourceName对应的数据
		for _, v := range resourceName {
			//匹配资源名称，过滤想要的资源加入到slice
			if v == value.Get("name").String() {
				var mysql Mysql
				mysql.Id = value.Get("id").String()
				mysql.Name = value.Get("name").String()
				mysql.Engine = value.Get("engine").String()
				mysql.Version = value.Get("version").String()
				mysql.Az = value.Get("az").String()
				mysql.Measuretype = value.Get("measureType").String()
				mysql.Instancespec = value.Get("instanceSpec").String()
				mysql.Volumesize = value.Get("volumeSize").String()

				sliceCloudPort = append(sliceCloudPort, mysql)
			}
		}

		if len(sliceCloudPort) == len(resourceName) {
			return false
		}

		//fmt.Println("test", cloudPort)
		return true
	})

	return sliceCloudPort
}

func ResourceInfoALL(body string) []Mysql {

	var sliceCloudPort []Mysql
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body|@pretty")

	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		var mysql Mysql
		mysql.Id = value.Get("id").String()
		mysql.Name = value.Get("name").String()
		mysql.Engine = value.Get("engine").String()
		mysql.Version = value.Get("version").String()
		mysql.Az = value.Get("az").String()
		mysql.Measuretype = value.Get("measureType").String()
		mysql.Instancespec = value.Get("instanceSpec").String()
		mysql.Volumesize = value.Get("volumeSize").String()

		sliceCloudPort = append(sliceCloudPort, mysql)
		//fmt.Println("test", cloudPort)
		return true
	})

	return sliceCloudPort
}
