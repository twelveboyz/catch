package ecs

import (
	"Catch/bump/internal/mapping"
	"Catch/bump/internal/utils"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"strconv"
	"sync"
)

func JsonParseResourceInfoALL(body string) []*ECloudServer {

	var sliceECloudServer []*ECloudServer
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.content|@pretty")

	//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		var eCloudServer ECloudServer
		eCloudServer.Id = value.Get("id").String()
		eCloudServer.Name = value.Get("name").String()
		eCloudServer.Region = mapping.AZConvert(value.Get("region").String())
		eCloudServer.VCpu = value.Get("vcpu").String()
		eCloudServer.VMemory = strconv.Itoa(int(value.Get("vmemory").Int()) / 1024)
		eCloudServer.VDisk = value.Get("vdisk").String()
		eCloudServer.BootVolumeType = mapping.DiskConvert(value.Get("bootVolumeType").String())
		eCloudServer.ImageOsType = value.Get("imageOsType").String()
		eCloudServer.ImageName = value.Get("imageName").String()
		eCloudServer.SpecsName = value.Get("specsName").String()
		eCloudServer.CreateTime = value.Get("createdTime").String()

		if eCloudServer.Id != "" {
			sliceECloudServer = append(sliceECloudServer, &eCloudServer)
		}
		//fmt.Println("test", cloudPort)
		return true
	})

	return sliceECloudServer
}

func JsonParseResourceByUUIDs(body string, uuids []string) []*ECloudServer {

	var sliceECloudServer []*ECloudServer
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.content|@pretty")

	for _, uuid := range uuids {

		//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
		//循环content下的所有元素
		JsonBody.ForEach(func(key, value gjson.Result) bool {
			var eCloudServer ECloudServer

			if uuid != value.Get("id").String() {
				return true
			}

			eCloudServer.Id = value.Get("id").String()
			eCloudServer.Name = value.Get("name").String()
			eCloudServer.Region = mapping.AZConvert(value.Get("region").String())
			eCloudServer.VCpu = value.Get("vcpu").String()
			eCloudServer.VMemory = strconv.Itoa(int(value.Get("vmemory").Int()) / 1024)
			eCloudServer.VDisk = value.Get("vdisk").String()
			eCloudServer.BootVolumeType = mapping.DiskConvert(value.Get("bootVolumeType").String())
			eCloudServer.ImageOsType = value.Get("imageOsType").String()
			eCloudServer.ImageName = value.Get("imageName").String()
			eCloudServer.SpecsName = value.Get("specsName").String()

			sliceECloudServer = append(sliceECloudServer, &eCloudServer)
			//fmt.Println("test", cloudPort)
			return false
		})

	}
	return sliceECloudServer
}

