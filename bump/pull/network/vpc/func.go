package vpc

import "fmt"

func SubNetUrl(routeId string) string {
	url := fmt.Sprintf("https://console.ecloud.10086.cn/api/web/routes/console-openstack-network/customer/v3/network/NetworkResps?page=1&size=1000&routerId=%s&networkConsole=true", routeId)
	return url
}
