package network_comparison

import (
	iinternal "Catch/bump/comparison/internal/utils"
	"Catch/bump/excel/cmdb/get"
	"Catch/bump/pull/network/directconnect"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func DcComparison(excel get.DirectConnectResource, consoles []directconnect.DirectConnect) {
	var mark = false
	logrus.Debugf("------------------------------ Row:%s ------------------------------", excel.Row)

	for _, console := range consoles {
		if excel.ResourceID == console.Id {
			mark = true
			iinternal.IfFieldComparison(" ID ", excel.Row, excel.ResourceID, console.Id)

			iinternal.IfFieldComparison("资源名称", excel.Row, excel.ResourceName, console.SpecialLineName)

			iinternal.IfFieldComparison("VPC", excel.Row, excel.VPC, console.VpcName)

			iinternal.IfFieldComparison("VPC子网", excel.Row, excel.VPCSubnet, strings.Join(console.VpsSubnets, "；"))

			iinternal.IfFieldComparison("用户子网", excel.Row, excel.UserSubnet, strings.Join(console.UserSubnets, "；"))

			iinternal.IfFieldComparison("接入节点地域", excel.Row, excel.AccessNodeRegion, console.Province+console.City+console.Borough)

			iinternal.IfFieldComparison("接入节点详细地址", excel.Row, excel.AccessNodeAddress, console.Province+console.City+console.Borough+console.Address)

			iinternal.IfFieldComparison("接入节点联系人", excel.Row, excel.AccessNodeContact, console.ContactName)

			iinternal.IfFieldComparison("接入节点客户经理", excel.Row, excel.AccessNodeAccountManager, console.ManagerName)

			iinternal.IfFieldComparison("带宽", excel.Row, excel.BandWidth, strings.ReplaceAll(console.SpecialLineBandwidth, "M", ""))

		}
	}
	if !mark {
		logrus.Warnf("未匹配到资源,Row=%s ID=%s Name=%s", excel.Row, excel.ResourceID, excel.ResourceName)
		rowInt, _ := strconv.Atoi(excel.Row)
		iinternal.MisMatchResourceCount[rowInt] = 1
	}
}
