package hbase

import (
	"Catch/bump/internal/bootstrap"
	"fmt"
)

type HBase struct {
	InstanceId   string
	InstanceName string
	Region       string
	Version      string
	ProductSd    int64 //表格存储版独有字段

	//Master节点
	MasterConfigVCore  int64
	MasterConfigMemory int64
	MasterConfigCount  int64

	//Core节点
	CoreConfigVCore  int64
	CoreConfigMemory int64
	CoreConfigCount  int64

	//多元索引节点
	EsConfigVCore         int64
	EsConfigMemory        int64
	EsConfigCount         int64
	StorageConfigType     string
	StorageConfigCapacity int64
}

var Url = fmt.Sprintf("https://console.ecloud.10086.cn/api/web/cloudhbase/v1/server/instances/pool/%s?limit=-1&offset=-1&keyword=", bootstrap.PoolId)

var Headers = map[string]string{
	"Accept":          "application/json, text/plain, */*",
	"Accept-Language": "zh-CN",
	"Connection":      "keep-alive",
	"Content-Type":    "application/json",
	"Cookie":          fmt.Sprintf("CMECLOUDTOKEN=%s", bootstrap.Token),
	"Host":            "console.ecloud.10086.cn",
	"pool-id":         bootstrap.PoolId,
}
