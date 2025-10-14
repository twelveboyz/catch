package utils

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"strings"
	"time"
)

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

func Cache() string {
	cache := ".cache"

	CreateFile(cache)

	cacheFile := ReadFile(cache)

	buf := bufio.NewReader(os.Stdin)
	fmt.Printf("Token (cache=%s) :", cacheFile)
	input, err := buf.ReadBytes('\n')
	if err != nil {
		fmt.Println("读取异常", err)
	}

	inputStr := string(input)
	inputStr = strings.TrimSpace(inputStr)

	if inputStr == "" {
		return ReadFile(cache)

	} else {
		f, err := os.OpenFile(cache, os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("打开文件失败", err)
		}

		defer func() {
			err = f.Close()
			if err != nil {
				fmt.Println("关闭失败", err)
			}
		}()

		err = f.Truncate(0)
		if err != nil {
			fmt.Println("清空cookie文件失败")
		}

		_, err = f.WriteString(inputStr)
		if err != nil {
			fmt.Println("写入失败", err)
		}
		return inputStr
	}
}

func ReadFile(filename string) string {
	f, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("读取失败", err)
	}
	return string(f)
}

func CreateFile(filename string) {
	if _, err := os.Stat(filename); err == nil {
		return
	}

	f, err := os.Create(filename)
	if err != nil {
		logrus.Println(err)
		return
	}
	f.Close()
}

func InputUUIDs() []string {
	fmt.Print("请输入UUID(使用英文句号隔开)：")
	input := bufio.NewReader(os.Stdin)

	str, _ := input.ReadString('\n')
	str = strings.TrimSpace(str)

	strSlice := strings.Split(str, ",")

	return strSlice
}

func InputProjectID() string {
	fmt.Print("请输入工程编号：")
	input := bufio.NewReader(os.Stdin)

	str, _ := input.ReadString('\n')
	str = strings.TrimSpace(str)

	return str
}

func InputChoose() string {
	for {
		inputReader := bufio.NewReader(os.Stdin)
		fmt.Print("1.根据云主机工程编号查询资源（d157_） 2.根据UUID查询资源（使用英文句号分割） [1 / 2]:")
		input, err := inputReader.ReadString('\n')
		if err != nil {
			log.Println("获取输入内容失败：", err)
		}
		input = strings.TrimSpace(input)

		if input == "1" || input == "2" {
			return input
		} else if input == "q" {
			os.Exit(0)
		} else {
			continue
		}
	}
}
