package daily

import (
	"Catch/internal"
	"fmt"
	"log"
	"time"
)

func Run() {
	//手动输入相应的目录路径
	root := internal.InputRoot()
	temp := "temp"
	fileName := "/资源摘要统计表.xlsx"
	filePath := temp + fileName
	sheetName := "Sheet1"

	//创建临时目录并把项目复制到临时目录中
	err := internal.CreateTempDirAndCopy(root, temp)
	if err != nil {
		log.Println(err)
		return
	}

	defer func() {
		err = internal.MoveFile(filePath, root+fileName)
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
	err = CreateXlsx(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	//日志
	logfile := internal.LogFunc(root)
	defer func() {
		if err := logfile.Close(); err != nil {
			log.Println("关闭日志文件失败")
		}
	}()

	log.Println("------ start review------")

	CompRun(root, filePath, sheetName)
	StorageRun(root, filePath, sheetName)
	NetworkRun(root, filePath, sheetName)
	SecurityRun(root, filePath, sheetName)
	PaaSRun(root, filePath, sheetName)

	WriteDailyNewspaperInfo(root, filePath)

	log.Printf("已在项目目录下生成\"%v\"\n", filePath)
	log.Println("------ end ------\n")
	time.Sleep(100 * time.Millisecond)
}
