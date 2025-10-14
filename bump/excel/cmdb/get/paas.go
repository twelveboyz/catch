package get

import "log"

type PAASResource struct {
	ResourceID          string
	ResourceName        string
	ResourceCategory    string
	ResourceSubCategory string
	Specification       string
	Architecture        string
	ImageVersion        string
	CloudPool           string
	Node                string
	PrivateIPv4         string
	PrivateIPv6         string
	PublicIPv4          string
	PortInfo            string
	Vpc                 string
	SecurityGroup       string
}

func (e *Excel) ParsePaaSResourceToStruct() []PAASResource {
	rows, err := e.File.GetRows(e.Sheet)
	if err != nil {
		log.Printf("filename=%s sheet=%s get rows error:%s", e.FileName, e.Sheet, err.Error())
	}

	//行内容长度不统一，填充为一样的长度
	newRows := PadField(e.TitleLine, rows)

	resourceIDIndex := e.CaptureColNumber(ResourceID, "y")
	resourceNameIndex := e.CaptureColNumber(ResourceName, "y")
	resourceCategoryIndex := e.CaptureColNumber(ResourceCategory, "y")
	resourceSubCategoryIndex := e.CaptureColNumber(ResourceSubCategory, "y")
	specificationIndex := e.CaptureColNumber(Specification, "y")
	architectureIndex := e.CaptureColNumber(Architecture, "y")
	imageVersionIndex := e.CaptureColNumber(ImageVersion, "y")
	cloudPoolIndex := e.CaptureColNumber(CloudPool, "y")
	nodeIndex := e.CaptureColNumber(Node, "n")
	privateIPv4Index := e.CaptureColNumber(PrivateIPv4, "y")
	privateIPv6Index := e.CaptureColNumber(PrivateIPv6, "y")
	publicIPv4Index := e.CaptureColNumber(PublicIPv4, "y")
	portInfoIndex := e.CaptureColNumber(PortInfo, "y")
	vpcIndex := e.CaptureColNumber(LowerVPC, "y")
	securityGroupIndex := e.CaptureColNumber(SecurityGroup, "y")

	//将每行数据解析成一个结构体，存放在切片中
	var prSlice []PAASResource
	for i, row := range newRows {
		//从那行开始复制，跳过前n行
		if i+1 < e.CopyLine {
			continue
		}

		// 填充 PAASResource 结构体的字段
		var paas PAASResource
		paas.ResourceID = GetFieldContent(row, resourceIDIndex)
		paas.ResourceName = GetFieldContent(row, resourceNameIndex)
		paas.ResourceCategory = GetFieldContent(row, resourceCategoryIndex)
		paas.ResourceSubCategory = GetFieldContent(row, resourceSubCategoryIndex)
		paas.Specification = GetFieldContent(row, specificationIndex)
		paas.Architecture = GetFieldContent(row, architectureIndex)
		paas.ImageVersion = GetFieldContent(row, imageVersionIndex)
		paas.CloudPool = GetFieldContent(row, cloudPoolIndex)
		paas.Node = GetFieldContent(row, nodeIndex)
		paas.PrivateIPv4 = GetFieldContent(row, privateIPv4Index)
		paas.PrivateIPv6 = GetFieldContent(row, privateIPv6Index)
		paas.PublicIPv4 = GetFieldContent(row, publicIPv4Index)
		paas.PortInfo = GetFieldContent(row, portInfoIndex)
		paas.Vpc = GetFieldContent(row, vpcIndex)
		paas.SecurityGroup = GetFieldContent(row, securityGroupIndex)

		// 将填充好的结构体添加到切片中
		prSlice = append(prSlice, paas)
	}

	return prSlice
}
