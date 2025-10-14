package get

import (
	"Catch/bump/internal/mapping"
	"errors"
	"log"
	"strconv"
	"strings"
)

type NetworkResource struct {
	LoadBalance          []LoadBalanceResource
	Eip                  []EipResource
	SharedTrafficPackage []SharedTrafficPackageResource
	IPv6                 []IPv6Resource
	DirectConnect        []DirectConnectResource
	NATGateway           []NATGatewayResource
	CloudPort            []CloudPortResource
	Other                []OtherResource
}

type LoadBalanceResource struct {
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
	IPv4PrivateAddress  string
	IPv6Address         string
	IPv4PublicAddress   string
	VPC                 string
	Subnet              string
	CloudPool           string
	Node                string
}

type EipResource struct {
	Row                    string
	Tag1                   string
	TagX                   string
	Tag2                   string
	Orderer                string
	ResourceID             string
	ResourceName           string
	ResourceCategory       string
	ResourceSubCategory    string
	IPv4PublicAddress      string
	IPv6Address            string
	BandWidth              string
	TrafficPacketDeduction string
	BindingResourceType    string
	BindingResourceName    string
	BindingResourceID      string
	CloudPool              string
	Node                   string
}

type SharedTrafficPackageResource struct {
	Row                 string
	Tag1                string
	TagX                string
	Tag2                string
	Orderer             string
	ResourceID          string
	ResourceName        string
	ResourceCategory    string
	ResourceSubCategory string
	TotalFlow           string
	Type                string
	DeductedIP          string
	CloudPool           string
	Node                string
}

type IPv6Resource struct {
	Row                 string
	Tag1                string
	TagX                string
	Tag2                string
	Orderer             string
	ResourceID          string
	ResourceName        string
	ResourceCategory    string
	ResourceSubCategory string
	IPv6Address         string
	BandWidth           string
	RelatedResources    string
	CloudPool           string
	Node                string
}

type DirectConnectResource struct {
	Row                      string
	Tag1                     string
	TagX                     string
	Tag2                     string
	Orderer                  string
	ResourceID               string
	ResourceName             string
	ResourceCategory         string
	ResourceSubCategory      string
	VPC                      string
	VPCSubnet                string
	UserSubnet               string
	AccessNodeRegion         string
	AccessNodeAddress        string
	AccessNodeContact        string
	AccessNodeAccountManager string
	BandWidth                string
	CloudPool                string
	Node                     string
}

type NATGatewayResource struct {
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
	VPC                 string
	BandWidth           string
	CloudPool           string
	Node                string
}

type CloudPortResource struct {
	Row                               string
	Tag1                              string
	TagX                              string
	Tag2                              string
	Orderer                           string
	ResourceID                        string //uuid
	ResourceName                      string //name
	ResourceCategory                  string //大类
	ResourceSubCategory               string //小类
	Specification                     string //规格
	BandWidth                         string //带宽
	IPType                            string //ip类型
	VPC                               string //vpcid
	AccessNodeRegion                  string //接入节点地域
	AccessNodeAddress                 string //接入节点详细地址
	AccessNodeContact                 string //接入节点联系人
	AccountManager                    string //客户经理
	VGWAndAggDockingIPMaster          string //vgw与汇聚设备对接IP(主）
	VGWAndAggDockingIPBackup          string //vgw与汇聚设备对接IP(备）
	AggAndVGWDockingIPMaster          string //汇聚设备与vgw对接IP(主）
	AggAndVGWDockingIPBackup          string //汇聚设备与vgw对接IP(备）
	AggAndVGWDockingVLANMaster        string //汇聚设备与vgw对接VLAN（主）
	AggAndVGWDockingVLANBackup        string //汇聚设备与vgw对接VLAN（备）
	AggAndCustomerDockingLogicalPort  string //汇聚设备与客户设备对接逻辑端口
	AggAndCustomerDockingPhysicalPort string //汇聚设备与客户设备对接物理端口
	Device1                           string //设备1位置
	Device2                           string //设备2位置
	CloudPool                         string //云池
	ResourcePool                      string //资源池
	Node                              string //节点
}

type OtherResource struct {
	Row                      string
	Tag1                     string
	TagX                     string
	Tag2                     string
	Orderer                  string
	ResourceID               string
	ResourceName             string
	ResourceCategory         string
	ResourceSubCategory      string
	Specification            string
	IPv4PrivateAddress       string
	Ipv6Address              string
	BandWidth                string
	VPC                      string
	NodeIP                   string
	EndpointNodeServiceName  string
	NetworkProductInstanceId string
	BindingResourceType      string
	BindingResourceName      string
	BindingResourceID        string
	CloudPool                string
	Node                     string
}

