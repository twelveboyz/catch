package vpc

import (
	"Catch/bump/internal/bootstrap"
	"fmt"
)

type VirtualPrivateCloud struct {
	Id                    string
	Name                  string
	InstanceName          string
	Region                string
	RouterId              string
	Scale                 string
	VpcExtraSpecification string
	NetWork               []*NetWork
}

type NetWork struct {
	Id              string
	Name            string
	Region          string
	VpcId           string
	RouterId        string
	NetworkTypeEnum string
	SubNet          []*SubNet
}

type SubNet struct {
	Id        string
	Name      string
	NetworkId string
	Cidr      string
	IpVersion string
	GatewayIp string
	Region    string
}

var Url = "https://console.ecloud.10086.cn/api/web/routes/console-openstack-network/customer/v3/vpc?visible=true&page=1&size=20"

var Headers = map[string]string{
	"Accept":          "application/json, text/plain, */*",
	"Accept-Language": "zh-CN",
	"Connection":      "keep-alive",
	"Content-Type":    "application/json",
	"Cookie":          fmt.Sprintf("CMECLOUDTOKEN=%s", bootstrap.Token),
	"Host":            "console.ecloud.10086.cn",
	"pool-id":         bootstrap.PoolId,
}
