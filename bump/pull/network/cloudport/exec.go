package cloudport

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

func ExecuteGetCPInfo() []CloudPort {
	logrus.Println("正在获取云端口控制台资源······")
	Body := utils.Get(Url, HeadersFun())

	InfoSet := ResourceInfoALL(Body)

	return InfoSet
}

func Execute() {
	/*--------------------------------------云端口--------------------------------------*/
	//获取response返回的数据
	cloudPortBody := utils.Get(Url, Headers)
	fmt.Println(cloudPortBody)

	//模拟所需数据
	var testData = []string{"cp_gz3_vpcep_nj_01", "cp_gz3_aisc_d_01"}

	cloudPortInfoSet := ResourceInfo(cloudPortBody, testData)
	for _, cloudPortInfo := range cloudPortInfoSet {
		fmt.Println(cloudPortInfo)
	}
}
