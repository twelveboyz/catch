package daily

import (
	"Catch/internal"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Excel struct {
	Filename     string
	Sheet        string
	CopyLine     int
	TitleLine    int
	ResourceType string
}

func (e *Excel) OpenFile() *excelize.File {
	f, err := excelize.OpenFile(e.Filename)
	if err != nil {
		log.Fatal("打开文件失败:", err)
	}

	return f
}

func CompRun(root string, filePath string, sheetName string) {
	/*----------------------------------------------计算----------------------------------------------*/
	ComPutFileName := internal.ConfigFilterFileName(root, ".xls", "计算资源")

	if ComPutFileName != "" {
		//创建excel对象

		comPutExcel := &Excel{Filename: ComPutFileName, Sheet: "", CopyLine: 0, TitleLine: 3, ResourceType: "计算资源"}
		//打开一个excel文件
		f := comPutExcel.OpenFile()
		defer func() {
			if err := f.Close(); err != nil {
				log.Fatal("文件关闭失败", err)
			}
		}()

		var resources []string

		sheets := comPutExcel.ForGetSheet(f)
		log.Println("获取有效的Sheet=", sheets)
		for _, sheet := range sheets {
			LineInt := DetermineCase(ComPutFileName, sheet)
			comPutExcel.CopyLine = LineInt
			log.Println("sheet=", sheet)
			comPutExcel.Sheet = sheet
			slice := comPutExcel.DailyComPutExec(f)
			resources = append(resources, slice...)
		}
		endString := comPutExcel.FormatString(resources)
		log.Println("resource=", resources)
		log.Println("endString=", endString)

		DailyGetRowInt := comPutExcel.GrabXlsxRow(filePath, sheetName)
		log.Printf("%v,获取到sheet存在%v数据\n\n", comPutExcel.ResourceType, DailyGetRowInt)
		comPutExcel.WriteDailyNewspaper(filePath, DailyGetRowInt, endString)
	} else {
		internal.Prompt("计算资源")
	}
}

func StorageRun(root string, filePath string, sheetName string) {
	/*----------------------------------------------系统盘----------------------------------------------*/
	var systemDiskResources []string
	systemDiskFileName := internal.ConfigFilterFileName(root, ".xls", "计算资源")

	if systemDiskFileName != "" {
		//创建excel对象

		systemDiskExcel := &Excel{Filename: systemDiskFileName, Sheet: "", CopyLine: 0, TitleLine: 3, ResourceType: "系统盘资源"}
		//打开一个excel文件
		f := systemDiskExcel.OpenFile()
		defer func() {
			if err := f.Close(); err != nil {
				log.Fatal("文件关闭失败", err)
			}
		}()

		sheets := systemDiskExcel.ForGetSheet(f)
		log.Println("获取有效的Sheet=", sheets)
		for _, sheet := range sheets {
			LineInt := DetermineCase(systemDiskFileName, sheet)
			systemDiskExcel.CopyLine = LineInt
			log.Println("sheet=", sheet)
			systemDiskExcel.Sheet = sheet
			slice := systemDiskExcel.DailySystemDiskExec(f)
			systemDiskResources = append(systemDiskResources, slice...)
		}
		endString := systemDiskExcel.FormatString(systemDiskResources)
		log.Println("resource=", systemDiskResources)
		log.Println("endString=", endString)

	} else {
		internal.Prompt("系统云硬盘资源")
	}

	/*----------------------------------------------存储----------------------------------------------*/
	var storageResources []string
	storageFileName := internal.ConfigFilterFileName(root, ".xls", "存储资源")

	if storageFileName != "" {
		//创建excel对象
		storageExcel := &Excel{Filename: storageFileName, Sheet: "", CopyLine: 0, TitleLine: 3, ResourceType: "存储资源"}
		//打开一个excel文件
		f := storageExcel.OpenFile()
		defer func() {
			if err := f.Close(); err != nil {
				log.Fatal("文件关闭失败", err)
			}
		}()

		sheets := storageExcel.ForGetSheet(f)
		log.Println("获取有效的Sheet=", sheets)
		for _, sheet := range sheets {
			LineInt := DetermineCase(storageFileName, sheet)
			storageExcel.CopyLine = LineInt
			log.Println("sheet=", sheet)
			storageExcel.Sheet = sheet
			slice := storageExcel.DailyStorageExec(f)
			storageResources = append(storageResources, slice...)
		}

		if len(systemDiskResources) > 0 {
			storageResources = append(storageResources, systemDiskResources...)
		}

		endString := storageExcel.FormatString(storageResources)
		log.Println("resource=", storageResources)
		log.Println("endString=", endString)

		DailyGetRowInt := storageExcel.GrabXlsxRow(filePath, sheetName)
		log.Printf("%v,获取到sheet存在%v数据\n\n", storageExcel.ResourceType, DailyGetRowInt)
		storageExcel.WriteDailyNewspaper(filePath, DailyGetRowInt, endString)

	} else if len(systemDiskResources) > 0 {
		storageExcel := &Excel{Filename: systemDiskFileName, Sheet: "", CopyLine: 0, TitleLine: 3, ResourceType: "存储资源"}
		storageResources = append(storageResources, systemDiskResources...)

		endString := storageExcel.FormatString(storageResources)
		log.Println("resource=", storageResources)
		log.Println("endString=", endString)

		DailyGetRowInt := storageExcel.GrabXlsxRow(filePath, sheetName)
		log.Printf("%v,获取到sheet存在%v数据\n\n", storageExcel.ResourceType, DailyGetRowInt)
		storageExcel.WriteDailyNewspaper(filePath, DailyGetRowInt, endString)
	} else {
		internal.Prompt("存储资源")
	}
}

func NetworkRun(root string, filePath string, sheetName string) {
	networkFileName := internal.ConfigFilterFileName(root, ".xls", "网络")

	if networkFileName != "" {
		//创建excel对象

		networkExcel := &Excel{Filename: networkFileName, Sheet: "", CopyLine: 0, TitleLine: 3, ResourceType: "网络资源"}
		//打开一个excel文件
		f := networkExcel.OpenFile()
		defer func() {
			if err := f.Close(); err != nil {
				log.Fatal("文件关闭失败", err)
			}
		}()

		var resources []string
		sheets := networkExcel.ForGetSheet(f)
		log.Println("获取有效的Sheet=", sheets)
		for _, sheet := range sheets {
			LineInt := DetermineCase(networkFileName, sheet)
			networkExcel.CopyLine = LineInt
			log.Println("sheet=", sheet)
			networkExcel.Sheet = sheet
			slice := networkExcel.DailyNetworkExec(f)
			resources = append(resources, slice...)
		}
		endString := networkExcel.FormatString(resources)
		log.Println("resource=", resources)
		log.Println("endString=", endString)

		DailyGetRowInt := networkExcel.GrabXlsxRow(filePath, sheetName)
		log.Printf("%v,获取到sheet存在%v数据\n\n", networkExcel.ResourceType, DailyGetRowInt)
		networkExcel.WriteDailyNewspaper(filePath, DailyGetRowInt, endString)
	} else {
		internal.Prompt("网络资源")
	}
}

func SecurityRun(root string, filePath string, sheetName string) {
	securityFileName := internal.ConfigFilterFileName(root, ".xls", "安全")

	if securityFileName != "" {
		//创建excel对象

		securityExcel := &Excel{Filename: securityFileName, Sheet: "", CopyLine: 0, TitleLine: 3, ResourceType: "安全资源"}
		//打开一个excel文件
		f := securityExcel.OpenFile()
		defer func() {
			if err := f.Close(); err != nil {
				log.Fatal("文件关闭失败", err)
			}
		}()

		var resources []string
		sheets := securityExcel.ForGetSheet(f)
		log.Println("获取有效的Sheet=", sheets)
		for _, sheet := range sheets {
			LineInt := DetermineCase(securityFileName, sheet)
			securityExcel.CopyLine = LineInt
			log.Println("sheet=", sheet)
			securityExcel.Sheet = sheet
			slice := securityExcel.DailySecurityExec(f)
			resources = append(resources, slice...)
		}
		endString := securityExcel.FormatString(resources)
		log.Println("resource=", resources)
		log.Println("endString=", endString)

		DailyGetRowInt := securityExcel.GrabXlsxRow(filePath, sheetName)
		log.Printf("%v,获取到sheet存在%v数据\n\n", securityExcel.ResourceType, DailyGetRowInt)
		securityExcel.WriteDailyNewspaper(filePath, DailyGetRowInt, endString)
	} else {
		internal.Prompt("安全资源")
	}
}

func PaaSRun(root string, filePath string, sheetName string) {
	paasFileName := internal.ConfigFilterFileName(root, ".xls", "PAAS")

	if paasFileName != "" {
		//创建excel对象

		paasExcel := &Excel{Filename: paasFileName, Sheet: "", CopyLine: 0, TitleLine: 3, ResourceType: "PAAS资源"}
		//打开一个excel文件
		f := paasExcel.OpenFile()
		defer func() {
			if err := f.Close(); err != nil {
				log.Fatal("文件关闭失败", err)
			}
		}()

		var resources []string
		sheets := paasExcel.ForGetSheet(f)
		log.Println("获取有效的Sheet=", sheets)
		for _, sheet := range sheets {
			LineInt := DetermineCase(paasFileName, sheet)
			paasExcel.CopyLine = LineInt
			log.Println("sheet=", sheet)
			paasExcel.Sheet = sheet
			slice := paasExcel.DailyPAASExec(f)
			resources = append(resources, slice...)
		}
		endString := paasExcel.FormatString(resources)
		log.Println("resource=", resources)
		log.Println("endString=", endString)

		DailyGetRowInt := paasExcel.GrabXlsxRow(filePath, sheetName)
		log.Printf("%v,获取到sheet存在%v数据\n\n", paasExcel.ResourceType, DailyGetRowInt)
		paasExcel.WriteDailyNewspaper(filePath, DailyGetRowInt, endString)
	} else {
		internal.Prompt("PAAS资源")
	}
}

func (e *Excel) DailyComPutExec(f *excelize.File) []string {
	//根据第一行标题匹配string，匹配成功后返回标题所在的列对于的切片下标，使用这个下标来获取该标题一列的数据
	resourceSubcategoryInt := e.CaptureColNumber(ResourceType, e.TitleLine, "n", f)
	commentsInt := e.CaptureColNumber(Comments, e.TitleLine, "y", f)

	//根据下标抓取一列的数据，例如: 下标=0就是A列、下标=18，则是对于Xlsx表中“S”列的数据
	resourceSubcategoryData := e.GetColCall(resourceSubcategoryInt, f)
	commentsData := e.GetColCall(commentsInt, f)

	set := Set{
		ResourceSubcategory: resourceSubcategoryData,
		Comments:            commentsData,
	}

	countMap := e.ComPutDailyNewspaper(set)

	var resources []string
	resources = append(resources, e.MapFormatToString(countMap))
	return resources
}

func (e *Excel) DailyStorageExec(f *excelize.File) []string {
	//根据第一行标题匹配string，匹配成功后返回标题所在的列对于的切片下标，使用这个下标来获取该标题一列的数据
	resourceSubcategoryInt := e.CaptureColNumber(ResourceType, e.TitleLine, "y", f)
	capacityInt := e.CaptureColNumber(Capacity, e.TitleLine, "y", f)
	commentsInt := e.CaptureColNumber(Comments, e.TitleLine, "y", f)

	//根据下标抓取一列的数据，例如: 下标=0就是A列、下标=18，则是对于Xlsx表中“S”列的数据
	resourceSubcategoryData := e.GetColCall(resourceSubcategoryInt, f)
	capacityData := e.GetColCall(capacityInt, f)
	commentsData := e.GetColCall(commentsInt, f)

	set := Set{
		ResourceSubcategory: resourceSubcategoryData,
		Capacity:            capacityData,
		Comments:            commentsData,
	}

	countMap := e.StorageDailyNewspaper(set)

	var resources []string
	resources = append(resources, e.MapFormatToString(countMap))
	return resources
}
func (e *Excel) DailySystemDiskExec(f *excelize.File) []string {
	//根据第一行标题匹配string，匹配成功后返回标题所在的列对于的切片下标，使用这个下标来获取该标题一列的数据
	resourceSubcategoryInt := e.CaptureColNumber(ResourceType, e.TitleLine, "n", f)
	specificationInt := e.CaptureColNumber(SystemDiskSpecification, e.TitleLine, "y", f)
	capacityInt := e.CaptureColNumber(SystemDiskCapacity, e.TitleLine, "y", f)
	resourceCommentsInt := e.CaptureColNumber(Comments, e.TitleLine, "y", f)

	//根据下标抓取一列的数据，例如: 下标=0就是A列、下标=18，则是对于Xlsx表中“S”列的数据
	resourceSubcategoryData := e.GetColCall(resourceSubcategoryInt, f)
	specificationData := e.GetColCall(specificationInt, f)
	capacityData := e.GetColCall(capacityInt, f)
	commentsData := e.GetColCall(resourceCommentsInt, f)

	set := Set{
		ResourceSubcategory:     resourceSubcategoryData,
		SystemDiskSpecification: specificationData,
		SystemDiskCapacity:      capacityData,
		Comments:                commentsData,
	}

	countMap := e.SystemDiskDailyNewspaper(set)

	var resources []string
	resources = append(resources, e.MapFormatToString(countMap))
	return resources
}

func (e *Excel) DailyNetworkExec(f *excelize.File) []string {
	//根据第一行标题匹配string，匹配成功后返回标题所在的列对于的切片下标，使用这个下标来获取该标题一列的数据
	resourceTypeInt := e.CaptureColNumber(ResourceType, e.TitleLine, "n", f)
	resourceSubcategoryInt := e.CaptureColNumber(ResourceSubcategory, e.TitleLine, "n", f)
	resourceCommentsInt := e.CaptureColNumber(Comments, e.TitleLine, "y", f)

	//根据下标抓取一列的数据，例如: 下标=-1就是A列、下标=18，则是对于Xlsx表中“S”列的数据
	resourceTypeData := e.GetColCall(resourceTypeInt, f)
	resourceSubcategoryData := e.GetColCall(resourceSubcategoryInt, f)
	resourceCommentsData := e.GetColCall(resourceCommentsInt, f)
	//fmt.Println("typeInt=", resourceTypeInt, "subInt=", resourceSubcategoryInt)
	//fmt.Println("typeData=", resourceTypeData, "subData=", resourceSubcategoryData)

	var set Set
	if len(resourceTypeData) > -1 && len(resourceSubcategoryData) > 0 {
		if resourceTypeData[0] != "NF" {
			//fmt.Println("resourceTypeData-in")
			set = Set{
				ResourceSubcategory: resourceTypeData,
				Comments:            resourceCommentsData,
			}
		} else if resourceSubcategoryData[0] != "NF" {
			//fmt.Println("resourceSubcategoryData-in")
			set = Set{
				ResourceSubcategory: resourceSubcategoryData,
				Comments:            resourceCommentsData,
			}
		}
	}

	//fmt.Println("tag0=", set)
	countMap := e.NetworkDailyNewspaper(set)

	var resources []string
	resources = append(resources, e.MapFormatToString(countMap))
	return resources
}

func (e *Excel) DailySecurityExec(f *excelize.File) []string {
	//根据第一行标题匹配string，匹配成功后返回标题所在的列对于的切片下标，使用这个下标来获取该标题一列的数据
	resourceSubcategoryInt := e.CaptureColNumber(ResourceType, e.TitleLine, "n", f)
	numberInt := e.CaptureColNumber(Number, e.TitleLine, "n", f)
	resourceCommentsInt := e.CaptureColNumber(Comments, e.TitleLine, "y", f)

	//根据下标抓取一列的数据，例如: 下标=-1就是A列、下标=18，则是对于Xlsx表中“S”列的数据
	resourceSubcategoryData := e.GetColCall(resourceSubcategoryInt, f)
	numberData := e.GetColCall(numberInt, f)
	CommentsData := e.GetColCall(resourceCommentsInt, f)

	set := Set{
		ResourceSubcategory: resourceSubcategoryData,
		Number:              numberData,
		Comments:            CommentsData,
	}

	countMap := e.SecurityDailyNewspaper(set)

	var resources []string
	resources = append(resources, e.MapFormatToString(countMap))
	return resources
}

func (e *Excel) DailyPAASExec(f *excelize.File) []string {
	//根据第一行标题匹配string，匹配成功后返回标题所在的列对于的切片下标，使用这个下标来获取该标题一列的数据
	resourceSubcategoryInt := e.CaptureColNumber(ResourceType, e.TitleLine, "n", f)
	commentsInt := e.CaptureColNumber(Comments, e.TitleLine, "y", f)

	//根据下标抓取一列的数据，例如: 下标=-1就是A列、下标=18，则是对于Xlsx表中“S”列的数据
	resourceSubcategoryData := e.GetColCall(resourceSubcategoryInt, f)
	commentsData := e.GetColCall(commentsInt, f)

	set := Set{
		ResourceSubcategory: resourceSubcategoryData,
		Comments:            commentsData,
	}

	countMap := e.PAASDailyNewspaper(set)

	var resources []string
	resources = append(resources, e.MapFormatToString(countMap))
	return resources
}

// GetColCall *传参说明
// colCallNumber = 根据切片下标，获取某一列数据的内容   f=xlsx文件
func (e *Excel) GetColCall(colCallNumber int, f *excelize.File) []string {
	var Slice []string
	//如果获取xlsx行数是空则捕获切片超出index的panic报错，原因是获取xlsx表格的行数内容是空的
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("%v %v: GetColCallFunc函数执行报错 Recovered from panic:%v , 下标是=%v,可能是行尾无数据导致", e.ResourceType, e.Sheet, r, colCallNumber)
		}
	}()

	rows, err := f.GetRows(e.Sheet)
	if err != nil {
		log.Fatal("读取Sheet失败:", err)
	}
	// 从colCallNumber第x行开始读取数据,并且把一列的数据保存到Slice切片中
	//log.Println("rows=", rows)
	for i, row := range rows {
		//log.Printf("排查：%v:%v %v", e.ResourceType, e.Sheet, len(row))
		//CopyLine默认从第三行开始复制,捕获的callNumber不是-1，row的长度要大于捕获的number 这里是防止OutOfIndex
		if i >= (e.CopyLine-1) && colCallNumber != -1 && len(row) > colCallNumber {
			//条件都满足将这列数据添加到
			Slice = append(Slice, row[colCallNumber])
			//log.Println(row[colCallNumber]) //排错 打印获取到的每一行数据
			//如果传参为-1，则默认填充NotFound(NF)
		} else if i >= (e.CopyLine-1) && colCallNumber == -1 {
			Slice = append(Slice, "NF")

			//如果下标小于len(row)长度则填充NotFoundIndex
		} else if i >= (e.CopyLine-1) && len(row) <= colCallNumber {
			Slice = append(Slice, "NFI")
		}
	}
	//网络资源 弹性公网IP  下标： 12 = [弹性公网IP 弹性公网IP]  -  网络资源 弹性公网IP  下标： -1 = [NF NF]
	//log.Println(xlsx.ResourceType, xlsx.Sheet, " 下标：", colCallNumber, "=", Slice) //*排错使用，获取捕抓的一列数据内容，打印切片中信息
	return Slice
}

