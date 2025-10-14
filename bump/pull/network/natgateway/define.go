package natgateway

import (
	"Catch/bump/internal/bootstrap"
	"fmt"
)

type NATGateway struct {
	Id         string
	Name       string
	Scale      string
	VpcId      string
	VpcName    string
	RouterId   string
	Bandwidth  string
	PeriodType string
	Pool       string
	Region     string
	Tag1       string
	TagX       string
	Tag2       string
	Orderer    string
}

var Url = "https://ecloud.10086.cn/api/web/routes/nat-console/customer/v3/natGateway?visible=true&page=1&size=10"

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

var tagUrl = "https://ecloud.10086.cn/api/web/routes/nat-order/acl/v3/mop?method=SYAN_UNHQ_queryOrderInfoExt&format=json&status=1"
