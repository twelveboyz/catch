package efs

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

func ExecuteGetEFSInfo(uuids, pools []string) []*ElasticFileStorage {
	logrus.Println("正在获取文件存储控制台资源······")
	var InfoSet []*ElasticFileStorage
	for _, pool := range pools {
		Body := utils.Get(Url, HeadersFun(pool))

		is := JsonParseResourceByUUIDs(Body, uuids)
		JoinPoolInfo(is, pool)
		JsonParseTag(is)

		InfoSet = append(InfoSet, is...)

	}
	return InfoSet
}
