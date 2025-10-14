package pfs

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

func ExecuteGetPFSInfo(uuids []string, pools []string) []*ParallelFileStorage {
	logrus.Println("正在获取并行文件存储控制台资源······")
	var infoSet []*ParallelFileStorage

	for _, pool := range pools {
		body := utils.Get(Url, HeadersFun(pool))
		//fmt.Println(Body)

		is := JsonParseResourceByUUID(body, uuids)
		JoinPoolInfo(is, pool)
		JsonParseTag(is)

		infoSet = append(infoSet, is...)
	}

	return infoSet
}
