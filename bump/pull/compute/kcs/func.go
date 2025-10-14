package kcs

import "fmt"

func GetNodeInfoUrl(clusterId string) string {
	url := fmt.Sprintf("https://console.ecloud.10086.cn/api/web/kcs/v2/clusters/%s/nodes?page=1&pageSize=10&role=&clusterType=notRegistered&labelInfo=&name=", clusterId)
	return url
}
