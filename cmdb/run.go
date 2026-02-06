package cmdb

import (
	"Catch/internal"
	"fmt"
	"github.com/sirupsen/logrus"
)

func Run() {
	//手动输入相应的目录路径
	root, projectName := internal.UserInputRoot()

	//日志
	logfile := internal.LogFunc(root)
	defer func() {
		if err := logfile.Close(); err != nil {
			logrus.Warnln("close log error:", err)
		}
	}()

	//定义临时目录和统计表格名称
	temp := "temp"
	projectFile := fmt.Sprintf("/%v.xlsx", projectName)
	countExcel := temp + projectFile

	//创建临时目录并把项目复制到临时目录中
	err := internal.CreateTempDirAndCopy(root, temp)
	if err != nil {
		logrus.Error(err)
		return
	}

	defer func() {
		err = internal.MoveAndClearFile(countExcel, root+projectFile, temp)
		if err != nil {
			logrus.Error(err)
			return
		}
	}()

	//初始化统计表格
	initCountExcelFile(countExcel, projectName)

	/*-------------------------------------------资源录入模板-------------------------------------------------*/
	fileName := internal.CMDBFilterFileName(root, "资源录入表", []string{"资源录入", "资源录入表"})

	if fileName != "" {
		resourceEntryExcel := NewExcel(fileName, SheetResourceEntry, 1)

		resultExcel := NewExcel(countExcel, "Sheet1", 1)
		resultExcel.GenerateStatisticalData(countExcel, resourceEntryExcel.ResourceStatistics())

	} else {
		internal.Prompt("资源录入表")
	}

	securityFileName := internal.CMDBFilterFileName(root, "订购评审纪要", []string{"安全资源", "安全"})

	if securityFileName != "" {
		securityResourceExcel := NewExcel(securityFileName, SheetSecurityProduct, 3)

		resultExcel := NewExcel(countExcel, "Sheet1", 1)
		resultExcel.GenerateStatisticalData(countExcel, securityResourceExcel.SecurityResourceStatistics())

	} else {
		internal.Prompt("安全资源")
	}

	/*--------------------------------------------设置单格式------------------------------------------------*/
	//自动设置列宽
	AutoWidth(countExcel)
	//自动设置边框
	SetCollStyle(countExcel)
	/*----------------------------------------------end---------------------------------------------------*/

}
