package cis

import "github.com/tidwall/gjson"

func ResourceInfoALL(body string) []ContainerImageService {

	var sliceContainerImageService []ContainerImageService
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.data|@pretty")

	//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		var containerImageService ContainerImageService
		containerImageService.Uid = value.Get("uid").String()
		containerImageService.InstanceName = value.Get("instance_name").String()
		containerImageService.Specification = Convert(value.Get("specification").String())
		containerImageService.BucketName = value.Get("bucket_name").String()
		containerImageService.Domain = value.Get("domain").String()
		containerImageService.PeriodType = value.Get("period_type").String()

		sliceContainerImageService = append(sliceContainerImageService, containerImageService)
		//fmt.Println("test", cloudPort)
		return true
	})

	return sliceContainerImageService
}
