package network_comparison

import (
	iinternal "Catch/bump/comparison/internal/utils"
	"Catch/bump/excel/cmdb/get"
	"Catch/bump/internal/mapping"
	"Catch/bump/pull/network/loadbalance"
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func SlbComparison(excel get.LoadBalanceResource, consoles []*loadbalance.LoadBalance) {
	var mark = false
	logrus.Debugf("------------------------------ Row:%s ------------------------------", excel.Row)

	for _, console := range consoles {
		if excel.ResourceID == console.Id {
			mark = true

			iinternal.IfFieldComparison("标签1", excel.Row, strings.TrimLeft(excel.Tag1, "_"), console.Tag1)

			iinternal.IfFieldComparison("标签X", excel.Row, strings.TrimLeft(excel.TagX, "_"), console.TagX)

			iinternal.IfFieldComparison("标签2", excel.Row, excel.Tag2, console.Tag2)

			iinternal.IfFieldComparison(" ID ", excel.Row, excel.ResourceID, console.Id)

			iinternal.IfFieldComparison("资源名称", excel.Row, excel.ResourceName, console.Name)

			if strings.Contains(excel.Specification, "性能保障型") {
				bandwidth := console.Bandwidth / 1024
				iinternal.IfFieldComparison("规格", excel.Row, excel.Specification, fmt.Sprintf("%s-%dG", console.Flavor, bandwidth))
			} else {
				iinternal.IfFieldComparison("规格", excel.Row, excel.Specification, console.Flavor)
			}

			switch console.IPVersion {
			case "4":
				iinternal.IfFieldComparison("IPv4内网地址", excel.Row, excel.IPv4PrivateAddress, console.Privateip)
			case "6":
				iinternal.IfFieldComparison("IPv6地址", excel.Row, excel.IPv6Address, console.Privateip)
			default:
				logrus.Errorf("未找到IPVersion对应版本,id=%s  ipversion=%s", excel.ResourceID, console.IPVersion)
			}

			if console.PublicIp != "" && console.IPVersion == "4" {
				iinternal.IfFieldComparison("公网IPv4地址", excel.Row, excel.IPv4PublicAddress, console.PublicIp)
			}

			iinternal.IfFieldComparison("VPC", excel.Row, excel.VPC, console.Vpcname)

			iinternal.IfFieldComparison("子网", excel.Row, excel.Subnet, console.SubnetName)

			iinternal.IfFieldComparison("资源池", excel.Row, excel.Node, mapping.PoolCodeToNameConvert(console.Pool))

		}
	}
	if !mark {
		logrus.Warnf("未匹配到资源,Row=%s ID=%s Name=%s", excel.Row, excel.ResourceID, excel.ResourceName)
		rowInt, _ := strconv.Atoi(excel.Row)
		iinternal.MisMatchResourceCount[rowInt] = 1
	}
}
