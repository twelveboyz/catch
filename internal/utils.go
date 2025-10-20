package internal

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// CMDBFilterFileName
// 根据传入两个key来匹配定位文件名称
func CMDBFilterFileName(root string, key1 string, key2 string) string {
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
		//获取文件相对的路径
		reFile := RegexFunc("CMDB录入表.*", file)

		if key2 == "PAAS" {
			upperFile := strings.ToUpper(reFile)
			//判断条件如果key1 key2匹配上则返回改文件绝对路径
			fileBool = strings.Contains(reFile, key1) && strings.Contains(upperFile, key2) && !strings.Contains(reFile, "~$")
		} else {
			fileBool = strings.Contains(reFile, key1) && strings.Contains(reFile, key2) && !strings.Contains(reFile, "~$")
		}

		if fileBool {
			log.Println("正则抓取的相对路径是：", reFile)
			log.Println("最终捕抓文件名称路径是：", file)
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
			log.Println("自动捕抓文件名称是：", file)
			return file
		}
	}
	return ""
}

// DailyFilterFileName
// 根据传入两个key来匹配定位文件名称
func DailyFilterFileName(root string, key1 string, key2 string) string {

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
		//获取文件相对的路径
		reFile := RegexFunc("资源订购配置.*", file)

		if key2 == "PAAS" {
			upperFile := strings.ToUpper(reFile)
			//判断条件如果key1 key2匹配上则返回改文件绝对路径
			fileBool = strings.Contains(reFile, key1) && strings.Contains(upperFile, key2) && !strings.Contains(reFile, "~$")
		} else {
			fileBool = strings.Contains(reFile, key1) && strings.Contains(reFile, key2) && !strings.Contains(reFile, "~$")
		}

		if fileBool {
			log.Println("正则抓取的相对路径是：", reFile)
			log.Println("最终捕抓文件名称路径是：", file)
			return file
		}
	}
	return ""
}

func RegexFunc(regex string, Str string) string {
	r, err := regexp.Compile(regex)
	if err != nil {
		fmt.Println("regexp.compile失败", err)
	}
	cat := r.FindString(Str)

	return cat
}

func InputRoot() string {
	for {
		inputReader := bufio.NewReader(os.Stdin)
		fmt.Print("目录路径：")
		input, err := inputReader.ReadString('\n')
		if err != nil {
			log.Println("获取输入内容失败：", err)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		input = strings.TrimSpace(input)

		fInfo, err := os.Stat(input)
		if err != nil {
			log.Println("输入的目录路径有误:", err)
			time.Sleep(100 * time.Millisecond)
			continue
		}

		//需要判断是否是目录
		dirBool := fInfo.IsDir()
		if !dirBool {
			log.Println("请选择项目根目录!")
			time.Sleep(100 * time.Millisecond)
			continue
		}

		return input
	}
}

func UserInputRoot() (string, string) {
	root := InputRoot()
	projectName := RegexFunc(RegexStr, root)
	if projectName == "" {
		projectName = RegexFunc(RegexHorizontalBar, root)
		projectName = strings.Replace(projectName, "-", "_", 1)
	}

	return root, projectName
}

func InputChoose() string {
	for {
		inputReader := bufio.NewReader(os.Stdin)
		fmt.Print("1.统计本批次交维信息 2.汇总本次建设摘要 3.汇总本次建设信息 4.资源字段检查 5.查询云主机字段，请输入(1 / 2 / 3 / 4 / 5): ")
		input, err := inputReader.ReadString('\n')
		if err != nil {
			log.Println("获取输入内容失败：", err)
		}
		input = strings.TrimSpace(input)

		if input == "1" || input == "2" || input == "3" || input == "4" || input == "5" {
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
	for i := 1; i <= 6; i++ {
		inputReader := bufio.NewReader(os.Stdin)
		switch i {
		case 1:
			fmt.Print("是否分批（默认为\"不分配\"，如分批请输入\"第x批\"）：")
			input, err := inputReader.ReadString('\n')
			input = strings.TrimSpace(input)
			if err != nil {
				log.Println("获取输入内容失败：", err)
			}
			if input == "" {
				InputMap["InBatches"] = "不分批"
			} else {
				InputMap["InBatches"] = input
			}
		case 2:
			fmt.Print("需求人：")
			input, err := inputReader.ReadString('\n')
			input = strings.TrimSpace(input)
			if err != nil {
				log.Println("获取输入内容失败：", err)
			}
			InputMap["DemandPerson"] = input

		case 3:
			fmt.Print("需求方：")
			input, err := inputReader.ReadString('\n')
			input = strings.TrimSpace(input)
			if err != nil {
				log.Println("获取输入内容失败：", err)
			}
			InputMap["Demander"] = input

		case 4:
			fmt.Print("OA发文时间：")
			input, err := inputReader.ReadString('\n')
			input = strings.TrimSpace(input)
			if err != nil {
				log.Println("获取输入内容失败：", err)
			}
			InputMap["OATime"] = input

		case 5:
			fmt.Print("主账号：")
			input, err := inputReader.ReadString('\n')
			input = strings.TrimSpace(input)
			if err != nil {
				log.Println("获取输入内容失败：", err)
			}
			InputMap["PrimaryAccount"] = input
		case 6:
			fmt.Print("新增子网信息：")
			input, err := inputReader.ReadString('\n')
			input = strings.TrimSpace(input)
			if err != nil {
				log.Println("获取输入内容失败：", err)
			}
			InputMap["SubnetInfo"] = input
		}
	}
	return InputMap
}

func Prompt(p string) {
	log.Printf("\n-----------------------------------------\n未匹配到%v的Excel表格,已跳过···\n-----------------------------------------\n", p)
}

func Tip() {
	fmt.Println(`
|------------------------------------------------------------------------------------------------------------------|
|version: 1.6.1
|注意事项：                                                                                 
|一、目录名称命名需要按照"项目ID_项目名称"，命名例子:                                               
|    - C284_短信平台黑名单模块资源扩容需                                                        
|    - 移动云-C284_短信平台黑名单模块资源扩容需求-交维材料                                         
|二、自动抓取Excel表格内容是根据文件夹和文件名称抓取
|    1.统计本批次交维信息
|    - 目录名称要包含: "CMDB录入表" , 文件名称要包含: "资源录入"         
|    2.汇总本次建设摘要
|    - 文件名称要包含: "计算资源"、"存储资源"、"网络"、"安全"、"PAAS"
|    3.汇总本次建设信息
|    - 文件名称要包含: "计算资源"、"存储资源"、"网络"、"安全"、"PAAS"
|三、移动云产品资源字段检查
|    已支持产品：云主机、裸金属、云硬盘、文件存储、并行文件存储、公网IP、ipv6、Nat网关、负载均衡、云端口(需主账号权限)、云专线
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
		log.Fatal("创建目录失败", err)
	}
	fileName = fmt.Sprintf("log/%v-%v-Catch日志.txt", time.Now().Format(fileTimeFormat), filepath.Base(project))

	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println("文件写入失败", err)
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)
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
