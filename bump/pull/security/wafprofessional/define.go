package wafProfessional

import (
	"Catch/bump/internal/bootstrap"
	"fmt"
)

type WafProfessional struct {
	InstanceId       string
	InstanceName     string
	Bandwidth        string //默认100Mb， 购买一个扩展包加50Mb，100Mb + (扩展包数量 * 50Mb)
	ResourcePoolName string
	ExclusivePackage string //是否独享
}

var Url = "https://console.ecloud.10086.cn/api/web/esp/wafc/portal/instance/list?pageSize=10&pageNum=1&version=1"

var Headers = map[string]string{
	"Accept":          "application/json, text/plain, */*",
	"Accept-Language": "zh-CN",
	"Connection":      "keep-alive",
	"Content-Type":    "application/json",
	"Cookie":          fmt.Sprintf("CMECLOUDTOKEN=%s", bootstrap.Token),
	"Host":            "console.ecloud.10086.cn",
}