func (e *Excel) ParseNetworkResourceToStruct() (NetworkResource, error) {
	var totalNetworkResource NetworkResource
	if strings.Contains(e.Sheet, "负载均衡") {
		totalNetworkResource.LoadBalance = e.ParseLBToStruct()

	} else if strings.Contains(e.Sheet, "公网IP") {
		totalNetworkResource.Eip = e.ParseEipToStruct()

	} else if strings.Contains(e.Sheet, "流量包") {
		totalNetworkResource.SharedTrafficPackage = e.ParseSharedTrafficPackageToStruct()

	} else if strings.Contains(e.Sheet, "IPv6") {
		totalNetworkResource.IPv6 = e.ParseIPv6ToStruct()

	} else if strings.Contains(e.Sheet, "云专线") {
		totalNetworkResource.DirectConnect = e.ParseDirectConnectToStruct()

	} else if strings.Contains(e.Sheet, "NAT网关") {
		totalNetworkResource.NATGateway = e.ParseNATGatewayToStruct()

	} else if strings.Contains(e.Sheet, "云端口") {
		totalNetworkResource.CloudPort = e.ParseCloudPortToStruct()

	} else if strings.Contains(e.Sheet, "其他") {
		totalNetworkResource.Other = e.ParseOtherResourceToStruct()

	} else {
		return totalNetworkResource, errors.New("未匹配到sheet=" + e.Sheet)
	}

	return totalNetworkResource, nil
}

// ParseLBToStruct
// 将excel表格中的每一行数据转成对应的结构体后加入到切片中
func (e *Excel) ParseLBToStruct() []LoadBalanceResource {
	rows, err := e.File.GetRows(e.Sheet)
	if err != nil {
		log.Printf("filename=%s sheet=%s get rows error:%s", e.FileName, e.Sheet, err.Error())
		return nil
	}

	// 行内容长度不统一，填充为一样的长度
	newRows := PadField(e.TitleLine, rows)

	// 获取各字段的列索引
	tag1Index := e.CaptureColNumber(Tag1, "n")
	tagXIndex := e.CaptureColNumber(TagX, "n")
	tag2Index := e.CaptureColNumber(Tag2, "n")
	resourceIDIndex := e.CaptureColNumber(ResourceID, "y")
	resourceNameIndex := e.CaptureColNumber(ResourceName, "y")
	resourceCategoryIndex := e.CaptureColNumber(ResourceCategory, "y")
	resourceSubCategoryIndex := e.CaptureColNumber(ResourceSubCategory, "y")
	specificationIndex := e.CaptureColNumber(Specification, "y")
	ipv4PrivateAddressIndex := e.CaptureColNumber(IPv4PrivateAddress, "y")
	ipv6AddressIndex := e.CaptureColNumber(IPv6Address, "y")
	ipv4PublicAddressIndex := e.CaptureColNumber(PublicIPv4Address, "y")
	vpcIndex := e.CaptureColNumber(VPC, "y")
	subnetIndex := e.CaptureColNumber(Subnet, "y")
	cloudPoolIndex := e.CaptureColNumber(CloudPool, "y")
	nodeIndex := e.CaptureColNumber(Node, "n")

	var lbSlice []LoadBalanceResource
	for i, row := range newRows {
		// 跳过前 n 行
		if i+1 < e.CopyLine {
			continue
		}

		// 初始化结构体并填充字段
		var lb LoadBalanceResource
		lb.Row = strconv.Itoa(i + 1)
		lb.Tag1 = GetFieldContent(row, tag1Index)
		lb.TagX = GetFieldContent(row, tagXIndex)
		lb.Tag2 = GetFieldContent(row, tag2Index)
		lb.ResourceID = GetFieldContent(row, resourceIDIndex)
		lb.ResourceName = GetFieldContent(row, resourceNameIndex)
		lb.ResourceCategory = GetFieldContent(row, resourceCategoryIndex)
		lb.ResourceSubCategory = GetFieldContent(row, resourceSubCategoryIndex)
		lb.Specification = GetFieldContent(row, specificationIndex)
		lb.IPv4PrivateAddress = GetFieldContent(row, ipv4PrivateAddressIndex)
		lb.IPv6Address = GetFieldContent(row, ipv6AddressIndex)
		lb.IPv4PublicAddress = GetFieldContent(row, ipv4PublicAddressIndex)
		lb.VPC = GetFieldContent(row, vpcIndex)
		lb.Subnet = GetFieldContent(row, subnetIndex)
		lb.CloudPool = GetFieldContent(row, cloudPoolIndex)
		lb.Node = GetFieldContent(row, nodeIndex)

		// 将填充好的结构体添加到切片中
		lbSlice = append(lbSlice, lb)
	}

	return lbSlice
}

