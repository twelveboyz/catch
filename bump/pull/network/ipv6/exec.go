package ipv6

import (
	"Catch/bump/internal/utils"
	"fmt"
	"github.com/sirupsen/logrus"
)

func ExecuteALL() {
	Body := utils.Get(Url, Headers)
	//fmt.Println(Body)

	InfoSet := ResourceInfoALL(Body)
	for _, info := range InfoSet {
		fmt.Println(info)
	}
}

func ExecuteGetIPv6Info(pools []string) []*IPv6 {
	logrus.Println("正在获取IPv6带宽控制台资源······")
	var InfoSet []*IPv6
	for _, pool := range pools {
		Body := utils.Get(Url, HeadersFun(pool))
		//log.Println(Body)

		is := ResourceInfoALL(Body)
		JoinPoolInfo(is, pool)

		InfoSet = append(InfoSet, is...)
	}

	return InfoSet
}
