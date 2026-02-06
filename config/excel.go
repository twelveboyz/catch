package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
	"path/filepath"
	"strings"
)

type Excel struct {
	Filename     string
	Sheet        string
	CopyLine     int
	TitleLine    int
	ResourceType string
	File         *excelize.File
}

func NewExcel(filename string, titleLine int, resourceType string) (*Excel, error) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, fmt.Errorf("NewExcel:打开文件失败:%w", err)
	}

	return &Excel{
		Filename:     filename,
		TitleLine:    titleLine,
		ResourceType: resourceType,
		File:         f,
	}, nil
}

func (e *Excel) OpenFile() *excelize.File {
	f, err := excelize.OpenFile(e.Filename)
	if err != nil {
		logrus.Fatal("打开文件失败:", err)
	}

	return f
}

func (e *Excel) ComPutExec(countExcelFileName string) {
	//根据第一行标题匹配string，匹配成功后返回标题所在的列对于的切片下标，使用这个下标来获取该标题一列的数据
	comPutResourceSubcategoryInt := e.CaptureColNumber(ResourceSubcategory, true)
	comPutSpecificationInt := e.CaptureColNumber(Specification, true)
	comPutCloudPoolInt := e.CaptureColNumber(CloudPool, true)
	comPutNodeInt := e.CaptureColNumber(Node, true)
	comPutAZInt := e.CaptureColNumber(AvailabilityZone, true)
	comPutSystemDiskSpecificationInt := e.CaptureColNumber(SystemDiskSpecification, true)
	comPutSystemDiskCapacityInt := e.CaptureColNumber(SystemDiskCapacity, true)

	//根据下标抓取一列的数据，例如: 下标=-1就是A列、下标=18，则是对于Xlsx表中“S”列的数据
	comPutResourceSubcategoryData := e.GetColCall(comPutResourceSubcategoryInt)
	comPutSpecificationData := e.GetColCall(comPutSpecificationInt)
	comPutCloudPoolData := e.GetColCall(comPutCloudPoolInt)
	comPutNodeData := e.GetColCall(comPutNodeInt)
	comPutAZData := e.GetColCall(comPutAZInt)
	comPutSystemDiskSpecificationData := e.GetColCall(comPutSystemDiskSpecificationInt)
	comPutSystemDiskCapacityData := e.GetColCall(comPutSystemDiskCapacityInt)

	comPutSet := Set{
		ResourceSubcategory:     comPutResourceSubcategoryData,
		Specification:           comPutSpecificationData,
		CloudPool:               comPutCloudPoolData,
		Node:                    comPutNodeData,
		AvailabilityZone:        comPutAZData,
		SystemDiskSpecification: comPutSystemDiskSpecificationData,
		SystemDiskCapacity:      comPutSystemDiskCapacityData,
	}

	//合并统计两列数据共出现的次数，并且返回一个[][]string类型的统计数据
	comPutCountsSlice := ComPutTotalResource(comPutSet)

	//将[][]切片转换为[]map
	comPutCountsMap := TotalResourceSliceToMap(comPutCountsSlice, e.ResourceType)

	/*
		for k, v := range comPutCountsMap {
			fmt.Println("测试输出结果：", k, v)
		}*/

	//打印FillInInformation数据
	logrus.Println("comPutCountsMap打印：", comPutCountsMap)

	//获取表格中是否有数据存在，如果有返回对应int，
	e.WriteDataToExcel(countExcelFileName, e.grabXlsxRow(countExcelFileName), comPutCountsMap)

}

