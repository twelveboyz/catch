package network_comparison

import (
	iinternal "Catch/bump/comparison/internal/utils"
	"Catch/bump/excel/cmdb/get"
	"Catch/bump/internal/mapping"
	"Catch/bump/pull/network/eip"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func EipComparison(excel get.EipResource, consoles []*eip.ElasticIP) {
	var mark = false
	logrus.Debugf("------------------------------ Row:%s ------------------------------", excel.Row)

	for _, console := range consoles {
		if excel.ResourceID == console.Id {
			mark = true
			iinternal.IfFieldComparison("标签1", excel.Row, strings.TrimLeft(excel.Tag1, "_"), console.Tag1)

			iinternal.IfFieldComparison("标签X", excel.Row, strings.TrimLeft(excel.TagX, "_"), console.TagX)

			iinternal.IfFieldComparison("标签2", excel.Row, excel.Tag2, console.Tag2)

			iinternal.IfFieldComparison(" ID ", excel.Row, excel.ResourceID, console.Id)

			iinternal.IfFieldComparison("资源名称", excel.Row, excel.ResourceName, console.ResourceName)

			iinternal.IfFieldComparison("IPv4公网地址", excel.Row, excel.IPv4PublicAddress, console.EipName)

			iinternal.IfFieldComparison("带宽", excel.Row, excel.BandWidth, console.BandwidthSize)

			iinternal.IfFieldContains("绑定资源类型", excel.Row, excel.BindingResourceType, console.BindType)

			iinternal.IfFieldComparison("绑定资源名称", excel.Row, excel.BindingResourceName, console.BindResourceName)

			iinternal.IfFieldComparison("绑定资源ID", excel.Row, excel.BindingResourceID, console.BindResourceId)

			iinternal.IfFieldComparison("资源池", excel.Row, excel.Node, mapping.PoolCodeToNameConvert(console.Pool))

		}
	}

	if !mark {
		logrus.Warnf("未匹配到资源,Row=%s ID=%s Name=%s", excel.Row, excel.ResourceID, excel.ResourceName)
		rowInt, _ := strconv.Atoi(excel.Row)
		iinternal.MisMatchResourceCount[rowInt] = 1
	}
}
