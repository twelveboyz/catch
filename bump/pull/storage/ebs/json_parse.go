package ebs

import (
	"Catch/bump/internal/mapping"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

func ResourceInfoALL(body string) []ElasticBlockStorage {

	var sliceElasticBlockStorage []ElasticBlockStorage
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.content|@pretty")

	//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		var elasticBlockStorage ElasticBlockStorage
		elasticBlockStorage.Id = value.Get("id").String()
		elasticBlockStorage.Name = value.Get("name").String()
		elasticBlockStorage.Region = mapping.AZConvert(value.Get("region").String())
		elasticBlockStorage.Size = value.Get("size").String()
		elasticBlockStorage.IsShare = isShareConvert(value.Get("isShare").String())
		elasticBlockStorage.Type = Convert(value.Get("type").String())
		elasticBlockStorage.ServerId = value.Get("attachSevers.0").Get("serverId").String()
		elasticBlockStorage.ServerName = value.Get("attachSevers.0").Get("serverName").String()

		//tags
		value.ForEach(func(key, value gjson.Result) bool {
			switch value.Get("key").String() {
			case "标签1":
				elasticBlockStorage.Tag1 = value.Get("value").String()
			case "标签X":
				elasticBlockStorage.TagX = value.Get("value").String()
			case "标签2":
				elasticBlockStorage.Tag2 = value.Get("value").String()
			case "订购人":
				elasticBlockStorage.Orderer = value.Get("value").String()
			}
			return true
		})

		sliceElasticBlockStorage = append(sliceElasticBlockStorage, elasticBlockStorage)
		//fmt.Println("test", cloudPort)
		return true
	})

	return sliceElasticBlockStorage
}

func JsonParseResourceByUUID(body string, uuids []string) []*ElasticBlockStorage {

	var sliceElasticBlockStorage []*ElasticBlockStorage
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.content|@pretty")

	for _, uuid := range uuids {

		//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
		//循环content下的所有元素
		JsonBody.ForEach(func(key, value gjson.Result) bool {
			if uuid != value.Get("id").String() {
				return true
			}

			var ebs ElasticBlockStorage
			ebs.Id = value.Get("id").String()
			ebs.Name = value.Get("name").String()
			ebs.Region = mapping.AZConvert(value.Get("region").String())
			ebs.Size = value.Get("size").String()
			ebs.IsShare = isShareConvert(value.Get("isShare").String())
			ebs.Type = Convert(value.Get("type").String())
			ebs.ServerId = value.Get("attachSevers.0").Get("serverId").String()
			ebs.ServerName = value.Get("attachSevers.0").Get("serverName").String()

			//tags
			value.Get("tags").ForEach(func(key, value gjson.Result) bool {
				switch value.Get("key").String() {
				case "标签1":
					ebs.Tag1 = value.Get("value").String()
				case "标签X":
					ebs.TagX = value.Get("value").String()
				case "标签2":
					ebs.Tag2 = value.Get("value").String()
				case "订购人":
					ebs.Orderer = value.Get("value").String()
				}
				return true
			})

			if ebs.Tag1 == "" || ebs.Tag2 == "" || ebs.TagX == "" || ebs.Orderer == "" {
				logrus.Warnf("未找到标签，id='%s' name='%s' ,标签1='%s' 标签X='%s' 标签2='%s' 订购人='%s'\n", ebs.Id, ebs.Name, ebs.Tag1, ebs.TagX, ebs.Tag2, ebs.Orderer)
			}

			sliceElasticBlockStorage = append(sliceElasticBlockStorage, &ebs)
			//fmt.Println("test", cloudPort)
			return false
		})
	}
	return sliceElasticBlockStorage
}