func (e *Excel) StorageExec(countExcelFileName string) {
	//根据第一行标题匹配string，匹配成功后返回标题所在的列对于的切片下标，使用这个下标来获取该标题一列的数据
	storageResourceSubcategoryInt := e.CaptureColNumber(ResourceSubcategory, true)
	storageSpecificationInt := e.CaptureColNumber(Specification, true)
	storageCloudPoolInt := e.CaptureColNumber(CloudPool, true)
	storageNodeInt := e.CaptureColNumber(Node, true)
	storageAZInt := e.CaptureColNumber(AvailabilityZone, true)
	storageCapacityInt := e.CaptureColNumber(Capacity, true)
	storageCommentsInt := e.CaptureColNumber(Comments, true)

	//根据下标抓取一列的数据，例如: 下标=0就是A列、下标=18，则是对于Xlsx表中“S”列的数据
	storageResourceSubcategoryData := e.GetColCall(storageResourceSubcategoryInt)
	storageSpecificationData := e.GetColCall(storageSpecificationInt)
	storageCloudPoolData := e.GetColCall(storageCloudPoolInt)
	storageNodeData := e.GetColCall(storageNodeInt)
	storageAZData := e.GetColCall(storageAZInt)
	storageCapacityData := e.GetColCall(storageCapacityInt)
	storageCommentsData := e.GetColCall(storageCommentsInt)

	storageSet := Set{
		ResourceSubcategory: storageResourceSubcategoryData,
		Specification:       storageSpecificationData,
		CloudPool:           storageCloudPoolData,
		Node:                storageNodeData,
		AvailabilityZone:    storageAZData,
		Capacity:            storageCapacityData,
		Comments:            storageCommentsData,
	}

	//合并统计N列数据共出现的次数，并且返回一个[][]string类型的统计数据
	storageCountsSlice := StorageTotalResource(storageSet)

	//将[][]切片转换为[]map
	storageCountsMap := TotalResourceSliceToMap(storageCountsSlice, e.ResourceType)

	//打印FillInInformation数据
	logrus.Println("storageCountsMap打印：", storageCountsMap)
	//获取表格中是否有数据存在，如果有返回对应int，

	e.WriteDataToExcel(countExcelFileName, e.grabXlsxRow(countExcelFileName), storageCountsMap)
}

func (e *Excel) SystemDiskExec(countExcelFileName string) {
	//根据第一行标题匹配string，匹配成功后返回标题所在的列对于的切片下标，使用这个下标来获取该标题一列的数据
	SystemDiskResourceSubcategoryInt := e.CaptureColNumber(ResourceSubcategory, true)
	SystemDiskSpecificationInt := e.CaptureColNumber(Specification, true)
	SystemDiskCloudPoolInt := e.CaptureColNumber(CloudPool, true)
	SystemDiskNodeInt := e.CaptureColNumber(Node, true)
	SystemDiskAZInt := e.CaptureColNumber(AvailabilityZone, true)
	SystemDiskSystemDiskSpecificationInt := e.CaptureColNumber(SystemDiskSpecification, true)
	SystemDiskSystemDiskCapacityInt := e.CaptureColNumber(SystemDiskCapacity, true)

	//根据下标抓取一列的数据，例如: 下标=0就是A列、下标=18，则是对于Xlsx表中“S”列的数据
	SystemDiskResourceSubcategoryData := e.GetColCall(SystemDiskResourceSubcategoryInt)
	SystemDiskSpecificationData := e.GetColCall(SystemDiskSpecificationInt)
	SystemDiskCloudPoolData := e.GetColCall(SystemDiskCloudPoolInt)
	SystemDiskNodeData := e.GetColCall(SystemDiskNodeInt)
	SystemDiskAZData := e.GetColCall(SystemDiskAZInt)
	SystemDiskSystemDiskSpecificationData := e.GetColCall(SystemDiskSystemDiskSpecificationInt)
	SystemDiskSystemDiskCapacityData := e.GetColCall(SystemDiskSystemDiskCapacityInt)

	SystemDiskSet := Set{
		ResourceSubcategory:     SystemDiskResourceSubcategoryData,
		Specification:           SystemDiskSpecificationData,
		CloudPool:               SystemDiskCloudPoolData,
		Node:                    SystemDiskNodeData,
		AvailabilityZone:        SystemDiskAZData,
		SystemDiskSpecification: SystemDiskSystemDiskSpecificationData,
		SystemDiskCapacity:      SystemDiskSystemDiskCapacityData,
	}

	//合并统计两列数据共出现的次数，并且返回一个[][]string类型的统计数据
	SystemDiskCountsSlice := SystemDiskTotalResource(SystemDiskSet)

	//将[][]切片转换为[]map
	SystemDiskCountsMap := TotalResourceSliceToMap(SystemDiskCountsSlice, e.ResourceType)

	/*
		for k, v := range SystemDiskCountsMap {
			fmt.Println("测试输出结果：", k, v)
		}*/

	//打印FillInInformation数据
	logrus.Println("SystemDiskCountsMap打印：", SystemDiskCountsMap)

	//获取表格中是否有数据存在，如果有返回对应int，
	e.WriteDataToExcel(countExcelFileName, e.grabXlsxRow(countExcelFileName), SystemDiskCountsMap)

}

