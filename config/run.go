package config

import (
	"Catch/internal"
	"github.com/sirupsen/logrus"
)

func Run() {
	//手动输入相应的目录路径
	root := internal.InputRoot()

	temp := "temp"
	countExcelFileName := temp + "/建设信息统计表.xlsx"
	dest := root + "/建设信息统计表.xlsx"

	//创建临时目录并把项目复制到临时目录中
	err := internal.CreateTempDirAndCopy(root, temp)
	if err != nil {
		logrus.Error(err)
		return
	}

	defer func() {
		err = internal.MoveFile(countExcelFileName, dest)
		if err != nil {
			logrus.Error(err)
			return
		}

		err = internal.CleanTempDir(temp)
		if err != nil {
			logrus.Error(err)
			return
		}
	}()

	//创建一个新的excel表格
	CreateXlsx(countExcelFileName)

	//插入excel表格表头数据
	Insert(countExcelFileName)

	//日志
	logfile := internal.LogFunc(root)
	defer func() {
		if err := logfile.Close(); err != nil {
			logrus.Println("关闭日志文件失败")
		}
	}()

	/*----------------------------------------------计算----------------------------------------------*/
	if comPutFileName := internal.ConfigFilterFileName(root, ".xls", "计算资源"); comPutFileName != "" {
		comPutExcel, err := NewExcel(comPutFileName, 3, "计算资源")
		if err != nil {
			logrus.Error(err)
			return
		}

		//获取表格所有的sheet
		for _, sheet := range comPutExcel.ForGetSheet() {
			logrus.Println("sheet=", sheet)
			comPutExcel.CopyLine = DetermineCase(comPutFileName, sheet)
			comPutExcel.Sheet = sheet
			comPutExcel.ComPutExec(countExcelFileName)
		}
	} else {
		internal.Prompt("计算资源")
	}

	/*----------------------------------------------系统盘----------------------------------------------*/
	if systemDiskFileName := internal.ConfigFilterFileName(root, ".xls", "计算资源"); systemDiskFileName != "" {
		//创建xlsx对象
		systemDiskexcel, err := NewExcel(systemDiskFileName, 3, "系统盘资源")
		if err != nil {
			logrus.Error(err)
			return
		}

		//获取表格所有的sheet
		for _, sheet := range systemDiskexcel.ForGetSheet() {
			logrus.Println("sheet=", sheet)
			systemDiskexcel.CopyLine = DetermineCase(systemDiskFileName, sheet)
			systemDiskexcel.Sheet = sheet
			systemDiskexcel.SystemDiskExec(countExcelFileName)
		}
	} else {
		internal.Prompt("系统盘资源")
	}

	/*----------------------------------------------存储------------------------------------------------*/
	if storageFileName := internal.ConfigFilterFileName(root, ".xls", "存储资源"); storageFileName != "" {

		storageExcel, err := NewExcel(storageFileName, 3, "存储资源")
		if err != nil {
			logrus.Error(err)
			return
		}

		//获取表格所有的sheet
		for _, sheet := range storageExcel.ForGetSheet() {
			logrus.Println("sheet=", sheet)
			storageExcel.CopyLine = DetermineCase(storageFileName, sheet)
			storageExcel.Sheet = sheet
			storageExcel.StorageExec(countExcelFileName)
		}
	} else {
		internal.Prompt("存储资源")
	}

	/*----------------------------------------------云网络---------------------------------------------------*/
	if networkFileName := internal.ConfigFilterFileName(root, ".xls", "网络"); networkFileName != "" {
		networkExcel, err := NewExcel(networkFileName, 3, "网络资源")
		if err != nil {
			logrus.Error(err)
			return
		}

		//获取表格所有的sheet
		for _, sheet := range networkExcel.ForGetSheet() {
			logrus.Println("sheet=", sheet)
			LineInt := DetermineCase(networkFileName, sheet)
			networkExcel.CopyLine = LineInt
			networkExcel.Sheet = sheet
			networkExcel.NetworkExec(countExcelFileName)
		}
	} else {
		internal.Prompt("网络资源")
	}

	/*----------------------------------------------云安全---------------------------------------------------*/
	securityFileName := internal.ConfigFilterFileName(root, ".xls", "安全")

	if securityFileName != "" {
		securityExcel, err := NewExcel(securityFileName, 3, "安全资源")
		if err != nil {
			logrus.Error(err)
			return
		}

		//获取表格所有的sheet
		for _, sheet := range securityExcel.ForGetSheet() {
			logrus.Println("sheet=", sheet)
			securityExcel.CopyLine = DetermineCase(securityFileName, sheet)
			securityExcel.Sheet = sheet
			securityExcel.SecurityExec(countExcelFileName)
		}

	} else {
		internal.Prompt("安全资源")
	}

	/*----------------------------------------------PAAS---------------------------------------------------*/
	if paasFileName := internal.ConfigFilterFileName(root, ".xls", "PAAS"); paasFileName != "" {

		PaasExcel, err := NewExcel(paasFileName, 3, "PAAS资源")
		if err != nil {
			logrus.Error(err)
			return
		}

		//获取表格所有的sheet
		for _, sheet := range PaasExcel.ForGetSheet() {
			logrus.Println("sheet=", sheet)
			PaasExcel.CopyLine = DetermineCase(paasFileName, sheet)
			PaasExcel.Sheet = sheet
			PaasExcel.PAASExec(countExcelFileName)
		}
	} else {
		internal.Prompt("PAAS资源")
	}

	/*--------------------------------------------设置单格式------------------------------------------------*/
	//自动设置列宽

	// 10 是列的长度，10 = K
	for i := 0; i < 11; i++ {
		var runeLetter = 'A' + int32(i)                      //每次循环A+1 = B +1  = C
		AutoWidth(countExcelFileName, string(runeLetter), i) //colStr = A 那colInt 必须等于0
	}
	//自动设置边框
	SetCollStyle(countExcelFileName)

	/*----------------------------------------------end---------------------------------------------------*/
	logrus.Printf("已在项目目录下生成\"%v\"\n", countExcelFileName)
}
