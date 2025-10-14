package execute

import (
	"Catch/bump/comparison/internal/utils"
	"Catch/bump/comparison/network_comparison"
	"Catch/bump/excel/cmdb/get"
	"Catch/bump/pull/network/cloudport"
	"Catch/bump/pull/network/directconnect"
	"Catch/bump/pull/network/eip"
	"Catch/bump/pull/network/ipv6"
	"Catch/bump/pull/network/loadbalance"
	"Catch/bump/pull/network/natgateway"
	"Catch/bump/pull/network/sharedtrafficpackage"
	"errors"
	"github.com/sirupsen/logrus"
)

func NetworkRun(root string) {
	networkExcel, err := initNetworkExcel(root)
	if err != nil {
		logrus.Infoln(err.Error())
		return
	}

	/*------------------------------循环网络表格的sheet---------------------------------------*/
	sheets := networkExcel.ForGetSheet()
	for _, sheet := range sheets {
		networkExcel.Sheet = sheet

		if err = networkExcel.EffectiveSheet(); err != nil {
			continue
		}

		var excelResources, err = networkExcel.ParseNetworkResourceToStruct()
		if err != nil {
			logrus.Errorf(err.Error())
			return
		}

		uuids, err := get.NetworkGetUUIDs(excelResources)
		if err != nil {
			logrus.Errorln(err.Error())
			return
		}

		pools, err := get.NetworkGetPool(excelResources)
		if err != nil {
			logrus.Errorln(err.Error())
			return
		} else if len(pools) == 0 {
			logrus.Errorln("pools is empty")
			continue
		}

		// 根据不同的资源类型进行处理
		switch {
		case len(excelResources.LoadBalance) > 0:
			processLoadBalanceResources(excelResources.LoadBalance, uuids, pools)

		case len(excelResources.Eip) > 0:
			processEipResources(excelResources.Eip, uuids, pools)

		case len(excelResources.SharedTrafficPackage) > 0:
			processSharedTrafficPackageResources(excelResources.SharedTrafficPackage, uuids, pools)

		case len(excelResources.IPv6) > 0:
			processIPv6Resources(excelResources.IPv6, pools)

		case len(excelResources.DirectConnect) > 0:
			processDirectConnectResources(excelResources.DirectConnect)

		case len(excelResources.NATGateway) > 0:
			processNATGatewayResources(excelResources.NATGateway, uuids, pools)

		case len(excelResources.CloudPort) > 0:
			processCloudPortResources(excelResources.CloudPort)

		case len(excelResources.Other) > 0:
			err = processOtherResources(excelResources.Other, uuids, pools)
			if err != nil {
				logrus.Traceln(err)
				continue
			}
		default:
			logrus.Errorf("未匹配到sheet: %s\n", sheet)
			return
		}

		if len(excelResources.IPv6) == 0 && len(excelResources.Eip) == 0 && len(excelResources.LoadBalance) == 0 && len(excelResources.NATGateway) == 0 && len(excelResources.Other) == 0 && len(excelResources.CloudPort) == 0 && len(excelResources.DirectConnect) == 0 && len(excelResources.SharedTrafficPackage) == 0 {
			return
		}

		utils.NetworkInfoSummaryPrint(sheet)
	}
}

// 定义处理不同资源类型的函数
func processLoadBalanceResources(resources []get.LoadBalanceResource, uuids, pools []string) {
	consoleSlb := loadbalance.ExecuteGetLoadBalanceInfo(uuids, pools)
	for _, resource := range resources {
		network_comparison.SlbComparison(resource, consoleSlb)
	}
}

func processEipResources(resources []get.EipResource, uuids, pools []string) {
	consoleEip := eip.ExecuteGetEIPInfo(uuids, pools)
	for _, resource := range resources {
		network_comparison.EipComparison(resource, consoleEip)
	}
}

func processSharedTrafficPackageResources(resources []get.SharedTrafficPackageResource, uuids, pools []string) {
	consoleStp := sharedtrafficpackage.ExecuteGetSTPInfo(uuids, pools)
	for _, resource := range resources {
		network_comparison.StpComparison(resource, consoleStp)
	}
}

func processIPv6Resources(resources []get.IPv6Resource, pools []string) {
	consoleIpv6 := ipv6.ExecuteGetIPv6Info(pools)
	for _, resource := range resources {
		network_comparison.IPv6Comparison(resource, consoleIpv6)
	}
}

func processDirectConnectResources(resources []get.DirectConnectResource) {
	consoleDC := directconnect.ExecuteGetResourceInfo()
	for _, resource := range resources {
		network_comparison.DcComparison(resource, consoleDC)
	}
}

func processNATGatewayResources(resources []get.NATGatewayResource, uuids, pools []string) {
	consoleNat := natgateway.ExecuteGetNatInfo(uuids, pools)
	for _, resource := range resources {
		network_comparison.NatComparison(resource, consoleNat)
	}
}

func processCloudPortResources(resources []get.CloudPortResource) {
	consoleCP := cloudport.ExecuteGetCPInfo()

	for _, resource := range resources {
		network_comparison.CPComparison(resource, consoleCP)
	}
}

func processOtherResources(resources []get.OtherResource, uuids, pools []string) error {
	for _, _ = range resources {

	}
	return errors.New("Not Supported")
}

func initNetworkExcel(root string) (*get.Excel, error) {
	networkFileName := get.GetFileName(root, ".xls", "网络")
	if networkFileName == "" {
		return nil, errors.New("未找到网络资源文件，已跳过···")
	}

	networkExcel := get.NewExcel(root, networkFileName, 2, 3)

	networkExcel.File = networkExcel.OpenFile(networkExcel.FileName)

	return networkExcel, nil
}
