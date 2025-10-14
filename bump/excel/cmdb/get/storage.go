package get

import (
	"Catch/bump/internal/mapping"
	"log"
	"strconv"
	"strings"
)

type StorageResource struct {
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
	Capacity            string
	CloudPool           string
	Node                string
	AvailabilityZone    string
	IsShare             string
	Mount               string
}

// ParseStorageResourceToStruct
// 将excel表格中的每一行数据转成对应的结构体后加入到切片中
func (e *Excel) ParseStorageResourceToStruct() []StorageResource {
	rows, err := e.File.GetRows(e.Sheet)
	if err != nil {
		log.Printf("filename=%s sheet=%s get rows error:%s", e.FileName, e.Sheet, err.Error())
	}

	//行内容长度不统一，填充为一样的长度
	newRows := PadField(e.TitleLine, rows)
	tag1Index := e.CaptureColNumber(Tag1, "n")
	tagXIndex := e.CaptureColNumber(TagX, "n")
	tag2Index := e.CaptureColNumber(Tag2, "n")
	resourceIDIndex := e.CaptureColNumber(ResourceID, "y")
	resourceNameIndex := e.CaptureColNumber(ResourceName, "y")
	resourceCategoryIndex := e.CaptureColNumber(ResourceCategory, "y")
	resourceSubCategoryIndex := e.CaptureColNumber(ResourceSubCategory, "y")
	specificationIndex := e.CaptureColNumber(Specification, "y")
	capacityIndex := e.CaptureColNumber(Capacity, "y")
	cloudPoolIndex := e.CaptureColNumber(CloudPool, "y")
	nodeIndex := e.CaptureColNumber(Node, "n")
	availabilityZoneIndex := e.CaptureColNumber(AvailabilityZone, "y")
	MountIndex := e.CaptureColNumber(Mount, "y")

	var isShareIndex int
	if strings.Contains(e.Sheet, "云硬盘") {
		isShareIndex = e.CaptureColNumber(IsShare, "y")
	}

	//将每行数据解析成一个结构体，存放在切片中
	var srSlice []StorageResource
	for i, row := range newRows {
		//从那行开始复制，跳过前n行
		if i+1 < e.CopyLine {
			continue
		}

		// 填充 StorageResource 结构体的字段
		var storage StorageResource
		storage.Row = strconv.Itoa(i + 1)
		storage.Tag1 = GetFieldContent(row, tag1Index)
		storage.TagX = GetFieldContent(row, tagXIndex)
		storage.Tag2 = GetFieldContent(row, tag2Index)
		storage.ResourceID = GetFieldContent(row, resourceIDIndex)
		storage.ResourceName = GetFieldContent(row, resourceNameIndex)
		storage.ResourceCategory = GetFieldContent(row, resourceCategoryIndex)
		storage.ResourceSubCategory = GetFieldContent(row, resourceSubCategoryIndex)
		storage.Specification = GetFieldContent(row, specificationIndex)
		storage.Capacity = GetFieldContent(row, capacityIndex)
		storage.CloudPool = GetFieldContent(row, cloudPoolIndex)
		storage.Node = GetFieldContent(row, nodeIndex)
		storage.AvailabilityZone = GetFieldContent(row, availabilityZoneIndex)
		storage.Mount = GetFieldContent(row, MountIndex)

		if strings.Contains(e.Sheet, "云硬盘") {
			storage.IsShare = GetFieldContent(row, isShareIndex)
		}

		// 将填充好的结构体添加到切片中
		srSlice = append(srSlice, storage)
	}

	return srSlice

}

func StorageGetUUIDs(excelResources []StorageResource) []string {
	var strSlice []string
	for _, excelResource := range excelResources {
		strSlice = append(strSlice, excelResource.ResourceID)
	}
	return strSlice
}

func StorageGetPool(excelResources []StorageResource) []string {
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
