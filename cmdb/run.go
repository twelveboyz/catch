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

	// 尝试从PDF提取信息
	var pdfInfo *PdfInfo
	pdfInfo, err = ExtractFromPDF(temp)
	if err != nil {
		logrus.Warnf("PDF提取失败: %v，将使用手动输入方式", err)
		pdfInfo = nil
	} else {
		logrus.Infof("PDF提取成功 - 申请人: %s, 需求方: %s, OA发文时间: %s, 主账号: %s",
			pdfInfo.Applicant, pdfInfo.DemandSide, pdfInfo.OAPublishTime, pdfInfo.PrimaryAccount)
	}

	//初始化统计表格（传入pdfInfo，可以为nil）
	initCountExcelFile(countExcel, projectName, pdfInfo)

	/*-------------------------------------------资源录入模板-------------------------------------------------*/
	fileName := internal.CMDBFilterFileName(root, "资源录入表", []string{"资源录入", "资源录入表"})

	if fileName != "" {
		resourceEntryExcel := NewExcel(fileName, "", 1)

		resultExcel := NewExcel(countExcel, "Sheet1", 1)
		resultExcel.GenerateStatisticalData(countExcel, resourceEntryExcel.ResourceStatistics())

	} else {
		internal.Prompt("资源录入表")
	}

	securityFileName := internal.CMDBFilterFileName(root, "订购评审纪要", []string{"安全资源", "安全"})

	if securityFileName != "" {
		securityResourceExcel := NewExcel(securityFileName, "", 3)

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
