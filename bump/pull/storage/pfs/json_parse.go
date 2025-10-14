package pfs

import (
	"Catch/bump/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"strings"
)

func ResourceInfoALL(body string) []ParallelFileStorage {

	var sliceParallelFileStorage []ParallelFileStorage
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.content|@pretty")

	//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		var parallelFileStorage ParallelFileStorage
		parallelFileStorage.ShareId = value.Get("shareId").String()
		parallelFileStorage.Name = value.Get("name").String()
		parallelFileStorage.PoolId = value.Get("poolId").String()
		parallelFileStorage.Region = value.Get("region").String()
		parallelFileStorage.Size = value.Get("size").String()
		parallelFileStorage.ShareType = Convert(value.Get("shareType").String())

		//处理json格式
		var e ExportLocation
		elJson := value.Get("exportLocation").String()
		elJson = strings.ReplaceAll(elJson, "\\", "")
		err := json.Unmarshal([]byte(elJson), &e)
		if err != nil {
			logrus.Warnln(err)
			parallelFileStorage.ExportLocation = "Fail"
		} else {
			parallelFileStorage.ExportLocation = fmt.Sprintf("IPv4：%s\nIPv6：%s\nSharePath：%s\n", e.IPv4, e.IPv6, e.SharePath)
		}

		sliceParallelFileStorage = append(sliceParallelFileStorage, parallelFileStorage)

		return true
	})

	return sliceParallelFileStorage
}

func JsonParseResourceByUUID(body string, uuids []string) []*ParallelFileStorage {

	var sliceParallelFileStorage []*ParallelFileStorage
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.content|@pretty")

	for _, uuid := range uuids {
		//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
		//循环content下的所有元素
		JsonBody.ForEach(func(key, value gjson.Result) bool {
			if uuid != value.Get("shareId").String() {
				return true
			}

			var parallelFileStorage ParallelFileStorage
			parallelFileStorage.ShareId = value.Get("shareId").String()
			parallelFileStorage.Name = value.Get("name").String()
			parallelFileStorage.PoolId = value.Get("poolId").String()
			parallelFileStorage.Region = value.Get("region").String()
			parallelFileStorage.Size = value.Get("size").String()
			parallelFileStorage.ShareType = Convert(value.Get("shareType").String())

			//处理json格式
			var e ExportLocation
			elJson := value.Get("exportLocation").String()
			elJson = strings.ReplaceAll(elJson, "\\", "")
			err := json.Unmarshal([]byte(elJson), &e)
			if err != nil {
				logrus.Warnln(err)
				parallelFileStorage.ExportLocation = "Fail"
			} else {
				parallelFileStorage.ExportLocation = fmt.Sprintf("IPv4：%s\nIPv6：%s\nSharePath：%s\n", e.IPv4, e.IPv6, e.SharePath)
			}

			sliceParallelFileStorage = append(sliceParallelFileStorage, &parallelFileStorage)

			return false
		})
	}
	return sliceParallelFileStorage
}

func JsonParseTag(pfss []*ParallelFileStorage) {
	body := utils.Post(tagUrl, HeadersNoPool(), strings.NewReader(payloadPfs))
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
