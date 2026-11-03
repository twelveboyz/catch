package internal

import (
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

func CreateExcel() *excelize.File {
	excel := excelize.NewFile()
	return excel
}

type Excel struct {
	Filename  string
	File      *excelize.File
	Sheet     string
	CopyLine  int
	TitleLine int
}

func NewExcel(filename string, sheet string, titleLine, copyLine int) (*Excel, error) {
	file, err := excelize.OpenFile(filename)
	if err != nil {
		logrus.Fatal("打开文件失败:", err)
		return nil, err
	}

	if sheet == "" {
		sheet = file.GetSheetName(0)
	}

	return &Excel{
		Filename:  filename,
		File:      file,
		Sheet:     sheet,
		CopyLine:  copyLine,
		TitleLine: titleLine,
	}, nil
}

func (e *Excel) CaptureColNumber(titleRow []string, title string, contains bool) int {
	var titleBool bool

	//rows是保存每一行数据的全部内容 1xxxx 2xxxx 3xxxxx
	//row是字符串，循环第一行的每个标题  例子：实例ID	实例名称	实例状态	可用区······
	for i, row := range titleRow {
		if contains {
			titleBool = strings.Contains(row, title)
		} else {
			titleBool = row == title
		}

		if titleBool {
			logrus.Debugf("%v表-sheet:%v：\"%v\"的下标是 %v", filepath.Base(e.Filename), e.Sheet, title, i)
			return i
		}
	}

	logrus.Warnf("提示：[%v表 sheet=%v] 没有找到\"%v\"的标题\n", filepath.Base(e.Filename), e.Sheet, title)

	return -1
}

func (e *Excel) SetSheet(sheet string) *Excel {
	e.Sheet = sheet
	return e
}

func (e *Excel) SetTitleLine(titleLine int) *Excel {
	e.TitleLine = titleLine
	return e
}

func (e *Excel) SetCopyLine(copyLine int) *Excel {
	e.CopyLine = copyLine
	return e
}