func (e *Excel) NetworkExec(countExcelFileName string) {
	//根据第一行标题匹配string，匹配成功后返回标题所在的列对于的切片下标，使用这个下标来获取该标题一列的数据
	resourceTypeInt := e.CaptureColNumber(ResourceSubcategory, true)
	resourceSubcategoryInt := e.CaptureColNumber(NetResourceSubcategory, true)
	specificationInt := e.CaptureColNumber(Specification, true)
	cloudPoolInt := e.CaptureColNumber(CloudPool, true)
	nodeInt := e.CaptureColNumber(Node, true)
	azInt := e.CaptureColNumber(AvailabilityZone, true)
	bandwidthInt := e.CaptureColNumber(Bandwidth, true)

	//根据下标抓取一列的数据，例如: 下标=0就是A列、下标=18，则是对于Xlsx表中“S”列的数据
	resourceTypeData := e.GetColCall(resourceTypeInt)
	resourceSubcategoryData := e.GetColCall(resourceSubcategoryInt)
	specificationData := e.GetColCall(specificationInt)
	cloudPoolData := e.GetColCall(cloudPoolInt)
	nodeData := e.GetColCall(nodeInt)
	azData := e.GetColCall(azInt)
	bandwidthData := e.GetColCall(bandwidthInt)

	var set Set
	if len(resourceTypeData) > 0 {
		if resourceTypeData[0] != "NF" {
			//fmt.Println("resourceTypeData-in")
			set = Set{
				ResourceSubcategory: resourceTypeData,
				Specification:       specificationData,
				CloudPool:           cloudPoolData,
				Node:                nodeData,
				AvailabilityZone:    azData,
				Bandwidth:           bandwidthData,
			}
		}
	} else if len(resourceSubcategoryData) > 0 {
		if resourceSubcategoryData[0] != "NF" {
			//fmt.Println("resourceSubcategoryData-in")
			set = Set{
				ResourceSubcategory: resourceSubcategoryData,
				Specification:       specificationData,
				CloudPool:           cloudPoolData,
				Node:                nodeData,
				AvailabilityZone:    azData,
				Bandwidth:           bandwidthData,
			}
		}
	}

	//合并统计两列数据共出现的次数，并且返回一个[][]string类型的统计数据
	networkCountsSlice := NetworkTotalResource(set)

	//将[][]切片转换为[]map
	networkCountsMap := TotalResourceSliceToMap(networkCountsSlice, e.ResourceType)

	//打印FillInInformation数据
	logrus.Println("networkCountsMap打印：", networkCountsMap)

	//获取表格中是否有数据存在，如果有返回对应int，
	e.WriteDataToExcel(countExcelFileName, e.grabXlsxRow(countExcelFileName), networkCountsMap)
}

func (e *Excel) SecurityExec(countExcelFileName string) {
	//根据第一行标题匹配string，匹配成功后返回标题所在的列对于的切片下标，使用这个下标来获取该标题一列的数据
	securityResourceSubcategoryInt := e.CaptureColNumber(ResourceSubcategory, true)
	securitySpecificationInt := e.CaptureColNumber(Specification, true)
	securityCloudPoolInt := e.CaptureColNumber(CloudPool, true)
	securityNodeInt := e.CaptureColNumber(Node, true)
	securityAZInt := e.CaptureColNumber(AvailabilityZone, true)
	securityNumberInt := e.CaptureColNumber(Number, true)

	//根据下标抓取一列的数据，例如: 下标=0就是A列、下标=18，则是对于Xlsx表中“S”列的数据
	securityResourceSubcategoryData := e.GetColCall(securityResourceSubcategoryInt)
	securitySpecificationData := e.GetColCall(securitySpecificationInt)
	securityCloudPoolData := e.GetColCall(securityCloudPoolInt)
	securityNodeData := e.GetColCall(securityNodeInt)
	securityAZData := e.GetColCall(securityAZInt)
	securityNumberData := e.GetColCall(securityNumberInt)

	securitySet := Set{
		ResourceSubcategory: securityResourceSubcategoryData,
		Specification:       securitySpecificationData,
		CloudPool:           securityCloudPoolData,
		Node:                securityNodeData,
		AvailabilityZone:    securityAZData,
		Number:              securityNumberData,
	}

	securityCountSlice := SecurityTotalResource(securitySet)
	securityCountMap := TotalResourceSliceToMap(securityCountSlice, e.ResourceType)

	//打印FillInInformation数据
	logrus.Println("securityCountMap打印：", securityCountMap)

	//获取表格中是否有数据存在，如果有返回对应int，
	e.WriteDataToExcel(countExcelFileName, e.grabXlsxRow(countExcelFileName), securityCountMap)

	/*	for _, m := range securityCountMapXlsxData
			fmt.Println("test=", m)
		}
		}*/

}

