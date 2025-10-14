package rocket

import (
	"Catch/bump/internal/bootstrap"
	"fmt"
)

type Rocket struct {
	Id                  string
	ResourceId          string
	AvailZone           string
	Type                string
	Tps                 string
	MaxTopicSize        string
	Storage             string
	IsExclusiveResource string
	PeriodType          string
}

var Url = "https://console.ecloud.10086.cn/api/web/mq/mq/resources"

var Headers = map[string]string{
	"Accept":          "application/json, text/plain, */*",
	"Accept-Language": "zh-CN",
	"Connection":      "keep-alive",
	"Content-Type":    "application/json",
	"Cookie":          fmt.Sprintf("CMECLOUDTOKEN=%s", bootstrap.Token),
	"Host":            "console.ecloud.10086.cn",
	"pool-id":         bootstrap.PoolId,
}
