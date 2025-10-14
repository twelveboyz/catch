package execute

import (
	"Catch/bump/excel/cmdb/get"
	"github.com/sirupsen/logrus"
)

func SecurityRun(root string) {
	securityFileName := get.GetFileName(root, ".xls", "安全")

	securityExcel := &get.Excel{
		Root:      root,
		FileName:  securityFileName,
		TitleLine: 2,
		CopyLine:  3,
	}

	securityExcel.File = securityExcel.OpenFile(securityExcel.FileName)

	sheets := securityExcel.ForGetSheet()
	for _, sheet := range sheets {
		securityExcel.Sheet = sheet
		s := securityExcel.ParseSecurityToStruct()
		for i, col := range s {
			logrus.Println(i+1, col)
		}
	}
}
