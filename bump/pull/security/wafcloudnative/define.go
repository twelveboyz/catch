package wafCloudNative

import (
	"Catch/bump/internal/bootstrap"
	"fmt"
)

type WafCloudNative struct {
	Id          string
	Description string
	IpNum       string
	Bandwidth   string
	PeriodType  string
}

var Url = "https://console.ecloud.10086.cn/api/web/esp/waf?page=0&size=10&isDelete=0"

var Headers = map[string]string{
	"Accept":          "application/json, text/plain, */*",
	"Accept-Language": "zh-CN",
	"Connection":      "keep-alive",
	"Content-Type":    "application/json",
	"Cookie":          fmt.Sprintf("CMECLOUDTOKEN=%s", bootstrap.Token),
	"Host":            "console.ecloud.10086.cn",
	"consolepoolid":   bootstrap.PoolId,
	"pool-id":         "CIDC-CORE-00",
}
