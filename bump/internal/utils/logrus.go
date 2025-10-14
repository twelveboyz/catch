package utils

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/windows"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var globalRow = new(string)

func init() {
	// 启用 Windows 控制台虚拟终端（VT）处理
	stdout := windows.Handle(os.Stdout.Fd())
	var originalMode uint32

	windows.GetConsoleMode(stdout, &originalMode)
	windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)

	logrus.SetLevel(logrus.DebugLevel)

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		DisableTimestamp: false,
	})

}

func LogFunc(project string) *os.File {
	var fileTimeFormat = "2006-01-02"
	var fileName string
	err := os.Mkdir("log", os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatal("创建目录失败", err)
	}
	fileName = fmt.Sprintf("log/%v-%v.log", time.Now().Format(fileTimeFormat), filepath.Base(project))

	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println("文件写入失败", err)
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	logrus.SetOutput(multiWriter)
	return logFile
}

func PrintDebugFFun(comment string, row string, excelResource string, consoleResource string) {
	logrus.Debugf("[%s]\t\trow:%s excel:%s console:%s\n", comment, row, excelResource, consoleResource)
}

func PrintWarnFFun(comment string, row string, excelResource string, consoleResource string) {
	if logrus.GetLevel().String() == "info" {
		if *globalRow == "" {
			*globalRow = row
			logrus.Infof("------------------------------ Row:%s ------------------------------", row)
		} else if *globalRow != row {
			*globalRow = row
			logrus.Infof("------------------------------ Row:%s ------------------------------", row)
		}
	}
	logrus.Warnf("[%s]\t\trow:%s excel:%s console:%s\n", comment, row, excelResource, consoleResource)
}

func GetLogLevel() logrus.Level {
	for {
		fmt.Print("日志等级 [debug / info] (default=debug): ")
		inp, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %s\n", err)
		}

		inp = strings.TrimSpace(inp)
		if inp == "" {
			return logrus.Level(5)
		} else if inp == "debug" {
			return logrus.Level(5)
		} else if inp == "info" {
			return logrus.Level(4)
		} else {
			continue
		}
	}
}
