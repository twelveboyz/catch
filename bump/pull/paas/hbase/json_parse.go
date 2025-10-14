package hbase

import "github.com/tidwall/gjson"

func ResourceInfoALL(body string) []HBase {

	var sliceHBase []HBase
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "data.resultList|@pretty")

	//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		var hbase HBase
		hbase.InstanceId = value.Get("instanceId").String()
		hbase.InstanceName = value.Get("instanceName").String()
		hbase.Region = value.Get("region").String()
		hbase.Version = value.Get("version").String()
		hbase.ProductSd = value.Get("productSd").Int()

		hbase.MasterConfigVCore = value.Get("masterConfig.vcore").Int()
		hbase.MasterConfigMemory = value.Get("masterConfig.memory").Int()
		hbase.MasterConfigCount = value.Get("masterConfig.count").Int()
		hbase.CoreConfigVCore = value.Get("coreConfig.vcore").Int()
		hbase.CoreConfigMemory = value.Get("coreConfig.memory").Int()
		hbase.CoreConfigCount = value.Get("coreConfig.count").Int()

		hbase.EsConfigVCore = value.Get("esConfig.vcore").Int()
		hbase.EsConfigMemory = value.Get("esConfig.memory").Int()
		hbase.EsConfigCount = value.Get("esConfig.count").Int()

		hbase.StorageConfigType = convert(value.Get("storageConfig.type").String())
		hbase.StorageConfigCapacity = value.Get("storageConfig.capacity").Int()
		sliceHBase = append(sliceHBase, hbase)
		//fmt.Println("test", xxx)
		return true
	})

	return sliceHBase
}
