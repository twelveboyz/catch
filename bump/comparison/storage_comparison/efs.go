package storage_comparison

import (
	iinternal "Catch/bump/comparison/internal/utils"
	"Catch/bump/excel/cmdb/get"
	"Catch/bump/internal/mapping"
	"Catch/bump/pull/storage/efs"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func EFSComparison(excel get.StorageResource, consoles []*efs.ElasticFileStorage) {
	var mark = false
	logrus.Debugf("------------------------------ Row:%s ------------------------------", excel.Row)
	for _, console := range consoles {

		if excel.ResourceID == console.ShareId {
			mark = true
			iinternal.IfFieldComparison("标签1", excel.Row, strings.TrimLeft(excel.Tag1, "_"), console.Tag1)

			iinternal.IfFieldComparison("标签X", excel.Row, strings.TrimLeft(excel.TagX, "_"), console.TagX)

			iinternal.IfFieldComparison("标签2", excel.Row, excel.Tag2, console.Tag2)

			iinternal.IfFieldComparison(" ID ", excel.Row, excel.ResourceID, console.ShareId)

			iinternal.IfFieldComparison("资源名称", excel.Row, excel.ResourceName, console.Name)

			iinternal.IfFieldComparison("规格", excel.Row, excel.Specification, console.ShareType)

			iinternal.IfFieldComparison("容量", excel.Row, excel.Capacity, console.Size)

			iinternal.IfFieldComparison("挂载点", excel.Row, strings.ReplaceAll(excel.Mount, "\n", ","), strings.ReplaceAll(console.Mount, "\n", ","))

			iinternal.IfFieldComparison("资源池", excel.Row, excel.Node, mapping.PoolCodeToNameConvert(console.Pool))

			iinternal.IfFieldComparison("可用区", excel.Row, excel.AvailabilityZone, mapping.AZConvert(console.Region))

		}
	}
	if !mark {
		logrus.Warnf("未匹配到资源,Row=%s ID=%s Name=%s", excel.Row, excel.ResourceID, excel.ResourceName)
		rowInt, _ := strconv.Atoi(excel.Row)
		iinternal.MisMatchResourceCount[rowInt] = 1
	}
}