// ParseEipToStruct
// 将excel表格中的每一行数据转成对应的结构体后加入到切片中
func (e *Excel) ParseEipToStruct() []EipResource {
	rows, err := e.File.GetRows(e.Sheet)
	if err != nil {
		log.Printf("filename=%s sheet=%s get rows error:%s", e.FileName, e.Sheet, err.Error())
		return nil
	}

	// 行内容长度不统一，填充为一样的长度
	newRows := PadField(e.TitleLine, rows)

	// 获取各字段的列索引
	tag1Index := e.CaptureColNumber(Tag1, "n")
	tagXIndex := e.CaptureColNumber(TagX, "n")
	tag2Index := e.CaptureColNumber(Tag2, "n")
	resourceIDIndex := e.CaptureColNumber(ResourceID, "y")
	resourceNameIndex := e.CaptureColNumber(ResourceName, "y")
	resourceCategoryIndex := e.CaptureColNumber(ResourceCategory, "y")
	resourceSubCategoryIndex := e.CaptureColNumber(ResourceSubCategory, "y")
	ipv4PublicAddressIndex := e.CaptureColNumber(IPv4PublicAddress, "y")
	ipv6AddressIndex := e.CaptureColNumber(IPv6Address, "y")
	bandwidthIndex := e.CaptureColNumber(Bandwidth, "y")
	trafficPacketDeductionIndex := e.CaptureColNumber(TrafficPacketDeduction, "y")
	bindingResourceTypeIndex := e.CaptureColNumber(BindingResourceType, "y")
	bindingResourceNameIndex := e.CaptureColNumber(BindingResourceName, "y")
	bindingResourceIDIndex := e.CaptureColNumber(BindingResourceID, "y")
	cloudPoolIndex := e.CaptureColNumber(CloudPool, "y")
	nodeIndex := e.CaptureColNumber(Node, "n")

	var eipSlice []EipResource
	for i, row := range newRows {
		// 跳过前 n 行
		if i+1 < e.CopyLine {
			continue
		}

		// 初始化结构体并填充字段
		var eip EipResource
		eip.Row = strconv.Itoa(i + 1)
		eip.Tag1 = GetFieldContent(row, tag1Index)
		eip.TagX = GetFieldContent(row, tagXIndex)
		eip.Tag2 = GetFieldContent(row, tag2Index)
		eip.ResourceID = GetFieldContent(row, resourceIDIndex)
		eip.ResourceName = GetFieldContent(row, resourceNameIndex)
		eip.ResourceCategory = GetFieldContent(row, resourceCategoryIndex)
		eip.ResourceSubCategory = GetFieldContent(row, resourceSubCategoryIndex)
		eip.IPv4PublicAddress = GetFieldContent(row, ipv4PublicAddressIndex)
		eip.IPv6Address = GetFieldContent(row, ipv6AddressIndex)
		eip.BandWidth = GetFieldContent(row, bandwidthIndex)
		eip.TrafficPacketDeduction = GetFieldContent(row, trafficPacketDeductionIndex)
		eip.BindingResourceType = GetFieldContent(row, bindingResourceTypeIndex)
		eip.BindingResourceName = GetFieldContent(row, bindingResourceNameIndex)
		eip.BindingResourceID = GetFieldContent(row, bindingResourceIDIndex)
		eip.CloudPool = GetFieldContent(row, cloudPoolIndex)
		eip.Node = GetFieldContent(row, nodeIndex)

		// 将填充好的结构体添加到切片中
		eipSlice = append(eipSlice, eip)
	}

	return eipSlice
}

// ParseSharedTrafficPackageToStruct
// 将excel表格中的每一行数据转成对应的结构体后加入到切片中
func (e *Excel) ParseSharedTrafficPackageToStruct() []SharedTrafficPackageResource {
	rows, err := e.File.GetRows(e.Sheet)
	if err != nil {
		log.Printf("filename=%s sheet=%s get rows error:%s", e.FileName, e.Sheet, err.Error())
		return nil
	}

	// 行内容长度不统一，填充为一样的长度
	newRows := PadField(e.TitleLine, rows)

	// 获取各字段的列索引
	tag1Index := e.CaptureColNumber(Tag1, "n")
	tagXIndex := e.CaptureColNumber(TagX, "n")
	tag2Index := e.CaptureColNumber(Tag2, "n")
	resourceIDIndex := e.CaptureColNumber(ResourceID, "y")
	resourceNameIndex := e.CaptureColNumber(ResourceName, "y")
	resourceCategoryIndex := e.CaptureColNumber(ResourceCategory, "y")
	resourceSubCategoryIndex := e.CaptureColNumber(ResourceSubCategory, "y")
	totalFlowIndex := e.CaptureColNumber(totalFlow, "y")
	typeIndex := e.CaptureColNumber(Type, "y")
	deductedIPIndex := e.CaptureColNumber(DeductedIP, "y")
	cloudPoolIndex := e.CaptureColNumber(CloudPool, "y")
	nodeIndex := e.CaptureColNumber(Node, "n")

	var stpSlice []SharedTrafficPackageResource
	for i, row := range newRows {
		// 跳过前 n 行
		if i+1 < e.CopyLine {
			continue
		}

		// 初始化结构体并填充字段
		var stp SharedTrafficPackageResource
		stp.Row = strconv.Itoa(i + 1)
		stp.Tag1 = GetFieldContent(row, tag1Index)
		stp.TagX = GetFieldContent(row, tagXIndex)
		stp.Tag2 = GetFieldContent(row, tag2Index)
		stp.ResourceID = GetFieldContent(row, resourceIDIndex)
		stp.ResourceName = GetFieldContent(row, resourceNameIndex)
		stp.ResourceCategory = GetFieldContent(row, resourceCategoryIndex)
		stp.ResourceSubCategory = GetFieldContent(row, resourceSubCategoryIndex)
		stp.TotalFlow = GetFieldContent(row, totalFlowIndex)
		stp.Type = GetFieldContent(row, typeIndex)
		stp.DeductedIP = GetFieldContent(row, deductedIPIndex)
		stp.CloudPool = GetFieldContent(row, cloudPoolIndex)
		stp.Node = GetFieldContent(row, nodeIndex)

		// 将填充好的结构体添加到切片中
		stpSlice = append(stpSlice, stp)
	}

	return stpSlice
}

