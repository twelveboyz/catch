package utils

import (
	"github.com/sirupsen/logrus"
	"sort"
)

var MisMatchResourceCount = make(map[int]int)

func SortMapKeyAndPrint(m map[int]int, s string) {
	
	var keys []int
	for k, _ := range m {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for i, k := range keys {
		if i+1 == len(keys) {
			logrus.Infof("%s第%d行，共有%d处不匹配\n\n", s, k, m[k])
		} else {
			logrus.Infof("%s第%d行，共有%d处不匹配\n", s, k, m[k])
		}
	}

}

func ClearMapData(m map[int]int) {
	for k, _ := range m {
		delete(m, k)
	}
}
