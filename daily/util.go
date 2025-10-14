package daily

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func PdfToText(root string) (string, string, error) {
	pdfName, err := findPdf(root)
	if err != nil {
		return "", "", fmt.Errorf("findPdf failed: %w", err)
	}
	src := filepath.Join(root, pdfName)

	dest := strings.ReplaceAll(pdfName, ".pdf", ".txt")

	err = cmdExec(src, dest)
	if err != nil {
		return "", "", fmt.Errorf("cmdExec failed: %w", err)
	}

	b, err := os.ReadFile(dest)
	if err != nil {
		return "", "", fmt.Errorf("ReadFile failed: %w", err)
	}

	//fmt.Println(content)
	return string(b), dest, nil
}

func findPdf(root string) (string, error) {
	dirEntry, err := os.ReadDir(root)
	if err != nil {
		return "", fmt.Errorf("ReadDir failed: %w", err)
	}

	fileName, err := MostSimilarFileName(dirEntry)
	if err != nil {
		return "", fmt.Errorf("MostSimilarFileName failed: %w", err)
	}

	return fileName, nil
}

func MostSimilarFileName(dirEntry []os.DirEntry) (string, error) {
	num := 0
	fileName := ""
	//如果num 1 直接返回文件
	for _, entry := range dirEntry {
		if !entry.IsDir() {
			if strings.HasSuffix(entry.Name(), ".pdf") {
				num++
				fileName = entry.Name()
			}
		}
	}

	if num == 1 {
		return fileName, nil
	} else if num == 0 {
		return "", fmt.Errorf("not found pdf fileName")
	}

	//如果num非1，则匹配最符合名称的pdf文件
	r := `\w\d{3}`
	re := regexp.MustCompile(r)

	for _, entry := range dirEntry {
		if !entry.IsDir() {
			if strings.HasSuffix(entry.Name(), ".pdf") {
				if strings.Contains(entry.Name(), "集中流程管理系统") {
					return entry.Name(), nil
				}
				if re.FindString(entry.Name()) != "" {
					return entry.Name(), nil
				}
			}
		}
	}

	return fileName, nil
}

func cmdExec(sourcePath, destination string) error {
	cmd := exec.Command("tools/pdftotext.exe", "-nopgbrk", "-raw", sourcePath, destination)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("pdftotext failed: %w", err)
	}
	return nil
}

func projectBatch(content string) string {
	//工单
	projectName := ` 工程名称：.*`
	re := regexp.MustCompile(projectName)
	str := re.FindString(content)
	//fmt.Println(str)
	pb := `[a-zA-Z][\d]+[-_]((\w*\d*)*[\p{Han}]+)+`
	re3 := regexp.MustCompile(pb)
	projectBatchStr := re3.FindString(str)
	//fmt.Println(projectBatchStr)

	return projectBatchStr
}

func publicationTime(content string) string {
	rt := `公有云资源订购部署信\s+息接收\s*(\d{4}-\d{2}-\d{2}){1}`

	re := regexp.MustCompile(rt)
	strSub := re.FindStringSubmatch(content)
	//fmt.Println("sl4", strSub[len(strSub)-1])

	if len(strSub) == 0 {
		return ""
	}

	return strSub[len(strSub)-1]
}

func DeployImplementers(content string) string {
	d := `公有云资源订购部署实\s+施\s*.*(\p{Han}{2,3}){1}`
	re := regexp.MustCompile(d)

	strSub := re.FindStringSubmatch(content)
	if len(strSub) == 0 {
		return ""
	}
	return strSub[len(strSub)-1]
}

func planApprover(content string) string {
	r := `公有云资源与技术方案[\s\S]*\n\s*(\p{Han}{2,3}){1}`
	re := regexp.MustCompile(r)
	strSub := re.FindStringSubmatch(content)
	if len(strSub) == 0 {
		return ""
	}
	return strSub[len(strSub)-1]
}

func oderNumber(content string) string {
	o := `工单号：\s*(\d+)`

	re := regexp.MustCompile(o)
	strSub := re.FindStringSubmatch(content)

	if len(strSub) == 0 {
		return ""
	}

	return strSub[len(strSub)-1]
}
