package cloudport

import (
	"github.com/tidwall/gjson"
)

func ResourceInfo(body string, resourceName []string) []CloudPort {

	var sliceCloudPort []CloudPort
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.content|@pretty")

	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		//fmt.Println(key) //打印循环次数

		//循环过滤resourceName对应的数据
		for _, v := range resourceName {
			//匹配资源名称，过滤想要的资源加入到slice
			if v == value.Get("particularlineName").String() {
				var cloudPort CloudPort
				cloudPort.Id = value.Get("id").String()
				cloudPort.ParticularLineName = value.Get("particularlineName").String()
				cloudPort.PortSpecification = value.Get("portSpecification").String()
				cloudPort.ParticularLineBandwidth = value.Get("particularlineBandwidth").String()
				cloudPort.IpType = value.Get("iptype").String()
				cloudPort.VpcId = value.Get("vpcId").String()

				cloudPort.AccessNodeRegion = value.Get("city").String() + value.Get("borough").String()
				cloudPort.AccessNodeAddress = value.Get("address").String()
				cloudPort.AccessNodeContact = value.Get("contactName").String()
				cloudPort.AccountManager = value.Get("managerName").String()

				cloudPort.VGWAndAggDockingIPMaster = value.Get("vgwConvergeDeviceIp").String()
				cloudPort.VGWAndAggDockingIPBackup = value.Get("vgwConvergeDeviceIpBack").String()

				cloudPort.AggAndVGWDockingIPMaster = value.Get("convergeDeviceVgwIp").String()
				cloudPort.AggAndVGWDockingIPBackup = value.Get("convergeDeviceVgwIpBack").String()

				cloudPort.AggAndVGWDockingVLANMaster = value.Get("vgwConvergeDeviceVlan").String()
				cloudPort.AggAndVGWDockingVLANBackup = value.Get("vgwConvergeDeviceVlanBack").String()

				cloudPort.AggAndCustomerDockingLogicalPort = value.Get("logicalPort").String()
				cloudPort.AggAndCustomerDockingPhysicalPort = value.Get("physicalPort").String()

				cloudPort.Device1 = value.Get("deviceLocationInfo1").String()
				cloudPort.Device2 = value.Get("deviceLocationInfo2").String()
				cloudPort.PoolName = value.Get("poolName").String()

				sliceCloudPort = append(sliceCloudPort, cloudPort)
			}
		}

		if len(sliceCloudPort) == len(resourceName) {
			return false
		}

		//fmt.Println("test", cloudPort)
		return true
	})

	return sliceCloudPort
}

func ResourceInfoALL(body string) []CloudPort {

	var sliceCloudPort []CloudPort
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.content|@pretty")

	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		//fmt.Println(key) //打印循环次数

		var cloudPort CloudPort
		cloudPort.Id = value.Get("id").String()
		cloudPort.ParticularLineName = value.Get("particularlineName").String()
		cloudPort.PortSpecification = value.Get("portSpecification").String()
		cloudPort.ParticularLineBandwidth = value.Get("particularlineBandwidth").String()
		cloudPort.IpType = value.Get("iptype").String()
		cloudPort.VpcId = value.Get("vpcId").String()

		cloudPort.AccessNodeRegion = value.Get("province").String() + value.Get("city").String() + value.Get("borough").String()
		cloudPort.AccessNodeAddress = value.Get("address").String()
		cloudPort.AccessNodeContact = value.Get("contactName").String()
		cloudPort.AccountManager = value.Get("managerName").String()

		cloudPort.VGWAndAggDockingIPMaster = value.Get("vgwConvergeDeviceIp").String()
		cloudPort.VGWAndAggDockingIPBackup = value.Get("vgwConvergeDeviceIpBack").String()

		cloudPort.AggAndVGWDockingIPMaster = value.Get("convergeDeviceVgwIp").String()
		cloudPort.AggAndVGWDockingIPBackup = value.Get("convergeDeviceVgwIpBack").String()

		cloudPort.AggAndVGWDockingVLANMaster = value.Get("vgwConvergeDeviceVlan").String()
		cloudPort.AggAndVGWDockingVLANBackup = value.Get("vgwConvergeDeviceVlanBack").String()

		cloudPort.AggAndCustomerDockingLogicalPort = value.Get("logicalPort").String()
		cloudPort.AggAndCustomerDockingPhysicalPort = value.Get("physicalPort").String()

		cloudPort.Device1 = value.Get("deviceLocationInfo1").String()
		cloudPort.Device2 = value.Get("deviceLocationInfo2").String()
		cloudPort.PoolName = value.Get("poolName").String()

		sliceCloudPort = append(sliceCloudPort, cloudPort)

		//fmt.Println("test", cloudPort)
		return true
	})

	return sliceCloudPort
}
