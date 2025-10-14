package config

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"strings"
)

// AutoXlsx 传参说明：xlsxFileName=xlsx的文件名称，n是传入需要写入的文件目前有多少行，sliceMap是聚合所有必要的值
func (e *Excel) AutoXlsx(xlsxFileName string, n int, sliceMap []map[string]string) {

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

	for i := 0; i < len(sliceMap); i++ {
		cell := fmt.Sprintf("A%v", i+n+1)
		switch e.ResourceType {
		case "计算资源":
			err = f.SetSheetRow(sheetName, cell, &[]interface{}{sliceMap[i]["CloudPool"], sliceMap[i]["Node"], sliceMap[i]["AvailabilityZone"], "云计算", sliceMap[i]["ResourceSubcategory"], sliceMap[i]["Specification"], "个", sliceMap[i]["Counts"]})

		case "系统盘资源":
			err = f.SetSheetRow(sheetName, cell, &[]interface{}{sliceMap[i]["CloudPool"], sliceMap[i]["Node"], sliceMap[i]["AvailabilityZone"], "云存储", "云硬盘", sliceMap[i]["SystemDiskSpecification"], "GB", sliceMap[i]["SystemDiskCapacity"], sliceMap[i]["Counts"], "系统盘"})

		case "存储资源":
			if sliceMap[i]["ResourceSubcategory"] == "文件存储" {
				err = f.SetSheetRow(sheetName, cell, &[]interface{}{sliceMap[i]["CloudPool"], sliceMap[i]["Node"], sliceMap[i]["AvailabilityZone"], "云存储", sliceMap[i]["ResourceSubcategory"], sliceMap[i]["Specification"], "GB", sliceMap[i]["Capacity"], sliceMap[i]["Counts"]})

			} else if sliceMap[i]["ResourceSubcategory"] == "对象存储" {
				/*s := fmt.Sprintf("%v ,共享容量%vG", sliceMap[i]["Specification"], sliceMap[i]["SingleCapacity"])*/
				err = f.SetSheetRow(sheetName, cell, &[]interface{}{sliceMap[i]["CloudPool"], sliceMap[i]["Node"], sliceMap[i]["AvailabilityZone"], "云存储", sliceMap[i]["ResourceSubcategory"], sliceMap[i]["Specification"], "GB", sliceMap[i]["SingleCapacity"], sliceMap[i]["Counts"]})
			} else {
				//除了文件存储类型，默认输出
				err = f.SetSheetRow(sheetName, cell, &[]interface{}{sliceMap[i]["CloudPool"], sliceMap[i]["Node"], sliceMap[i]["AvailabilityZone"], "云存储", sliceMap[i]["ResourceSubcategory"], sliceMap[i]["Specification"], "GB", sliceMap[i]["Capacity"], sliceMap[i]["Counts"]})
			}

		case "网络资源":
			//主要处理网络资源中有些是没有规格的，如果没有规格则判断规格如果是NF则不打印规格的字段
			if strings.Contains(sliceMap[i]["Specification"], "NF") {
				sliceMap[i]["Specification"] = ""
			}
			if strings.Contains(sliceMap[i]["Bandwidth"], "NF") || sliceMap[i]["Bandwidth"] == "" {
				sliceMap[i]["Bandwidth"] = ""
			} else {
				sliceMap[i]["Bandwidth"] = fmt.Sprintf("%s Mbps", sliceMap[i]["Bandwidth"])
			}
			err = f.SetSheetRow(sheetName, cell, &[]interface{}{sliceMap[i]["CloudPool"], sliceMap[i]["Node"], "", "云网络", sliceMap[i]["ResourceSubcategory"], sliceMap[i]["Specification"], "个", sliceMap[i]["Counts"], "", sliceMap[i]["Bandwidth"]})

		case "安全资源":
			if sliceMap[i]["AvailabilityZone"] == "" {
				err = f.SetSheetRow(sheetName, cell, &[]interface{}{sliceMap[i]["CloudPool"], sliceMap[i]["Node"], "", "云安全", sliceMap[i]["ResourceSubcategory"], sliceMap[i]["Specification"], "个", sliceMap[i]["Counts"]})
			} else {
				err = f.SetSheetRow(sheetName, cell, &[]interface{}{sliceMap[i]["CloudPool"], sliceMap[i]["Node"], sliceMap[i]["AvailabilityZone"], "云安全", sliceMap[i]["ResourceSubcategory"], sliceMap[i]["Specification"], "个", sliceMap[i]["Counts"]})
			}

		case "PAAS资源":
			err = f.SetSheetRow(sheetName, cell, &[]interface{}{sliceMap[i]["CloudPool"], sliceMap[i]["Node"], sliceMap[i]["AvailabilityZone"], "PAAS", sliceMap[i]["ResourceSubcategory"], sliceMap[i]["Specification"], "个", sliceMap[i]["Counts"]})
		default:
			log.Println("写入数据时没找到对应资源")
		}

		if err != nil {
			log.Println("填写失败", err)
		}

		err = f.SaveAs(xlsxFileName)
		if err != nil {
			log.Println(err)
		}
	}
}

/*-------------------------------------------------------------------------------------------------*/

// AutoWidth 参数 colStr列的字母例如：A，colIntA列对于的下标例如：0
func AutoWidth(FileName string, colStr string, colInt int) {
	sheetName := "Sheet1"
	f, _ := excelize.OpenFile(FileName)
	rows, _ := f.GetRows(sheetName)

	//n 是从第几行开始计算列的长度
	widthFloat := WidthCalculation(rows, colInt, 1)

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
