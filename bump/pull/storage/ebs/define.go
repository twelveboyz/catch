package ebs

import (
	"Catch/bump/internal/bootstrap"
	"fmt"
)

type ElasticBlockStorage struct {
	Id         string
	Name       string
	Size       string
	Type       string
	IsShare    string
	ServerId   string
	ServerName string
	Pool       string
	Region     string
	Tag1       string
	TagX       string
	Tag2       string
	Orderer    string
}

var Url = "https://console.ecloud.10086.cn/api/web/routes/console-openstack-volume/acl/v3/volume/volume/list/with/server?page=1&size=200"

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