func (e *Excel) PAASExec(countExcelFileName string) {
	//根据第一行标题匹配string，匹配成功后返回标题所在的列对于的切片下标，使用这个下标来获取该标题一列的数据
	//PAASResourceCategoryInt := e.CaptureColNumber(ResourceCategory, true)
	paasResourceSubcategoryInt := e.CaptureColNumber(ResourceSubcategory, true)
	paasSpecificationInt := e.CaptureColNumber(Specification, true)
	paasCloudPoolInt := e.CaptureColNumber(CloudPool, true)
	paasNodeInt := e.CaptureColNumber(Node, true)
	paasAZInt := e.CaptureColNumber(AvailabilityZone, true)

	//根据下标抓取一列的数据，例如: 下标=0就是A列、下标=18，则是对于Xlsx表中“S”列的数据
	//PAASResourceCategoryData := e.GetColCall(PAASResourceCategoryInt)
	paasResourceSubcategoryData := e.GetColCall(paasResourceSubcategoryInt)
	paasSpecificationData := e.GetColCall(paasSpecificationInt)
	paasCloudPoolData := e.GetColCall(paasCloudPoolInt)
	paasNodeData := e.GetColCall(paasNodeInt)
	paasAZData := e.GetColCall(paasAZInt)

	PAASSet := Set{
		//ResourceCategory:    PAASResourceCategoryData,
		ResourceSubcategory: paasResourceSubcategoryData,
		Specification:       paasSpecificationData,
		CloudPool:           paasCloudPoolData,
		Node:                paasNodeData,
		AvailabilityZone:    paasAZData,
	}

	//合并统计两列数据共出现的次数，并且返回一个[][]string类型的统计数据
	PAASCountsSlice := PAASTotalResource(PAASSet)

	//将[][]切片转换为[]map
	PAASCountsMap := TotalResourceSliceToMap(PAASCountsSlice, "")

	/*
		for k, v := range PAASCountsMap {
			fmt.Println("测试输出结果：", k, v)
		}*/

	//打印FillInInformation数据
	logrus.Println("PAASCountsMap打印：", PAASCountsMap)

	//获取表格中是否有数据存在，如果有返回对应int，
	e.WriteDataToExcel(countExcelFileName, e.grabXlsxRow(countExcelFileName), PAASCountsMap)
}

// GetColCall *传参说明
// colCallNumber = 根据切片下标，获取某一列数据的内容   f=xlsx文件
func (e *Excel) GetColCall(colCallNumber int) []string {
	var Slice []string
	//如果获取xlsx行数是空则捕获切片超出index的panic报错，原因是获取xlsx表格的行数内容是空的
	defer func() {
		if r := recover(); r != nil {
			logrus.Fatalf("%v %v: GetColCallFunc函数执行报错 Recovered from panic:%v , 下标是=%v,可能是行尾无数据导致", e.ResourceType, e.Sheet, r, colCallNumber)
		}
	}()

	rows, err := e.File.GetRows(e.Sheet)
	if err != nil {
		logrus.Fatal("读取Sheet失败:", err)
	}
	// 从colCallNumber第x行开始读取数据,并且把一列的数据保存到Slice切片中
	//fmt.Println(rows)
	for i, row := range rows {
		//logrus.Printf("排查：%v:%v %v", xlsx.ResourceType, xlsx.Sheet, len(row))
		//CopyLine默认从第三行开始复制,捕获的callNumber不是-1，row的长度要大于捕获的number 这里是防止OutOfIndex
		if i >= (e.CopyLine-1) && colCallNumber != -1 && len(row) > colCallNumber {
			//条件都满足将这列数据添加到
			Slice = append(Slice, row[colCallNumber])

			//logrus.Println(row[colCallNumber]) //排错 打印获取到的每一行数据
			//如果传参为-1，则默认填充NotFound(N)
		} else if i >= (e.CopyLine-1) && colCallNumber == -1 {
			Slice = append(Slice, "NF")

			//如果下标小于len(row)长度则填充NotFoundIndex
		} else if i >= (e.CopyLine-1) && len(row) <= colCallNumber {
			Slice = append(Slice, "NFI")
		}
	}
	//网络资源 弹性公网IP  下标： 12 = [弹性公网IP 弹性公网IP]  -  网络资源 弹性公网IP  下标： -1 = [NF NF]
	//logrus.Println(xlsx.ResourceType, xlsx.Sheet, " 下标：", colCallNumber, "=", Slice) //*排错使用，获取捕抓的一列数据内容，打印切片中信息
	return Slice
}

