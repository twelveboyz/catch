package efs

import (
	"Catch/bump/internal/mapping"
	"Catch/bump/internal/utils"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"strings"
)

func ResourceInfoALL(body string) []*ElasticFileStorage {

	var sliceElasticFileStorage []*ElasticFileStorage
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.content|@pretty")

	//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		var elasticFileStorage ElasticFileStorage
		elasticFileStorage.ShareId = value.Get("shareId").String()
		elasticFileStorage.Name = value.Get("name").String()
		elasticFileStorage.PoolId = value.Get("poolId").String()
		elasticFileStorage.Region = mapping.AZConvert(value.Get("region").String())
		elasticFileStorage.Size = value.Get("size").String()
		elasticFileStorage.ShareType = Convert(value.Get("shareType").String())
		sliceElasticFileStorage = append(sliceElasticFileStorage, &elasticFileStorage)
		//fmt.Println("test", cloudPort)
		return true
	})

	return sliceElasticFileStorage
}

func JsonParseResourceByUUIDs(body string, uuids []string) []*ElasticFileStorage {

	var sliceElasticFileStorage []*ElasticFileStorage
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.content|@pretty")

	for _, uuid := range uuids {
		//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
		//循环content下的所有元素
		JsonBody.ForEach(func(key, value gjson.Result) bool {
			if uuid != value.Get("shareId").String() {
				return true
			}

			var elasticFileStorage ElasticFileStorage
			elasticFileStorage.ShareId = value.Get("shareId").String()
			elasticFileStorage.Name = value.Get("name").String()
			elasticFileStorage.PoolId = value.Get("poolId").String()
			elasticFileStorage.Region = mapping.AZConvert(value.Get("region").String())
			elasticFileStorage.Size = value.Get("size").String()
			elasticFileStorage.ShareType = Convert(value.Get("shareType").String())

			mount, err := MountInfoFormat(value.Get("exportLocation").String())
			if err != nil {
				elasticFileStorage.Mount = "parse err"
				logrus.Warnln("mount parse err：" + err.Error())
			} else {
				elasticFileStorage.Mount = mount
			}

			sliceElasticFileStorage = append(sliceElasticFileStorage, &elasticFileStorage)
			//fmt.Println("test", cloudPort)
			return false
		})
	}
	return sliceElasticFileStorage
}

func JsonParseTag(pfss []*ElasticFileStorage) {
	body := utils.Post(tagUrl, HeadersNoPool(), strings.NewReader(payloadEfs))
	JsonBody := gjson.Get(body, "body.content|@pretty")
	for _, pfs := range pfss {
		JsonBody.ForEach(func(key, value gjson.Result) bool {
			//跳过不匹配的uuid
			if pfs.ShareId != value.Get("resourceId").String() {
				return true
			}

			value.Get("tag").ForEach(func(key, value gjson.Result) bool {
				switch value.Get("tagKey").String() {
				case "标签1":
					pfs.Tag1 = value.Get("tagValue").String()
				case "标签X":
					pfs.TagX = value.Get("tagValue").String()
				case "标签2":
					pfs.Tag2 = value.Get("tagValue").String()
				case "订购人":
					pfs.Orderer = value.Get("tagValue").String()
				}
				return true
			})

			/*	value.Get("tag").ForEach(func(key, value gjson.Result) bool {
				logrus.Warnln("key=", value.Get("tagKey").String())
				switch value.Get("tagKey").String() {
				case "标签1":
					pfs.Tag1 = value.Get("tagValue").String()
					logrus.Warnln("tag1=", value.Get("tagKey").String())
				case "标签X":
					pfs.TagX = value.Get("tagValue").String()
					logrus.Warnln("tagX=", value.Get("tagKey").String())
				case "标签2":
					pfs.Tag2 = value.Get("tagValue").String()
					logrus.Warnln("tag2=", value.Get("tagKey").String())
				case "订购人":
					pfs.Orderer = value.Get("tagValue").String()
					logrus.Warnln("Orderer=", value.Get("tagKey").String())
				}
				return true
			})*/

			return false
		})
	}
}
