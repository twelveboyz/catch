package cmdb

import (
	"Catch/internal"
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/xuri/excelize/v2"
)

type Excel struct {
	Filename  string
	File      *excelize.File
	Sheet     string
	CopyLine  int
	TitleLine int
}

func NewExcel(fileName, sheet string, titleLine int) *Excel {
	//打开一个excel文件
	f := OpenFile(fileName)
	if sheet == "" {
		sheet = f.GetSheetName(0)
	}

	return &Excel{
		Filename:  fileName,
		Sheet:     sheet,
		CopyLine:  titleLine + 1,
		TitleLine: titleLine,
		File:      f,
	}
}

func OpenFile(fileName string) *excelize.File {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		logrus.Fatal("打开文件失败:", err)
	}

	return f
}

func (e *Excel) ResourceStatistics() []ResourceCountInfo {
	//根据第一行标题匹配string，匹配成功后返回标题所在的列对于的切片下标，使用这个下标来获取该标题一列的数据
	resourceCategoryInt := e.CaptureColNumber(ResourceCategory, e.TitleLine, "y")
	resourceSubcategoryInt := e.CaptureColNumber(ResourceSubcategory, e.TitleLine, "y")
	specificationInt := e.CaptureColNumber(Specification, e.TitleLine, "y")
	nodeNameInt := e.CaptureColNumber(NodeName, e.TitleLine, "n")
	commentInt := e.CaptureColNumber(Comments, e.TitleLine, "n")

	//根据下标抓取一列的数据，例如: 下标=0就是A列、下标=18，则是对于Xlsx表中“S”列的数据
	resourceCategoryData := e.GetColCall(resourceCategoryInt)
	resourceSubcategoryData := e.GetColCall(resourceSubcategoryInt)
	specificationData := e.GetColCall(specificationInt)
	nodeData := e.GetColCall(nodeNameInt)
	commentData := e.GetColCall(commentInt)

	//针对弹性计算资源大类，增加系统盘的统计
	for index, resourceCategory := range resourceCategoryData {
		if resourceCategory == "弹性计算" || resourceCategory == "KCS" {
			resourceCategoryData = append(resourceCategoryData, "云存储-系统盘")

			//获取资源小类“移动云-云硬盘”
			cloudPool, _ := SplitCloudPortAndResourceSubcategory(resourceSubcategoryData[index])
			resourceSubcategoryData = append(resourceSubcategoryData, fmt.Sprintf("%s-云硬盘", cloudPool))
			//获取注释中的系统盘规格
			systemDiskSpec := systemDiskRegexp(commentData[index])
			specificationData = append(specificationData, systemDiskSpec)

			nodeData = append(nodeData, nodeData[index])
			commentData = append(commentData, commentData[index])
		}
	}

	set := Set{
		ResourceCategory:    resourceCategoryData,
		ResourceSubcategory: resourceSubcategoryData,
		Specification:       specificationData,
		Node:                nodeData,
		Comments:            commentData,
	}

	//合并统计两列数据共出现的次数，并且返回一个[][]string类型的统计数据

	return slicesToStruct(TotalResource(set))
}

func (e *Excel) SecurityResourceStatistics() []ResourceCountInfo {
	//根据第一行标题匹配string，匹配成功后返回标题所在的列对于的切片下标，使用这个下标来获取该标题一列的数据
	resourceSubcategoryInt := e.CaptureColNumber(ResourceType, e.TitleLine, "y")
	specificationInt := e.CaptureColNumber(Specification, e.TitleLine, "y")
	nodeNameInt := e.CaptureColNumber(Node, e.TitleLine, "y")
	number := e.CaptureColNumber(Number, e.TitleLine, "y")

	//根据下标抓取一列的数据，例如: 下标=0就是A列、下标=18，则是对于Xlsx表中“S”列的数据
	resourceSubcategoryData := e.GetColCall(resourceSubcategoryInt)
	resourceCategoryData := e.fillInColumn(len(resourceSubcategoryData), "云安全")
	specificationData := e.GetColCall(specificationInt)
	nodeData := e.GetColCall(nodeNameInt)
	numberData := e.GetColCall(number)

	set := Set{
		ResourceCategory:    resourceCategoryData,
		ResourceSubcategory: resourceSubcategoryData,
		Specification:       specificationData,
		Node:                nodeData,
		Comments:            numberData,
	}

	//合并统计两列数据共出现的次数，并且返回一个[][]string类型的统计数据
	return slicesToStruct(TotalResource(set))
}