// CaptureColNumber Title=列标题  TitleLine=标题在哪一行 f=xlsx文件
func (e *Excel) CaptureColNumber(title string, titleLine int, contains string, f *excelize.File) int {
	var titleBool bool

	//获取xlsx表格中所有内容
	rows, err := f.GetRows(e.Sheet)
	if err != nil {
		log.Fatal("获取行内容失败：", err)
	}

	//rows是保存每一行数据的全部内容 1xxxx 2xxxx 3xxxxx
	//row是字符串，循环第一行的每个标题  例子：实例ID	实例名称	实例状态	可用区······

	var hasSkippedRemark bool
	if len(rows) >= titleLine {
		for i, row := range rows[titleLine-1] {
			if contains == "y" {
				titleBool = strings.Contains(row, title)
				if title == "备注" && !hasSkippedRemark && titleBool {
					hasSkippedRemark = true
					continue // 继续下一次循环，不执行下面的代码
				}

			} else if contains == "n" {
				titleBool = row == title

			}
			if titleBool {
				log.Printf("%v表-sheet:%v：\"%v\"的下标是 %v", filepath.Base(e.Filename), e.Sheet, title, i)
				return i

			}
		}
	} else if len(rows) < titleLine {
		log.Printf("提示：%s-%s,行数小于%d行，可能是空白的sheet！\n", filepath.Base(e.Filename), e.Sheet, e.TitleLine)
	}

	log.Printf("提示：[%v表 sheet=%v] 没有找到\"%v\"的标题\n", filepath.Base(e.Filename), e.Sheet, title)

	return -1
}

