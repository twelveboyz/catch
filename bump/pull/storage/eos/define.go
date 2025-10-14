package eos

import (
	"Catch/bump/internal/bootstrap"
	"fmt"
)

type ElasticObjectStorage struct {
	Id           string
	Name         string
	StorageClass string
	Size         string
	Pool         string
	Region       string
}

var Url = "https://console.ecloud.10086.cn/api/web/eosbeijing4/eos-console/customer/v3/eosBucket"

var Headers = map[string]string{
	"Accept":          "application/json, text/plain, */*",
	"Accept-Language": "zh-CN",
	"Connection":      "keep-alive",
	"Content-Type":    "application/json",
	"Cookie":          fmt.Sprintf("CMECLOUDTOKEN=%s", bootstrap.Token),
	"Host":            "console.ecloud.10086.cn",
	"pool-id":         "CIDC-RP-25", //这个资源池固定是苏州
}