// fillInColumn 填充列数据
func (e *Excel) fillInColumn(columnNumber int, data string) []string {
	var dataSlice = make([]string, 0, columnNumber)
	for i := 0; i < columnNumber; i++ {
		dataSlice = append(dataSlice, data)
	}

	return dataSlice
}

func systemDiskRegexp(content string) string {
	r := regexp.MustCompile(`[(（](:?[\p{Han}[a-zA-Z]+)[）)]`).FindStringSubmatch(content)

	if len(r) == 0 {
		return ""
	}

	return r[1]
}

func (e *Excel) GenerateStatisticalData(countExcel string, cris []ResourceCountInfo) {
	//获取表格中是否有数据存在，如果有返回对应int，
	getRowInt := e.GrabXlsxRow(countExcel, "Sheet1")
	logrus.Printf("写入数据时获取到sheet存在%d条数据\n", getRowInt)

	//写入结果统计数据到Excel表格中
	e.writeToResultExcel(getRowInt, cris)
}

// writeToResultExcel 传参说明：n是文件已经存在的行数，resourceCountInfos是聚合所有必要的值
func (e *Excel) writeToResultExcel(n int, resourceCountInfos []ResourceCountInfo) {

	var num = new(int)
	*num = n
	for _, rci := range resourceCountInfos {

		e.writeCompute(num, rci)
		e.writeSystemDisk(num, rci)
		e.writeCloudStorage(num, rci)
		e.writeCloudNetwork(num, rci)
		e.writeCloudPaaS(num, rci)
		e.writeOther(num, rci)
		e.writeSecurity(num, rci)

	}

	err := e.File.Save()
	if err != nil {
		logrus.Println(err)
	}
}

func (e *Excel) writeCompute(num *int, rci ResourceCountInfo) {
	resourceCategory := rci.resourceCategory
	_ = rci.resourceSubcategory
	nodeName := strings.ReplaceAll(rci.NodeName, "_", "-")

	if resourceCategory == "弹性计算" || resourceCategory == "KCS" {
		*num++
		cell := fmt.Sprintf("A%d", *num)

		err := e.File.SetSheetRow(e.Sheet, cell, &[]interface{}{rci.CloudPool, nodeName, "", rci.resourceCategory, rci.resourceSubcategory, rci.Specification, "个", rci.Counts})
		if err != nil {
			logrus.Errorf("WriteCompute %s写入数据失败：%v", cell, err)
			return
		}

		logrus.Printf("%s写入数据：%v, %v, %v, %v, %v, %v, %v, %v\n", cell, rci.CloudPool, nodeName, "", rci.resourceCategory, rci.resourceSubcategory, rci.Specification, "个", rci.Counts)
	}

}

func (e *Excel) writeSystemDisk(num *int, rci ResourceCountInfo) {
	resourceCategory := rci.resourceCategory
	_ = rci.resourceSubcategory
	nodeName := strings.ReplaceAll(rci.NodeName, "_", "-")

	if resourceCategory == "云存储-系统盘" {
		*num++
		cell := fmt.Sprintf("A%d", *num)

		err := e.File.SetSheetRow(e.Sheet, cell, &[]interface{}{rci.CloudPool, nodeName, "", "云存储", "云硬盘", rci.Specification, "GB", rci.Capacity, rci.Counts, "系统盘"})
		if err != nil {
			logrus.Errorf("WriteSystemDisk %s写入数据失败：%v", cell, err)
			return
		}

		logrus.Printf("%s写入数据：%v, %v, %v, %v, %v, %v, %v, %v, %v, %v\n", cell, rci.CloudPool, nodeName, "", "云存储", "云硬盘", rci.Specification, "GB", rci.Capacity, rci.Counts, "系统盘")

	}

}

