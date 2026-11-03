package internal

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/windows"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// CMDBFilterFileName
// 根据传入两个key来匹配定位文件名称
func CMDBFilterFileName(root string, DirKey string, FileKey []string) string {

	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		var fileBool bool
		//获取文件相对的路径

		reFile := RegexFunc(RegexMatchFileName(DirKey), file)

		if reFile == "" {
			continue
		} else {
			//fmt.Println("reFile:", reFile)
		}

		//循环Key只要匹配中一个则认为符合条件
		for _, key := range FileKey {
			if strings.Contains(reFile, key) && strings.Contains(reFile, ".xls") && !strings.Contains(reFile, "~$") {
				fileBool = true
				break
			}
		}

		if fileBool {
			logrus.Println("捕抓文件名称路径：", file)
			return file
		}
	}
	return ""
}

// ConfigFilterFileName
// 根据传入两个key来匹配定位文件名称
func ConfigFilterFileName(root string, key1 string, key2 string) string {
	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		var fileBool bool
		var reFile = filepath.Base(file)
		//fmt.Println("file=", reFile, file)
		if key2 == "PAAS" {
			UpperFile := strings.ToUpper(reFile)
			fileBool = strings.Contains(reFile, key1) && strings.Contains(UpperFile, key2) && !strings.Contains(reFile, "~$")
		} else {
			fileBool = strings.Contains(reFile, key1) && strings.Contains(reFile, key2) && !strings.Contains(reFile, "~$")
		}
		if fileBool {
			logrus.Println("自动捕抓文件名称是：", file)
			return file
		}
	}
	return ""
}

func RegexFunc(regex string, Str string) string {
	r := regexp.MustCompile(regex)

	file := r.FindStringSubmatch(Str)

	if len(file) == 0 {
		return ""
	} else if len(file) == 2 {
		return file[1]
	} else {
		return ""
	}

}

