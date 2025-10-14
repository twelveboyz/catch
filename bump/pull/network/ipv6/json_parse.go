package ipv6

import (
	"github.com/tidwall/gjson"
)

func ResourceInfoALL(body string) []*IPv6 {

	var sliceIPv6 []*IPv6
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.content|@pretty")

	//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		var ipv6 IPv6
		ipv6.NsQosPolicyId = value.Get("nsQosPolicyId").String()
		ipv6.MixedId = value.Get("mixedId").String()
		ipv6.IpAddress = value.Get("ipAddress").String()
		ipv6.BandWidthSize = value.Get("bandWidthSize").String()
		ipv6.BandwidthType = value.Get("bandwidthType").String()
		ipv6.BindResourceId = value.Get("resourceId").String()
		ipv6.BindResourceName = value.Get("resourceName").String()

		sliceIPv6 = append(sliceIPv6, &ipv6)
		return true
	})

	return sliceIPv6
}
