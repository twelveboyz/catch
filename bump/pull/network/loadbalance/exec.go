package loadbalance

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

func ExecuteGetLoadBalanceInfo(uuids []string, pools []string) []*LoadBalance {
	logrus.Println("正在获取负载均衡控制台资源······")
	var InfoSet []*LoadBalance
	for _, pool := range pools {
		Body := utils.Get(Url, HeadersFun(pool))

		is := ResourceMatchByUUID(Body, uuids)
		JoinPoolInfo(is, pool)
		JsonParseTag(is)

		InfoSet = append(InfoSet, is...)
	}

	return InfoSet
}