func (e *Excel) writeCloudStorage(num *int, rci ResourceCountInfo) {
	resourceCategory := rci.resourceCategory
	resourceSubcategory := rci.resourceSubcategory
	nodeName := strings.ReplaceAll(rci.NodeName, "_", "-")

	if resourceCategory == "云存储" {
		*num++
		cell := fmt.Sprintf("A%d", *num)

		var err error
		if strings.Contains(resourceSubcategory, "对象存储") {
			err = e.File.SetSheetRow(e.Sheet, cell, &[]interface{}{rci.CloudPool, nodeName, "", rci.resourceCategory, rci.resourceSubcategory, rci.Specification, "GB", rci.SingleCapacity, rci.Counts, rci.comment})

		} else {
			err = e.File.SetSheetRow(e.Sheet, cell, &[]interface{}{rci.CloudPool, nodeName, "", rci.resourceCategory, rci.resourceSubcategory, rci.Specification, "GB", rci.Capacity, rci.Counts})

		}

		if err != nil {
			logrus.Errorf("WriteCloudStorage %s写入数据失败：%v", cell, err)
			return
		}

		logrus.Printf("%s写入数据：%v, %v, %v, %v, %v, %v, %v, %v, %v\n", cell, rci.CloudPool, nodeName, "", rci.resourceCategory, rci.resourceSubcategory, rci.Specification, "GB", rci.Capacity, rci.Counts)
	}

}

func (e *Excel) writeCloudNetwork(num *int, rci ResourceCountInfo) {
	resourceCategory := rci.resourceCategory
	_ = rci.resourceSubcategory
	nodeName := strings.ReplaceAll(rci.NodeName, "_", "-")

	if resourceCategory == "云网络" {
		*num++
		cell := fmt.Sprintf("A%d", *num)

		if rci.comment == "NF" || rci.comment == "NFI" {
			rci.comment = ""
		}
		err := e.File.SetSheetRow(e.Sheet, cell, &[]interface{}{rci.CloudPool, nodeName, "", rci.resourceCategory, rci.resourceSubcategory, rci.Specification, "个", rci.Counts, "", rci.comment})
		if err != nil {
			logrus.Errorf("WriteCloudNetwork %s写入数据失败：%v", cell, err)
			return
		}

		logrus.Printf("%s写入数据：%v, %v, %v, %v, %v, %v, %v, %v, %v, %v\n", cell, rci.CloudPool, nodeName, "", rci.resourceCategory, rci.resourceSubcategory, rci.Specification, "个", rci.Counts, "", rci.comment)

	}

}

func (e *Excel) writeCloudPaaS(num *int, rci ResourceCountInfo) {
	resourceCategory := rci.resourceCategory
	_ = rci.resourceSubcategory
	nodeName := strings.ReplaceAll(rci.NodeName, "_", "-")

	if resourceCategory == "云PAAS" {
		*num++
		cell := fmt.Sprintf("A%d", *num)

		err := e.File.SetSheetRow(e.Sheet, cell, &[]interface{}{rci.CloudPool, nodeName, "", rci.resourceCategory, rci.resourceSubcategory, rci.Specification, "个", rci.Counts})
		if err != nil {
			logrus.Errorf("WriteCloudPaaS %s写入数据失败：%v", cell, err)
			return
		}

		logrus.Printf("%s写入数据：%v, %v, %v, %v, %v, %v, %v, %v\n", cell, rci.CloudPool, nodeName, "", rci.resourceCategory, rci.resourceSubcategory, rci.Specification, "个", rci.Counts)

	}
}

