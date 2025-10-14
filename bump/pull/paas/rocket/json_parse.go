package rocket

import "github.com/tidwall/gjson"

func ResourceInfoALL(body string) []Rocket {

	var sliceRocket []Rocket
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body|@pretty")

	//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		var rocket Rocket
		rocket.Id = value.Get("id").String()
		rocket.ResourceId = value.Get("resourceId").String()
		rocket.AvailZone = value.Get("availZone").String()
		rocket.Type = value.Get("type").String()
		rocket.Tps = value.Get("tps").String()
		rocket.MaxTopicSize = value.Get("maxTopicSize").String()
		rocket.Storage = value.Get("storage").String()
		rocket.IsExclusiveResource = value.Get("isExclusiveResource").String()
		rocket.PeriodType = value.Get("periodType").String()

		sliceRocket = append(sliceRocket, rocket)

		return true
	})

	return sliceRocket
}
