package cmdb

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
	"log"
	"strings"
)

// WriteToResultExcel 传参说明：xlsxFileName=xlsx的文件名称，n是传入需要写入的文件目前有多少行，sliceMap是聚合所有必要的值
func (e *Excel) writeToResultExcel(n int, resourceCountInfos []ResourceCountInfo) {
	var num = new(int)
	var externalNum = new(int)

	for i, rci := range resourceCountInfos {
		*num = i + 1 + n + *externalNum
		cell := fmt.Sprintf("A%d", *num)

		e.WriteCompute(cell, rci)
		e.WriteSystemDisk(num, externalNum, rci)
		e.WriteCloudStorage(cell, rci)
		e.WriteCloudNetwork(cell, rci)
		e.WriteCloudPaaS(cell, rci)

	}

	err := e.File.SaveAs(e.Filename)
	if err != nil {
		log.Println(err)
	}
}

func (e *Excel) WriteCompute(cell string, rci ResourceCountInfo) {
	resourceCategory := rci.resourceCategory
	_ = rci.resourceSubcategory

	if resourceCategory == "弹性计算" || resourceCategory == "KCS" {

		err := e.File.SetSheetRow(e.Sheet, cell, &[]interface{}{rci.CloudPool, rci.NodeName, "", rci.resourceCategory, rci.resourceSubcategory, rci.Specification, "个", "", rci.Counts})
		if err != nil {
			logrus.Errorf("WriteCompute %s写入数据失败：%v", cell, err)
			return
		}

		logrus.Printf("%s写入数据：%v, %v, %v, %v, %v, %v, %v, %v, %v\n", cell, rci.CloudPool, rci.NodeName, "", rci.resourceCategory, rci.resourceSubcategory, rci.Specification, "个", "", rci.Counts)
	}

}

func (e *Excel) WriteSystemDisk(num, externalNum *int, rci ResourceCountInfo) {
	resourceCategory := rci.resourceCategory
	_ = rci.resourceSubcategory

	if resourceCategory == "弹性计算" || resourceCategory == "KCS" {
		*externalNum++
		*num++
		nCell := fmt.Sprintf("A%d", *num)

		err := e.File.SetSheetRow(e.Sheet, nCell, &[]interface{}{rci.CloudPool, rci.NodeName, "", "云存储", "云硬盘", rci.comment, "GB", rci.Capacity, rci.Counts, "系统盘"})
		if err != nil {
			logrus.Errorf("WriteSystemDisk %s写入数据失败：%v", nCell, err)
			return
		}

		logrus.Printf("%s写入数据：%v, %v, %v, %v, %v, %v, %v, %v, %v, %v\n", nCell, rci.CloudPool, rci.NodeName, "", "云存储", "云硬盘", rci.comment, "GB", rci.Capacity, rci.Counts, "系统盘")

	}

}

func (e *Excel) WriteCloudStorage(cell string, rci ResourceCountInfo) {

	resourceCategory := rci.resourceCategory
	resourceSubcategory := rci.resourceSubcategory

	if resourceCategory == "云存储" {
		var err error
		if strings.Contains(resourceSubcategory, "对象存储") {
			err = e.File.SetSheetRow(e.Sheet, cell, &[]interface{}{rci.CloudPool, rci.NodeName, "", rci.resourceCategory, rci.resourceSubcategory, rci.Specification, "GB", rci.SingleCapacity, rci.Counts, rci.comment})

		} else {
			err = e.File.SetSheetRow(e.Sheet, cell, &[]interface{}{rci.CloudPool, rci.NodeName, "", rci.resourceCategory, rci.resourceSubcategory, rci.Specification, "GB", rci.Capacity, rci.Counts})

		}

		if err != nil {
			logrus.Errorf("WriteCloudStorage %s写入数据失败：%v", cell, err)
			return
		}

		logrus.Printf("%s写入数据：%v, %v, %v, %v, %v, %v, %v, %v, %v\n", cell, rci.CloudPool, rci.NodeName, "", rci.resourceCategory, rci.resourceSubcategory, rci.Specification, "GB", rci.Capacity, rci.Counts)
	}

}

func (e *Excel) WriteCloudNetwork(cell string, rci ResourceCountInfo) {
	resourceCategory := rci.resourceCategory
	_ = rci.resourceSubcategory

	if resourceCategory == "云网络" {
		if rci.comment == "NF" || rci.comment == "NFI" {
			rci.comment = ""
		}
		err := e.File.SetSheetRow(e.Sheet, cell, &[]interface{}{rci.CloudPool, rci.NodeName, "", rci.resourceCategory, rci.resourceSubcategory, rci.Specification, "个", rci.Counts, "", rci.comment})
		if err != nil {
			logrus.Errorf("WriteCloudNetwork %s写入数据失败：%v", cell, err)
			return
		}

		logrus.Printf("%s写入数据：%v, %v, %v, %v, %v, %v, %v, %v, %v, %v\n", cell, rci.CloudPool, rci.NodeName, "", rci.resourceCategory, rci.resourceSubcategory, rci.Specification, "个", rci.Counts, "", rci.comment)

	}

}

func (e *Excel) WriteCloudPaaS(cell string, rci ResourceCountInfo) {
	resourceCategory := rci.resourceCategory
	_ = rci.resourceSubcategory

	if resourceCategory == "云PAAS" {

		err := e.File.SetSheetRow(e.Sheet, cell, &[]interface{}{rci.CloudPool, rci.NodeName, "", rci.resourceCategory, rci.resourceSubcategory, rci.Specification, "个", rci.Counts})
		if err != nil {
			logrus.Errorf("WriteCloudPaaS %s写入数据失败：%v", cell, err)
			return
		}

		logrus.Printf("%s写入数据：%v, %v, %v, %v, %v, %v, %v, %v\n", cell, rci.CloudPool, rci.NodeName, "", rci.resourceCategory, rci.resourceSubcategory, rci.Specification, "个", rci.Counts)

	}

}

/*-------------------------------------------------------------------------------------------------*/

// AutoWidth 参数 colStr列的字母例如：A，colIntA列对于的下标例如：0
func AutoWidth(FileName string, colStr string, colInt int) {
	sheetName := "Sheet1"
	f, _ := excelize.OpenFile(FileName)
	rows, _ := f.GetRows(sheetName)

	//n 是从第几行开始计算列的长度
	widthFloat := WidthCalculation(rows, colInt, 8)

	//fmt.Println("test=", colStr, colInt, widthFloat) //拍错打印，打印列和计算后返回的宽度

	//设置列宽，sheet , colStr =A列到A列， widthFloat = 宽度
	err := f.SetColWidth(sheetName, colStr, colStr, widthFloat)
	if err != nil {
		log.Println("设置宽度失败", err)
	}

	err = f.Save()
	if err != nil {
		log.Println("保存失败", err)
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
func SetCollStyle(FileName string) {
	f, _ := excelize.OpenFile(FileName)
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
		log.Println("获取xlsx表格数据失败：", err)
	}
	styleId, err := f.NewStyle(&borderStyle)
	if err != nil {
		log.Println("创建边宽风格失败：", err)
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
		log.Println("设置边框失败", err)
	}

	err = f.Save()
	if err != nil {
		log.Println("保存失败", err)
	}
}
