package kafka

import (
	"Catch/bump/internal/bootstrap"
	"fmt"
)

type Kafka struct {
	Id                 string
	Name               string
	Region             string
	InstanceVersion    string
	SpecType           string
	ConnectAddress     string
	Ipv6ConnectAddress string
	Cpu                string
	Memory             string
	VolumeSize         int64
	StorageCapacity    int64
	MessageSizeMax     string
	Brokers            string
}

var Url = fmt.Sprintf("https://console.ecloud.10086.cn/api/web/ekafka/ekafka/v1/tenants/tenantsMockId/clusters?includeEkafka=true&status=!RELEASED&regionId=%s&limit=50&offset=1&name=&id=", bootstrap.PoolId)

var Headers = map[string]string{
	"Accept":          "application/json, text/plain, */*",
	"Accept-Language": "zh-CN",
	"Connection":      "keep-alive",
	"Content-Type":    "application/json",
	"Cookie":          fmt.Sprintf("CMECLOUDTOKEN=%s", bootstrap.Token),
	"Host":            "console.ecloud.10086.cn",
	"pool-id":         bootstrap.PoolId,
}
