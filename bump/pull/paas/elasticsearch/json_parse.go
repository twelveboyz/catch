package elasticsearch

import "github.com/tidwall/gjson"

func ResourceInfoALL(body string) []ElasticSearch {

	var sliceElasticSearch []ElasticSearch
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "data.items|@pretty")

	//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		var elasticSearch ElasticSearch
		elasticSearch.Id = value.Get("id").String()
		elasticSearch.ZoneZoneName = value.Get("zone.zoneName").String()
		elasticSearch.AppType = value.Get("appType").String()
		elasticSearch.Alias = value.Get("alias").String()

		var server Server
		serversJsonBody := value.Get("servs")
		serversJsonBody.ForEach(func(key, value gjson.Result) bool {
			if value.Get("type.code").String() != "infini-gateway" {
				server.TypeCode = value.Get("type.code").String()
				server.ArchId = value.Get("arch.id").String()
				server.ArchName = value.Get("arch.name").String()
				server.ArchMode = value.Get("arch.mode").String()
				server.ArchShardCnt = value.Get("arch.shardCnt").String()
				server.ArchUnitCnt = value.Get("arch.unitCnt").String()
				server.ScaleName = value.Get("scale.name").String()
				server.ScaleCpuCnt = value.Get("scale.cpuCnt").String()
				server.ScaleMemSize = value.Get("scale.memSize").String()
				server.DiskTypeDisplay = value.Get("diskType.display").String()
				server.DataSize = value.Get("dataSize").String()

				elasticSearch.Servers = append(elasticSearch.Servers, server)
			}
			return true
		})

		sliceElasticSearch = append(sliceElasticSearch, elasticSearch)
		//fmt.Println("test", xxx)
		return true
	})

	return sliceElasticSearch
}
