package internal

import (
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
	"path/filepath"
	"strings"
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

// GetColCall *传参说明
// colCallNumber = 根据切片下标，获取某一列数据的内容   f=xlsx文件
func (e *Excel) GetColCall(colCallNumber int) []string {
	var Slice []string
	//如果获取xlsx行数是空则捕获切片超出index的panic报错，原因是获取xlsx表格的行数内容是空的
	defer func() {
		if r := recover(); r != nil {
			logrus.Fatalf("%v %v: GetColCallFunc函数执行报错 Recovered from panic:%v , 下标是=%v,可能是行尾无数据导致", e.Sheet, r, colCallNumber)
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