// CaptureColNumber Title=列标题  TitleLine=标题在哪一行 f=xlsx文件
func (e *Excel) CaptureColNumber(title string, contains bool) int {
	var titleBool bool

	//获取xlsx表格中所有内容
	rows, err := e.File.GetRows(e.Sheet)
	if err != nil {
		logrus.Fatal("获取行内容失败：", err)
	}

	//rows是保存每一行数据的全部内容 1xxxx 2xxxx 3xxxxx
	//row是字符串，循环第一行的每个标题  例子：实例ID	实例名称	实例状态	可用区······
	for i, row := range rows[e.TitleLine-1] {
		if contains {
			titleBool = strings.Contains(row, title)
		} else {
			titleBool = row == title
		}

		if titleBool {
			logrus.Printf("%v表-sheet:%v：\"%v\"的下标是 %v", filepath.Base(e.Filename), e.Sheet, title, i)
			return i
		}
	}

	logrus.Printf("提示：[%v表 sheet=%v] 没有找到\"%v\"的标题\n", filepath.Base(e.Filename), e.Sheet, title)

	return -1
}

func (e *Excel) ForGetSheet() []string {
	var slice []string

	fMap := e.File.GetSheetMap()
	logrus.Println("fMap=", fMap)
	for k, v := range fMap {
		b := !strings.Contains(v, "excelhidesheetname")
		visibleBool, _ := e.File.GetSheetVisible(v)
		//logrus.Println("可见性 sheet=", v, "bool=", visibleBool)
		if b && visibleBool {
			//logrus.Printf("sheet=%v，key: %v value: %v", xlsx.ResourceType, k, v) //排错使用 查看获取到的每一个sheet名称
			//xlsx.Sheet = fMap[i]

			//sheet工作表行数要大于三行才视为有效sheet
			rows, err := e.File.GetRows(fMap[k])
			if err != nil {
				logrus.Println("获取行失败：", err)
			}

			if len(rows) >= 3 {
				slice = append(slice, fMap[k])
			}
		}
	}
	return slice
}

func CreateXlsx(xlsxFileName string) {
	f := excelize.NewFile()

	defer func() {
		if err := f.Close(); err != nil {
			logrus.Println("CreateXlsx,关闭xlsx失败:", err)
		}
	}()

	err := f.SaveAs(xlsxFileName)
	if err != nil {
		logrus.Println("CreateXlsx,保存文件失败: ", err)
	}
}

func Insert(xlsxFileName string) {
	// 创建一个新的工作簿。
	f, err := excelize.OpenFile(xlsxFileName)
	if err != nil {
		logrus.Println("打开文件失败", err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			logrus.Println("关闭xlsx文件失败", err)
		}
	}()

	// 使用默认的工作表 Sheet1。
	sheetName := "Sheet1"
	err = f.SetSheetRow(sheetName, "A1", &[]interface{}{"云池", "节点", "可用区", "资源大类", "资源小类", "规格", "单位", "资源统计", "卷/桶数量", "备注"})

	if err != nil {
		logrus.Println("填写失败", err)
	}
	err = f.SaveAs(xlsxFileName)
	if err != nil {
		logrus.Println(err)
	}
}

func (e *Excel) grabXlsxRow(xlsxFileName string) int {
	f, err := excelize.OpenFile(xlsxFileName)
	defer func() {
		if err = f.Close(); err != nil {
			logrus.Println("关闭xlsx失败", err)
		}
	}()
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		logrus.Println("获取行失败", err)
	}
	logrus.Printf("%v,写入数据时获取到sheet存在%v数据\n", e.ResourceType, len(rows))
	return len(rows)
}

// DetermineCase 判断是否存在填写案列的行
func DetermineCase(fileName string, sheet string) int {
	f, _ := excelize.OpenFile(fileName)

	rows, err := f.GetRows(sheet)
	if err != nil {
		logrus.Println("读取xlsx内容失败", err)
	}

	if len(rows) >= 4 && len(rows[3]) != 0 {
		if rows[3][0] == "填写案例" {
			return 5
		} else {
			return 4
		}
	} else {
		return 4
	}
}
