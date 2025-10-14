package sharedtrafficpackage

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

func ExecuteGetSTPInfo(uuids []string, pools []string) []SharedTrafficPackage {
	var InfoSet []SharedTrafficPackage
	for _, p := range pools {
		Body := utils.Get(Url, HeadersFun(p))
		//log.Println(Body)

		is := ResourceMatchByUUID(Body, uuids)

		InfoSet = append(InfoSet, is...)
	}

	logrus.Println("共享流量包获取到的数据：", len(InfoSet))

	return InfoSet
}
