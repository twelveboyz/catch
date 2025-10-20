package cmdb

import (
	"Catch/internal"
	"fmt"
	"log"
)

func Run() {
	//手动输入相应的目录路径

	root, projectName := internal.UserInputRoot()

	//定义临时目录和统计表格名称
	temp := "temp"
	projectFile := fmt.Sprintf("/%v.xlsx", projectName)
	countExcel := temp + projectFile

	//创建临时目录并把项目复制到临时目录中
	err := internal.CreateTempDirAndCopy(root, temp)
	if err != nil {
		log.Println(err)
		return
	}

	defer func() {
		err = internal.MoveFile(countExcel, root+projectFile)
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
	Insert(countExcel, projectName, messages)

	//日志
	logfile := internal.LogFunc(root)
	defer func() {
		if err := logfile.Close(); err != nil {
			log.Println("关闭日志文件失败")
		}
	}()

	log.Println("------ start CMDB ------")
	/*-------------------------------------------资源录入模板-------------------------------------------------*/
	fileName := internal.CMDBFilterFileName(root, "CMDB录入表", "资源录入")

	if fileName != "" {
		excel := NewExcel(fileName, "资源录入")

		resultExcel := NewExcel(countExcel, "Sheet1")
		resultExcel.GenerateStatisticalFile(countExcel, excel.ResourceStatistics())

	} else {
		internal.Prompt("资源录入模板")
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
}