func (e *Excel) writeOther(num *int, rci ResourceCountInfo) {
	resourceCategory := rci.resourceCategory
	_ = rci.resourceSubcategory
	nodeName := strings.ReplaceAll(rci.NodeName, "_", "-")

	if resourceCategory == "其他资源" {
		*num++
		cell := fmt.Sprintf("A%d", *num)

		err := e.File.SetSheetRow(e.Sheet, cell, &[]interface{}{rci.CloudPool, nodeName, "", rci.resourceCategory, rci.resourceSubcategory, rci.Specification, "个", rci.Counts})
		if err != nil {
			logrus.Errorf("WriteOther %s写入数据失败：%v", cell, err)
			return
		}

		logrus.Printf("%s写入数据：%v, %v, %v, %v, %v, %v, %v, %v\n", cell, rci.CloudPool, nodeName, "", rci.resourceCategory, rci.resourceSubcategory, rci.Specification, "个", rci.Counts)

	}
}

func (e *Excel) writeSecurity(num *int, rci ResourceCountInfo) {
	resourceCategory := rci.resourceCategory
	_ = rci.resourceSubcategory
	nodeName := strings.ReplaceAll(rci.NodeName, "_", "-")

	if resourceCategory == "云安全" {
		*num++
		cell := fmt.Sprintf("A%d", *num)

		err := e.File.SetSheetRow(e.Sheet, cell, &[]interface{}{rci.CloudPool, nodeName, "", rci.resourceCategory, rci.resourceSubcategory, rci.Specification, "个", rci.Counts})
		if err != nil {
			logrus.Errorf("WriteOther %s写入数据失败：%v", cell, err)
			return
		}

		logrus.Printf("%s写入数据：%v, %v, %v, %v, %v, %v, %v, %v\n", cell, rci.CloudPool, nodeName, "", rci.resourceCategory, rci.resourceSubcategory, rci.Specification, "个", rci.Counts)

	}
}

/*-------------------------------------------------------------------------------------------------*/

func AutoWidth(countExcel string) {
	// 10 是列的长度，10 = J
	for i := 0; i < 10; i++ {
		var runeLetter = 'A' + int32(i)              //每次循环A+1 = B +1  = C
		autoWidth(countExcel, string(runeLetter), i) //colStr = A 那colInt 必须等于0
	}
}

// AutoWidth 参数 colStr列的字母例如：A，colIntA列对于的下标例如：0
func autoWidth(FileName string, colStr string, colInt int) {
	sheetName := "Sheet1"
	f, _ := excelize.OpenFile(FileName)
	rows, _ := f.GetRows(sheetName)

	//n 是从第几行开始计算列的长度
	widthFloat := WidthCalculation(rows, colInt, 8)

	//fmt.Println("test=", colStr, colInt, widthFloat) //拍错打印，打印列和计算后返回的宽度

	//设置列宽，sheet , colStr =A列到A列， widthFloat = 宽度
	err := f.SetColWidth(sheetName, colStr, colStr, widthFloat)
	if err != nil {
		logrus.Println("设置宽度失败", err)
	}

	err = f.Save()
	if err != nil {
		logrus.Println("保存失败", err)
	}
}

// WidthCalculation 参数rows是数据，col是某一列的下标，例如A 对于下标是0 ,n 从第几行开始获取内容的长度
func WidthCalculation(rows [][]string, col int, n int) float64 {
	//获取
	var f float64 //存储计算宽度后返回的浮点
	var count int //用来统计列中最大的长度
	var ii = 0    //循环行使用
	for _, row := range rows {
		ii++
		//从第7行开始获取数据
		if ii >= n {
			//每行填充10个空数据，避免切片Out Of Index
			row = append(row, "", "", "", "", "", "", "", "", "", "")
			//fmt.Println("排错=", row[col]) //排错使用，打印获取到的每一列内容

			//比较如果该数据大于已经保存的数据则更新数据
			if len(row[col]) > count {
				count = len(row[col])
			}
		}
	}

	//根据内容长度来算出大概的宽度
	f = float64(count) * 0.75
	if f > 60 {
		f = 60
	}

	//fmt.Println(col, "-----------end", count, f) //排错使用，最大长度、计算后宽度
	return f
}

