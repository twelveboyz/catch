package vpc

import (
	"Catch/bump/internal/mapping"
	"Catch/bump/internal/utils"
	"sync"

	"github.com/tidwall/gjson"
)

func ResourceInfoALL(body string) []*VirtualPrivateCloud {

	var sliceVirtualPrivateCloud []*VirtualPrivateCloud
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.content|@pretty")

	//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		var virtualPrivateCloud VirtualPrivateCloud
		virtualPrivateCloud.Id = value.Get("id").String()
		virtualPrivateCloud.Name = value.Get("name").String()
		virtualPrivateCloud.InstanceName = value.Get("instanceName").String()
		virtualPrivateCloud.Region = value.Get("region").String()
		virtualPrivateCloud.RouterId = value.Get("routerId").String()
		virtualPrivateCloud.Scale = value.Get("scale").String()
		virtualPrivateCloud.VpcExtraSpecification = value.Get("vpcExtraSpecification").String()

		sliceVirtualPrivateCloud = append(sliceVirtualPrivateCloud, &virtualPrivateCloud)
		//fmt.Println("test", cloudPort)
		return true
	})

	return sliceVirtualPrivateCloud
}

// NetWorkAndSubNetInfo 传参vpcs传入所有获取到的vpc内容，在通过routeid发起第二次http请求获取vpc下的所有子网详情
func NetWorkAndSubNetInfo(vpcs []*VirtualPrivateCloud) []*VirtualPrivateCloud {
	var wg sync.WaitGroup
	wg.Add(len(vpcs))

	for _, vpc := range vpcs {
		go func() {
			defer wg.Done()
			url := SubNetUrl(vpc.RouterId)
			response := utils.Get(url, Headers)

			//获取Network内容
			jsonBody := gjson.Get(response, "body.content")
			jsonBody.ForEach(func(key, value gjson.Result) bool {
				var netWork NetWork
				netWork.Id = value.Get("id").String()
				netWork.Name = value.Get("name").String()
				netWork.Region = value.Get("region").String()
				netWork.VpcId = value.Get("vpcId").String()
				netWork.RouterId = value.Get("routerId").String()
				netWork.NetworkTypeEnum = value.Get("networkTypeEnum").String()

				// 将NetWork信息加入到vpc中
				vpc.NetWork = append(vpc.NetWork, &netWork)

				// 循环subnets获取子网数据
				content := value.Get("subnets")
				content.ForEach(func(key, value gjson.Result) bool {
					var subNet SubNet
					subNet.Id = value.Get("id").String()
					subNet.Name = value.Get("name").String()
					subNet.Region = mapping.AZConvert(value.Get("region").String())
					subNet.NetworkId = value.Get("networkId").String()
					subNet.IpVersion = value.Get("ipVersion").String()
					subNet.Cidr = value.Get("cidr").String()
					subNet.GatewayIp = value.Get("gatewayIp").String()

					// 将Subnet内容加入到Network中
					netWork.SubNet = append(netWork.SubNet, &subNet)
					return true
				})
				return true
			})
		}()

	}

	wg.Wait()
	return vpcs
}