// ParseIPv6ToStruct
// 将excel表格中的每一行数据转成对应的结构体后加入到切片中
func (e *Excel) ParseIPv6ToStruct() []IPv6Resource {
	rows, err := e.File.GetRows(e.Sheet)
	if err != nil {
		log.Printf("filename=%s sheet=%s get rows error:%s", e.FileName, e.Sheet, err.Error())
		return nil
	}

	// 行内容长度不统一，填充为一样的长度
	newRows := PadField(e.TitleLine, rows)

	// 获取各字段的列索引
	tag1Index := e.CaptureColNumber(Tag1, "n")
	tagXIndex := e.CaptureColNumber(TagX, "n")
	tag2Index := e.CaptureColNumber(Tag2, "n")
	resourceIDIndex := e.CaptureColNumber(ResourceID, "y")
	resourceNameIndex := e.CaptureColNumber(ResourceName, "y")
	resourceCategoryIndex := e.CaptureColNumber(ResourceCategory, "y")
	resourceSubCategoryIndex := e.CaptureColNumber(ResourceSubCategory, "y")
	ipv6AddressIndex := e.CaptureColNumber(IPv6Address, "y")
	bandwidthIndex := e.CaptureColNumber(Bandwidth, "y")
	relatedResourcesIndex := e.CaptureColNumber(RelatedResources, "y")
	cloudPoolIndex := e.CaptureColNumber(CloudPool, "y")
	nodeIndex := e.CaptureColNumber(Node, "n")

	var ipv6Slice []IPv6Resource
	for i, row := range newRows {
		// 跳过前 n 行
		if i+1 < e.CopyLine {
			continue
		}

		// 初始化结构体并填充字段
		var ipv6 IPv6Resource
		ipv6.Row = strconv.Itoa(i + 1)
		ipv6.Tag1 = GetFieldContent(row, tag1Index)
		ipv6.TagX = GetFieldContent(row, tagXIndex)
		ipv6.Tag2 = GetFieldContent(row, tag2Index)
		ipv6.ResourceID = GetFieldContent(row, resourceIDIndex)
		ipv6.ResourceName = GetFieldContent(row, resourceNameIndex)
		ipv6.ResourceCategory = GetFieldContent(row, resourceCategoryIndex)
		ipv6.ResourceSubCategory = GetFieldContent(row, resourceSubCategoryIndex)
		ipv6.IPv6Address = GetFieldContent(row, ipv6AddressIndex)
		ipv6.BandWidth = GetFieldContent(row, bandwidthIndex)
		ipv6.RelatedResources = GetFieldContent(row, relatedResourcesIndex)
		ipv6.CloudPool = GetFieldContent(row, cloudPoolIndex)
		ipv6.Node = GetFieldContent(row, nodeIndex)

		// 将填充好的结构体添加到切片中
		ipv6Slice = append(ipv6Slice, ipv6)
	}

	return ipv6Slice
}