func JsonParseResourceNetworkInfoByPools(ecss []*ECloudServer, pool string) {
	var wg sync.WaitGroup
	wg.Add(len(ecss))

	for _, ecs := range ecss {
		utils.GoPool.Go(func() {
			defer wg.Done()
			portUrl := PortUrl(ecs.Id)
			jsonBody := utils.Get(portUrl, HeadersFun(pool))

			contents := gjson.Get(jsonBody, "body.content")
			contents.ForEach(func(key, value gjson.Result) bool {
				var nic NIC
				nic.Id = value.Get("id").String()
				nic.Name = value.Get("name").String()
				nic.MacAddress = value.Get("macAddress").String()
				nic.VpcId = value.Get("vpcId").String()
				nic.VpcName = value.Get("vpcName").String()
				nic.SubnetName = value.Get("subnetName").String()
				nic.NetworkId = value.Get("networkId").String()
				nic.ResourceId = value.Get("resourceId").String()
				nic.Type = value.Get("type").String()
				nic.Region = value.Get("region").String()

				//body.content.fixedIpResps
				fixedIpJsonBody := value.Get("fixedIpResps")
				fixedIpJsonBody.ForEach(func(key, value gjson.Result) bool {
					var fixedIp FixedIp
					fixedIp.IpVersion = value.Get("ipVersion").String()
					fixedIp.PortId = value.Get("portId").String()
					fixedIp.IpAddress = value.Get("ipAddress").String()
					fixedIp.VpcId = value.Get("vpcId").String()
					fixedIp.VpcName = value.Get("vpcName").String()
					fixedIp.RouterId = value.Get("routerId").String()
					fixedIp.SubnetId = value.Get("subnetId").String()
					fixedIp.SubnetName = value.Get("subnetName").String()
					fixedIp.SubnetCidr = value.Get("subnetCidr").String()
					fixedIp.PortName = value.Get("portName").String()

					nic.FixedIp = append(nic.FixedIp, fixedIp)
					return true
				})

				//body.content.securityGroupResps
				securityGroupJsonBody := value.Get("securityGroupResps")
				securityGroupJsonBody.ForEach(func(key, value gjson.Result) bool {
					var securityGroup SecurityGroup
					securityGroup.PortId = value.Get("portId").String()
					securityGroup.Id = value.Get("id").String()
					securityGroup.Name = value.Get("name").String()

					nic.SecurityGroup = append(nic.SecurityGroup, securityGroup)
					return true
				})

				ecs.NIC = append(ecs.NIC, &nic)
				return true
			})
		})
	}

	wg.Wait()
}

func JsonParseTags(ecss []*ECloudServer) error {
	var tagError error
	var wg sync.WaitGroup
	wg.Add(len(ecss))

	for _, ecs := range ecss {
		utils.GoPool.Go(func() {
			defer wg.Done()
			jsonBody := utils.Get(TagUrl(ecs.Id), TagHeaders())

			body := gjson.Get(jsonBody, "body")
			body.ForEach(func(key, value gjson.Result) bool {
				switch value.Get("tagKey").String() {
				case "标签1":
					ecs.Tag1 = value.Get("tagValue").String()
				case "标签X":
					ecs.TagX = value.Get("tagValue").String()
				case "标签2":
					ecs.Tag2 = value.Get("tagValue").String()
				case "订购人":
					ecs.Orderer = value.Get("tagValue").String()
				}
				return true
			})

			/*		body.ForEach(func(key, value gjson.Result) bool {
					logrus.Warnln("key=", value.Get("tagKey").String())
					switch value.Get("tagKey").String() {
					case "标签1":
						ecs.Tag1 = value.Get("tagValue").String()
						logrus.Warnln("tag1=", value.Get("tagKey").String())
					case "标签X":
						ecs.TagX = value.Get("tagValue").String()
						logrus.Warnln("tagX=", value.Get("tagKey").String())
					case "标签2":
						ecs.Tag2 = value.Get("tagValue").String()
						logrus.Warnln("tag2=", value.Get("tagKey").String())
					case "订购人":
						ecs.Orderer = value.Get("tagValue").String()
						logrus.Warnln("Orderer=", value.Get("tagKey").String())
					}
					return true
				})*/

			if ecs.Tag1 == "" || ecs.Tag2 == "" || ecs.TagX == "" || ecs.Orderer == "" {
				if tagError == nil {
					tagError = errors.New(fmt.Sprintf("未找到标签，id='%s' name='%s' ,标签1='%s' 标签X='%s' 标签2='%s' 订购人='%s'\n", ecs.Id, ecs.Name, ecs.Tag1, ecs.TagX, ecs.Tag2, ecs.Orderer))
				} else {
					tagError = errors.New(tagError.Error() + fmt.Sprintf("未找到标签，id='%s' name='%s' ,标签1='%s' 标签X='%s' 标签2='%s' 订购人='%s'\n", ecs.Id, ecs.Name, ecs.Tag1, ecs.TagX, ecs.Tag2, ecs.Orderer))
				}
			}
		})
	}

	wg.Wait()

	if tagError != nil {
		return tagError
	}

	return nil
}
