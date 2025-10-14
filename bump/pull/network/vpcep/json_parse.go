package vpcep

import "github.com/tidwall/gjson"

func ResourceInfoALL(body string) []VpcEndPoint {

	var sliceVpcEndPoint []VpcEndPoint
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.content|@pretty")

	//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		var vpcEndPoint VpcEndPoint
		vpcEndPoint.VpcEpId = value.Get("vpcepId").String()
		vpcEndPoint.SpecialLineId = value.Get("speciallineId").String()
		vpcEndPoint.VpcEpName = value.Get("vpcepName").String()
		vpcEndPoint.VpcEpServiceName = value.Get("vpcepServiceName").String()
		vpcEndPoint.ProductName = value.Get("productName").String()
		vpcEndPoint.Region = value.Get("region").String()
		vpcEndPoint.RegionName = value.Get("regionName").String()
		sliceVpcEndPoint = append(sliceVpcEndPoint, vpcEndPoint)
		//fmt.Println("test", cloudPort)
		return true
	})

	return sliceVpcEndPoint
}
