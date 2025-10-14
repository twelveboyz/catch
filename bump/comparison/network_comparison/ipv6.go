package network_comparison

import (
	iinternal "Catch/bump/comparison/internal/utils"
	"Catch/bump/excel/cmdb/get"
	"Catch/bump/internal/mapping"
	"Catch/bump/pull/network/ipv6"
	"github.com/sirupsen/logrus"
	"strconv"
)

func IPv6Comparison(excel get.IPv6Resource, consoles []*ipv6.IPv6) {
	var mark = false
	logrus.Debugf("------------------------------ Row:%s ------------------------------", excel.Row)

	for _, console := range consoles {
		if excel.IPv6Address == console.IpAddress {
			mark = true
			if console.MixedId != "" {
				iinternal.IfFieldComparison(" ID ", excel.Row, excel.ResourceID, console.IpAddress)
			} else {
				iinternal.IfFieldComparison(" ID ", excel.Row, excel.ResourceID, console.NsQosPolicyId)
			}

			iinternal.IfFieldComparison("IPv6", excel.Row, excel.IPv6Address, console.IpAddress)

			iinternal.IfFieldComparison("带宽", excel.Row, excel.BandWidth, console.BandWidthSize)

			iinternal.IfFieldComparison("关联资源", excel.Row, excel.RelatedResources, console.BindResourceName)

			iinternal.IfFieldComparison("资源池", excel.Row, excel.Node, mapping.PoolCodeToNameConvert(console.Pool))

		}
	}
	if !mark {
		logrus.Warnf("未匹配到资源,Row=%s ID=%s Name=%s", excel.Row, excel.ResourceID, excel.ResourceName)
		rowInt, _ := strconv.Atoi(excel.Row)
		iinternal.MisMatchResourceCount[rowInt] = 1
	}
}
