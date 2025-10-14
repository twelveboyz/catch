package config

import (
	"Catch/internal"
	"log"
)

func Run() {
	//手动输入相应的目录路径
	root := internal.InputRoot()

	temp := "temp"
	countXlsx := temp + "/建设信息统计表.xlsx"
	dest := root + "/建设信息统计表.xlsx"

	//创建临时目录并把项目复制到临时目录中
	err := internal.CreateTempDirAndCopy(root, temp)
	if err != nil {
		log.Println(err)
		return
	}


	defer func() {
		err = internal.MoveFile(countXlsx, dest)
		if err != nil {
			log.Println(err)
			return
		}

		err = internal.CleanTempDir(temp)
		if err != nil {
			log.Println(err)
			return
		}
	}()

	//创建一个新的excel表格
	CreateXlsx(countXlsx)

	//插入excel表格表头数据
	Insert(countXlsx)

	//日志
	logfile := internal.LogFunc(root)
	defer func() {
		if err := logfile.Close(); err != nil {
			log.Println("关闭日志文件失败")
		}
	}()

	/*----------------------------------------------计算----------------------------------------------*/
	ComPutFileName := internal.ConfigFilterFileName(root, ".xls", "计算资源")

	if ComPutFileName != "" {
		//创建xlsx对象
		comPutExcel := &Excel{Filename: ComPutFileName, Sheet: "", CopyLine: 0, TitleLine: 3, ResourceType: "计算资源"}
		//打开一个xlsx文件
		f := comPutExcel.OpenFile()
		defer func() {
			if err := f.Close(); err != nil {
				log.Fatal("文件关闭失败", err)
			}
		}()

		//获取表格所有的sheet
		sheets := comPutExcel.ForGetSheet(f)
		log.Println("获取有效的Sheet=", sheets)
		for _, sheet := range sheets {
			log.Println("sheet=", sheet)
			LineInt := DetermineCase(ComPutFileName, sheet)
			comPutExcel.CopyLine = LineInt
			comPutExcel.Sheet = sheet
			comPutExcel.ComPutExec(root, countXlsx, f)
		}

	} else {
		internal.Prompt("计算资源")
	}

	/*----------------------------------------------系统盘----------------------------------------------*/
	SystemDiskFileName := internal.ConfigFilterFileName(root, ".xls", "计算资源")

	if ComPutFileName != "" {
		//创建xlsx对象
		systemDiskexcel := &Excel{Filename: SystemDiskFileName, Sheet: "", CopyLine: 0, TitleLine: 3, ResourceType: "系统盘资源"}
		//打开一个xlsx文件
		f := systemDiskexcel.OpenFile()
		defer func() {
			if err := f.Close(); err != nil {
				log.Fatal("文件关闭失败", err)
			}
		}()

		//获取表格所有的sheet
		sheets := systemDiskexcel.ForGetSheet(f)
		log.Println("获取有效的Sheet=", sheets)
		for _, sheet := range sheets {
			log.Println("sheet=", sheet)
			LineInt := DetermineCase(SystemDiskFileName, sheet)
			systemDiskexcel.CopyLine = LineInt
			systemDiskexcel.Sheet = sheet
			systemDiskexcel.SystemDiskExec(root, countXlsx, f)
		}

	} else {
		internal.Prompt("系统盘资源")
	}

	/*----------------------------------------------存储------------------------------------------------*/
	storageFileName := internal.ConfigFilterFileName(root, ".xls", "存储资源")

	if storageFileName != "" {
		storageExcel := &Excel{Filename: storageFileName, Sheet: "", CopyLine: 0, TitleLine: 3, ResourceType: "存储资源"}
		f := storageExcel.OpenFile()
		defer func() {
			if err := f.Close(); err != nil {
				log.Fatal("文件关闭失败", err)
			}
		}()

		//获取表格所有的sheet
		sheets := storageExcel.ForGetSheet(f)
		log.Println("获取有效的Sheet=", sheets)
		for _, sheet := range sheets {
			log.Println("sheet=", sheet)
			LineInt := DetermineCase(storageFileName, sheet)
			storageExcel.CopyLine = LineInt
			storageExcel.Sheet = sheet
			storageExcel.StorageExec(root, countXlsx, f)
		}

	} else {
		internal.Prompt("存储资源")
	}

	/*----------------------------------------------云网络---------------------------------------------------*/
	NetworkFileName := internal.ConfigFilterFileName(root, ".xls", "网络")

	if NetworkFileName != "" {
		networkExcel := &Excel{Filename: NetworkFileName, Sheet: "", CopyLine: 0, TitleLine: 3, ResourceType: "网络资源"}

		f := networkExcel.OpenFile()
		defer func() {
			if err := f.Close(); err != nil {
				log.Fatal("文件关闭失败", err)
			}
		}()
		//获取表格所有的sheet
		sheets := networkExcel.ForGetSheet(f)
		log.Println("获取有效的Sheet=", sheets)
		for _, sheet := range sheets {
			log.Println("sheet=", sheet)
			LineInt := DetermineCase(NetworkFileName, sheet)
			networkExcel.CopyLine = LineInt
			networkExcel.Sheet = sheet
			networkExcel.NetworkExec(root, countXlsx, f)
		}
	} else {
		internal.Prompt("云网络资源")
	}

	/*----------------------------------------------云安全---------------------------------------------------*/
	SecurityFileName := internal.ConfigFilterFileName(root, ".xls", "安全")

	if SecurityFileName != "" {
		securityExcel := &Excel{Filename: SecurityFileName, Sheet: "", CopyLine: 0, TitleLine: 3, ResourceType: "安全资源"}

		f := securityExcel.OpenFile()
		defer func() {
			if err := f.Close(); err != nil {
				log.Fatal("文件关闭失败", err)
			}
		}()

		//获取表格所有的sheet
		sheets := securityExcel.ForGetSheet(f)
		log.Println("获取有效的Sheet=", sheets)
		for _, sheet := range sheets {
			log.Println("sheet=", sheet)
			LineInt := DetermineCase(SecurityFileName, sheet)
			securityExcel.CopyLine = LineInt
			securityExcel.Sheet = sheet
			securityExcel.SecurityExec(root, countXlsx, f)
		}

	} else {
		internal.Prompt("云安全资源")
	}

	/*----------------------------------------------PAAS---------------------------------------------------*/
	PaasFileName := internal.ConfigFilterFileName(root, ".xls", "PAAS")

	if PaasFileName != "" {
		PaasExcel := &Excel{Filename: PaasFileName, Sheet: "", CopyLine: 0, TitleLine: 3, ResourceType: "PAAS资源"}

		f := PaasExcel.OpenFile()
		defer func() {
			if err := f.Close(); err != nil {
				log.Fatal("文件关闭失败", err)
			}
		}()

		//获取表格所有的sheet
		sheets := PaasExcel.ForGetSheet(f)
		log.Println("获取有效的Sheet=", sheets)
		for _, sheet := range sheets {
			log.Println("sheet=", sheet)
			LineInt := DetermineCase(PaasFileName, sheet)
			PaasExcel.CopyLine = LineInt
			PaasExcel.Sheet = sheet
			PaasExcel.PAASExec(root, countXlsx, f)
		}

	} else {
		internal.Prompt("PAAS")
	}

	/*--------------------------------------------设置单格式------------------------------------------------*/
	//自动设置列宽

	// 10 是列的长度，10 = J
	for i := 0; i < 10; i++ {
		var runeLetter = 'A' + int32(i)             //每次循环A+1 = B +1  = C
		AutoWidth(countXlsx, string(runeLetter), i) //colStr = A 那colInt 必须等于0
	}
	//自动设置边框
	SetCollStyle(countXlsx)

	/*----------------------------------------------end---------------------------------------------------*/
	log.Printf("已在项目目录下生成\"%v\"\n", countXlsx)
}
