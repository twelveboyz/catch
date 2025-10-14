package eip

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

func ExecuteGetEIPInfo(uuids []string, pools []string) []*ElasticIP {
	logrus.Println("正在获取弹性公网IP控制台资源······")
	var InfoSet []*ElasticIP
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
