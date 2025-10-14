package mysql

import (
	"Catch/bump/internal/bootstrap"
	"fmt"
)

type Mysql struct {
	Id           string
	Name         string
	Engine       string
	Version      string
	Az           string
	Measuretype  string
	Instancespec string
	Volumesize   string
}

var Url = `https://console.ecloud.10086.cn/api/web/mysql/apps/v1/clusters?engine=mysql&limit=10&offset=1&filter=%7B%22az%22:%22%22,%22arch%22:%22%22,%22version%22:%22%22,%22status%22:%22%22,%22network_type%22:%22%22,%22non_standard%22:%22%22%7D&tagIds=&Rs-Source=console`

var Headers = map[string]string{
	"Accept":          "application/json, text/plain, */*",
	"Accept-Language": "zh-CN",
	"Connection":      "keep-alive",
	"Content-Type":    "application/json",
	"Cookie":          fmt.Sprintf("CMECLOUDTOKEN=%s", bootstrap.Token),
	"Host":            "console.ecloud.10086.cn",
	"pool-id":         "CIDC-RP-26",
}
