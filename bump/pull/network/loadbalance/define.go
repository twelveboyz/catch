package loadbalance

import (
	"Catch/bump/internal/bootstrap"
	"Catch/bump/internal/mapping"
	"fmt"
)

type LoadBalance struct {
	Id         string
	Name       string
	Bandwidth  int64
	Flavor     string
	Vpcname    string
	Routerid   string
	Networkid  string
	Subnetid   string
	IPVersion  string
	Privateip  string
	PublicIp   string
	SubnetName string
	Pool       string
	Region     string
	Tag1       string
	TagX       string
	Tag2       string
	Orderer    string
}

var Url = "https://ecloud.10086.cn/api/web/routes/lb-console-openstack-lb/acl/v3/loadBalancer/loadBalancers?page=1&pageSize=50&visible=true&orderBy=&showVirtualLoadBalancers=false&isMain=true&loadbalanceQuery=%7B%7D"

var Headers = map[string]string{
	"Accept":          "application/json, text/plain, */*",
	"Accept-Language": "zh-CN",
	"Connection":      "keep-alive",
	"Content-Type":    "application/json",
	"Cookie":          fmt.Sprintf("CMECLOUDTOKEN=%s", bootstrap.Token),
	"Host":            "console.ecloud.10086.cn",
	"pool-id":         "CIDC-RP-26",
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

var tagUrl = `https://ecloud.10086.cn/api/web/routes/lb-console-openstack-lb/acl/v3/loadBalancer/`

func Convert(s string) string {
	switch s {
	case "2":
		return mapping.SLB_FLAVOR_2
	case "3":
		return mapping.SLB_FLAVOR_3
	case "4":
		return mapping.SLB_FLAVOR_4
	case "5":
		return mapping.SLB_FLAVOR_5
	case "6":
		return mapping.SLB_FLAVOR_6
	case "21":
		return mapping.SLB_FLAVOR_21
	case "30":
		return mapping.SLB_FLAVOR_30
	default:
		return "NotFount=" + s
	}
}
