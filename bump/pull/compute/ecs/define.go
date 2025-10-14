package ecs

import (
	"Catch/bump/internal/bootstrap"
	"fmt"
)

type ECloudServer struct {
	Id             string
	Name           string
	VCpu           string
	VMemory        string
	VDisk          string
	BootVolumeType string
	ImageOsType    string
	ImageName      string
	SpecsName      string
	CreateTime     string
	NIC            []*NIC
	Pool           string
	Region         string
	Tag1           string
	TagX           string
	Tag2           string
	Orderer        string
}

type NIC struct {
	Id            string //portId
	Name          string //网卡名称
	MacAddress    string
	VpcId         string
	VpcName       string
	SubnetName    string
	NetworkId     string //子网ID
	ResourceId    string //云主机UUID
	Type          string
	Region        string
	FixedIp       []FixedIp
	SecurityGroup []SecurityGroup
}

type FixedIp struct {
	PortId     string
	IpVersion  string
	IpAddress  string
	SubnetCidr string //子网网段
	VpcId      string
	VpcName    string
	RouterId   string
	SubnetId   string //子网编码 分IPv4 和IPV6 编码
	SubnetName string //子网名称
	PortName   string //网卡名称
}

type SecurityGroup struct {
	PortId string
	Id     string
	Name   string
}

type ExcelECS struct {
	Id             string
	Name           string
	SpecsName      string
	VCpu           string
	VMemory        string
	ImageName      string
	BootVolumeType string
	VDisk          string
	NIC            []*ExcelNic
}

type ExcelNic struct {
	VnName        string
	VpcName       string
	SubnetName    string
	IpV4Address   string
	V4Mask        string
	IpV6Address   string
	V6Mask        string
	SecurityGroup string
}

var Url = `https://console.ecloud.10086.cn/api/web/routes/ec-console-web/v3/server/with/network?serverTypes=VM&productTypes=NORMAL&productTypes=AUTOSCALING&productTypes=VO&productTypes=CDN&productTypes=PAAS_MASTER&productTypes=PAAS_SLAVE&productTypes=EMR&productTypes=MSE&visible=true&optimized=true&page=1&size=200`

func HeadersFun(pool string) map[string]string {
	return map[string]string{
		"Accept":          "application/json, text/plain, */*",
		"Accept-Language": "zh-CN",
		"Connection":      "keep-alive",
		"Content-Type":    "application/json",
		"Cookie":          fmt.Sprintf("CMECLOUDTOKEN=%s", bootstrap.Token),
		"Host":            "console.ecloud.10086.cn",
		"pool-id":         pool,
	}
}

var Headers = map[string]string{
	"Accept":          "application/json, text/plain, */*",
	"Accept-Language": "zh-CN",
	"Connection":      "keep-alive",
	"Content-Type":    "application/json",
	"Cookie":          fmt.Sprintf("CMECLOUDTOKEN=%s", bootstrap.Token),
	"Host":            "console.ecloud.10086.cn",
	"pool-id":         bootstrap.PoolId,
}

func PortUrl(uuid string) string {
	return fmt.Sprintf("https://console.ecloud.10086.cn/api/web/routes/console-openstack-network/customer/v3/port?resourceId=%s", uuid)
}

func TagUrl(uuid string) string {
	return fmt.Sprintf("https://ecloud.10086.cn/api/web/tag/user/v4/tags/resource/tags?resourceId=%s&resourceType=CIDC-RT-VM", uuid)
}

func TagHeaders() map[string]string {
	return map[string]string{
		"Accept":          "application/json, text/plain, */*",
		"Accept-Language": "zh-CN",
		"Connection":      "keep-alive",
		"Content-Type":    "application/json",
		"Cookie":          fmt.Sprintf("CMECLOUDTOKEN=%s", bootstrap.Token),
		"Host":            "console.ecloud.10086.cn",
	}
}
