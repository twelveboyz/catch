package network_comparison

import (
	"Catch/bump/excel/cmdb/get"
	"Catch/bump/pull/network/sharedtrafficpackage"
	"fmt"
	"github.com/sirupsen/logrus"
)

func StpComparison(excel get.SharedTrafficPackageResource, consoles []sharedtrafficpackage.SharedTrafficPackage) {
	logrus.Debugf("------------------------------ Row:%s ------------------------------", excel.Row)
	for _, console := range consoles {
		fmt.Println(console)
	}
}
