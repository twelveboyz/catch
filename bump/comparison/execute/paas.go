package execute

import (
	"Catch/bump/excel/cmdb/get"
	"github.com/sirupsen/logrus"
)

func PAASRun(root string) {
	paasFileName := get.GetFileName(root, ".xls", "PAAS")

	paasExcel := &get.Excel{
		Root:      root,
		FileName:  paasFileName,
		TitleLine: 2,
		CopyLine:  3,
	}

	paasExcel.File = paasExcel.OpenFile(paasExcel.FileName)

	sheets := paasExcel.ForGetSheet()
	for _, sheet := range sheets {
		paasExcel.Sheet = sheet
		s := paasExcel.ParsePaaSResourceToStruct()
		for i, col := range s {
			logrus.Println(i+1, col)
		}
	}
}
