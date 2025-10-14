package sharedtrafficpackage

import (
	"Catch/bump/internal/bootstrap"
	"fmt"
)

type SharedTrafficPackage struct {
	Id   string
	Uuid string
	Name string
	Spec string
	Size int64 //TB
}

//var Url = "https://console.ecloud.10086.cn/api/web/routes/console-shared-network-traffic-packet/customer/v3/SharedNetworkTrafficPacket/global?page=1&size=20"

var Url = "https://console.ecloud.10086.cn/api/web/routes/console-shared-network-traffic-packet/customer/v3/SharedNetworkTrafficPacket?page=1&size=10000"

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
