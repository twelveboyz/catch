package ecs

import (
	"Catch/bump/internal/bootstrap"
	"Catch/bump/internal/mapping"
	"Catch/bump/internal/utils"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"strings"
)

func ExecuteALL() []*ECloudServer {
	Body := utils.Get(Url, Headers)
	//log.Println(Body)

	InfoSet := JsonParseResourceInfoALL(Body)
	//JsonParseResourceNetworkInfoByPools(InfoSet,pool)
	err := JsonParseTags(InfoSet)
	if err != nil {
		logrus.Warnln(err)
	}

	for i, info := range InfoSet {
		log.Println("循环次数：", i+1)
		log.Println(info.Id, info.Name, info.Region, info.SpecsName, info.ImageName, info.ImageOsType, info.VCpu, info.VMemory, info.VDisk)

		for _, nic := range info.NIC {
			log.Println(nic.Id, nic.Name, nic.Type, nic.ResourceId, nic.NetworkId)

			for _, fixedIp := range nic.FixedIp {
				log.Println(fixedIp)
			}

			for _, securityGroup := range nic.SecurityGroup {
				log.Println(securityGroup)
			}
			log.Println()
		}
		log.Println()
	}

	logrus.Println("云主机共发起请求：", len(InfoSet))

	return InfoSet
}

func ExecuteGetEcsInfo(uuids []string, pools []string) []*ECloudServer {
	logrus.Println("正在获取云主机控制台资源······")
	var InfoSet []*ECloudServer
	for _, pool := range pools {
		Body := utils.Get(Url, HeadersFun(pool))

		is := JsonParseResourceByUUIDs(Body, uuids)
		JoinPoolInfo(is, pool)
		JsonParseResourceNetworkInfoByPools(is, pool)
		err := JsonParseTags(is)
		if err != nil {
			logrus.Warnln(err)
		}

		InfoSet = append(InfoSet, is...)
	}

	return InfoSet
}

func ExecuteMatchByName() {
	projectID := utils.InputProjectID()

	for _, pool := range bootstrap.PoolIDs {
		Body := utils.Get(Url, HeadersFun(pool))
		//fmt.Println(HeadersFun(pool))
		//log.Println(Body)

		infoSet := JsonParseResourceInfoALL(Body)

		JoinPoolInfo(infoSet, pool)

		JsonParseResourceNetworkInfoByPools(infoSet, pool)

		PrintEcsInfoByProjectID(infoSet, projectID)

	}

}

func ExecuteMatchByID() {
	uuids := utils.InputUUIDs()

	for _, pool := range bootstrap.PoolIDs {
		Body := utils.Get(Url, HeadersFun(pool))

		infoSet := JsonParseResourceInfoALL(Body)

		JoinPoolInfo(infoSet, pool)

		JsonParseResourceNetworkInfoByPools(infoSet, pool)

		PrintEcsInfoByUUIDs(infoSet, uuids)

	}

}

func PrintEcsInfoByUUIDs(ECSInfoSet []*ECloudServer, uuids []string) {
	for _, uuid := range uuids {
		//uuid不匹配直接跳过
		for _, info := range ECSInfoSet {
			if uuid != info.Id {
				continue
			}

			fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t", info.Id, info.Name, "弹性计算", "云主机", info.SpecsName, info.VCpu, info.VMemory, "", info.ImageName, "运行中", "未交维", "移动云", mapping.PoolCodeToNameConvert(info.Pool), info.Region, info.VDisk, info.BootVolumeType, "按月计费", info.CreateTime, "")

			for _, nic := range info.NIC {
				fmt.Printf("%s\t%s\t%s\t%s\t", nic.Name, "", nic.VpcName, nic.SubnetName)

				for _, fixedIp := range nic.FixedIp {
					if fixedIp.IpVersion == "4" {
						mask := utils.CidrToMask(fixedIp.SubnetCidr)
						fmt.Printf("%s\t%s\t", fixedIp.IpAddress, mask)
					}
				}
				for _, fixedIp := range nic.FixedIp {
					if fixedIp.IpVersion == "6" {
						fmt.Printf("%s\t%s\t", fixedIp.IpAddress, "/128")
					}
				}

				var sgSlice []string
				for _, securityGroup := range nic.SecurityGroup {
					sgSlice = append(sgSlice, securityGroup.Name)
				}

				fmt.Print(strings.Join(sgSlice, ","))

			}
			break
		}
		fmt.Println()
	}
}

