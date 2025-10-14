package eip

import (
	"Catch/bump/internal/utils"
	"github.com/tidwall/gjson"
	"strconv"
	"sync"
)

func ResourceInfoALL(body string) []ElasticIP {

	var sliceElasticIP []ElasticIP
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.content|@pretty")

	//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		var elasticIP ElasticIP
		elasticIP.Id = value.Get("id").String()
		elasticIP.EipName = value.Get("name").String()
		elasticIP.BandwidthSize = strconv.Itoa(int(value.Get("bandwidthSize").Int() / 1024))
		elasticIP.BindType = Convert(value.Get("bindType").String())
		elasticIP.BandwidthType = value.Get("bandwidthType").String()
		elasticIP.BindResourceId = value.Get("resourceId").String()
		elasticIP.BindResourceName = value.Get("resourceName").String()

		sliceElasticIP = append(sliceElasticIP, elasticIP)
		//fmt.Println("test", cloudPort)
		return true
	})

	return sliceElasticIP
}

func ResourceMatchByUUID(body string, uuids []string) []*ElasticIP {

	var sliceElasticIP []*ElasticIP
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.content|@pretty")
	for _, uuid := range uuids {
		//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
		//循环content下的所有元素
		JsonBody.ForEach(func(key, value gjson.Result) bool {
			var elasticIP ElasticIP

			if uuid != value.Get("id").String() {
				return true
			}

			elasticIP.Id = value.Get("id").String()
			elasticIP.ResourceName = value.Get("description").String()
			elasticIP.EipName = value.Get("name").String()
			elasticIP.BandwidthSize = strconv.Itoa(int(value.Get("bandwidthSize").Int() / 1024))
			elasticIP.BindType = Convert(value.Get("bindType").String())
			elasticIP.BandwidthType = value.Get("bandwidthType").String()
			elasticIP.BindResourceId = value.Get("resourceId").String()
			elasticIP.BindResourceName = value.Get("resourceName").String()

			sliceElasticIP = append(sliceElasticIP, &elasticIP)
			//fmt.Println("test", cloudPort)
			return false
		})
	}
	return sliceElasticIP
}

func JsonParseTag(eips []*ElasticIP) {
	var wg sync.WaitGroup
	wg.Add(len(eips))

	for _, eip := range eips {
		go func() {
			defer wg.Done()
			jsonBody := utils.Get(tagUrl+eip.Id, HeadersNoPool())

			gjson.Get(jsonBody, "body").ForEach(func(key, value gjson.Result) bool {
				switch value.Get("tagKey").String() {
				case "标签1":
					eip.Tag1 = value.Get("tagValue").String()
				case "标签X":
					eip.TagX = value.Get("tagValue").String()
				case "标签2":
					eip.Tag2 = value.Get("tagValue").String()
				case "订购人":
					eip.Orderer = value.Get("tagValue").String()
				}
				return true
			})
		}()
	}

	wg.Wait()
}
