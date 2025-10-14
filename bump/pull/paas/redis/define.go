package redis

import (
	"Catch/bump/internal/bootstrap"
	"fmt"
	"github.com/tidwall/gjson"
)

type Redis struct {
	Id                string
	Name              string
	Version           string
	Type              string
	Shards            string
	Memory            string
	ChargeType        string
	RedisRegion       []gjson.Result
	RedisRegionFormat []gjson.Result
	SlaveCount        string
}

var Url = "https://console.ecloud.10086.cn/api/web/redis/business/customer/v3/redis?page=1&size=100&name=&id=&version=&redisRegion=&tagIds=&checkGFS="

var Headers = map[string]string{
	"Accept":          "application/json, text/plain, */*",
	"Accept-Language": "zh-CN",
	"Connection":      "keep-alive",
	"Content-Type":    "application/json",
	"Cookie":          fmt.Sprintf("CMECLOUDTOKEN=%s", bootstrap.Token),
	"Host":            "console.ecloud.10086.cn",
	"pool-id":         bootstrap.PoolId,
}
