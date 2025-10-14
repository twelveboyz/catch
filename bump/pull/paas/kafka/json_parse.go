package kafka

import (
	"Catch/bump/internal/mapping"
	"github.com/tidwall/gjson"
)

func ResourceInfoALL(body string) []Kafka {

	var sliceKafka []Kafka
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "instances|@pretty")

	//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		var kafka Kafka
		kafka.Id = value.Get("id").String()
		kafka.Name = value.Get("name").String()
		kafka.Region = mapping.AZConvert(value.Get("region").String())
		kafka.InstanceVersion = value.Get("instanceVersion").String()
		kafka.SpecType = convert(value.Get("specType").String())
		kafka.ConnectAddress = value.Get("connectAddress").String()
		kafka.Ipv6ConnectAddress = value.Get("ipv6ConnectAddress").String()
		kafka.Cpu = value.Get("cpu").String()
		kafka.Memory = value.Get("memory").String()
		kafka.VolumeSize = value.Get("volumeSize").Int()
		kafka.StorageCapacity = value.Get("storageCapacity").Int()
		kafka.MessageSizeMax = value.Get("messageSizeMax").String()
		kafka.Brokers = value.Get("brokers").String()
		sliceKafka = append(sliceKafka, kafka)
		//fmt.Println("test", cloudPort)
		return true
	})

	return sliceKafka
}
