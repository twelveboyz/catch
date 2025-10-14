package directconnect

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

func ExecuteGetResourceInfo() []DirectConnect {
	logrus.Println("正在获取云专线控制台资源······")
	Body := utils.Get(Url, HeadersFun())
	//fmt.Println(Body)

	InfoSet := ResourceInfoALL(Body)

	return InfoSet
}
