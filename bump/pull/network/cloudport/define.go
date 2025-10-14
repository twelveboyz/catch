package cloudport

import (
	"Catch/bump/internal/bootstrap"
	"fmt"
)

type CloudPort struct {
	Id                      string //uuid
	ParticularLineName      string //name
	PortSpecification       string //规格
	ParticularLineBandwidth string //带宽
	BandwidthUnit           string //带宽单位
	IpType                  string //ip类型
	VpcId                   string //vpcid
	VpcName                 string //vpcname

	AccessNodeRegion  string //接入节点地域,city+borough
	AccessNodeAddress string //接入节点详细地址,address
	AccessNodeContact string //接入节点联系人,contactName
	AccountManager    string //客户经理,managerName

	VGWAndAggDockingIPMaster string //vgw与汇聚设备对接IP(主）,vgwConvergeDeviceIp
	VGWAndAggDockingIPBackup string //vgw与汇聚设备对接IP(备）,vgwConvergeDeviceIpBack

	AggAndVGWDockingIPMaster string //汇聚设备与vgw对接IP(主）,convergeDeviceVgwIp
	AggAndVGWDockingIPBackup string //汇聚设备与vgw对接IP(备）,convergeDeviceVgwIpBack

	AggAndVGWDockingVLANMaster string //汇聚设备与vgw对接VLAN（主）,vgwConvergeDeviceVlan
	AggAndVGWDockingVLANBackup string //汇聚设备与vgw对接VLAN（备）,vgwConvergeDeviceVlanBack

	AggAndCustomerDockingLogicalPort  string //汇聚设备与客户设备对接逻辑端口,logicalPort
	AggAndCustomerDockingPhysicalPort string //汇聚设备与客户设备对接物理端口,physicalPort

	Device1  string //设备1位置,deviceLocationInfo1
	Device2  string //设备2位置,deviceLocationInfo2
	PoolName string //资源池
	Node     string //节点
}

var Url = "https://console.ecloud.10086.cn/api/web/cloud-port/server/v1/particularLine?page=1&size=100&applyStatus=&resourceType=safe&localFlag=0&&particularLineName="

func HeadersFun() map[string]string {
	return map[string]string{
		"Accept":          "application/json, text/plain, */*",
		"Accept-Language": "zh-CN",
		"Connection":      "keep-alive",
		"Content-Type":    "application/json",
		"Cookie":          fmt.Sprintf("CMECLOUDTOKEN=%s", bootstrap.Token),
		"Host":            "console.ecloud.10086.cn",
		"zone":            "SLINE",
	}
}

var Headers = map[string]string{
	"Accept":          "application/json, text/plain, */*",
	"Accept-Language": "zh-CN",
	"Connection":      "keep-alive",
	"Content-Type":    "application/json",
	"Cookie":          fmt.Sprintf("CMECLOUDTOKEN=%s", bootstrap.Token),
	"Host":            "console.ecloud.10086.cn",
	"zone":            "SLINE",
}
