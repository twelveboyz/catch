package cmdb

import (
	"Catch/internal"
	"fmt"
	"log"
	"strings"
	"time"
)

var TL = 2
var CL = 3

func Run() {
	//手动输入相应的目录路径
	root := internal.InputRoot()
	project := internal.RegexFunc(internal.RegexStr, root)
	if project == "" {
		project = internal.RegexFunc(internal.RegexHorizontalBar, root)
		project = strings.Replace(project, "-", "_", 1)
	}

	temp := "temp"
	projectName := fmt.Sprintf("/%v.xlsx", project)
	countExcel := temp + projectName



	//创建临时目录并把项目复制到临时目录中
	err := internal.CreateTempDirAndCopy(root, temp)
	if err != nil {
		log.Println(err)
		return
	}


	defer func ()  {
		err = internal.MoveFile(countExcel, root+projectName)
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
	err = CreateXlsx(countExcel)
	if err != nil {
		log.Println(err)
		return
	}

	//打印
	messages := internal.InputMessages()
	//插入excel表格表头数据
	Insert(countExcel, project, messages)

	//日志
	logfile := internal.LogFunc(root)
	defer func() {
		if err := logfile.Close(); err != nil {
			log.Println("关闭日志文件失败")
		}
	}()

	log.Println("------ start CMDB ------")
	/*----------------------------------------------计算----------------------------------------------*/
	ComPutFileName := internal.CMDBFilterFileName(root, "CMDB录入表", "计算资源")

	if ComPutFileName != "" {
		//创建excel对象
		comPutExcel := &Excel{Filename: ComPutFileName, Sheet: "", CopyLine: CL, TitleLine: TL, ResourceType: "计算资源"}
		//打开一个excel文件
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
			comPutExcel.Sheet = sheet
			comPutExcel.ComPutExec(root, countExcel, messages, f)

		}

	} else {
		internal.Prompt("计算资源")
	}

	/*----------------------------------------------系统盘----------------------------------------------*/
	SystemDiskFileName := internal.CMDBFilterFileName(root, "CMDB录入表", "计算资源")

	if ComPutFileName != "" {
		//创建excel对象
		SystemDiskExcel := &Excel{Filename: SystemDiskFileName, Sheet: "", CopyLine: CL, TitleLine: TL, ResourceType: "系统盘资源"}
		//打开一个excel文件
		f := SystemDiskExcel.OpenFile()
		defer func() {
			if err := f.Close(); err != nil {
				log.Fatal("文件关闭失败", err)
			}
		}()

		//获取表格所有的sheet
		sheets := SystemDiskExcel.ForGetSheet(f)
		log.Println("获取有效的Sheet=", sheets)
		for _, sheet := range sheets {
			log.Println("sheet=", sheet)
			SystemDiskExcel.Sheet = sheet
			SystemDiskExcel.SystemDiskExec(root, countExcel, messages, f)
		}

	} else {
		internal.Prompt("系统盘资源")
	}

	/*----------------------------------------------存储------------------------------------------------*/
	storageFileName := internal.CMDBFilterFileName(root, "CMDB录入表", "存储资源")

	if storageFileName != "" {
		storageExcel := &Excel{Filename: storageFileName, Sheet: "", CopyLine: CL, TitleLine: TL, ResourceType: "存储资源"}
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
			storageExcel.Sheet = sheet
			storageExcel.StorageExec(root, countExcel, messages, f)
		}

	} else {
		internal.Prompt("存储资源")
	}

	/*----------------------------------------------云网络---------------------------------------------------*/
	NetworkFileName := internal.CMDBFilterFileName(root, "CMDB录入表", "网络")

	if NetworkFileName != "" {
		networkExcel := &Excel{Filename: NetworkFileName, Sheet: "", CopyLine: CL, TitleLine: TL, ResourceType: "网络资源"}

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
			networkExcel.Sheet = sheet
			networkExcel.NetworkExec(root, countExcel, messages, f)

		}

	} else {
		internal.Prompt("云网络资源")
	}

	/*----------------------------------------------云安全---------------------------------------------------*/
	securityFileName := internal.CMDBFilterFileName(root, "CMDB录入表", "安全")

	if securityFileName != "" {
		securityExcel := &Excel{Filename: securityFileName, Sheet: "", CopyLine: CL, TitleLine: TL, ResourceType: "安全资源"}

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
			securityExcel.Sheet = sheet
			securityExcel.SecurityExec(root, countExcel, messages, f)
		}

	} else {
		internal.Prompt("云安全资源")
	}

	/*----------------------------------------------PAAS---------------------------------------------------*/
	PaasFileName := internal.CMDBFilterFileName(root, "CMDB录入表", "PAAS")

	if PaasFileName != "" {
		PaasExcel := &Excel{Filename: PaasFileName, Sheet: "", CopyLine: CL, TitleLine: TL, ResourceType: "PAAS资源"}

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
			PaasExcel.Sheet = sheet
			PaasExcel.PAASExec(root, countExcel, messages, f)
		}

	} else {
		internal.Prompt("PAAS")
	}

	/*--------------------------------------------设置单格式------------------------------------------------*/
	//自动设置列宽

	// 10 是列的长度，10 = J
	for i := 0; i < 10; i++ {
		var runeLetter = 'A' + int32(i)              //每次循环A+1 = B +1  = C
		AutoWidth(countExcel, string(runeLetter), i) //colStr = A 那colInt 必须等于0
	}
	//自动设置边框
	SetCollStyle(countExcel)

	/*----------------------------------------------end---------------------------------------------------*/
	log.Printf("已在项目目录下生成\"%v\"\n", countExcel)
	log.Println("------ end ------\n")
	time.Sleep(100 * time.Millisecond)
}