func (e *Excel) ForGetSheet(f *excelize.File) []string {
	var slice []string
	fMap := f.GetSheetMap()
	log.Println("fMap=", fMap)
	for k, v := range fMap {
		b := !strings.Contains(v, "excelhidesheetname")
		visibleBool, _ := f.GetSheetVisible(v)
		//log.Println("可见性 sheet=", v, "bool=", visibleBool)
		if b && visibleBool {
			//log.Printf("sheet=%v，key: %v value: %v", xlsx.ResourceType, k, v) //排错使用 查看获取到的每一个sheet名称

			//sheet工作表行数要大于三行才视为有效sheet
			rows, err := f.GetRows(fMap[k])
			if err != nil {
				log.Println("获取行失败：", err)
			}

			if len(rows) >= 3 {
				slice = append(slice, fMap[k])
			}
		}
	}
	return slice
}

func CreateXlsx(xlsxFileName string) error {
	f := excelize.NewFile()

	defer func() {
		if err := f.Close(); err != nil {
			log.Println("CreateXlsx,关闭xlsx失败:", err)
		}
	}()

	_, err := f.NewSheet("Sheet1")
	if err != nil {
		log.Println("创建sheet失败", err)
	}

	err = f.SaveAs(xlsxFileName)
	if err != nil {
		log.Println("CreateXlsx,创建文件保存失败: ", err)
	}
	return err
}

