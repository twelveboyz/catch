package privatedns

import (
	"github.com/tidwall/gjson"
)

func ResourceInfoALL(body string) []PrivateDNS {

	var slicePrivateDNS []PrivateDNS
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "data.data|@pretty")

	//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		var privateDNS PrivateDNS
		privateDNS.InstanceId = value.Get("instanceId").String()
		privateDNS.VDnsId = value.Get("vdnsId").String()
		privateDNS.VDnsName = value.Get("vdnsName").String()
		privateDNS.Ipv4 = value.Get("ipv4").String()
		privateDNS.Ipv6 = value.Get("ipv6").String()
		privateDNS.IpType = value.Get("ipType").String()
		privateDNS.PackageType = value.Get("packageType").String()

		slicePrivateDNS = append(slicePrivateDNS, privateDNS)
		return true
	})

	return slicePrivateDNS
}
