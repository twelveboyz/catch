package ironic

import (
	"Catch/bump/internal/bootstrap"
	"fmt"
)

type Ironic struct {
	Id             string
	Name           string
	VCpu           string
	VMemory        string
	VDisk          string
	BootVolumeType string
	ImageOsType    string
	ImageName      string
	SpecsName      string
	PrivateIp      string
	VpcName        string
	SubnetName     string
	PortName       string
	NIC            []*NIC
	CreateTime     string
	Pool           string
	Region         string
	Tag1           string
	TagX           string
	Tag2           string
	Orderer        string
}

type NIC struct {
	Id            string
	Name          string
	Type          string
	Region        string
	MacAddress    string
	ResourceId    string
	ResourceName  string
	VpcId         string
	VpcName       string
	RouterId      string
	NetworkId     string
	SubnetName    string
	FixedIp       []FixedIp
	SecurityGroup []SecurityGroup
}

type FixedIp struct {
	PortId       string
	PortName     string
	IpVersion    string
	IpAddress    string
	SubnetCidr   string
	SubnetId     string
	SubnetName   string
	ResourceId   string
	ResourceName string
	VpcId        string
	VpcName      string
	RouterId     string
}

type SecurityGroup struct {
	PortId string
	Id     string
	Name   string
}

var Url = "https://console.ecloud.10086.cn/api/web/routes/bcec-console-backend-bms/acl/v3/BareMetal/with/network?page=1&size=200&serverTypes=IRONIC&serverTypes=EBM&productTypes=NORMAL&productTypes=PAAS_MASTER&productTypes=PAAS_SLAVE&createSource=WEB_DEFAULT&visible=true"

func PortUrl(uuid string) string {
	return fmt.Sprintf("https://console.ecloud.10086.cn/api/web/routes/console-openstack-network/customer/v3/port?resourceId=%s", uuid)

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

func TagUrl(uuid string) string {
	return fmt.Sprintf("https://ecloud.10086.cn/api/web/tag/user/v4/tags/resource/tags?resourceId=%s&resourceType=CIDC-RT-IRONIC", uuid)
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
