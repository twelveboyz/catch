package execute

import (
	computeComparison "Catch/bump/comparison/compute_comparison"
	"Catch/bump/comparison/internal/utils"
	"Catch/bump/excel/cmdb/get"
	"Catch/bump/pull/compute/ecs"
	"Catch/bump/pull/compute/ironic"
	"errors"
	"github.com/sirupsen/logrus"
	"strings"
	"sync"
)

func ComputeRun(root string) {
	//获取文件名
	comPuteExcel, err := InitComputeExcel(root)
	if err != nil {
		logrus.Infoln(err.Error())
		return
	}
	/*------------------------------循环计算表格的sheet---------------------------------------*/
	sheets := comPuteExcel.ForGetSheet()
	for _, sheet := range sheets {
		comPuteExcel.Sheet = sheet

		//获取excel表格中需要对比的数据
		excelResources := comPuteExcel.ParseComputeResourceToStruct()

		//获取excel表格中的uuid
		uuids := get.ComputeGetUUIDs(excelResources)

		//获取excel表格中pools
		pools := get.ComputeGetPool(excelResources)

		//并发拉取控制台资源信息，并且仅获取uuid匹配的控制台资源信息
		ecsConsoles, ironicConsoles := GoroutineGetConsole(uuids, pools)

		//匹配excel表格中的资源小类，和控制台中的资源进行对比
		for _, er := range excelResources {
			//匹配资源小类为云主机则对比控制台云主机资源
			if strings.Contains(er.ResourceSubCategory, "云主机") {
				computeComparison.EcsComparison(er, ecsConsoles)
			}

			//匹配资源小类为裸金属则对比控制台裸金属资源
			if strings.Contains(er.ResourceSubCategory, "裸金属") {
				computeComparison.IronicComparison(er, ironicConsoles)
			}
		}

		if len(excelResources) == 0 {
			return
		}
		utils.ComputeInfoSummaryPrint(sheet)
	}

}

func InitComputeExcel(root string) (*get.Excel, error) {
	computeFileName := get.GetFileName(root, ".xls", "计算资源")
	if computeFileName == "" {
		return nil, errors.New("未找到计算资源，已跳过···")
	}

	comPuteExcel := get.NewExcel(root, computeFileName, 2, 3)

	comPuteExcel.File = comPuteExcel.OpenFile(comPuteExcel.FileName)

	return comPuteExcel, nil
}

func GoroutineGetConsole(uuids []string, pools []string) (ecsConsoles []*ecs.ECloudServer, ironicConsoles []*ironic.Ironic) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		ecsConsoles = ecs.ExecuteGetEcsInfo(uuids, pools)
		wg.Done()
	}()

	go func() {
		ironicConsoles = ironic.ExecuteGetIronicInfo(uuids, pools)
		wg.Done()
	}()

	wg.Wait()
	return
}
