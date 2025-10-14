package natgateway

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

func ExecuteGetNatInfo(uuids []string, pools []string) []*NATGateway {
	logrus.Println("正在获取NAT网关控制台资源······")
	var InfoSet []*NATGateway
	for _, pool := range pools {
		Body := utils.Get(Url, HeadersFun(pool))
		//log.Println(Body)

		is := ResourceMatchByUUID(Body, uuids)
		JoinPoolInfo(is, pool)
		JsonParseTag(is)

		InfoSet = append(InfoSet, is...)
	}

	return InfoSet
}
