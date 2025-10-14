package get

import "log"

type SecurityResource struct {
	ResourceID              string
	ResourceName            string
	ResourceCategory        string
	ResourceSubCategory     string
	Specification           string
	RelatedComputeResources string
	CloudPool               string
	Node                    string
}

func (e *Excel) ParseSecurityToStruct() []SecurityResource {
	rows, err := e.File.GetRows(e.Sheet)
	if err != nil {
		log.Printf("filename=%s sheet=%s get rows error:%s", e.FileName, e.Sheet, err.Error())
		return nil
	}
	// 行内容长度不统一，填充为一样的长度
	newRows := PadField(e.TitleLine, rows)
	// 获取各字段的列索引
	resourceIDIndex := e.CaptureColNumber(ResourceID, "y")
	resourceNameIndex := e.CaptureColNumber(ResourceName, "y")
	resourceCategoryIndex := e.CaptureColNumber(ResourceCategory, "y")
	resourceSubCategoryIndex := e.CaptureColNumber(ResourceSubCategory, "y")
	specificationIndex := e.CaptureColNumber(Specification, "y")
	relatedComputeResourcesIndex := e.CaptureColNumber(RelatedComputeResources, "y")
	cloudPoolIndex := e.CaptureColNumber(CloudPool, "y")
	nodeIndex := e.CaptureColNumber(Node, "y")

	var securitySlice []SecurityResource
	for i, row := range newRows {
		// 跳过前 n 行
		if i+1 < e.CopyLine {
			continue
		}
		// 初始化结构体并填充字段
		var security SecurityResource
		security.ResourceID = GetFieldContent(row, resourceIDIndex)
		security.ResourceName = GetFieldContent(row, resourceNameIndex)
		security.ResourceCategory = GetFieldContent(row, resourceCategoryIndex)
		security.ResourceSubCategory = GetFieldContent(row, resourceSubCategoryIndex)
		security.Specification = GetFieldContent(row, specificationIndex)
		security.RelatedComputeResources = GetFieldContent(row, relatedComputeResourcesIndex)
		security.CloudPool = GetFieldContent(row, cloudPoolIndex)
		security.Node = GetFieldContent(row, nodeIndex)
		// 将填充好的结构体添加到切片中
		securitySlice = append(securitySlice, security)
	}
	return securitySlice
}