func InputRoot() string {
	//命令行模式下直接使用传入的目录，不再读取 stdin（合法性由入口统一校验）
	if o := FromCLI(); o != nil {
		return o.Root
	}

	for {
		inputReader := bufio.NewReader(os.Stdin)
		fmt.Print("目录路径：")
		input, err := inputReader.ReadString('\n')
		if err != nil {
			logrus.Println("获取输入内容失败：", err)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		input = strings.TrimSpace(input)

		fInfo, err := os.Stat(input)
		if err != nil {
			logrus.Println("输入的目录路径有误:", err)
			time.Sleep(100 * time.Millisecond)
			continue
		}

		//需要判断是否是目录
		dirBool := fInfo.IsDir()
		if !dirBool {
			logrus.Println("请选择项目根目录!")
			time.Sleep(100 * time.Millisecond)
			continue
		}

		return input
	}
}

func UserInputRoot() (string, string) {
	root := InputRoot()
	projectName := regexp.MustCompile(RegexProjectBatch).FindString(root)
	projectName = strings.Replace(projectName, "-", "_", 1)

	return root, projectName
}

func InputChoose() string {
	for {
		inputReader := bufio.NewReader(os.Stdin)
		fmt.Print("1.统计本批次交维信息 2.汇总本次建设摘要 3.汇总本次建设信息，请输入(1 / 2 / 3): ")
		input, err := inputReader.ReadString('\n')
		if err != nil {
			logrus.Println("获取输入内容失败：", err)
		}
		input = strings.TrimSpace(input)

		if input == "1" || input == "2" || input == "3" || input == "4" {
			return input
		} else if input == "q" {
			os.Exit(0)
		} else {
			continue
		}
	}
}

func InputMessages() map[string]string {
	var InputMap = make(map[string]string)
	//命令行模式下直接从参数取值
	if o := FromCLIWithDefaults(); o != nil {
		InputMap["InBatches"] = o.InBatches
		return InputMap
	}

	inputReader := bufio.NewReader(os.Stdin)

	fmt.Print("是否分批（默认为\"不分配\"，如分批请输入\"第x批\"）：")
	input, err := inputReader.ReadString('\n')
	input = strings.TrimSpace(input)
	if err != nil {
		logrus.Println("获取输入内容失败：", err)
	}
	if input == "" {
		InputMap["InBatches"] = "不分批"
	} else {
		InputMap["InBatches"] = input
	}

	return InputMap
}

func InputMessagesWithDefaults(defaults map[string]string) map[string]string {
	var InputMap = make(map[string]string)
	//命令行模式下直接从参数取值
	if o := FromCLIWithDefaults(); o != nil {
		InputMap["InBatches"] = o.InBatches
		return InputMap
	}

	inputReader := bufio.NewReader(os.Stdin)

	defaultValue := "不分批"
	if val, ok := defaults["InBatches"]; ok {
		defaultValue = val
	}
	fmt.Printf("是否分批（默认为\"%s\"，如分批请输入\"第x批\"）：", defaultValue)
	input, err := inputReader.ReadString('\n')
	input = strings.TrimSpace(input)
	if err != nil {
		logrus.Println("获取输入内容失败：", err)
	}
	if input == "" {
		InputMap["InBatches"] = defaultValue
	} else {
		InputMap["InBatches"] = input
	}

	return InputMap
}

func Prompt(p string) {
	logrus.Printf("\n-----------------------------------------\n未匹配到%v的Excel表格,已跳过···\n-----------------------------------------\n", p)
}

func Tip() {
	fmt.Println(`
|------------------------------------------------------------------------------------------------------------------|
|version: 1.6.8
|注意事项：                                                                                 
|一、目录名称命名需要按照"项目ID_项目名称"，命名例子:                                               
|    - C284_短信平台黑名单模块资源扩容需                                                        
|    - 移动云-C284_短信平台黑名单模块资源扩容需求-交维材料                                         
|二、自动抓取Excel表格内容是根据文件夹和文件名称抓取
|    1：统计本批次交维信息
|    - 目录名称要包含: "资源录入表" , 文件名称要包含: "资源录入" 
|    - 目录名称要包含: "订购评审纪要" , 文件名称要包含: "安全资源"
|    2：汇总本次建设摘要
|    - 文件名称要包含: "计算资源"、"存储资源"、"网络"、"安全"、"PAAS"
|    3：汇总本次建设信息
|    - 文件名称要包含: "计算资源"、"存储资源"、"网络"、"安全"、"PAAS"
|                                                                                               deployer: Linzheng
|------------------------------------------------------------------------------------------------------------------|`,
	)
	fmt.Println()
}

func LogFunc(project string) *os.File {
	var fileTimeFormat = "2006-01-02"
	var fileName string
	err := os.Mkdir("log", os.ModePerm)
	if err != nil && !os.IsExist(err) {
		logrus.Fatal("创建目录失败", err)
	}

	fileName = fmt.Sprintf("log/%v-%v-日志.txt", time.Now().Format(fileTimeFormat), filepath.Base(project))

	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println("文件写入失败", err)
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	logrus.SetOutput(multiWriter)
	return logFile
}

func CreateTempDirAndCopy(projectPath, tempDir string) error {
	fs := os.DirFS(projectPath)
	err := os.CopyFS(tempDir, fs)
	if err != nil {
		cleanErr := CleanTempDir(tempDir)
		if cleanErr != nil {
			return fmt.Errorf("copy failed: %v (cleanup also failed: %v)", err, cleanErr)
		}
		return fmt.Errorf("copy failed: %v", err)
	}
	return nil
}

func MoveAndClearFile(source, destination, clearTempDir string) error {
	err := MoveFile(source, destination)
	if err != nil {
		return fmt.Errorf("move err:%s", err)
	}

	err = CleanTempDir(clearTempDir)
	if err != nil {
		return fmt.Errorf("remove err:%s", err)
	}

	return nil
}

func CleanTempDir(tempDir string) error {
	err := os.RemoveAll(tempDir)
	if err != nil {
		return fmt.Errorf("remove err:%s", err)
	}
	return nil
}

func MoveFile(source, destination string) error {

	// 打开源文件
	srcFile, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("无法打开源文件: %v", err)
	}
	defer srcFile.Close()

	// 创建目标文件
	destFile, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("无法创建目标文件: %v", err)
	}
	defer destFile.Close()

	// 复制文件内容
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return fmt.Errorf("复制文件内容失败: %v", err)
	}

	// 确保数据写入磁盘
	err = destFile.Sync()
	if err != nil {
		return fmt.Errorf("同步文件失败: %v", err)
	}

	return nil
}

func AllSame[T comparable](slice []T) bool {
	if len(slice) == 0 {
		return true
	}

	first := slice[0]
	for _, element := range slice[1:] {
		if element != first {
			return false
		}
	}
	return true
}

func init() {
	// 启用 Windows 控制台虚拟终端（VT）处理
	stdout := windows.Handle(os.Stdout.Fd())
	var originalMode uint32

	windows.GetConsoleMode(stdout, &originalMode)
	windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)

	logrus.SetLevel(logrus.InfoLevel)

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		FullTimestamp:    true,
		DisableTimestamp: false,
		TimestampFormat:  "2006-01-02 15:04:05",
	})
}
