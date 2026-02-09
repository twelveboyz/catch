package cmdb

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// PdfInfo 存储从PDF中提取的信息
type PdfInfo struct {
	Applicant        string // 申请人
	DemandSide       string // 需求方
	OAPublishTime    string // OA发文时间
	PrimaryAccount   string // 主账号
}

// FindPdfInReviewDir 在 "1、订购评审纪要" 目录下查找PDF文件
func FindPdfInReviewDir(root string) (string, error) {
	reviewDir := filepath.Join(root, "1、订购评审纪要")

	// 检查目录是否存在
	dirInfo, err := os.Stat(reviewDir)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("review directory not found: %s", reviewDir)
		}
		return "", fmt.Errorf("failed to check review directory: %w", err)
	}

	if !dirInfo.IsDir() {
		return "", fmt.Errorf("review path is not a directory: %s", reviewDir)
	}

	// 读取目录内容
	dirEntry, err := os.ReadDir(reviewDir)
	if err != nil {
		return "", fmt.Errorf("failed to read review directory: %w", err)
	}

	// 查找PDF文件
	fileName, err := mostSimilarFileName(dirEntry)
	if err != nil {
		return "", fmt.Errorf("failed to find PDF in review directory: %w", err)
	}

	return filepath.Join(reviewDir, fileName), nil
}

// mostSimilarFileName 从目录条目中查找最相似的PDF文件名
func mostSimilarFileName(dirEntry []os.DirEntry) (string, error) {
	num := 0
	fileName := ""
	// 如果只有一个PDF文件，直接返回
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
		return "", fmt.Errorf("no PDF file found")
	}

	// 如果有多个PDF，按优先级匹配
	r := regexp.MustCompile(`[A-Z]\d{3}`)

	for _, entry := range dirEntry {
		if !entry.IsDir() {
			if strings.HasSuffix(entry.Name(), ".pdf") {
				// 优先选择包含"集中流程管理系统"的文件
				if strings.Contains(entry.Name(), "集中流程管理系统") {
					return entry.Name(), nil
				}
				// 其次选择匹配项目ID模式的文件（如 E020.pdf）
				if r.FindString(entry.Name()) != "" {
					return entry.Name(), nil
				}
			}
		}
	}

	return fileName, nil
}

// convertPdfToText 使用 tools/pdftotext.exe 将PDF转换为文本
func convertPdfToText(pdfPath string) (string, error) {
	// 生成临时文本文件路径
	txtPath := strings.ReplaceAll(pdfPath, ".pdf", ".txt")

	// 执行 pdftotext 命令
	cmd := exec.Command("tools/pdftotext.exe", "-nopgbrk", "-raw", pdfPath, txtPath)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("pdftotext failed: %w", err)
	}

	// 读取转换后的文本内容
	content, err := os.ReadFile(txtPath)
	if err != nil {
		return "", fmt.Errorf("failed to read text file: %w", err)
	}

	// 删除临时文本文件
	_ = os.Remove(txtPath)

	return string(content), nil
}

// ExtractApplicant 从PDF内容中提取申请人
// 匹配格式： "申请人： 黄昌焕"
func ExtractApplicant(content string) string {
	r := regexp.MustCompile(`申请人：\s*(\p{Han}+)`)
	matches := r.FindStringSubmatch(content)
	if len(matches) >= 2 {
		return matches[1]
	}
	return ""
}

// ExtractDemandSide 从PDF内容中提取需求方
// 匹配格式： "所属独立业务X： X_研发管理平台" -> "研发管理平台"
// 或 "所属独立业务X： 集中运维平台" -> "集中运维平台"
func ExtractDemandSide(content string) string {
	// 首先尝试匹配 X_ 前缀的格式（注意：可能是数字或字母X）
	// 使用非贪婪匹配，只捕获业务名称（到空格或换行为止）
	r := regexp.MustCompile(`所属独立业务[\dX]+：\s*X_(\S+)`)
	matches := r.FindStringSubmatch(content)
	if len(matches) >= 2 {
		return strings.TrimSpace(matches[1])
	}

	// 如果没有 X_ 前缀，直接提取业务名称
	r = regexp.MustCompile(`所属独立业务[\dX]+：\s*(\S+)`)
	matches = r.FindStringSubmatch(content)
	if len(matches) >= 2 {
		return strings.TrimSpace(matches[1])
	}

	return ""
}

// ExtractOAPublishTime 从PDF内容中提取OA发文时间
// 匹配格式： "公有云资源订购部署信息接收 2023-10-15"
// 或审批记录中的日期
func ExtractOAPublishTime(content string) string {
	// 首先尝试匹配 "公有云资源订购部署信息接收" 后的日期
	r := regexp.MustCompile(`公有云资源订购部署信息接收\s*(\d{4}-\d{2}-\d{2})`)
	matches := r.FindStringSubmatch(content)
	if len(matches) >= 2 {
		return matches[1]
	}

	// 如果没有找到，尝试在审批记录中查找日期
	// 查找 "公有云资源订购部署信息接收" 相关段落中的日期
	r = regexp.MustCompile(`公有云资源订购部署.*?(\d{4}-\d{2}-\d{2})`)
	matches = r.FindStringSubmatch(content)
	if len(matches) >= 2 {
		return matches[1]
	}

	// 尝试查找 "发文时间" 或 "OA发文" 相关的日期
	r = regexp.MustCompile(`(?:发文时间|OA发文|发文).*?(\d{4}-\d{2}-\d{2})`)
	matches = r.FindStringSubmatch(content)
	if len(matches) >= 2 {
		return matches[1]
	}

	// 最后尝试：查找文档中最早的日期（通常是发文时间）
	r = regexp.MustCompile(`(\d{4}-\d{2}-\d{2})`)
	allMatches := r.FindAllStringSubmatch(content, -1)
	if len(allMatches) > 0 {
		// 返回找到的第一个日期
		return allMatches[0][1]
	}

	return ""
}

// ExtractPrimaryAccount 从PDF内容中提取主账号
// 匹配格式： "订购帐号： cmic_jsglpt" 或 "订购账号： test_account"
func ExtractPrimaryAccount(content string) string {
	// 尝试匹配 "订购帐号：" 或 "订购账号："
	r1 := regexp.MustCompile(`订购帐号[：:]\s*(\w+)`)
	matches := r1.FindStringSubmatch(content)
	if len(matches) >= 2 {
		return matches[1]
	}

	r2 := regexp.MustCompile(`订购账号[：:]\s*(\w+)`)
	matches = r2.FindStringSubmatch(content)
	if len(matches) >= 2 {
		return matches[1]
	}

	return ""
}

// ExtractFromPDF 从指定目录的PDF文件中提取信息
func ExtractFromPDF(root string) (*PdfInfo, error) {
	// 1. 查找PDF文件
	pdfPath, err := FindPdfInReviewDir(root)
	if err != nil {
		return nil, fmt.Errorf("failed to find PDF: %w", err)
	}

	// 2. 转换PDF为文本
	content, err := convertPdfToText(pdfPath)
	if err != nil {
		return nil, fmt.Errorf("failed to convert PDF: %w", err)
	}

	// 3. 提取各个字段
	info := &PdfInfo{
		Applicant:      ExtractApplicant(content),
		DemandSide:     ExtractDemandSide(content),
		OAPublishTime:  ExtractOAPublishTime(content),
		PrimaryAccount: ExtractPrimaryAccount(content),
	}

	return info, nil
}
