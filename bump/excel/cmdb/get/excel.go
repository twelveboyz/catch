package get

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Excel struct {
	Root      string
	FileName  string
	File      *excelize.File
	Title     string
	TitleLine int
	CopyLine  int
	Sheet     string
}

func NewExcel(root, filename string, titleLine, copyLine int) *Excel {
	return &Excel{
		Root:      root,
		FileName:  filename,
		TitleLine: titleLine,
		CopyLine:  copyLine,
	}
}

func (e *Excel) OpenFile(FileName string) *excelize.File {
	file, err := excelize.OpenFile(FileName)
	if err != nil {
		fmt.Println("open file err:", err)
		os.Exit(1)
	}

	return file
}

func (e *Excel) EffectiveSheet() error {
	rows, err := e.File.GetRows(e.Sheet)
	if err != nil {
		return err
	}

	if len(rows) <= 2 {
		return errors.New("len < 2")
	}

	return nil
}

func (e *Excel) ForGetSheet() []string {
	var slice []string
	fMap := e.File.GetSheetMap()
	logrus.Traceln("fMap=", fMap)
	for k, v := range fMap {
		//过滤掉excelhidesheetname的sheet
		b := !strings.Contains(v, "excelhidesheetname")

		//过滤掉隐藏的sheet
		visibleBool, _ := e.File.GetSheetVisible(v)
		//log.Println("可见性 sheet=", v, "bool=", visibleBool) //排错使用，打印sheet可见性，如果隐藏则过滤
		if b && visibleBool {
			//log.Printf("sheet=%v，key: %v value: %v", xlsx.ResourceType, k, v) //排错使用 查看获取到的每一个sheet名称
			//xlsx.Sheet = fMap[i]
			slice = append(slice, fMap[k])
		}
	}
	logrus.Traceln("有效sheet=", slice)
	return slice
}

func (e *Excel) CaptureColNumber(titleName string, contains string) int {
	var titleBool bool

	//获取xlsx表格中所有内容
	rows, err := e.File.GetRows(e.Sheet)
	if err != nil {
		log.Fatal("获取行内容失败：", err)
	}

	//rows是保存每一行数据的全部内容 1xxxx 2xxxx 3xxxxx
	//row是字符串，循环第一行的每个标题  例子：实例ID	实例名称	实例状态	可用区······
	if len(rows) >= e.TitleLine {
		for i, row := range rows[e.TitleLine-1] {
			if contains == "y" {
				titleBool = strings.Contains(row, titleName)
			} else if contains == "n" {
				titleBool = row == titleName
			}

			if titleBool {
				logrus.Tracef("%v表-sheet:%v：'%v'的下标是 %v\n", filepath.Base(e.FileName), e.Sheet, titleName, i)
				return i
			}
		}
	} else if len(rows) < e.TitleLine {
		logrus.Warnf("提示：%s-%s,行数小于%d行，可能是空白的sheet！\n", filepath.Base(e.FileName), e.Sheet, e.TitleLine)
	}

	logrus.Warnf("提示：[%v表 sheet=%v] 没有找到'%v'的标题\n", filepath.Base(e.FileName), e.Sheet, titleName)

	return -1
}

// PadField 填充行长度
func PadField(titleLine int, rows [][]string) [][]string {
	var rowLen = len(rows[titleLine-1]) //获取标题行的长度, -1是因为切片下标从0开始
	var newRows [][]string

	for _, row := range rows {
		//row长度如果小于rowLen则使用空字符串补齐
		if len(row) < rowLen {
			fillInLen := rowLen - len(row)
			for i := 0; i < fillInLen; i++ {
				row = append(row, "")
			}
		}

		newRows = append(newRows, row)
	}

	return newRows
}

func GetFieldContent(row []string, index int) string {
	if index == -1 {
		return "NF"
	}
	return row[index]
}
