package wafProfessional

import "github.com/tidwall/gjson"

func ResourceInfoALL(body string) []WafProfessional {

	var sliceWafProfessional []WafProfessional
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.list|@pretty")

	//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		var wafProfessional WafProfessional
		wafProfessional.InstanceId = value.Get("instanceId").String()
		wafProfessional.InstanceName = value.Get("instanceName").String()
		wafProfessional.Bandwidth = BandwidthCount(value.Get("bandwidth").Int())
		wafProfessional.ResourcePoolName = value.Get("resourcePoolName").String()
		wafProfessional.ExclusivePackage = value.Get("exclusivePackage").String()

		sliceWafProfessional = append(sliceWafProfessional, wafProfessional)
		//fmt.Println("test", cloudPort)
		return true
	})

	return sliceWafProfessional
}