// ParseDirectConnectToStruct
// 将excel表格中的每一行数据转成对应的结构体后加入到切片中
func (e *Excel) ParseDirectConnectToStruct() []DirectConnectResource {
	rows, err := e.File.GetRows(e.Sheet)
	if err != nil {
		log.Printf("filename=%s sheet=%s get rows error:%s", e.FileName, e.Sheet, err.Error())
		return nil
	}
	// 行内容长度不统一，填充为一样的长度
	newRows := PadField(e.TitleLine, rows)
	// 获取各字段的列索引
	tag1Index := e.CaptureColNumber(Tag1, "n")
	tagXIndex := e.CaptureColNumber(TagX, "n")
	tag2Index := e.CaptureColNumber(Tag2, "n")
	resourceIDIndex := e.CaptureColNumber(ResourceID, "y")
	resourceNameIndex := e.CaptureColNumber(ResourceName, "y")
	resourceCategoryIndex := e.CaptureColNumber(ResourceCategory, "y")
	resourceSubCategoryIndex := e.CaptureColNumber(ResourceSubCategory, "y")
	vpcIndex := e.CaptureColNumber(VPC, "y")
	subnetIndex := e.CaptureColNumber(Subnet, "y")
	userSubnetIndex := e.CaptureColNumber(UserSubnet, "y")
	accessNodeRegionIndex := e.CaptureColNumber(AccessNodeRegion, "y")
	accessNodeAddressIndex := e.CaptureColNumber(AccessNodeAddress, "y")
	accessNodeContactIndex := e.CaptureColNumber(AccessNodeContact, "y")
	accessNodeAccountManagerIndex := e.CaptureColNumber(AccessNodeAccountManager, "y")
	bandwidthIndex := e.CaptureColNumber(Bandwidth, "y")
	cloudPoolIndex := e.CaptureColNumber(CloudPool, "y")
	nodeIndex := e.CaptureColNumber(Node, "n")
	var directConnectSlice []DirectConnectResource
	for i, row := range newRows {
		// 跳过前 n 行
		if i+1 < e.CopyLine {
			continue
		}
		// 初始化结构体并填充字段
		var directConnect DirectConnectResource
		directConnect.Row = strconv.Itoa(i + 1)
		directConnect.Tag1 = GetFieldContent(row, tag1Index)
		directConnect.TagX = GetFieldContent(row, tagXIndex)
		directConnect.Tag2 = GetFieldContent(row, tag2Index)
		directConnect.ResourceID = GetFieldContent(row, resourceIDIndex)
		directConnect.ResourceName = GetFieldContent(row, resourceNameIndex)
		directConnect.ResourceCategory = GetFieldContent(row, resourceCategoryIndex)
		directConnect.ResourceSubCategory = GetFieldContent(row, resourceSubCategoryIndex)
		directConnect.VPC = GetFieldContent(row, vpcIndex)
		directConnect.VPCSubnet = GetFieldContent(row, subnetIndex)
		directConnect.UserSubnet = GetFieldContent(row, userSubnetIndex)
		directConnect.AccessNodeRegion = GetFieldContent(row, accessNodeRegionIndex)
		directConnect.AccessNodeAddress = GetFieldContent(row, accessNodeAddressIndex)
		directConnect.AccessNodeContact = GetFieldContent(row, accessNodeContactIndex)
		directConnect.AccessNodeAccountManager = GetFieldContent(row, accessNodeAccountManagerIndex)
		directConnect.BandWidth = GetFieldContent(row, bandwidthIndex)
		directConnect.CloudPool = GetFieldContent(row, cloudPoolIndex)
		directConnect.Node = GetFieldContent(row, nodeIndex)
		// 将填充好的结构体添加到切片中
		directConnectSlice = append(directConnectSlice, directConnect)
	}
	return directConnectSlice
}

// ParseNATGatewayToStruct
// 将excel表格中的每一行数据转成对应的结构体后加入到切片中
func (e *Excel) ParseNATGatewayToStruct() []NATGatewayResource {
	rows, err := e.File.GetRows(e.Sheet)
	if err != nil {
		log.Printf("filename=%s sheet=%s get rows error:%s", e.FileName, e.Sheet, err.Error())
		return nil
	}

	// 行内容长度不统一，填充为一样的长度
	newRows := PadField(e.TitleLine, rows)

	// 获取各字段的列索引
	tag1Index := e.CaptureColNumber(Tag1, "n")
	tagXIndex := e.CaptureColNumber(TagX, "n")
	tag2Index := e.CaptureColNumber(Tag2, "n")
	resourceIDIndex := e.CaptureColNumber(ResourceID, "y")
	resourceNameIndex := e.CaptureColNumber(ResourceName, "y")
	resourceCategoryIndex := e.CaptureColNumber(ResourceCategory, "y")
	resourceSubCategoryIndex := e.CaptureColNumber(ResourceSubCategory, "y")
	specificationIndex := e.CaptureColNumber(Specification, "y")
	vpcIndex := e.CaptureColNumber(VPC, "y")
	bandwidthIndex := e.CaptureColNumber(Bandwidth, "y")
	cloudPoolIndex := e.CaptureColNumber(CloudPool, "y")
	nodeIndex := e.CaptureColNumber(Node, "n")

	var natSlice []NATGatewayResource
	for i, row := range newRows {
		// 跳过前 n 行
		if i+1 < e.CopyLine {
			continue
		}

		// 初始化结构体并填充字段
		var nat NATGatewayResource
		nat.Row = strconv.Itoa(i + 1)
		nat.Tag1 = GetFieldContent(row, tag1Index)
		nat.TagX = GetFieldContent(row, tagXIndex)
		nat.Tag2 = GetFieldContent(row, tag2Index)
		nat.ResourceID = GetFieldContent(row, resourceIDIndex)
		nat.ResourceName = GetFieldContent(row, resourceNameIndex)
		nat.ResourceCategory = GetFieldContent(row, resourceCategoryIndex)
		nat.ResourceSubCategory = GetFieldContent(row, resourceSubCategoryIndex)
		nat.Specification = GetFieldContent(row, specificationIndex)
		nat.VPC = GetFieldContent(row, vpcIndex)
		nat.BandWidth = GetFieldContent(row, bandwidthIndex)
		nat.CloudPool = GetFieldContent(row, cloudPoolIndex)
		nat.Node = GetFieldContent(row, nodeIndex)

		// 将填充好的结构体添加到切片中
		natSlice = append(natSlice, nat)
	}

	return natSlice
}

