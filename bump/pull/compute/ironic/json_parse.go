package ironic

import (
	"Catch/bump/internal/mapping"
	"Catch/bump/internal/utils"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"strconv"
	"sync"
)

func JsonParseResourceInfoALL(body string) []*Ironic {

	var sliceIronic []*Ironic
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.content|@pretty")

	//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		var ironic Ironic
		ironic.Id = value.Get("id").String()
		ironic.Name = value.Get("name").String()
		ironic.Region = mapping.AZConvert(value.Get("region").String())
		ironic.VCpu = value.Get("vcpu").String()
		memory := int(value.Get("vmemory").Int())
		ironic.VMemory = strconv.Itoa(memory / 1024)
		ironic.VDisk = value.Get("vdisk").String()
		ironic.BootVolumeType = mapping.DiskConvert(value.Get("bootVolumeType").String())
		ironic.ImageOsType = value.Get("imageOsType").String()
		ironic.ImageName = value.Get("imageName").String()
		ironic.SpecsName = value.Get("specsName").String()
		ironic.PrivateIp = value.Get("portDetail.#.privateIp").String()
		ironic.PortName = value.Get("portDetail.#.portName").String()
		ironic.VpcName = value.Get("portDetail.#.vpcName").String()
		ironic.SubnetName = value.Get("portDetail.#.subnetName").String()
		ironic.CreateTime = value.Get("createdTime").String()
		sliceIronic = append(sliceIronic, &ironic)
		//fmt.Println("test", cloudPort)
		return true
	})

	return sliceIronic
}
func JsonParseResourceNetworkInfo(ironics []*Ironic) {
	var wg sync.WaitGroup
	wg.Add(len(ironics))

	for _, ironic := range ironics {
		go func() {
			defer wg.Done()
			portUrl := PortUrl(ironic.Id)
			jsonBody := utils.Get(portUrl, Headers)

			contents := gjson.Get(jsonBody, "body.content")
			contents.ForEach(func(key, value gjson.Result) bool {
				var port NIC
				port.Id = value.Get("id").String()
				port.Name = value.Get("name").String()
				port.Type = value.Get("type").String()
				port.Region = value.Get("region").String()
				port.MacAddress = value.Get("macAddress").String()
				port.ResourceId = value.Get("resourceId").String()
				port.ResourceName = value.Get("resourceName").String()
				port.VpcId = value.Get("vpcId").String()
				port.VpcName = value.Get("vpcName").String()
				port.RouterId = value.Get("routerId").String()
				port.NetworkId = value.Get("networkId").String()
				port.SubnetName = value.Get("subnetName").String()

				//body.content.fixedIpResps
				fixedIpJsonBody := value.Get("fixedIpResps")
				fixedIpJsonBody.ForEach(func(key, value gjson.Result) bool {
					var fixedIp FixedIp
					fixedIp.PortId = value.Get("portId").String()
					fixedIp.PortName = value.Get("portName").String()
					fixedIp.IpVersion = value.Get("ipVersion").String()
					fixedIp.IpAddress = value.Get("ipAddress").String()
					fixedIp.SubnetCidr = value.Get("subnetCidr").String()
					fixedIp.SubnetId = value.Get("subnetId").String()
					fixedIp.SubnetName = value.Get("subnetName").String()
					fixedIp.ResourceId = value.Get("resourceId").String()
					fixedIp.ResourceName = value.Get("resourceName").String()
					fixedIp.VpcId = value.Get("vpcId").String()
					fixedIp.VpcName = value.Get("vpcName").String()
					fixedIp.RouterId = value.Get("routerId").String()

					port.FixedIp = append(port.FixedIp, fixedIp)
					return true
				})

				//body.content.securityGroupResps
				securityGroupJsonBody := value.Get("securityGroupResps")
				securityGroupJsonBody.ForEach(func(key, value gjson.Result) bool {
					var securityGroup SecurityGroup
					securityGroup.PortId = value.Get("portId").String()
					securityGroup.Id = value.Get("id").String()
					securityGroup.Name = value.Get("name").String()

					port.SecurityGroup = append(port.SecurityGroup, securityGroup)
					return true
				})

				ironic.NIC = append(ironic.NIC, &port)
				return true
			})
		}()
	}

	wg.Wait()
}

func JsonParseResourceByUUIDs(body string, uuids []string) []*Ironic {

	var sliceIronic []*Ironic
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.content|@pretty")

	for _, uuid := range uuids {

		//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
		//循环content下的所有元素
		JsonBody.ForEach(func(key, value gjson.Result) bool {
			var ironic Ironic

			if uuid != value.Get("id").String() {
				return true
			}

			ironic.Id = value.Get("id").String()
			ironic.Name = value.Get("name").String()
			ironic.Region = mapping.AZConvert(value.Get("region").String())
			ironic.VCpu = value.Get("vcpu").String()
			memory := int(value.Get("vmemory").Int())
			ironic.VMemory = strconv.Itoa(memory / 1024)
			ironic.VDisk = value.Get("vdisk").String()
			ironic.BootVolumeType = mapping.DiskConvert(value.Get("bootVolumeType").String())
			ironic.ImageOsType = value.Get("imageOsType").String()
			ironic.ImageName = value.Get("imageName").String()
			ironic.SpecsName = value.Get("specsName").String()
			ironic.PrivateIp = value.Get("portDetail.#.privateIp").String()
			ironic.PortName = value.Get("portDetail.#.portName").String()
			ironic.VpcName = value.Get("portDetail.#.vpcName").String()
			ironic.SubnetName = value.Get("portDetail.#.subnetName").String()

			sliceIronic = append(sliceIronic, &ironic)
			//fmt.Println("test", cloudPort)
			return false
		})
	}

	return sliceIronic
}

