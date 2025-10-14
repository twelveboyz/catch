package storage_comparison

import (
	iinternal "Catch/bump/comparison/internal/utils"
	"Catch/bump/excel/cmdb/get"
	"Catch/bump/pull/storage/eos"
	"github.com/sirupsen/logrus"
	"strconv"
)

func EOSComparison(excel get.StorageResource, consoles []*eos.ElasticObjectStorage) {
	var mark = false
	logrus.Debugf("------------------------------ Row:%s ------------------------------", excel.Row)
	for _, console := range consoles {

		if excel.ResourceName == console.Name {
			mark = true
			iinternal.IfFieldComparison("", excel.Row, excel.ResourceName, console.Name)

			iinternal.IfFieldComparison("", excel.Row, excel.Specification, console.StorageClass)

			iinternal.IfFieldComparison("", excel.Row, excel.Node, console.Region)

		}
	}
	if !mark {
		logrus.Warnf("未匹配到资源,Row=%s ID=%s Name=%s", excel.Row, excel.ResourceID, excel.ResourceName)
		rowInt, _ := strconv.Atoi(excel.Row)
		iinternal.MisMatchResourceCount[rowInt] = 1
	}
}
