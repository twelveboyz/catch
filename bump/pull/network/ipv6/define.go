package ipv6

import (
	"Catch/bump/internal/bootstrap"
	"fmt"
)

type IPv6 struct {
	NsQosPolicyId    string
	MixedId          string
	IpAddress        string
	BandWidthSize    string
	BandwidthType    string
	BindResourceId   string
	BindResourceName string
	Pool             string
}

var Url = "https://console.ecloud.10086.cn/api/web/routes/lb-console-openstack-lb/acl/v3/loadBalancer/ipAddressList/loadBalancer?page=1&size=200&ipVersion=V6"

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
