package eos

import (
	"Catch/bump/internal/utils"
	"fmt"
	"github.com/sirupsen/logrus"
)

func ExecuteALLAndPrint() {

	Body := utils.Get(Url, Headers)
	//fmt.Println(Body)

	InfoSet := JsonParseResourceInfo(Body)
	for _, info := range InfoSet {
		fmt.Println(info)
	}

	for _, info := range InfoSet {
		fmt.Println(info)
	}
}

func Execute() []*ElasticObjectStorage {
	logrus.Println("正在获取对象存储控制台资源······")
	Body := utils.Get(Url, Headers)
	//fmt.Println(Body)

	InfoSet := JsonParseResourceInfo(Body)
	return InfoSet

}
