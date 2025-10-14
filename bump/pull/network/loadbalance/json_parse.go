package loadbalance

import (
	"Catch/bump/internal/utils"
	"github.com/tidwall/gjson"
	"sync"
)

func ResourceInfoALL(body string) []LoadBalance {

	var sliceLoadBalance []LoadBalance
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.content|@pretty")

	//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		var loadBalance LoadBalance
		loadBalance.Id = value.Get("id").String()
		loadBalance.Name = value.Get("name").String()
		loadBalance.Region = value.Get("region").String()
		loadBalance.Bandwidth = value.Get("bandwidth").Int()
		loadBalance.Flavor = Convert(value.Get("flavor").String())
		loadBalance.Privateip = value.Get("privateIp").String()
		loadBalance.Vpcname = value.Get("vpcName").String()
		loadBalance.Routerid = value.Get("routerId").String()
		loadBalance.Networkid = value.Get("networkId").String()
		loadBalance.Subnetid = value.Get("subnetId").String()
		loadBalance.SubnetName = value.Get("subnetName").String()
		loadBalance.IPVersion = value.Get("ipVersion").String()
		loadBalance.PublicIp = value.Get("publicIp").String()

		sliceLoadBalance = append(sliceLoadBalance, loadBalance)
		//fmt.Println("test", cloudPort)
		return false
	})

	return sliceLoadBalance
}

func ResourceMatchByUUID(body string, uuids []string) []*LoadBalance {

	var sliceLoadBalance []*LoadBalance
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.content|@pretty")

	for _, uuid := range uuids {
		//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
		//循环content下的所有元素
		JsonBody.ForEach(func(key, value gjson.Result) bool {
			var loadBalance LoadBalance

			if uuid != value.Get("id").String() {
				return true
			}

			loadBalance.Id = value.Get("id").String()
			loadBalance.Name = value.Get("name").String()
			loadBalance.Region = value.Get("region").String()
			loadBalance.Bandwidth = value.Get("bandwidth").Int()
			loadBalance.Flavor = Convert(value.Get("flavor").String())
			loadBalance.Privateip = value.Get("privateIp").String()
			loadBalance.Vpcname = value.Get("vpcName").String()
			loadBalance.Routerid = value.Get("routerId").String()
			loadBalance.Networkid = value.Get("networkId").String()
			loadBalance.Subnetid = value.Get("subnetId").String()
			loadBalance.SubnetName = value.Get("subnetName").String()
			loadBalance.IPVersion = value.Get("ipVersion").String()
			loadBalance.PublicIp = value.Get("publicIp").String()

			sliceLoadBalance = append(sliceLoadBalance, &loadBalance)
			//fmt.Println("test", cloudPort)
			return false
		})
	}
	return sliceLoadBalance
}

func JsonParseTag(lbs []*LoadBalance) {
	var wg sync.WaitGroup
	wg.Add(len(lbs))

	for _, lb := range lbs {
		go func() {
			defer wg.Done()
			jsonBody := utils.Get(tagUrl+lb.Id, HeadersFun(lb.Pool))

			gjson.Get(jsonBody, "body.tags").ForEach(func(key, value gjson.Result) bool {
				switch value.Get("tagKey").String() {
				case "标签1":
					lb.Tag1 = value.Get("tagValue").String()
				case "标签X":
					lb.TagX = value.Get("tagValue").String()
				case "标签2":
					lb.Tag2 = value.Get("tagValue").String()
				case "订购人":
					lb.Orderer = value.Get("tagValue").String()
				}
				return true
			})
		}()
	}

	wg.Wait()
}
