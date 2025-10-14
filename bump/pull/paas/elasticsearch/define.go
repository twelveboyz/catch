package elasticsearch

import (
	"Catch/bump/internal/bootstrap"
	"fmt"
)

type ElasticSearch struct {
	Id           string
	ZoneZoneName string
	AppType      string
	Alias        string
	Servers      []Server
}

type Server struct {
	TypeCode        string
	ArchId          string
	ArchName        string
	ArchMode        string
	ArchShardCnt    string
	ArchUnitCnt     string
	ScaleName       string
	ScaleCpuCnt     string
	ScaleMemSize    string
	DiskTypeDisplay string
	DataSize        string
}

var Url = "https://ecloud.10086.cn/api/web/elasticsearch/infini-scale-api/customer/serv_groups/elasticsearch/list"

var Payload = `{"tags":[],"searchContent":"","healthStatus":"","runningState":"","appType":"","meta":{"page":1,"pageSize":100}}`

var Headers = map[string]string{
	"Accept":          "application/json, text/plain, */*",
	"Accept-Language": "zh-CN",
	"Connection":      "keep-alive",
	"Content-Type":    "application/json",
	"Cookie":          fmt.Sprintf("CMECLOUDTOKEN=%s", bootstrap.Token),
	"Host":            "console.ecloud.10086.cn",
	"pool-id":         bootstrap.PoolId,
}
