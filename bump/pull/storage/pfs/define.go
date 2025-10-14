package pfs

import (
	"Catch/bump/internal/bootstrap"
	"fmt"
)

type ParallelFileStorage struct {
	ShareId        string
	Name           string
	Size           string
	PoolId         string
	ShareType      string
	ExportLocation string
	Region         string
	Pool           string
	Tag1           string
	TagX           string
	Tag2           string
	Orderer        string
}

type ExportLocation struct {
	IPv4      string `json:"ipv4_address"`
	IPv6      string `json:"ipv6_address"`
	SharePath string `json:"share_path"`
}

var Url = "https://console.ecloud.10086.cn/api/web/routes/console-openstack-share/customer/v3/pfs?page=1&size=10"

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

var payloadPfs = `{"page":1,"size":1000,"poolIds":[],"resourceTypes":["CIDC-RT-PFS"]}`

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
