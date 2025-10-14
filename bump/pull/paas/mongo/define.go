package mongo

import (
	"Catch/bump/internal/bootstrap"
	"fmt"
)

type Mongo struct {
	Id                string
	Name              string
	AvailableAreaName string
	StorageTotal      int //shard存储空间
	Version           string
	Engine            string
	Type              string
	BackupStorage     int //备份存储

	//Mongos
	MongosCount         string
	MongosOsSpecType    string
	MongosOsSpecCpu     string
	MongosOsSpecMemory  string
	MongosOsSpecStorage string

	//Shard
	ShardCount          string
	MongodOsSpecType    string
	MongodOsSpecCpu     string
	MongodOsSpecMemory  string
	MongodOsSpecStorage string

	//ConfigServer
	ConfigServerCount       string
	ConfigServerSpecType    string
	ConfigServerSpecCpu     string
	ConfigServerSpecMemory  string
	ConfigServerSpecStorage string
}

var Url = "https://console.ecloud.10086.cn/api/web/dds/api/mdb/cluster?filter=%7B%22where%22:%7B%7D,%22size%22:10,%22page%22:1,%22sort%22:%5B%22createAt%20desc%22%5D%7D"

var Headers = map[string]string{
	"Accept":          "application/json, text/plain, */*",
	"Accept-Language": "zh-CN",
	"Connection":      "keep-alive",
	"Content-Type":    "application/json",
	"Cookie":          fmt.Sprintf("CMECLOUDTOKEN=%s", bootstrap.Token),
	"Host":            "console.ecloud.10086.cn",
	"pool-id":         bootstrap.PoolId,
}
