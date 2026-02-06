package internal

import "fmt"

//var RegexStr = `[\w][\d]+_((?:([\w]*[\d]*)*)?[\p{Han}]+(?:([\w]*[\d]*)*)?)+`
//var RegexHorizontalBar = `[\w][\d]+-((?:([\w]*[\d]*)*)?[\p{Han}]+(?:([\w]*[\d]*)*)?)+`

var RegexProjectBatch = `\w\d{2,3}[-_](\d*\w*\p{Han}+\d*\w*)+`
var RegexProjectInBatches = `\w\d{2,3}[-_](\d*\w*\p{Han}+\d*\w*)+[-]?第.批.*`
var RegexProjectBatchLoose = `\w\d{2,3}[-_]?.*`

func RegexMatchFileName(dir string) string {
	return fmt.Sprintf("%s[/\\\\](:?.*\\.xlsx)", dir)
}