// ParseCloudPortToStruct
// 将excel表格中的每一行数据转成对应的结构体后加入到切片中
func (e *Excel) ParseCloudPortToStruct() []CloudPortResource {
	rows, err := e.File.GetRows(e.Sheet)
	if err != nil {
		log.Printf("filename=%s sheet=%s get rows error:%s", e.FileName, e.Sheet, err.Error())
		return nil
	}

	// 行内容长度不统一，填充为一样的长度
	newRows := PadField(e.TitleLine, rows)

	// 获取各字段的列索引
	tag1Index := e.CaptureColNumber(Tag1, "n")
	tagXIndex := e.CaptureColNumber(TagX, "n")
	tag2Index := e.CaptureColNumber(Tag2, "n")
	resourceIDIndex := e.CaptureColNumber(ResourceID, "y")
	resourceNameIndex := e.CaptureColNumber(ResourceName, "y")
	resourceCategoryIndex := e.CaptureColNumber(ResourceCategory, "y")
	resourceSubCategoryIndex := e.CaptureColNumber(ResourceSubCategory, "y")
	specificationIndex := e.CaptureColNumber(Specification, "y")
	bandwidthIndex := e.CaptureColNumber(Bandwidth, "y")
	ipTypeIndex := e.CaptureColNumber(IPType, "y")
	vpcIndex := e.CaptureColNumber(VPC, "y")
	accessNodeRegionIndex := e.CaptureColNumber(AccessNodeRegion, "y")
	accessNodeAddressIndex := e.CaptureColNumber(AccessNodeAddress, "y")
	accessNodeContactIndex := e.CaptureColNumber(AccessNodeContact, "y")
	accountManagerIndex := e.CaptureColNumber(AccountManager, "y")
	vgwAndAggDockingIPMasterIndex := e.CaptureColNumber(VGWAndAggDockingIPMaster, "y")
	vgwAndAggDockingIPBackupIndex := e.CaptureColNumber(VGWAndAggDockingIPBackup, "y")
	aggAndVGWDockingIPMasterIndex := e.CaptureColNumber(AggAndVGWDockingIPMaster, "y")
	aggAndVGWDockingIPBackupIndex := e.CaptureColNumber(AggAndVGWDockingIPBackup, "y")
	aggAndVGWDockingVLANMasterIndex := e.CaptureColNumber(AggAndVGWDockingVLANMaster, "y")
	aggAndVGWDockingVLANBackupIndex := e.CaptureColNumber(AggAndVGWDockingVLANBackup, "y")
	aggAndCustomerDockingLogicalPortIndex := e.CaptureColNumber(AggAndCustomerDockingLogicalPort, "y")
	aggAndCustomerDockingPhysicalPortIndex := e.CaptureColNumber(AggAndCustomerDockingPhysicalPort, "y")
	device1Index := e.CaptureColNumber(Device1, "y")
	device2Index := e.CaptureColNumber(Device2, "y")
	cloudPoolIndex := e.CaptureColNumber(CloudPool, "y")
	resourcePoolIndex := e.CaptureColNumber(ResourcePool, "n")
	nodeIndex := e.CaptureColNumber(Node, "n")

	var cloudPortSlice []CloudPortResource
	for i, row := range newRows {
		// 跳过前 n 行
		if i+1 < e.CopyLine {
			continue
		}

		// 初始化结构体并填充字段
		var cloudPort CloudPortResource
		cloudPort.Row = strconv.Itoa(i + 1)
		cloudPort.Tag1 = GetFieldContent(row, tag1Index)
		cloudPort.TagX = GetFieldContent(row, tagXIndex)
		cloudPort.Tag2 = GetFieldContent(row, tag2Index)
		cloudPort.ResourceID = GetFieldContent(row, resourceIDIndex)
		cloudPort.ResourceName = GetFieldContent(row, resourceNameIndex)
		cloudPort.ResourceCategory = GetFieldContent(row, resourceCategoryIndex)
		cloudPort.ResourceSubCategory = GetFieldContent(row, resourceSubCategoryIndex)
		cloudPort.Specification = GetFieldContent(row, specificationIndex)
		cloudPort.BandWidth = GetFieldContent(row, bandwidthIndex)
		cloudPort.IPType = GetFieldContent(row, ipTypeIndex)
		cloudPort.VPC = GetFieldContent(row, vpcIndex)
		cloudPort.AccessNodeRegion = GetFieldContent(row, accessNodeRegionIndex)
		cloudPort.AccessNodeAddress = GetFieldContent(row, accessNodeAddressIndex)
		cloudPort.AccessNodeContact = GetFieldContent(row, accessNodeContactIndex)
		cloudPort.AccountManager = GetFieldContent(row, accountManagerIndex)
		cloudPort.VGWAndAggDockingIPMaster = GetFieldContent(row, vgwAndAggDockingIPMasterIndex)
		cloudPort.VGWAndAggDockingIPBackup = GetFieldContent(row, vgwAndAggDockingIPBackupIndex)
		cloudPort.AggAndVGWDockingIPMaster = GetFieldContent(row, aggAndVGWDockingIPMasterIndex)
		cloudPort.AggAndVGWDockingIPBackup = GetFieldContent(row, aggAndVGWDockingIPBackupIndex)
		cloudPort.AggAndVGWDockingVLANMaster = GetFieldContent(row, aggAndVGWDockingVLANMasterIndex)
		cloudPort.AggAndVGWDockingVLANBackup = GetFieldContent(row, aggAndVGWDockingVLANBackupIndex)
		cloudPort.AggAndCustomerDockingLogicalPort = GetFieldContent(row, aggAndCustomerDockingLogicalPortIndex)
		cloudPort.AggAndCustomerDockingPhysicalPort = GetFieldContent(row, aggAndCustomerDockingPhysicalPortIndex)
		cloudPort.Device1 = GetFieldContent(row, device1Index)
		cloudPort.Device2 = GetFieldContent(row, device2Index)
		cloudPort.CloudPool = GetFieldContent(row, cloudPoolIndex)
		cloudPort.ResourcePool = GetFieldContent(row, resourcePoolIndex)
		cloudPort.Node = GetFieldContent(row, nodeIndex)

		// 将填充好的结构体添加到切片中
		cloudPortSlice = append(cloudPortSlice, cloudPort)
	}

	return cloudPortSlice
}

