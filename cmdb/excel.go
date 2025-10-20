package cmdb

import (
	"log"
	"path/filepath"
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

func NewExcel(fileName, sheet string) *Excel {
	//打开一个excel文件
	return &Excel{
		Filename:  fileName,
		Sheet:     sheet,
		CopyLine:  2,
		TitleLine: 1,
		File:      OpenFile(fileName),
	}
}

func OpenFile(fileName string) *excelize.File {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		log.Fatal("打开文件失败:", err)
	}

	return f
}

func (e *Excel) OpenFile() *excelize.File {
	f, err := excelize.OpenFile(e.Filename)
	if err != nil {
		log.Fatal("打开文件失败:", err)
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

	computeSet := Set{
		ResourceCategory:    resourceCategoryData,
		ResourceSubcategory: resourceSubcategoryData,
		Specification:       specificationData,
		Node:                nodeData,
		Comments:            commentData,
	}

	//合并统计两列数据共出现的次数，并且返回一个[][]string类型的统计数据
	comPutCountsSlice := TotalResource(computeSet)

	return slicesToStruct(comPutCountsSlice)

	/*	for k, v := range countsMap {
		log.Println("TotalResourceSliceToMap打印结果：", k, v)
	}*/

}

func (e *Excel) GenerateStatisticalFile(countExcel string, cris []ResourceCountInfo) {
	//获取表格中是否有数据存在，如果有返回对应int，
	getRowInt := e.GrabXlsxRow(countExcel, "Sheet1")
	log.Printf("%v,写入数据时获取到sheet存在数据\n\n", getRowInt)

	e.writeToResultExcel(getRowInt, cris)
}

// GetColCall *传参说明
// colCallNumber = 根据切片下标，获取某一列数据的内容   f=xlsx文件
func (e *Excel) GetColCall(colCallNumber int) []string {
	var Slice []string
	//如果获取xlsx行数是空则捕获切片超出index的panic报错，原因是获取xlsx表格的行数内容是空的
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("%v %v: GetColCallFunc函数执行报错 Recovered from panic:%v , 下标是=%v,可能是行尾无数据导致", e.Sheet, r, colCallNumber)
		}
	}()

	rows, err := e.File.GetRows(e.Sheet)
	if err != nil {
		log.Fatal("读取Sheet失败:", err)
	}
	// 从colCallNumber第x行开始读取数据,并且把一列的数据保存到Slice切片中
	//fmt.Println(rows)
	for i, row := range rows {
		//log.Printf("排查：%v:%v %v", xlsx.ResourceType, xlsx.Sheet, len(row))
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
func (e *Excel) CaptureColNumber(title string, titleLine int, contains string) int {
	var titleBool bool

	/*	if xlsx.Sheet == "" {
		fMap := f.GetSheetMap()
		xlsx.Sheet = fMap[1]
		log.Printf("%v表格自动获取第一个Sheet：%v\n", xlsx.ResourceType, fMap[1])
	}*/

	//获取xlsx表格中所有内容
	rows, err := e.File.GetRows(e.Sheet)
	if err != nil {
		log.Fatal("获取行内容失败：", err)
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

func (e *Excel) ForGetSheet() []string {
	var slice []string

	fMap := e.File.GetSheetMap()
	log.Println("fMap=", fMap)
	for k, v := range fMap {
		//过滤掉excelhidesheetname的sheet
		b := !strings.Contains(v, "excelhidesheetname")

		//过滤掉隐藏的sheet
		visibleBool, _ := e.File.GetSheetVisible(v)
		//log.Println("可见性 sheet=", v, "bool=", visibleBool) //排错使用，打印sheet可见性，如果隐藏则过滤
		if b && visibleBool {
			//log.Printf("sheet=%v，key: %v value: %v", xlsx.ResourceType, k, v) //排错使用 查看获取到的每一个sheet名称
			//xlsx.Sheet = fMap[i]

			//sheet工作表行数要大于三行才视为有效sheet
			rows, err := e.File.GetRows(fMap[k])
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
	/*	if err != nil {
		log.Println("文件保存失败: ", err)
	}*/
	return err
}

func Insert(xlsxFileName string, projectName string, m map[string]string) {
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
	err = f.SetSheetRow(sheetName, "A1", &[]interface{}{"需求人：", m["DemandPerson"]})
	err = f.SetSheetRow(sheetName, "A2", &[]interface{}{"需求方：", m["Demander"]})
	err = f.SetSheetRow(sheetName, "A3", &[]interface{}{"OA发文时间：", m["OATime"]})
	err = f.SetSheetRow(sheetName, "A4", &[]interface{}{"是否分批：", m["InBatches"]})
	err = f.SetSheetRow(sheetName, "A5", &[]interface{}{"主账号：", m["PrimaryAccount"]})
	err = f.SetSheetRow(sheetName, "A6", &[]interface{}{"项目名称：", projectName})
	err = f.SetSheetRow(sheetName, "A7", &[]interface{}{"新增子网信息：", m["SubnetInfo"]})
	err = f.SetSheetRow(sheetName, "A8", &[]interface{}{"云池", "节点", "可用区", "资源大类", "资源小类", "规格", "单位", "资源统计", "卷/桶数量", "备注"})

	if err != nil {
		log.Println("填写失败", err)
	}
	err = f.SaveAs(xlsxFileName)
	if err != nil {
		log.Println(err)
	}
}

// GrabXlsxRow 获取excel表格已存在几行内容，返回行数
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
