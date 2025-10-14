package get

import (
	"Catch/bump/internal/mapping"
	"log"
	"strconv"
)

type ComputeResource struct {
	Row                 string
	Tag1                string
	TagX                string
	Tag2                string
	Orderer             string
	ResourceID          string
	ResourceName        string
	ResourceCategory    string
	ResourceSubCategory string
	Specification       string
	CPU                 string
	Memory              string
	InternalBandwidth   string
	Image               string
	CloudPool           string
	Node                string
	AvailabilityZone    string
	SystemDiskSize      string
	SystemDiskSpec      string
	BillingMethod string
	NIC1          Nic
	NIC2          Nic
	NIC3          Nic
}

type Nic struct {
	Name               string
	MAC                string
	VPC                string
	Subnet             string
	IPv4PrivateAddress string
	IPv4PrivateMask    string
	IPv6Address        string
	IPv6Mask           string
	SecurityGroup      string
}

// ParseComputeResourceToStruct
// 将excel表格中的每一行数据转成对应的结构体后加入到切片中
func (e *Excel) ParseComputeResourceToStruct() []ComputeResource {
	rows, err := e.File.GetRows(e.Sheet)
	if err != nil {
		log.Printf("filename=%s sheet=%s get rows error:%s", e.FileName, e.Sheet, err.Error())
	}

	//行内容长度不统一，填充为一样的长度
	newRows := PadField(e.TitleLine, rows)

	//获取每个标题的index
	tag1Index := e.CaptureColNumber(Tag1, "n")
	tagXIndex := e.CaptureColNumber(TagX, "n")
	tag2Index := e.CaptureColNumber(Tag2, "n")
	resourceIDIndex := e.CaptureColNumber(ResourceID, "y")
	resourceNameIndex := e.CaptureColNumber(ResourceName, "y")
	resourceCategoryIndex := e.CaptureColNumber(ResourceCategory, "y")
	resourceSubCategoryIndex := e.CaptureColNumber(ResourceSubCategory, "y")
	specificationIndex := e.CaptureColNumber(Specification, "y")
	cpuIndex := e.CaptureColNumber(CPU, "y")
	memoryIndex := e.CaptureColNumber(Memory, "y")
	internalBandwidthIndex := e.CaptureColNumber(InternalBandwidth, "y")
	imageIndex := e.CaptureColNumber(Image, "y")
	cloudPoolIndex := e.CaptureColNumber(CloudPool, "y")
	nodeIndex := e.CaptureColNumber(Node, "y")
	availabilityZoneIndex := e.CaptureColNumber(AvailabilityZone, "y")
	systemDiskSizeIndex := e.CaptureColNumber(SystemDiskSize, "y")
	systemDiskSpecIndex := e.CaptureColNumber(SystemDiskSpec, "y")
	billingMethodIndex := e.CaptureColNumber(BillingMethod, "y")
	nicNameIndex := e.CaptureColNumber(NICName, "y")
	nicMACIndex := e.CaptureColNumber(NICMAC, "y")
	nicVPCIndex := e.CaptureColNumber(NICVPC, "y")
	nicSubnetIndex := e.CaptureColNumber(NICSubnet, "y")
	nicIPv4PrivateAddressIndex := e.CaptureColNumber(NICIPv4PrivateAddress, "y")
	nicIPv4PrivateMaskIndex := e.CaptureColNumber(NICIPv4PrivateMask, "y")
	nicIPv6AddressIndex := e.CaptureColNumber(NICIPv6Address, "y")
	nicIPv6MaskIndex := e.CaptureColNumber(NICIPv6Mask, "y")
	nicSecurityGroupIndex := e.CaptureColNumber(NICSecurityGroup, "y")

	//将每行数据解析成一个结构体，存放在切片中
	var crSlice []ComputeResource
	for i, row := range newRows {
		//从那行开始复制，跳过前n行
		if i+1 < e.CopyLine {
			continue
		}

		if GetFieldContent(row, resourceSubCategoryIndex) == "" {
			continue
		}

		var cr ComputeResource
		cr.Row = strconv.Itoa(i + 1)
		cr.Tag1 = GetFieldContent(row, tag1Index)
		cr.TagX = GetFieldContent(row, tagXIndex)
		cr.Tag2 = GetFieldContent(row, tag2Index)
		cr.ResourceID = GetFieldContent(row, resourceIDIndex)
		cr.ResourceName = GetFieldContent(row, resourceNameIndex)
		cr.ResourceCategory = GetFieldContent(row, resourceCategoryIndex)
		cr.ResourceSubCategory = GetFieldContent(row, resourceSubCategoryIndex)
		cr.Specification = GetFieldContent(row, specificationIndex)
		cr.CPU = GetFieldContent(row, cpuIndex)
		cr.Memory = GetFieldContent(row, memoryIndex)
		cr.InternalBandwidth = GetFieldContent(row, internalBandwidthIndex)
		cr.Image = GetFieldContent(row, imageIndex)
		cr.CloudPool = GetFieldContent(row, cloudPoolIndex)
		cr.Node = GetFieldContent(row, nodeIndex)
		cr.AvailabilityZone = GetFieldContent(row, availabilityZoneIndex)
		cr.SystemDiskSize = GetFieldContent(row, systemDiskSizeIndex)
		cr.SystemDiskSpec = GetFieldContent(row, systemDiskSpecIndex)
		cr.BillingMethod = GetFieldContent(row, billingMethodIndex)

		// 填充网卡1信息
		cr.NIC1.Name = GetFieldContent(row, nicNameIndex)
		cr.NIC1.MAC = GetFieldContent(row, nicMACIndex)
		cr.NIC1.VPC = GetFieldContent(row, nicVPCIndex)
		cr.NIC1.Subnet = GetFieldContent(row, nicSubnetIndex)
		cr.NIC1.IPv4PrivateAddress = GetFieldContent(row, nicIPv4PrivateAddressIndex)
		cr.NIC1.IPv4PrivateMask = GetFieldContent(row, nicIPv4PrivateMaskIndex)
		cr.NIC1.IPv6Address = GetFieldContent(row, nicIPv6AddressIndex)
		cr.NIC1.IPv6Mask = GetFieldContent(row, nicIPv6MaskIndex)
		cr.NIC1.SecurityGroup = GetFieldContent(row, nicSecurityGroupIndex)

		// 填充网卡2信息
		cr.NIC2.Name = GetFieldContent(row, nicNameIndex+9)
		cr.NIC2.MAC = GetFieldContent(row, nicMACIndex+9)
		cr.NIC2.VPC = GetFieldContent(row, nicVPCIndex+9)
		cr.NIC2.Subnet = GetFieldContent(row, nicSubnetIndex+9)
		cr.NIC2.IPv4PrivateAddress = GetFieldContent(row, nicIPv4PrivateAddressIndex+9)
		cr.NIC2.IPv4PrivateMask = GetFieldContent(row, nicIPv4PrivateMaskIndex+9)
		cr.NIC2.IPv6Address = GetFieldContent(row, nicIPv6AddressIndex+9)
		cr.NIC2.IPv6Mask = GetFieldContent(row, nicIPv6MaskIndex+9)
		cr.NIC2.SecurityGroup = GetFieldContent(row, nicSecurityGroupIndex+9)

		// 填充网卡3信息
		cr.NIC3.Name = GetFieldContent(row, nicNameIndex+18)
		cr.NIC3.MAC = GetFieldContent(row, nicMACIndex+18)
		cr.NIC3.VPC = GetFieldContent(row, nicVPCIndex+18)
		cr.NIC3.Subnet = GetFieldContent(row, nicSubnetIndex+18)
		cr.NIC3.IPv4PrivateAddress = GetFieldContent(row, nicIPv4PrivateAddressIndex+18)
		cr.NIC3.IPv4PrivateMask = GetFieldContent(row, nicIPv4PrivateMaskIndex+18)
		cr.NIC3.IPv6Address = GetFieldContent(row, nicIPv6AddressIndex+18)
		cr.NIC3.IPv6Mask = GetFieldContent(row, nicIPv6MaskIndex+18)
		cr.NIC3.SecurityGroup = GetFieldContent(row, nicSecurityGroupIndex+18)

		crSlice = append(crSlice, cr)
	}
	return crSlice
}

func ComputeGetUUIDs(excelResources []ComputeResource) []string {
	var strSlice []string
	for _, excelResource := range excelResources {
		strSlice = append(strSlice, excelResource.ResourceID)
	}
	return strSlice
}

func ComputeGetPool(excelResources []ComputeResource) []string {
	var tuple = make(map[string]struct{})
	var pools []string
	for _, e := range excelResources {
		if _, exist := tuple[e.Node]; !exist {
			tuple[e.Node] = struct{}{}
		}
	}

	for k, _ := range tuple {
		pools = append(pools, mapping.PoolNameToCodeConvert(k))
	}
	return pools
}
