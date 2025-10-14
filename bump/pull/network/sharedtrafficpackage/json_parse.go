package sharedtrafficpackage

import "github.com/tidwall/gjson"

func ResourceInfoALL(body string) []SharedTrafficPackage {

	var sliceSharedTrafficPackage []SharedTrafficPackage
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.content|@pretty")

	//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		var sharedTrafficPackage SharedTrafficPackage
		sharedTrafficPackage.Id = value.Get("id").String()
		sharedTrafficPackage.Uuid = value.Get("uuid").String()
		sharedTrafficPackage.Name = value.Get("name").String()
		sharedTrafficPackage.Spec = Convert(value.Get("spec").String())
		sharedTrafficPackage.Size = value.Get("size").Int() / 1024 / 1024 / 1024 / 1024

		sliceSharedTrafficPackage = append(sliceSharedTrafficPackage, sharedTrafficPackage)
		//fmt.Println("test", cloudPort)
		return true
	})

	return sliceSharedTrafficPackage
}

func ResourceMatchByUUID(body string, uuids []string) []SharedTrafficPackage {

	var sliceSharedTrafficPackage []SharedTrafficPackage
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.content|@pretty")

	for _, uuid := range uuids {
		//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
		//循环content下的所有元素
		JsonBody.ForEach(func(key, value gjson.Result) bool {
			var sharedTrafficPackage SharedTrafficPackage
			if uuid != value.Get("id").String() {
				return true
			}
			sharedTrafficPackage.Id = value.Get("id").String()
			sharedTrafficPackage.Uuid = value.Get("uuid").String()
			sharedTrafficPackage.Name = value.Get("name").String()
			sharedTrafficPackage.Spec = Convert(value.Get("spec").String())
			sharedTrafficPackage.Size = value.Get("size").Int() / 1024 / 1024 / 1024 / 1024

			sliceSharedTrafficPackage = append(sliceSharedTrafficPackage, sharedTrafficPackage)
			//fmt.Println("test", cloudPort)
			return false
		})
	}
	return sliceSharedTrafficPackage
}