// SetCollStyle 设置单元格边框
func SetCollStyle(countExcel string) {
	f, _ := excelize.OpenFile(countExcel)
	borderStyle := excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
	}

	sheetName := "Sheet1"
	rows, err := f.GetRows(sheetName)
	if err != nil {
		logrus.Println("获取xlsx表格数据失败：", err)
	}
	styleId, err := f.NewStyle(&borderStyle)
	if err != nil {
		logrus.Println("创建边宽风格失败：", err)
	}

	//自动获取
	var ACellInt int
	for i, row := range rows {
		if row[0] == "云池" {
			ACellInt = i + 1 //因为下标是从0开始所以加1保证数字行数正确
		}
	}

	//最前一列 默认是A  加上获取到的行数
	ACell := fmt.Sprintf("A%d", ACellInt)
	//最后一列 默认是J列 这里固定J列加上行数
	JCell := fmt.Sprintf("J%d", len(rows))

	//设置单元格边框范围 例如 [A7 ~ J30]
	err = f.SetCellStyle("Sheet1", ACell, JCell, styleId)
	if err != nil {
		logrus.Println("设置边框失败", err)
	}

	err = f.Save()
	if err != nil {
		logrus.Println("保存失败", err)
	}

	logrus.Printf("已在项目目录下生成\"%v\"\n", countExcel)

}