func (e *Excel) GrabXlsxRow(xlsxFileName string, sheet string) int {
	f, err := excelize.OpenFile(xlsxFileName)
	defer func() {
		if err = f.Close(); err != nil {
			log.Println("关闭xlsx失败", err)
		}
	}()
	rows, err := f.GetRows(sheet)
	if err != nil {
		log.Println("获取行失败", err)
	}
	return len(rows)
}

// FormatString 将切片的数据合并成string =对象存储 * 10 (81920)、云硬盘 * 12 (6000)、文件存储 * 1 (2000)
func (e *Excel) FormatString(resources []string) string {
	var parts []string

	for _, r := range resources {
		if r != "" {
			parts = append(parts, r)
		}
	}

	s := strings.Join(parts, "、")
	return s
}

// MapFormatToString 将map数据每一个变成string内容
func (e *Excel) MapFormatToString(m map[string]*AggregatedData) string {
	var resourceCount string
	var slice []string

	for k, v := range m {
		if v.Comments == "扩容" {
			switch e.ResourceType {
			case "存储资源":
				if strings.Contains(k, "对象存储") {
					sliceK := strings.SplitN(k, "#&", 2)
					slice = append(slice, fmt.Sprintf("%s * %d (%dGB)(扩容)", sliceK[0], v.TotalCounts, v.CapacityCounts))

				} else {
					slice = append(slice, fmt.Sprintf("%s * %d (%dGB)(扩容)", k, v.TotalCounts, v.CapacityCounts))

				}
			case "系统盘资源":
				slice = append(slice, fmt.Sprintf("%s * %d (%dGB 系统盘)(扩容)", "云硬盘", v.TotalCounts, v.CapacityCounts))

			default:
				slice = append(slice, fmt.Sprintf("%s * %d (扩容)", k, v.TotalCounts))
			}

		} else {
			switch e.ResourceType {
			case "存储资源":
				if strings.Contains(k, "对象存储") {
					sliceK := strings.SplitN(k, "#&", 2)
					slice = append(slice, fmt.Sprintf("%s * %d (%dGB)", sliceK[0], v.TotalCounts, v.CapacityCounts))

				} else {
					slice = append(slice, fmt.Sprintf("%s * %d (%dGB)", k, v.TotalCounts, v.CapacityCounts))

				}
			case "系统盘资源":
				slice = append(slice, fmt.Sprintf("%s * %d (%dGB 系统盘)", "云硬盘", v.TotalCounts, v.CapacityCounts))

			default:
				slice = append(slice, fmt.Sprintf("%s * %d", k, v.TotalCounts))
			}
		}
	}

	resourceCount = strings.Join(slice, "、")

	//fmt.Println("resourceCount=", resourceCount)
	return resourceCount
}

