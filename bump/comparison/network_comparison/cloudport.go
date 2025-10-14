package network_comparison

import (
	iinternal "Catch/bump/comparison/internal/utils"
	"Catch/bump/excel/cmdb/get"
	"Catch/bump/pull/network/cloudport"
	"github.com/sirupsen/logrus"
	"strconv"
)

func CPComparison(excel get.CloudPortResource, consoles []cloudport.CloudPort) {
	var mark = false
	logrus.Debugf("------------------------------ Row:%s ------------------------------", excel.Row)

	for _, console := range consoles {
		if excel.ResourceID == console.Id {
			mark = true
			iinternal.IfFieldComparison(" ID ", excel.Row, excel.ResourceID, console.Id)

			iinternal.IfFieldComparison("资源名称", excel.Row, excel.ResourceName, console.ParticularLineName)

			iinternal.IfFieldComparison("规格", excel.Row, excel.Specification, console.PortSpecification)

			iinternal.IfFieldComparison("带宽", excel.Row, excel.BandWidth, console.ParticularLineBandwidth)

			iinternal.IfFieldComparison("ip类型", excel.Row, excel.IPType, console.IpType)

			iinternal.IfFieldComparison("VPC_ID", excel.Row, excel.VPC, console.VpcId)

			iinternal.IfFieldComparison("接入节点地域", excel.Row, excel.AccessNodeRegion, console.AccessNodeRegion)
			iinternal.IfFieldComparison("接入节点详细地址", excel.Row, excel.AccessNodeAddress, console.AccessNodeAddress)
			iinternal.IfFieldComparison("接入节点联系人", excel.Row, excel.AccessNodeContact, console.AccessNodeContact)
			iinternal.IfFieldComparison("客户经理", excel.Row, excel.AccountManager, console.AccountManager)

			iinternal.IfFieldComparison("vgw与汇聚设备对接IP(主)", excel.Row, excel.VGWAndAggDockingIPMaster, console.VGWAndAggDockingIPMaster)
			iinternal.IfFieldComparison("vgw与汇聚设备对接IP(备)", excel.Row, excel.VGWAndAggDockingIPBackup, console.VGWAndAggDockingIPBackup)

			iinternal.IfFieldComparison("汇聚设备与vgw对接IP(主)", excel.Row, excel.AggAndVGWDockingIPMaster, console.AggAndVGWDockingIPMaster)
			iinternal.IfFieldComparison("汇聚设备与vgw对接IP(备)", excel.Row, excel.AggAndVGWDockingIPBackup, console.AggAndVGWDockingIPBackup)

			iinternal.IfFieldComparison("汇聚设备与vgw对接VLAN(主)", excel.Row, excel.AggAndVGWDockingVLANMaster, console.AggAndVGWDockingVLANMaster)
			iinternal.IfFieldComparison("汇聚设备与vgw对接VLAN(备)", excel.Row, excel.AggAndVGWDockingVLANBackup, console.AggAndVGWDockingVLANBackup)

			iinternal.IfFieldComparison("汇聚设备与客户设备对接逻辑端口", excel.Row, excel.AggAndCustomerDockingLogicalPort, console.AggAndCustomerDockingLogicalPort)
			iinternal.IfFieldComparison("汇聚设备与客户设备对接物理端口", excel.Row, excel.AggAndCustomerDockingPhysicalPort, console.AggAndCustomerDockingPhysicalPort)

			iinternal.IfFieldComparison("设备1位置", excel.Row, excel.Device1, console.Device1)
			iinternal.IfFieldComparison("设备2位置", excel.Row, excel.Device2, console.Device2)

			iinternal.IfFieldComparison("资源池", excel.Row, excel.ResourcePool, console.PoolName)
			iinternal.IfFieldComparison("节点", excel.Row, excel.Node, console.PoolName)

		}
	}
	if !mark {
		logrus.Warnf("未匹配到资源,Row=%s ID=%s Name=%s", excel.Row, excel.ResourceID, excel.ResourceName)
		rowInt, _ := strconv.Atoi(excel.Row)
		iinternal.MisMatchResourceCount[rowInt] = 1
	}
}