func PrintEcsInfoByProjectID(EcsInfoSet []*ECloudServer, projectID string) {
	for _, info := range EcsInfoSet {
		if !strings.HasPrefix(info.Name, projectID) {
			continue
		}

		fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t", info.Id, info.Name, "弹性计算", "云主机", info.SpecsName, info.VCpu, info.VMemory, "", info.ImageName, "运行中", "未交维", "移动云", mapping.PoolCodeToNameConvert(info.Pool), info.Region, info.VDisk, info.BootVolumeType, "按月计费", info.CreateTime, "")

		for _, nic := range info.NIC {
			fmt.Printf("%s\t%s\t%s\t%s\t", nic.Name, "", nic.VpcName, nic.SubnetName)

			for _, fixedIp := range nic.FixedIp {
				if fixedIp.IpVersion == "4" {
					mask := utils.CidrToMask(fixedIp.SubnetCidr)
					fmt.Printf("%s\t%s\t", fixedIp.IpAddress, mask)
				}
			}
			for _, fixedIp := range nic.FixedIp {
				if fixedIp.IpVersion == "6" {
					fmt.Printf("%s\t%s\t", fixedIp.IpAddress, "/128")
				}
			}

			var sgSlice []string
			for _, securityGroup := range nic.SecurityGroup {
				sgSlice = append(sgSlice, securityGroup.Name)
			}
			fmt.Print(strings.Join(sgSlice, ","))

			break
		}
		fmt.Println()
	}
}

func EcsInfoByProjectIDToStruct(EcsInfoSet []*ECloudServer, projectID string) []*ExcelECS {
	var excelEcsSlice []*ExcelECS

	for _, info := range EcsInfoSet {
		if !strings.HasPrefix(info.Name, projectID) {
			continue
		}

		var excelEcs ExcelECS
		excelEcs.Id = info.Id
		excelEcs.Name = info.Name
		excelEcs.SpecsName = info.SpecsName
		excelEcs.VCpu = info.VCpu
		excelEcs.VMemory = info.VMemory
		excelEcs.ImageName = info.ImageName
		excelEcs.BootVolumeType = info.BootVolumeType
		excelEcs.VDisk = info.VDisk

		for _, nic := range info.NIC {
			var excelNic ExcelNic
			excelNic.VnName = nic.Name
			excelNic.VpcName = nic.VpcName
			excelNic.SubnetName = nic.SubnetName

			for _, fixedIp := range nic.FixedIp {
				if fixedIp.IpVersion == "4" {
					mask := utils.CidrToMask(fixedIp.SubnetCidr)
					excelNic.IpV4Address = fixedIp.IpAddress
					excelNic.V4Mask = mask
				}
			}

			for _, fixedIp := range nic.FixedIp {
				if fixedIp.IpVersion == "6" {
					excelNic.IpV6Address = fixedIp.IpAddress
					excelNic.V6Mask = "/128"
				}
			}

			for _, securityGroup := range nic.SecurityGroup {
				excelNic.SecurityGroup = securityGroup.Name
			}
			//将网络信息加入到excelEcs中
			excelEcs.NIC = append(excelEcs.NIC, &excelNic)
		}
		//将云主机信息加入到excelEcsSlice中
		excelEcsSlice = append(excelEcsSlice, &excelEcs)
	}

	return excelEcsSlice
}

func EcsInfoByUUIDsToStruct(ECSInfoSet []*ECloudServer, uuids []string) []*ExcelECS {
	var excelEcsSlice []*ExcelECS

	for _, uuid := range uuids {
		//uuid不匹配直接跳过
		for _, info := range ECSInfoSet {
			if uuid != info.Id {
				continue
			}
			var excelEcs ExcelECS
			excelEcs.Id = info.Id
			excelEcs.Name = info.Name
			excelEcs.SpecsName = info.SpecsName
			excelEcs.VCpu = info.VCpu
			excelEcs.VMemory = info.VMemory
			excelEcs.ImageName = info.ImageName
			excelEcs.BootVolumeType = info.BootVolumeType
			excelEcs.VDisk = info.VDisk

			for _, nic := range info.NIC {
				var excelNic ExcelNic
				excelNic.VnName = nic.Name
				excelNic.VpcName = nic.VpcName
				excelNic.SubnetName = nic.SubnetName

				for _, fixedIp := range nic.FixedIp {
					if fixedIp.IpVersion == "4" {
						mask := utils.CidrToMask(fixedIp.SubnetCidr)
						excelNic.IpV4Address = fixedIp.IpAddress
						excelNic.V4Mask = mask
					}
				}

				for _, fixedIp := range nic.FixedIp {
					if fixedIp.IpVersion == "6" {
						excelNic.IpV6Address = fixedIp.IpAddress
						excelNic.V6Mask = "/128"
					}
				}

				for _, securityGroup := range nic.SecurityGroup {
					excelNic.SecurityGroup = securityGroup.Name
				}
				//将网络信息加入到excelEcs中
				excelEcs.NIC = append(excelEcs.NIC, &excelNic)
			}
			//将云主机信息加入到excelEcsSlice中
			excelEcsSlice = append(excelEcsSlice, &excelEcs)

		}
		break
	}

	return excelEcsSlice
}
