package resourceCount

import (
	"Catch/internal"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func loopFindProjectBatchResources() ([]string, error) {
	root, _ := internal.UserInputRoot()
	found, notfound, err := findDirProjectFiles(root)
	if err != nil {
		return nil, fmt.Errorf("LoopFindProjectBatchResources error: %w", err)
	}

	//二次查找未找到xlsx文件的目录下的子目录
	var towNotFoundDirs []string
	for _, notfoundDir := range notfound {
		logrus.Println("----------------------------------------------------------------")
		towFound, _, err := findDirProjectFiles(notfoundDir)
		if err != nil {
			return nil, fmt.Errorf("LoopFindProjectBatchResources error: %w", err)
		} else if len(towFound) == 0 {
			towNotFoundDirs = append(towNotFoundDirs, notfoundDir)
		}
		found = append(found, towFound...)
	}

	if len(towNotFoundDirs) != 0 {
		logrus.Errorf("not found xlsx files total:%v directories:\n%v", len(towNotFoundDirs), strings.Join(towNotFoundDirs, "\n"))
	}

	logrus.Infoln("Total xlsx files found:", len(found))

	return found, nil
}

// findDirProjectFiles 查找指定目录下的项目批次目录及其xlsx文件, 返回找到的xlsx文件路径和未找到xlsx文件的目录路径
func findDirProjectFiles(root string) ([]string, []string, error) {
	//获取所有的项目批次目录
	projectBatchDirs, err := fetchProjectBatchDirs(root)
	if err != nil {
		return nil, nil, fmt.Errorf("findDirProjectFiles error: %w", err)
	}

	//获取所有项目批次目录下的xlsx文件
	xlsxFiles, notFoundDirs, err := loopFindProjectBatchXlsxFiles(projectBatchDirs)
	if err != nil {
		return nil, nil, fmt.Errorf("findDirProjectFiles error: %w", err)
	}

	//读取xlsx文件的资源内容，并统计资源数量
	return xlsxFiles, notFoundDirs, nil

}

// fetchProjectBatchDirs 获取指定目录下的所有项目批次目录
func fetchProjectBatchDirs(root string) ([]string, error) {
	dirEntry, err := os.ReadDir(root)
	if err != nil {
		return nil, fmt.Errorf("fetchProjectBatchDirs error: %w", err)
	}

	var projectBatchDirs []string
	var skippedDirs []string

	for _, entry := range dirEntry {
		if entry.IsDir() && isProjectBatchDirLoose(entry.Name()) {
			projectBatchDirs = append(projectBatchDirs, filepath.Join(root, entry.Name()))
			logrus.Println("Found project batch directory:", entry.Name())
		} else {
			skippedDirs = append(skippedDirs, entry.Name())
		}
	}

	if len(skippedDirs) != 0 {
		logrus.Println("Skipping non-project batch directory:", root, len(skippedDirs), skippedDirs)
	}

	if len(projectBatchDirs) != 0 {
		logrus.Println("Total project batch directories found:", len(projectBatchDirs))
	}

	return projectBatchDirs, nil
}

// loopFindProjectBatchXlsxFiles 循环查找多个目录下的xlsx文件
func loopFindProjectBatchXlsxFiles(dirPaths []string) ([]string, []string, error) {
	var xlsxFiles []string
	var notFoundDirs []string

	//遍历每个目录路径，查找xlsx文件
	for _, dirPath := range dirPaths {
		xlsxFile, err := findProjectBatchXlsxFile(dirPath)
		if err != nil && os.IsNotExist(err) {
			return nil, nil, fmt.Errorf("loopFindProjectBatchXlsxFiles error: %w", err)

		} else if err != nil { //未找到xlsx文件的目录
			notFoundDirs = append(notFoundDirs, dirPath)
			continue
		}

		xlsxFiles = append(xlsxFiles, xlsxFile)
		logrus.Println("Found xlsx file:", xlsxFile)
	}

	if len(notFoundDirs) != 0 {
		logrus.Warnf("Directories without xlsx files:\n"+
			"total:%v\n"+
			"%v",
			len(notFoundDirs), strings.Join(notFoundDirs, "\n"))
	}

	return xlsxFiles, notFoundDirs, nil
}

// findProjectBatchXlsxFiles 查找指定目录下的xlsx文件
func findProjectBatchXlsxFile(dirPath string) (string, error) {
	dirEntry, err := os.ReadDir(dirPath)
	if err != nil {
		return "", fmt.Errorf("findProjectBatchXlsxFiles error: %w", err)
	}
	//遍历目录下的所有文件

	for _, entry := range dirEntry {
		if entry.IsDir() {
			continue
		}
		if filepath.Ext(entry.Name()) == ".xlsx" && isProjectBatchDir(entry.Name()) {
			xlsxFile := filepath.Join(dirPath, entry.Name())
			return xlsxFile, nil
		}
	}

	return "", fmt.Errorf("no xlsx file found in directory: %s", dirPath)
}

func isProjectBatchDir(dirName string) bool {
	regex := regexp.MustCompile(internal.RegexProjectBatch)
	if regex.FindString(dirName) != "" {
		return true
	}

	return false
}

func isProjectBatchDirLoose(dirName string) bool {
	regex := regexp.MustCompile(internal.RegexProjectBatchLoose)
	if regex.FindString(dirName) != "" {
		return true
	}

	return false
}