func (e *Excel) WriteDailyNewspaper(xlsxFileName string, n int, endString string) {
	// 创建一个新的工作簿。
	f, err := excelize.OpenFile(xlsxFileName)
	if err != nil {
		log.Println("打开文件失败", err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Println("关闭xlsx文件失败", err)
		}
	}()

	// 使用默认的工作表 Sheet1。
	sheetName := "Sheet1"

	//从那个单元格开始写入
	cell := fmt.Sprintf("E%d", n+1)

	switch e.ResourceType {
	case "计算资源":

		err = f.SetSheetRow(sheetName, cell, &[]interface{}{e.ResourceType, endString, Crp[0], Crp[1], Crp[2], Crp[3], Crp[4], Crp[5], Crp[6]})
		if err != nil {
			log.Println("写入数据失败", err)
		}
	case "存储资源":
		err = f.SetSheetRow(sheetName, cell, &[]interface{}{e.ResourceType, endString, Storp[0], Storp[1], Storp[2], Storp[3], Storp[4], Storp[5], Storp[6]})
		if err != nil {
			log.Println("写入数据失败", err)
		}
	case "系统盘资源":
		err = f.SetSheetRow(sheetName, cell, &[]interface{}{e.ResourceType, endString, Storp[0], Storp[1], Storp[2], Storp[3], Storp[4], Storp[5], Storp[6]})
		if err != nil {
			log.Println("写入数据失败", err)
		}
	case "网络资源":
		err = f.SetSheetRow(sheetName, cell, &[]interface{}{e.ResourceType, endString, Nrp[0], Nrp[1], Nrp[2], Nrp[3], Nrp[4], Nrp[5], Nrp[6]})
		if err != nil {
			log.Println("写入数据失败", err)
		}
	case "安全资源":
		err = f.SetSheetRow(sheetName, cell, &[]interface{}{e.ResourceType, endString, Secrp[0], Secrp[1], Secrp[2], Secrp[3], Secrp[4], Secrp[5], Secrp[6]})
		if err != nil {
			log.Println("写入数据失败", err)
		}
	case "PAAS资源":
		err = f.SetSheetRow(sheetName, cell, &[]interface{}{e.ResourceType, endString, Prp[0], Prp[1], Prp[2], Prp[3], Prp[4], Prp[5], Prp[6]})
		if err != nil {
			log.Println("写入数据失败", err)
		}
	}

	//保存
	err = f.Save()
	if err != nil {
		log.Println("文件保存失败", err)
	}

}

