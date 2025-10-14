package eip

import (
	"Catch/bump/internal/bootstrap"
	"fmt"
)

type ElasticIP struct {
	Id               string
	ResourceName     string
	EipName          string //公网ip名称
	BindType         string //绑定类型 云主机、负载均衡、nat网关
	BindResourceId   string //绑定的uuid
	BindResourceName string //绑定的名称
	BandwidthSize    string //带宽
	BandwidthType    string //共享（shared） 独享（exclusive）
	Pool             string
	Tag1             string
	TagX             string
	Tag2             string
	Orderer          string
}

var Url = "https://console.ecloud.10086.cn/api/web/routes/console-openstack-network/acl/v3/floatingip/floatingips?page=1&size=20&frozen=&bound=&occupy=&preAllocatedForHY=true&showFipTypes=INTELLIGENT%2CUNDERLAY_FIP&sharedIp=false&bindType=&queryWord="

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

var tagUrl = "https://ecloud.10086.cn/api/web/tag/user/v4/tags/resource/tags?resourceType=CIDC-RT-IP&resourceId="

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
