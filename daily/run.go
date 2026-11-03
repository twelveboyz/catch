package daily

import (
	"Catch/internal"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func Run() error {
	//手动输入相应的目录路径
	//root := internal.InputRoot()
	root, projectName := internal.UserInputRoot()

	temp := "temp"
	fileName := fmt.Sprintf("周报/%s-资源摘要统计表.xlsx", projectName)
	sheetName := "Sheet1"

	_ = os.Mkdir("周报", os.ModePerm)

	//创建临时目录并把项目复制到临时目录中
	err := internal.CreateTempDirAndCopy(root, temp)
	if err != nil {
		logrus.Error(err)
		return err
	}

	defer func() {
		err = internal.CleanTempDir(temp)
		if err != nil {
			logrus.Error(err)
			return
		}
	}()

	//创建一个新的excel表格
	err = CreateXlsx(fileName)
	if err != nil {
		logrus.Error(err)
		return err
	}

	//日志
	logfile := internal.LogFunc(root)
	defer func() {
		if err := logfile.Close(); err != nil {
			logrus.Println("关闭日志文件失败")
		}
	}()

	logrus.Println("------ start review------")

	CompRun(root, fileName, sheetName)
	StorageRun(root, fileName, sheetName)
	NetworkRun(root, fileName, sheetName)
	SecurityRun(root, fileName, sheetName)
	PaaSRun(root, fileName, sheetName)

	WriteDailyNewspaperInfo(root, fileName)

	logrus.Printf("已在项目目录下生成\"%v\"\n", fileName)
	logrus.Println("------ end ------\n")
	time.Sleep(100 * time.Millisecond)

	return nil
}
