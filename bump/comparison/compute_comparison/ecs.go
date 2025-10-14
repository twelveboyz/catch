package compute_comparison

import (
	iinternal "Catch/bump/comparison/internal/utils"
	"Catch/bump/excel/cmdb/get"
	"Catch/bump/internal/mapping"
	"Catch/bump/internal/utils"
	"Catch/bump/pull/compute/ecs"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func EcsComparison(excel get.ComputeResource, consoleEcss []*ecs.ECloudServer) {
	var mark = false
	logrus.Debugf("------------------------------ Row:%s ------------------------------", excel.Row)
	for _, consoleEcs := range consoleEcss {

		if excel.ResourceID == consoleEcs.Id {
			mark = true
			iinternal.IfFieldComparison("标签1", excel.Row, strings.TrimLeft(excel.Tag1, "_"), consoleEcs.Tag1)

			iinternal.IfFieldComparison("标签X", excel.Row, strings.TrimLeft(excel.TagX, "_"), consoleEcs.TagX)

			iinternal.IfFieldComparison("标签2", excel.Row, excel.Tag2, consoleEcs.Tag2)

			iinternal.IfFieldComparison(" ID ", excel.Row, excel.ResourceID, consoleEcs.Id)

			iinternal.IfFieldComparison("Name", excel.Row, excel.ResourceName, consoleEcs.Name)

			iinternal.IfFieldComparison("规格", excel.Row, excel.Specification, consoleEcs.SpecsName)

			iinternal.IfFieldComparison("镜像", excel.Row, excel.Image, consoleEcs.ImageName)

			iinternal.IfFieldComparison("CPU", excel.Row, excel.CPU, consoleEcs.VCpu)

			iinternal.IfFieldComparison("Memory", excel.Row, excel.Memory, consoleEcs.VMemory)

			iinternal.IfFieldComparison("系统盘规格", excel.Row, excel.SystemDiskSpec, consoleEcs.BootVolumeType)

			iinternal.IfFieldComparison("系统盘大小", excel.Row, excel.SystemDiskSize, consoleEcs.VDisk)

			for i, eNicName := range []string{excel.NIC1.Name, excel.NIC2.Name, excel.NIC3.Name} {

				if eNicName == "" {
					continue
				}

				var nicMark bool
				for _, n := range consoleEcs.NIC {
					//NIC1---------------
					if eNicName == n.Name && i == 0 {
						nicMark = true
						iinternal.IfFieldComparison("网卡名称_1", excel.Row, excel.NIC1.Name, n.Name)

						for _, f := range n.FixedIp {

							iinternal.IfFieldComparison("VPC_1", excel.Row, excel.NIC1.VPC, f.VpcName)

							iinternal.IfFieldComparison("子网_1", excel.Row, excel.NIC1.Subnet, f.SubnetName)

							iinternal.IfFieldComparison("IPv4_1", excel.Row, excel.NIC1.IPv4PrivateAddress, f.IpAddress, map[string]string{f.IpVersion: "4"})

							//将console的24位后缀转成子网掩码
							mask := utils.CidrToMask(f.SubnetCidr)
							iinternal.IfFieldComparison("IPv4Mask_1", excel.Row, excel.NIC1.IPv4PrivateMask, mask, map[string]string{f.IpVersion: "4"})

							iinternal.IfFieldComparison("IPv6_1", excel.Row, excel.NIC1.IPv6Address, f.IpAddress, map[string]string{f.IpVersion: "6"})

							//consoleV6Cidr := utils.CidrRegex(f.SubnetCidr)
							//iinternal.IfFieldComparison("",excel.Row,excel.NIC1.IPv6Mask, "/"+consoleV6Cidr, map[string]string{f.IpVersion: "6"})
						}

						excelSgs := strings.Split(excel.NIC1.SecurityGroup, ",")
						for _, esg := range excelSgs {
							for _, sg := range n.SecurityGroup {
								if esg == sg.Name {
									iinternal.IfFieldComparison("安全组_1", excel.Row, esg, sg.Name)
								}
							}
						}

						//NIC1---------------

						//NIC2---------------
					} else if eNicName == n.Name && i == 1 {
						nicMark = true
						iinternal.IfFieldComparison("网卡名称_2", excel.Row, excel.NIC2.Name, n.Name)

						for _, f := range n.FixedIp {

							iinternal.IfFieldComparison("VPC_2", excel.Row, excel.NIC2.VPC, f.VpcName)

							iinternal.IfFieldComparison("子网_2", excel.Row, excel.NIC2.Subnet, f.SubnetName)

							iinternal.IfFieldComparison("IPv4_2", excel.Row, excel.NIC2.IPv4PrivateAddress, f.IpAddress, map[string]string{f.IpVersion: "4"})

							//将console的24位后缀转成子网掩码
							mask := utils.CidrToMask(f.SubnetCidr)
							iinternal.IfFieldComparison("IPv4Mask_2", excel.Row, excel.NIC2.IPv4PrivateMask, mask, map[string]string{f.IpVersion: "4"})

							iinternal.IfFieldComparison("IPv6_2", excel.Row, excel.NIC2.IPv6Address, f.IpAddress, map[string]string{f.IpVersion: "6"})

							//consoleV6Cidr := utils.CidrRegex(f.SubnetCidr)
							//iinternal.IfFieldComparison("",excel.Row,excel.NIC2.IPv6Mask, "/"+consoleV6Cidr, map[string]string{f.IpVersion: "6"})
						}

						excelSgs := strings.Split(excel.NIC2.SecurityGroup, ",")
						for _, esg := range excelSgs {
							for _, sg := range n.SecurityGroup {
								if esg == sg.Name {
									iinternal.IfFieldComparison("安全组_2", excel.Row, esg, sg.Name)
								}
							}
						}
						//NIC2---------------

						//NIC3---------------
					} else if eNicName == n.Name && i == 2 {
						nicMark = true
						iinternal.IfFieldComparison("网卡名称_3", excel.Row, excel.NIC3.Name, n.Name)

						for _, f := range n.FixedIp {

							iinternal.IfFieldComparison("VPC_3", excel.Row, excel.NIC3.VPC, f.VpcName)

							iinternal.IfFieldComparison("子网_3", excel.Row, excel.NIC3.Subnet, f.SubnetName)

							iinternal.IfFieldComparison("IPv4_3", excel.Row, excel.NIC3.IPv4PrivateAddress, f.IpAddress, map[string]string{f.IpVersion: "4"})

							//将console的34位后缀转成子网掩码
							mask := utils.CidrToMask(f.SubnetCidr)

							iinternal.IfFieldComparison("IPv4Mask_3", excel.Row, excel.NIC3.IPv4PrivateMask, mask, map[string]string{f.IpVersion: "4"})

							iinternal.IfFieldComparison("IPv6_3", excel.Row, excel.NIC3.IPv6Address, f.IpAddress, map[string]string{f.IpVersion: "6"})

							//consoleV6Cidr := utils.CidrRegex(f.SubnetCidr)
							//iinternal.IfFieldComparison("",excel.Row,excel.NIC3.IPv6Mask, "/"+consoleV6Cidr, map[string]string{f.IpVersion: "6"})
						}

						excelSgs := strings.Split(excel.NIC3.SecurityGroup, ",")
						for _, esg := range excelSgs {
							for _, sg := range n.SecurityGroup {
								if esg == sg.Name {
									iinternal.IfFieldComparison("安全组_3", excel.Row, esg, sg.Name)
								}
							}
						}
					}
					//NIC3---------------
				}
				if !nicMark {
					var nicList []string
					for _, nic := range consoleEcs.NIC {
						nicList = append(nicList, nic.Name)
					}
					logrus.Warnf("网卡%d %s,未匹配到控制台网卡名称:%v,", i+1, eNicName, nicList)
				}
			}

			iinternal.IfFieldComparison("资源池", excel.Row, excel.Node, mapping.PoolCodeToNameConvert(consoleEcs.Pool))

			iinternal.IfFieldComparison("可用区", excel.Row, excel.AvailabilityZone, consoleEcs.Region)
		}
	}

	if !mark {
		logrus.Warnf("未匹配到资源,Row=%s ID=%s Name=%s", excel.Row, excel.ResourceID, excel.ResourceName)
		rowInt, _ := strconv.Atoi(excel.Row)
		iinternal.MisMatchResourceCount[rowInt] = 1
	}
}