// ParseOtherResourceToStruct
// 将excel表格中的每一行数据转成对应的结构体后加入到切片中
func (e *Excel) ParseOtherResourceToStruct() []OtherResource {
	rows, err := e.File.GetRows(e.Sheet)
	if err != nil {
		log.Printf("filename=%s sheet=%s get rows error:%s", e.FileName, e.Sheet, err.Error())
		return nil
	}

	// 行内容长度不统一，填充为一样的长度
	newRows := PadField(e.TitleLine, rows)

	// 获取各字段的列索引
	tag1Index := e.CaptureColNumber(Tag1, "n")
	tagXIndex := e.CaptureColNumber(TagX, "n")
	tag2Index := e.CaptureColNumber(Tag2, "n")
	resourceIDIndex := e.CaptureColNumber(ResourceID, "y")
	resourceNameIndex := e.CaptureColNumber(ResourceName, "y")
	resourceCategoryIndex := e.CaptureColNumber(ResourceCategory, "y")
	resourceSubCategoryIndex := e.CaptureColNumber(ResourceSubCategory, "y")
	specificationIndex := e.CaptureColNumber(Specification, "y")
	ipv4PrivateAddressIndex := e.CaptureColNumber(IPv4PrivateAddress, "y")
	ipv6AddressIndex := e.CaptureColNumber(IPv6Address, "y")
	bandwidthIndex := e.CaptureColNumber(Bandwidth, "y")
	vpcIndex := e.CaptureColNumber(VPC, "y")
	nodeIPIndex := e.CaptureColNumber(NodeIP, "y")
	endpointNodeServiceNameIndex := e.CaptureColNumber(EndpointNodeServiceName, "y")
	networkProductInstanceIDIndex := e.CaptureColNumber(NetworkProductInstanceID, "y")
	bindingResourceTypeIndex := e.CaptureColNumber(BindingResourceType, "y")
	bindingResourceNameIndex := e.CaptureColNumber(BindingResourceName, "y")
	bindingResourceIDIndex := e.CaptureColNumber(BindingResourceID, "y")
	cloudPoolIndex := e.CaptureColNumber(CloudPool, "y")
	nodeIndex := e.CaptureColNumber(Node, "n")

	var otherSlice []OtherResource
	for i, row := range newRows {
		// 跳过前 n 行
		if i+1 < e.CopyLine {
			continue
		}

		// 初始化结构体并填充字段
		var other OtherResource
		other.Row = strconv.Itoa(i + 1)
		other.Tag1 = GetFieldContent(row, tag1Index)
		other.TagX = GetFieldContent(row, tagXIndex)
		other.Tag2 = GetFieldContent(row, tag2Index)
		other.ResourceID = GetFieldContent(row, resourceIDIndex)
		other.ResourceName = GetFieldContent(row, resourceNameIndex)
		other.ResourceCategory = GetFieldContent(row, resourceCategoryIndex)
		other.ResourceSubCategory = GetFieldContent(row, resourceSubCategoryIndex)
		other.Specification = GetFieldContent(row, specificationIndex)
		other.IPv4PrivateAddress = GetFieldContent(row, ipv4PrivateAddressIndex)
		other.Ipv6Address = GetFieldContent(row, ipv6AddressIndex)
		other.BandWidth = GetFieldContent(row, bandwidthIndex)
		other.VPC = GetFieldContent(row, vpcIndex)
		other.NodeIP = GetFieldContent(row, nodeIPIndex)
		other.EndpointNodeServiceName = GetFieldContent(row, endpointNodeServiceNameIndex)
		other.NetworkProductInstanceId = GetFieldContent(row, networkProductInstanceIDIndex)
		other.BindingResourceType = GetFieldContent(row, bindingResourceTypeIndex)
		other.BindingResourceName = GetFieldContent(row, bindingResourceNameIndex)
		other.BindingResourceID = GetFieldContent(row, bindingResourceIDIndex)
		other.CloudPool = GetFieldContent(row, cloudPoolIndex)
		other.Node = GetFieldContent(row, nodeIndex)

		// 将填充好的结构体添加到切片中
		otherSlice = append(otherSlice, other)
	}

	return otherSlice
}

