package utils

import (
	"github.com/sirupsen/logrus"
)

func PrintEnd(i int) {
	if i == 0 {
		return
	}
	if logrus.GetLevel().String() == "info" {
		logrus.Infof("-------------------------------------------------------------------\n")
	} else {
		logrus.Debugf("-------------------------------------------------------------------\n")
	}
}

func ComputeInfoSummaryPrint(sheet string) {
	PrintEnd(len(MisMatchResourceCount))
	SortMapKeyAndPrint(MisMatchResourceCount, "计算资源-"+sheet)
	if len(MisMatchResourceCount) == 0 {
		logrus.Println("|-------------------------------|")
		logrus.Printf("|计算资源-%s已通过验证！\n", sheet)
		logrus.Printf("|-------------------------------|\n\n")
	} else {
		ClearMapData(MisMatchResourceCount)
	}
}

func StorageInfoSummaryPrint(sheet string) {
	PrintEnd(len(MisMatchResourceCount))
	SortMapKeyAndPrint(MisMatchResourceCount, "存储资源-"+sheet)
	if len(MisMatchResourceCount) == 0 {
		logrus.Println("|-------------------------------|")
		logrus.Printf("|存储资源-%s已通过验证！\n", sheet)
		logrus.Printf("|-------------------------------|\n\n")
	} else {
		ClearMapData(MisMatchResourceCount)
	}

}

func NetworkInfoSummaryPrint(sheet string) {
	PrintEnd(len(MisMatchResourceCount))
	SortMapKeyAndPrint(MisMatchResourceCount, "网络资源-"+sheet)

	if len(MisMatchResourceCount) == 0 {
		logrus.Println("|-------------------------------|")
		logrus.Printf("|网络资源-%s已通过验证！\n", sheet)
		logrus.Printf("|-------------------------------|\n\n")
	} else {
		ClearMapData(MisMatchResourceCount)
	}
}
