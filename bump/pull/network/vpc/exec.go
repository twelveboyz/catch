package vpc

import (
	"Catch/bump/internal/utils"
	"fmt"
)

func ExecuteALL() {
	virtualPrivateCloudBody := utils.Get(Url, Headers)
	//fmt.Println(cloudPortBody)

	virtualPrivateCloudInfoSet := ResourceInfoALL(virtualPrivateCloudBody)

	virtualPrivateCloudInfoSet = NetWorkAndSubNetInfo(virtualPrivateCloudInfoSet)

	//打印vpc-network-subnet
	for _, virtualPrivateCloudInfo := range virtualPrivateCloudInfoSet {
		//输出VPC内容
		fmt.Println(virtualPrivateCloudInfo)
		fmt.Println("子网数:", len(virtualPrivateCloudInfo.NetWork))

		for _, net := range virtualPrivateCloudInfo.NetWork {
			//输出Network内容
			fmt.Printf("network=(vpc=%v,id=%v,name=%v,region=%v,type=%v);\n", virtualPrivateCloudInfo.Name, net.Id, net.Name, net.Region, net.NetworkTypeEnum)

			for _, sub := range net.SubNet {
				//输出SubNet内容
				fmt.Printf("subnet=(network=%v, id:%v, name:%v, region:%v, cidr:%v, gatewayIp:%v, ipVersion:%v);\n", net.Name, sub.Id, sub.Name, sub.Region, sub.Cidr, sub.GatewayIp, sub.IpVersion)
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