func NetworkGetUUIDs(excelResources NetworkResource) ([]string, error) {
	var strSlice []string

	switch {
	case len(excelResources.LoadBalance) > 0:
		for _, excelResource := range excelResources.LoadBalance {
			strSlice = append(strSlice, excelResource.ResourceID)
		}

	case len(excelResources.Eip) > 0:
		for _, excelResource := range excelResources.Eip {
			strSlice = append(strSlice, excelResource.ResourceID)
		}

	case len(excelResources.SharedTrafficPackage) > 0:
		for _, excelResource := range excelResources.SharedTrafficPackage {
			strSlice = append(strSlice, excelResource.ResourceID)
		}

	case len(excelResources.IPv6) > 0:
		for _, excelResource := range excelResources.IPv6 {
			strSlice = append(strSlice, excelResource.ResourceID)
		}

	case len(excelResources.DirectConnect) > 0:
		for _, excelResource := range excelResources.DirectConnect {
			strSlice = append(strSlice, excelResource.ResourceID)
		}

	case len(excelResources.NATGateway) > 0:
		for _, excelResource := range excelResources.NATGateway {
			strSlice = append(strSlice, excelResource.ResourceID)
		}

	case len(excelResources.CloudPort) > 0:
		for _, excelResource := range excelResources.CloudPort {
			strSlice = append(strSlice, excelResource.ResourceID)
		}

	case len(excelResources.Other) > 0:
		for _, excelResource := range excelResources.Other {
			strSlice = append(strSlice, excelResource.ResourceID)
		}

	default:
		return []string{}, errors.New("NetworkGetUUIDs未匹配到数据")
	}

	return strSlice, nil
}

func NetworkGetPool(excelResources NetworkResource) ([]string, error) {
	var tuple = make(map[string]struct{})
	var pools []string

	switch {
	case len(excelResources.LoadBalance) > 0:

		for _, e := range excelResources.LoadBalance {
			if _, exist := tuple[e.Node]; !exist {
				tuple[e.Node] = struct{}{}
			}
		}

	case len(excelResources.Eip) > 0:
		for _, e := range excelResources.Eip {
			if _, exist := tuple[e.Node]; !exist {
				tuple[e.Node] = struct{}{}
			}
		}

	case len(excelResources.SharedTrafficPackage) > 0:
		for _, e := range excelResources.SharedTrafficPackage {
			if _, exist := tuple[e.Node]; !exist {
				tuple[e.Node] = struct{}{}
			}
		}

	case len(excelResources.IPv6) > 0:
		for _, e := range excelResources.IPv6 {
			if _, exist := tuple[e.Node]; !exist {
				tuple[e.Node] = struct{}{}
			}
		}

	case len(excelResources.DirectConnect) > 0:
		for _, e := range excelResources.DirectConnect {
			if _, exist := tuple[e.Node]; !exist {
				tuple[e.Node] = struct{}{}
			}
		}

	case len(excelResources.NATGateway) > 0:
		for _, e := range excelResources.NATGateway {
			if _, exist := tuple[e.Node]; !exist {
				tuple[e.Node] = struct{}{}
			}
		}

	case len(excelResources.CloudPort) > 0:
		for _, e := range excelResources.CloudPort {
			if _, exist := tuple[e.Node]; !exist {
				tuple[e.Node] = struct{}{}
			}
		}

	case len(excelResources.Other) > 0:
		for _, e := range excelResources.Other {
			if _, exist := tuple[e.Node]; !exist {
				tuple[e.Node] = struct{}{}
			}
		}
	default:
		return []string{}, errors.New("NetworkGetPool未匹配到数据")
	}

	for k, _ := range tuple {
		pools = append(pools, mapping.PoolNameToCodeConvert(k))
	}
	return pools, nil
}
