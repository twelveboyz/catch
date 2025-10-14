package ironic

import (
	"Catch/bump/internal/bootstrap"
	"Catch/bump/internal/mapping"
	"Catch/bump/internal/utils"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"strings"
)

func ExecuteGetIronicInfo(uuids []string, pools []string) []*Ironic {
	logrus.Println("正在获取裸金属控制台资源······")
	var InfoSet []*Ironic
	for _, pool := range pools {
		Body := utils.Get(Url, HeadersFun(pool))
		//log.Println(Body)

		//InfoSet := JsonParseResourceInfoALL(Body)
		is := JsonParseResourceByUUIDs(Body, uuids)
		JoinPoolInfo(is, pool)
		JsonParseResourceNetworkInfoByPool(is, pool)
		err := JsonParseTags(is)
		if err != nil {
			logrus.Warnln(err)
		}

		InfoSet = append(InfoSet, is...)
	}

	return InfoSet
}

func ExecuteALL() []*Ironic {
	Body := utils.Get(Url, Headers)
	//fmt.Println(Body)

	InfoSet := JsonParseResourceInfoALL(Body)
	JsonParseResourceNetworkInfo(InfoSet)
	err := JsonParseTags(InfoSet)
	if err != nil {
		logrus.Warnln(err)
	}

	for _, info := range InfoSet {

		fmt.Println(info.Id, info.Name, info.VCpu, info.VMemory, info.VDisk, info.SpecsName, info.ImageName, info.ImageOsType)

		for _, port := range info.NIC {
			fmt.Println(port.Id, port.Name, port.VpcId, port.VpcName, port)

			for _, fixedIp := range port.FixedIp {
				fmt.Println(fixedIp)
			}

			for _, sg := range port.SecurityGroup {
				fmt.Println(sg)
			}

		}
		fmt.Println()
	}

	log.Println("裸金属共发起请求：", len(InfoSet))

	return InfoSet
}

func ExecuteALLMatch(contains string) []*Ironic {
	Body := utils.Get(Url, Headers)
	//fmt.Println(Body)

	InfoSet := JsonParseResourceInfoALL(Body)
	JsonParseResourceNetworkInfo(InfoSet)
	err := JsonParseTags(InfoSet)
	if err != nil {
		logrus.Warnln(err)
	}

	for _, info := range InfoSet {
		if !strings.HasPrefix(info.Name, contains) {
			continue
		}

		//fmt.Println(info.Id, info.Name, info.MacAddress, info.VCpu, info.VMemory, info.VDisk, info.SpecsName, info.ImageName, info.ImageOsType)

		//fmt.Println(info.Name, info.VpcName, info.SubnetName, info.PortName)

		for _, nic := range info.NIC {
			fmt.Printf("%s %s %s %s ", info.Name, nic.Name, nic.VpcName, nic.SubnetName)

			for _, fixedIp := range nic.FixedIp {
				if fixedIp.IpVersion == "4" {
					fmt.Printf("%s %s ", fixedIp.IpAddress, fixedIp.SubnetCidr)
				}
			}
			for _, fixedIp := range nic.FixedIp {
				if fixedIp.IpVersion == "6" {
					fmt.Printf("%s %s ", fixedIp.IpAddress, fixedIp.SubnetCidr)
				}
			}

			for _, securityGroup := range nic.SecurityGroup {
				fmt.Print(securityGroup.Name)
			}
			fmt.Println()
		}
		//fmt.Println()
	}

	fmt.Println("共发起请求：", len(InfoSet))

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

		JsonParseResourceNetworkInfoByPool(infoSet, pool)

		PrintEcsInfoByProjectID(infoSet, projectID)

	}

}

func ExecuteMatchByID() {
	uuids := utils.InputUUIDs()

	for _, pool := range bootstrap.PoolIDs {
		Body := utils.Get(Url, HeadersFun(pool))

		infoSet := JsonParseResourceInfoALL(Body)

		JoinPoolInfo(infoSet, pool)

		JsonParseResourceNetworkInfoByPool(infoSet, pool)

		PrintEcsInfoByUUIDs(infoSet, uuids)

	}

}

func PrintEcsInfoByUUIDs(ECSInfoSet []*Ironic, uuids []string) {
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

func PrintEcsInfoByProjectID(EcsInfoSet []*Ironic, projectID string) {
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
