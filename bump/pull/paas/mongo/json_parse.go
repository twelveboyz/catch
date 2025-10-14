package mongo

import "github.com/tidwall/gjson"

func ResourceInfoALL(body string) []Mongo {

	var sliceMongo []Mongo
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "data.items|@pretty")

	//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		var mongo Mongo
		mongo.Id = value.Get("id").String()
		mongo.Name = value.Get("name").String()
		mongo.AvailableAreaName = value.Get("availableAreaName").String()
		storage := int(value.Get("storage.total").Int())
		mongo.StorageTotal = storage / 1024 / 1024 / 1024
		mongo.Version = value.Get("spec.version").String()
		mongo.Engine = value.Get("spec.engine").String()
		mongo.Type = value.Get("spec.type").String()
		mongo.BackupStorage = int(value.Get("spec.backupStorage").Int())

		mongo.MongosCount = value.Get("spec.sharding.mongosCount").String()
		mongo.MongosOsSpecType = value.Get("spec.sharding.mongosOsSpec.type").String()
		mongo.MongosOsSpecCpu = value.Get("spec.sharding.mongosOsSpec.cpu").String()
		mongo.MongosOsSpecMemory = value.Get("spec.sharding.mongosOsSpec.memory").String()
		mongo.MongosOsSpecStorage = value.Get("spec.sharding.mongosOsSpec.storage").String()

		mongo.ShardCount = value.Get("spec.sharding.shardCount").String()
		mongo.MongodOsSpecType = value.Get("spec.sharding.mongodOsSpec.type").String()
		mongo.MongodOsSpecCpu = value.Get("spec.sharding.mongodOsSpec.cpu").String()
		mongo.MongodOsSpecMemory = value.Get("spec.sharding.mongodOsSpec.memory").String()
		mongo.MongodOsSpecStorage = value.Get("spec.sharding.mongodOsSpec.storage").String()

		mongo.ConfigServerCount = value.Get("spec.sharding.configServerCount").String()
		mongo.ConfigServerSpecType = value.Get("spec.sharding.configServerSpec.type").String()
		mongo.ConfigServerSpecCpu = value.Get("spec.sharding.configServerSpec.cpu").String()
		mongo.ConfigServerSpecMemory = value.Get("spec.sharding.configServerSpec.memory").String()
		mongo.ConfigServerSpecStorage = value.Get("spec.sharding.configServerSpec.storage").String()

		sliceMongo = append(sliceMongo, mongo)
		//fmt.Println("test", xxx)
		return true
	})

	return sliceMongo
}
