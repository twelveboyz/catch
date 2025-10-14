package wafCloudNative

import "github.com/tidwall/gjson"

func ResourceInfoALL(body string) []WafCloudNative {

	var sliceWafCloudNative []WafCloudNative
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.content|@pretty")

	//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		var wafCloudNative WafCloudNative
		wafCloudNative.Id = value.Get("id").String()
		wafCloudNative.Description = value.Get("description").String()
		wafCloudNative.IpNum = value.Get("ipNum").String()
		wafCloudNative.Bandwidth = value.Get("bandwidth").String()
		wafCloudNative.PeriodType = value.Get("periodType").String()

		sliceWafCloudNative = append(sliceWafCloudNative, wafCloudNative)
		//fmt.Println("test", cloudPort)
		return true
	})

	return sliceWafCloudNative
}
