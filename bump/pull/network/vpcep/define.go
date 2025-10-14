package vpcep

import (
	"Catch/bump/internal/bootstrap"
	"fmt"
)

type VpcEndPoint struct {
	VpcEpId          string
	SpecialLineId    string
	VpcEpName        string
	VpcEpServiceName string
	ProductName      string
	Region           string
	RegionName       string
}

var Url = "https://ecloud.10086.cn/api/web/vpcep-service/vpcep/v2/ep/list?page=1&size=20"

var Headers = map[string]string{
	"Accept":          "application/json, text/plain, */*",
	"Accept-Language": "zh-CN",
	"Connection":      "keep-alive",
	"Content-Type":    "application/json",
	"Cookie":          fmt.Sprintf("CMECLOUDTOKEN=%s", bootstrap.Token),
	"Host":            "console.ecloud.10086.cn",
	"poolid":          bootstrap.PoolId,
}
