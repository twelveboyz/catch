package directconnect

import (
	"github.com/tidwall/gjson"
)

func ResourceInfoALL(body string) []DirectConnect {

	var sliceCloudSpecialLines []DirectConnect
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.content|@pretty")

	//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		var cloudSpecialLines DirectConnect
		cloudSpecialLines.Id = value.Get("id").String()
		cloudSpecialLines.NetSideId = value.Get("netSideId").String()
		cloudSpecialLines.SpecialLineName = value.Get("speciallineName").String()
		cloudSpecialLines.PoolId = value.Get("poolId").String()
		cloudSpecialLines.IpType = value.Get("iptype").String()
		cloudSpecialLines.SpecialLineBandwidth = value.Get("speciallineBandwidth").String()
		cloudSpecialLines.VpcId = value.Get("vpcId").String()
		cloudSpecialLines.VpcName = value.Get("vpcName").String()

		var vpcSubnets []string
		for _, vpcSubnet := range value.Get("vpsSubnets").Array() {
			vpcSubnets = append(vpcSubnets, vpcSubnet.String())
		}

		var userSubnets []string
		for _, userSubnet := range value.Get("userSubnets").Array() {
			userSubnets = append(userSubnets, userSubnet.String())
		}

		cloudSpecialLines.VpsSubnets = vpcSubnets
		cloudSpecialLines.UserSubnets = userSubnets

		cloudSpecialLines.Province = value.Get("province").String()
		cloudSpecialLines.City = value.Get("city").String()
		cloudSpecialLines.Borough = value.Get("borough").String()
		cloudSpecialLines.Address = value.Get("address").String()
		cloudSpecialLines.ContactName = value.Get("contactName").String()
		cloudSpecialLines.ManagerName = value.Get("managerName").String()

		sliceCloudSpecialLines = append(sliceCloudSpecialLines, cloudSpecialLines)
		//fmt.Println("test", cloudPort)
		return true
	})

	return sliceCloudSpecialLines
}
