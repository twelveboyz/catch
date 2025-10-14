package redis

import "github.com/tidwall/gjson"

func ResourceInfoALL(body string) []Redis {

	var sliceRedis []Redis
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.content|@pretty")

	//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		var redis Redis
		redis.Id = value.Get("id").String()
		redis.Name = value.Get("name").String()
		redis.Version = value.Get("version").String()
		redis.Type = convert(value.Get("type").String())
		redis.Shards = value.Get("shards").String()
		redis.Memory = value.Get("memory").String()
		redis.ChargeType = value.Get("chargeType").String()
		redis.RedisRegion = value.Get("redisRegion").Array()
		redis.RedisRegionFormat = value.Get("redisRegionFormat").Array()
		redis.SlaveCount = value.Get("slaveCount").String()

		sliceRedis = append(sliceRedis, redis)
		//fmt.Println("test", cloudPort)
		return true
	})

	return sliceRedis
}
