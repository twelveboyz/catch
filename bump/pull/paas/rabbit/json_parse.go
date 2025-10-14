package rabbit

import "github.com/tidwall/gjson"

func ResourceInfoALL(body string) []Rabbit {

	var sliceRabbit []Rabbit
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "data|@pretty")

	//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		var rabbit Rabbit
		rabbit.InstanceId = value.Get("instanceId").String()
		rabbit.ClusterName = value.Get("clusterName").String()
		rabbit.Region = convert(value.Get("region").String())
		rabbit.VCore = value.Get("vcore").String()
		rabbit.Memory = value.Get("memory").String()
		rabbit.InstanceType = value.Get("instanceType").String()
		rabbit.Version = value.Get("version").String()
		rabbit.StorageCapacity = value.Get("storageCapacity").String()
		sliceRabbit = append(sliceRabbit, rabbit)
		//fmt.Println("test", xxx)
		return true
	})

	return sliceRabbit
}
