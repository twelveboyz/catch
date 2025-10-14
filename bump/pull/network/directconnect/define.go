package directconnect

import (
	"Catch/bump/internal/bootstrap"
	"fmt"
)

type DirectConnect struct {
	Id                   string
	NetSideId            string
	SpecialLineName      string
	PoolId               string
	IpType               string
	SpecialLineBandwidth string
	VpcId                string
	VpcName              string
	VpsSubnets           []string
	UserSubnets          []string
	Province             string
	City                 string
	Borough              string
	Address              string
	ContactName          string
	ManagerName          string
}

var Url = "https://console.ecloud.10086.cn/api/web/ceno/v2/service/cloudExpresses/specialLines?page=1&size=100&cloudConnType=SAME_ACCOUNT_CONN&pending=0&&strategy=9&bargainingMode=1"

var Headers = map[string]string{
	"Accept":          "application/json, text/plain, */*",
	"Accept-Language": "zh-CN",
	"Connection":      "keep-alive",
	"Content-Type":    "application/json",
	"Cookie":          fmt.Sprintf("CMECLOUDTOKEN=%s", bootstrap.Token),
	"Host":            "console.ecloud.10086.cn",
	"zone":            "SLINE",
}

func HeadersFun() map[string]string {
	return map[string]string{
		"Accept":          "application/json, text/plain, */*",
		"Accept-Language": "zh-CN",
		"Connection":      "keep-alive",
		"Content-Type":    "application/json",
		"Cookie":          fmt.Sprintf("CMECLOUDTOKEN=%s", bootstrap.Token),
		"Host":            "console.ecloud.10086.cn",
		"zone":            "SLINE",
	}
}