// WriteDailyNewspaperInfo 写入工单标题、审批人、实施人、发布时间、交维工单
func WriteDailyNewspaperInfo(root, filePath string) {
	// 创建一个新的工作簿。
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Println("打开文件失败", err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Println("关闭xlsx文件失败", err)
		}
	}()

	sheetName := "Sheet1"

	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Println("获取行失败", err)
		return
	}

	content, txtPath, err := PdfToText(root)
	if err != nil {
		log.Println("PdfToText failed: %w", err)
		return
	}

	ot := projectBatch(content)
	pt := publicationTime(content)
	pa := planApprover(content)
	di := DeployImplementers(content)
	on := fmt.Sprintf("1.%s发起交维，交维工单：%s", time.Now().Format("2006/01/02"), oderNumber(content))

	if err = f.SetSheetRow(sheetName, "A1", &[]interface{}{ot, pt, pa, di}); err != nil {
		log.Println("写入数据失败", err)
		return
	}

	if err = f.SetCellStr(sheetName, "N1", on); err != nil {
		log.Println("写入数据失败", err)
		return
	}

	//合并单元格
	for _, c := range []string{"A", "B", "C", "D", "N"} {
		topLeftCell := fmt.Sprintf("%s1", c)
		bottomRightCell := fmt.Sprintf("%s%d", c, len(rows))
		if err = f.MergeCell(sheetName, topLeftCell, bottomRightCell); err != nil {
			log.Println("合并单元格失败", err)
		}
	}

	if err = f.Save(); err != nil {
		log.Println("文件保存失败", err)
	}

	//删除txt文件
	if err = os.Remove(txtPath); err != nil {
		log.Println("删除txt文件失败", err)
	}

}

// DetermineCase 判断是否存在填写案列的行
func DetermineCase(fileName string, sheet string) int {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		log.Println("打开文件失败：", err)
	}

	rows, err := f.GetRows(sheet)
	if err != nil {
		log.Println("读取xlsx内容失败：", err)
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
