package privatedns

import (
	"Catch/bump/internal/bootstrap"
	"fmt"
)

type PrivateDNS struct {
	InstanceId  string
	VDnsId      string
	VDnsName    string
	Ipv4        string
	Ipv6        string
	IpType      string
	PackageType string
}

var Url = "https://ecloud.10086.cn/api/web/privatedns/api/v1/privatedns/vdns/list?pageSize=20&page=1"

var Headers = map[string]string{
	"Accept":          "application/json, text/plain, */*",
	"Accept-Language": "zh-CN",
	"Connection":      "keep-alive",
	"Content-Type":    "application/json",
	"Cookie":          fmt.Sprintf("CMECLOUDTOKEN=%s", bootstrap.Token),
	"Host":            "console.ecloud.10086.cn",
}

var payload string = `{"regionId":"CIDC-RP-26"}`
