package eos

import (
	"github.com/tidwall/gjson"
)

func JsonParseResourceInfo(body string) []*ElasticObjectStorage {
	var sliceElasticObjectStorage []*ElasticObjectStorage
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body|@pretty")

	//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		var elasticObjectStorage ElasticObjectStorage
		elasticObjectStorage.Id = value.Get("bucket.owner.id").String()
		elasticObjectStorage.Name = value.Get("bucket.name").String()
		elasticObjectStorage.StorageClass = value.Get("bucket.storageClass").String()
		elasticObjectStorage.Region = PoolConvert(value.Get("region").String())

		sliceElasticObjectStorage = append(sliceElasticObjectStorage, &elasticObjectStorage)
		//fmt.Println("test", cloudPort)
		return true
	})
	return sliceElasticObjectStorage
}