func JsonParseResourceNetworkInfoByPool(ironics []*Ironic, pool string) {
	var wg sync.WaitGroup
	wg.Add(len(ironics))

	for _, ironic := range ironics {
		go func() {
			defer wg.Done()
			portUrl := PortUrl(ironic.Id)
			jsonBody := utils.Get(portUrl, HeadersFun(pool))

			contents := gjson.Get(jsonBody, "body.content")
			contents.ForEach(func(key, value gjson.Result) bool {
				var port NIC
				port.Id = value.Get("id").String()
				port.Name = value.Get("name").String()
				port.Type = value.Get("type").String()
				port.Region = value.Get("region").String()
				port.MacAddress = value.Get("macAddress").String()
				port.ResourceId = value.Get("resourceId").String()
				port.ResourceName = value.Get("resourceName").String()
				port.VpcId = value.Get("vpcId").String()
				port.VpcName = value.Get("vpcName").String()
				port.RouterId = value.Get("routerId").String()
				port.NetworkId = value.Get("networkId").String()
				port.SubnetName = value.Get("subnetName").String()

				//body.content.fixedIpResps
				fixedIpJsonBody := value.Get("fixedIpResps")
				fixedIpJsonBody.ForEach(func(key, value gjson.Result) bool {
					var fixedIp FixedIp
					fixedIp.PortId = value.Get("portId").String()
					fixedIp.PortName = value.Get("portName").String()
					fixedIp.IpVersion = value.Get("ipVersion").String()
					fixedIp.IpAddress = value.Get("ipAddress").String()
					fixedIp.SubnetCidr = value.Get("subnetCidr").String()
					fixedIp.SubnetId = value.Get("subnetId").String()
					fixedIp.SubnetName = value.Get("subnetName").String()
					fixedIp.ResourceId = value.Get("resourceId").String()
					fixedIp.ResourceName = value.Get("resourceName").String()
					fixedIp.VpcId = value.Get("vpcId").String()
					fixedIp.VpcName = value.Get("vpcName").String()
					fixedIp.RouterId = value.Get("routerId").String()

					port.FixedIp = append(port.FixedIp, fixedIp)
					return true
				})

				//body.content.securityGroupResps
				securityGroupJsonBody := value.Get("securityGroupResps")
				securityGroupJsonBody.ForEach(func(key, value gjson.Result) bool {
					var securityGroup SecurityGroup
					securityGroup.PortId = value.Get("portId").String()
					securityGroup.Id = value.Get("id").String()
					securityGroup.Name = value.Get("name").String()

					port.SecurityGroup = append(port.SecurityGroup, securityGroup)
					return true
				})

				ironic.NIC = append(ironic.NIC, &port)
				return true
			})
		}()
	}

	wg.Wait()
}

func JsonParseTags(ironics []*Ironic) error {
	var tagError error
	var wg sync.WaitGroup
	wg.Add(len(ironics))

	for _, ironic := range ironics {
		utils.GoPool.Go(func() {
			defer wg.Done()
			jsonBody := utils.Get(TagUrl(ironic.Id), TagHeaders())

			body := gjson.Get(jsonBody, "body")
			body.ForEach(func(key, value gjson.Result) bool {
				switch value.Get("tagKey").String() {
				case "标签1":
					ironic.Tag1 = value.Get("tagValue").String()
				case "标签X":
					ironic.TagX = value.Get("tagValue").String()
				case "标签2":
					ironic.Tag2 = value.Get("tagValue").String()
				case "订购人":
					ironic.Orderer = value.Get("tagValue").String()
				}
				return true

				/*				switch value.Get("tagKey").String() {
								case "标签1":
									ironic.Tag1 = value.Get("tagValue").String()
									logrus.Warnln("tag1=", value.Get("tagKey").String())
								case "标签X":
									ironic.TagX = value.Get("tagValue").String()
									logrus.Warnln("tagX=", value.Get("tagKey").String())
								case "标签2":
									ironic.Tag2 = value.Get("tagValue").String()
									logrus.Warnln("tag2=", value.Get("tagKey").String())
								case "订购人":
									ironic.Orderer = value.Get("tagValue").String()
									logrus.Warnln("Orderer=", value.Get("tagKey").String())
								}
								return true*/
			})

			if ironic.Tag1 == "" || ironic.Tag2 == "" || ironic.TagX == "" || ironic.Orderer == "" {
				if tagError == nil {
					tagError = errors.New(fmt.Sprintf("未找到标签，id='%s' name='%s' ,标签1='%s' 标签X='%s' 标签2='%s' 订购人='%s'\n", ironic.Id, ironic.Name, ironic.Tag1, ironic.TagX, ironic.Tag2, ironic.Orderer))
				} else {
					tagError = errors.New(tagError.Error() + fmt.Sprintf("未找到标签，id='%s' name='%s' ,标签1='%s' 标签X='%s' 标签2='%s' 订购人='%s'\n", ironic.Id, ironic.Name, ironic.Tag1, ironic.TagX, ironic.Tag2, ironic.Orderer))
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
