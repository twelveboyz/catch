package natgateway

import (
	"Catch/bump/internal/utils"
	"fmt"
	"github.com/tidwall/gjson"
	"strings"
	"sync"
)

func ResourceInfoALL(body string) []*NATGateway {

	var sliceNATGateway []*NATGateway
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.content|@pretty")

	//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		var natGateway NATGateway
		natGateway.Id = value.Get("id").String()
		natGateway.Name = value.Get("name").String()
		natGateway.Region = value.Get("region").String()
		natGateway.Scale = Convert(value.Get("scale").String())
		natGateway.VpcId = value.Get("vpcId").String()
		natGateway.VpcName = value.Get("vpcName").String()
		natGateway.RouterId = value.Get("routerId").String()
		natGateway.Bandwidth = value.Get("bandwidth").String()
		natGateway.PeriodType = value.Get("periodType").String()

		sliceNATGateway = append(sliceNATGateway, &natGateway)
		//fmt.Println("test", cloudPort)
		return true
	})

	return sliceNATGateway
}

func ResourceMatchByUUID(body string, uuids []string) []*NATGateway {

	var sliceNATGateway []*NATGateway
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.content|@pretty")

	for _, uuid := range uuids {

		//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
		//循环content下的所有元素
		JsonBody.ForEach(func(key, value gjson.Result) bool {
			var natGateway NATGateway
			if uuid != value.Get("id").String() {
				return true
			}
			natGateway.Id = value.Get("id").String()
			natGateway.Name = value.Get("name").String()
			natGateway.Region = value.Get("region").String()
			natGateway.Scale = Convert(value.Get("scale").String())
			natGateway.VpcId = value.Get("vpcId").String()
			natGateway.VpcName = value.Get("vpcName").String()
			natGateway.RouterId = value.Get("routerId").String()
			natGateway.Bandwidth = value.Get("bandwidth").String()
			natGateway.PeriodType = value.Get("periodType").String()

			sliceNATGateway = append(sliceNATGateway, &natGateway)
			//fmt.Println("test", cloudPort)
			return false
		})
	}
	return sliceNATGateway
}

func JsonParseTag(nats []*NATGateway) {
	var wg sync.WaitGroup
	wg.Add(len(nats))

	for _, nat := range nats {
		go func() {
			defer wg.Done()
			jsonBody := utils.Post(tagUrl, HeadersFun(nat.Pool), strings.NewReader(fmt.Sprintf("{\"instanceId\":\"%s\"}", nat.Id)))

			gjson.Get(jsonBody, "body.customerRemark.tags").ForEach(func(key, value gjson.Result) bool {
				switch value.Get("key").String() {
				case "标签1":
					nat.Tag1 = value.Get("value").String()
				case "标签X":
					nat.TagX = value.Get("value").String()
				case "标签2":
					nat.Tag2 = value.Get("value").String()
				case "订购人":
					nat.Orderer = value.Get("value").String()
				}
				return true
			})

			/*	gjson.Get(jsonBody, "body.customerRemark.tags").ForEach(func(key, value gjson.Result) bool {
				logrus.Infoln(value.Get("key").String())
				switch value.Get("key").String() {
				case "标签1":
					nat.Tag1 = value.Get("value").String()
					logrus.Infoln("标签1=", value.Get("value").String())
				case "标签X":
					nat.TagX = value.Get("value").String()
					logrus.Infoln("标签X=", value.Get("value").String())
				case "标签2":
					nat.Tag2 = value.Get("value").String()
					logrus.Infoln("标签2=", value.Get("value").String())
				case "订购人":
					nat.Orderer = value.Get("value").String()
					logrus.Infoln("订购人=", value.Get("value").String())
				}
				return true
			})*/
		}()
	}

	wg.Wait()
}
