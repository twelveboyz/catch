package get

import (
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func GetFileName(root string, key1 string, key2 string) string {

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
		re := regexp.MustCompile(`CMDB录入表.*`)
		reFile := re.FindString(file)

		if key2 == "PAAS" {
			upperFile := strings.ToUpper(reFile)
			//判断条件如果key1 key2匹配上则返回改文件绝对路径
			fileBool = strings.Contains(reFile, key1) && strings.Contains(upperFile, key2) && !strings.Contains(reFile, "~$")
		} else {
			fileBool = strings.Contains(reFile, key1) && strings.Contains(reFile, key2) && !strings.Contains(reFile, "~$")
		}

		if fileBool {
			logrus.Println("捕抓文件名称路径是：", file)
			return file
		}
	}
	return ""
}
