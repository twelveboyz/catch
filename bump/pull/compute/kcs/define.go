package kcs

import (
	"Catch/bump/internal/bootstrap"
	"fmt"
)

type KubernetesCloudServer struct {
	ClusterId             string
	ClusterName           string
	Pool                  string
	Region                string
	NodeQuota             string
	KubeVersion           string
	ClusterRuntime        string
	ClusterRuntimeVersion string
	PeriodType            string
	VpcName               string
	Os                    string
	Nodes                 []Nodes
}

type Nodes struct {
	ServerId   string
	NodeId     string
	Name       string
	Cpu        string
	Memory     string
	OsImage    string
	Os         string
	ServerType string
	VmType     string
	SpecsName  string
	InstanceId string
}

var Url = "https://console.ecloud.10086.cn/api/web/kcs/v2/clusters?page=0"

var Headers = map[string]string{
	"Accept":          "application/json, text/plain, */*",
	"Accept-Language": "zh-CN",
	"Connection":      "keep-alive",
	"Content-Type":    "application/json",
	"Cookie":          fmt.Sprintf("CMECLOUDTOKEN=%s", bootstrap.Token),
	"Host":            "console.ecloud.10086.cn",
	"pool-id":         bootstrap.PoolId,
}