// GetColCall *传参说明
// colCallNumber = 根据切片下标，获取某一列数据的内容   f=xlsx文件
func (e *Excel) GetColCall(colCallNumber int) []string {
	var Slice []string
	//如果获取xlsx行数是空则捕获切片超出index的panic报错，原因是获取xlsx表格的行数内容是空的
	defer func() {
		if r := recover(); r != nil {
			logrus.Fatalf("%v %v: GetColCallFunc函数执行报错 Recovered from panic:%v , 下标=%v,可能是行尾无数据导致", filepath.Base(e.Filename), e.Sheet, r, colCallNumber)
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
			//如果传参为-1，则默认填充NotFound(NF)
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
func (e *Excel) CaptureColNumber(title string, titleLine int, contains string) int {
	var titleBool bool

	//获取xlsx表格中所有内容
	rows, err := e.File.GetRows(e.Sheet)
	if err != nil {
		logrus.Fatal("获取行内容失败：", err)
	}

	//rows是保存每一行数据的全部内容 1xxxx 2xxxx 3xxxxx
	//row是字符串，循环第一行的每个标题  例子：实例ID	实例名称	实例状态	可用区······
	if len(rows) >= titleLine {
		for i, row := range rows[titleLine-1] {
			if contains == "y" {
				titleBool = strings.Contains(row, title)
			} else if contains == "n" {
				titleBool = row == title
			}

			if titleBool {
				logrus.Printf("%v表-sheet:%v：\"%v\"的下标是 %v", filepath.Base(e.Filename), e.Sheet, title, i)
				return i
			}
		}
	} else if len(rows) < titleLine {
		logrus.Printf("提示：%s-%s,行数小于%d行，可能是空白的sheet！\n", filepath.Base(e.Filename), e.Sheet, e.TitleLine)
	}

	logrus.Printf("提示：[%v表 sheet=%v] 没有找到\"%v\"的标题\n", filepath.Base(e.Filename), e.Sheet, title)

	return -1
}

func (e *Excel) LoopSheet() []string {
	var slice []string

	fMap := e.File.GetSheetMap()
	logrus.Println("fMap=", fMap)
	for k, v := range fMap {
		//过滤掉excelhidesheetname的sheet
		b := !strings.Contains(v, "excelhidesheetname")

		//过滤掉隐藏的sheet
		visibleBool, _ := e.File.GetSheetVisible(v)
		//logrus.Println("可见性 sheet=", v, "bool=", visibleBool) //排错使用，打印sheet可见性，如果隐藏则过滤
		if b && visibleBool {
			//logrus.Printf("sheet=%v，key: %v value: %v", xlsx.ResourceType, k, v) //排错使用 查看获取到的每一个sheet名称
			//xlsx.Sheet = fMap[i]

			//sheet工作表行数要大于三行才视为有效sheet
			rows, err := e.File.GetRows(fMap[k])
			logrus.Println("获取行失败：", err)

			if err != nil {
			}

			if len(rows) >= 3 {
				slice = append(slice, fMap[k])
			}
		}
	}
	return slice
}

func initCountExcelFile(xlsxFileName string, projectName string, pdfInfo *PdfInfo) {
	messages := internal.InputMessages()

	f := excelize.NewFile()

	defer func() {
		if err := f.Close(); err != nil {
			logrus.Println("CreateXlsx,关闭xlsx失败:", err)
		}
	}()

	// 使用默认的工作表 Sheet1。
	sheetName := "Sheet1"

	// 填充标签到A列，值到B列
	// A列：标签
	f.SetSheetRow(sheetName, "A1", &[]interface{}{"需求人："})
	f.SetSheetRow(sheetName, "A2", &[]interface{}{"需求方："})
	f.SetSheetRow(sheetName, "A3", &[]interface{}{"OA发文时间："})
	f.SetSheetRow(sheetName, "A4", &[]interface{}{"是否分批："})
	f.SetSheetRow(sheetName, "A5", &[]interface{}{"主账号："})
	f.SetSheetRow(sheetName, "A6", &[]interface{}{"项目名称："})
	f.SetSheetRow(sheetName, "A7", &[]interface{}{"新增子网信息："})
	f.SetSheetRow(sheetName, "A8", &[]interface{}{"云池", "节点", "可用区", "资源大类", "资源小类", "规格", "单位", "资源统计", "卷/桶数量", "备注"})

	// B列：值 - 优先使用PDF提取的值，如果为空则使用手动输入的值
	var demandPerson, demander, oaTime, primaryAccount string
	if pdfInfo != nil {
		demandPerson = pdfInfo.Applicant
		demander = pdfInfo.DemandSide
		oaTime = pdfInfo.OAPublishTime
		primaryAccount = pdfInfo.PrimaryAccount
	}

	// 如果PDF提取的值为空，使用手动输入的值作为默认值
	if demandPerson == "" {
		demandPerson = messages["DemandPerson"]
	}
	if demander == "" {
		demander = messages["Demander"]
	}
	if oaTime == "" {
		oaTime = messages["OATime"]
	}
	if primaryAccount == "" {
		primaryAccount = messages["PrimaryAccount"]
	}

	f.SetSheetRow(sheetName, "B1", &[]interface{}{demandPerson})
	f.SetSheetRow(sheetName, "B2", &[]interface{}{demander})
	f.SetSheetRow(sheetName, "B3", &[]interface{}{oaTime})
	f.SetSheetRow(sheetName, "B4", &[]interface{}{messages["InBatches"]})
	f.SetSheetRow(sheetName, "B5", &[]interface{}{primaryAccount})
	f.SetSheetRow(sheetName, "B6", &[]interface{}{projectName})
	f.SetSheetRow(sheetName, "B7", &[]interface{}{messages["SubnetInfo"]})

	if err := f.SaveAs(xlsxFileName); err != nil {
		logrus.Println("保存Excel失败", err)
	}
}

// GrabXlsxRow 获取excel表格已存在几行内容，返回行数
func (e *Excel) GrabXlsxRow(xlsxFileName string, sheet string) int {
	f, err := excelize.OpenFile(xlsxFileName)
	defer func() {
		if err = f.Close(); err != nil {
			logrus.Println("关闭xlsx失败", err)
		}
	}()
	rows, err := f.GetRows(sheet)
	if err != nil {
		logrus.Println("获取行失败", err)
	}
	return len(rows)
}

// DetermineCase 判断是否存在填写案列的行
func DetermineCase(fileName string, sheet string) int {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		logrus.Println("打开文件失败：", err)
	}

	rows, err := f.GetRows(sheet)
	if err != nil {
		logrus.Println("读取xlsx内容失败：", err)
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
