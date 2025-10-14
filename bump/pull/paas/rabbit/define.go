package rabbit

import (
	"Catch/bump/internal/bootstrap"
	"fmt"
)

type Rabbit struct {
	InstanceId      string
	ClusterName     string
	Region          string
	VCore           string
	Memory          string
	InstanceType    string
	Version         string
	StorageCapacity string
}

var Url = fmt.Sprintf("https://console.ecloud.10086.cn/api/web/mq/cloudamqp/server/instances/list/pool/%s?pageSize=10&pageNum=1&clusterName=&instanceId=", bootstrap.PoolId)

var Headers = map[string]string{
	"Accept":          "application/json, text/plain, */*",
	"Accept-Language": "zh-CN",
	"Connection":      "keep-alive",
	"Content-Type":    "application/json",
	"Cookie":          fmt.Sprintf("CMECLOUDTOKEN=%s", bootstrap.Token),
	"Host":            "console.ecloud.10086.cn",
	"pool-id":         bootstrap.PoolId,
}
