package cis

import (
	"Catch/bump/internal/bootstrap"
	"fmt"
)

type ContainerImageService struct {
	Uid           string
	InstanceName  string
	Specification string
	BucketName    string
	Domain        string
	PeriodType    string
}

var Url = "https://console.ecloud.10086.cn/api/web/cis/ecis/instance/api/v1/instances"

var Headers = map[string]string{
	"Accept":          "application/json, text/plain, */*",
	"Accept-Language": "zh-CN",
	"Connection":      "keep-alive",
	"Content-Type":    "application/json",
	"Cookie":          fmt.Sprintf("CMECLOUDTOKEN=%s", bootstrap.Token),
	"Host":            "console.ecloud.10086.cn",
	"pool-id":         bootstrap.PoolId,
	"region":          bootstrap.PoolId,
}
