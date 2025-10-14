package ebs

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

	fmt.Println("共发起请求：", len(InfoSet))
}

func ExecuteGetEBSInfo(uuids []string, pools []string) []*ElasticBlockStorage {
	logrus.Println("正在获取云硬盘控制台资源······")
	var InfoSet []*ElasticBlockStorage
	for _, pool := range pools {
		Body := utils.Get(Url, HeadersFun(pool))

		is := JsonParseResourceByUUID(Body, uuids)
		JoinPoolInfo(is, pool)

		InfoSet = append(InfoSet, is...)
	}

	return InfoSet
}
