package internal

import (
	"fmt"
	"regexp"
)

//var RegexStr = `[\w][\d]+_((?:([\w]*[\d]*)*)?[\p{Han}]+(?:([\w]*[\d]*)*)?)+`
//var RegexHorizontalBar = `[\w][\d]+-((?:([\w]*[\d]*)*)?[\p{Han}]+(?:([\w]*[\d]*)*)?)+`

var RegexProjectBatch = `\w\d{2,3}[-_](\d*\w*\p{Han}+\d*\w*)+`
var RegexProjectInBatches = `\w\d{2,3}[-_](\d*\w*\p{Han}+\d*\w*)+[-]?第.批.*`
var RegexProjectBatchLoose = `\w\d{2,3}[-_]?.*`

// RegexInBatches 从目录路径中提取"第X批"，要求批次前带 - 或 _ 分隔符
var RegexInBatches = `[-_](第.{1,2}批)`

// DetectBatch 从目录路径中提取批次，未匹配时返回空串
func DetectBatch(root string) string {
	m := regexp.MustCompile(RegexInBatches).FindStringSubmatch(root)
	if len(m) == 2 {
		return m[1]
	}

	return ""
}

func RegexMatchFileName(dir string) string {
	return fmt.Sprintf("%s[/\\\\](:?.*\\.xlsx)", dir)
}
