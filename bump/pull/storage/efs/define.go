package efs

import (
	"Catch/bump/internal/bootstrap"
	"fmt"
)

type ElasticFileStorage struct {
	ShareId   string
	Name      string
	Size      string
	PoolId    string
	ShareType string
	Mount     string
	Region    string
	Pool      string
	Tag1      string
	TagX      string
	Tag2      string
	Orderer   string
}

var Url = "https://console.ecloud.10086.cn/api/web/routes/console-openstack-share/customer/v3/nas?page=1&size=500&shareStatus=&shareOPType=&protocol=&autoResize=true"

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

var tagUrl = `https://console.ecloud.10086.cn/api/web/routes/api/web/rms/api/list`

var payloadEfs = `{"page":1,"size":1000,"poolIds":[],"resourceTypes":["CIDC-RT-NAS"]}`

func HeadersNoPool() map[string]string {
	return map[string]string{
		"Accept":          "application/json, text/plain, */*",
		"Accept-Language": "zh-CN",
		"Connection":      "keep-alive",
		"Content-Type":    "application/json",
		"Cookie":          fmt.Sprintf("CMECLOUDTOKEN=%s", bootstrap.Token),
		"Host":            "console.ecloud.10086.cn",
	}
}
