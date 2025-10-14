package execute

import (
	"Catch/bump/comparison/internal/utils"
	"Catch/bump/comparison/storage_comparison"
	"Catch/bump/excel/cmdb/get"
	"Catch/bump/pull/storage/ebs"
	"Catch/bump/pull/storage/efs"
	"Catch/bump/pull/storage/eos"
	"Catch/bump/pull/storage/pfs"
	"errors"
	"github.com/sirupsen/logrus"
	"strings"
)

func StorageRun(root string) {
	storageExcel, err := initStorageExcel(root)
	if err != nil {
		logrus.Infoln(err.Error())
		return
	}

	sheets := storageExcel.ForGetSheet()

	for _, sheet := range sheets {
		storageExcel.Sheet = sheet

		if err = storageExcel.EffectiveSheet(); err != nil {
			continue
		}

		excelResources := storageExcel.ParseStorageResourceToStruct()

		uuids := get.StorageGetUUIDs(excelResources)

		pools := get.StorageGetPool(excelResources)

		switch sheet {
		case "云硬盘":
			infoSet := ebs.ExecuteGetEBSInfo(uuids, pools)

			for _, er := range excelResources {
				storage_comparison.EbsComparison(er, infoSet)
			}
		case "对象存储":
			infoSet := eos.Execute()

			for _, er := range excelResources {
				storage_comparison.EOSComparison(er, infoSet)
			}

		case "文件存储":
			efsSet := efs.ExecuteGetEFSInfo(uuids, pools)
			pfsSet := pfs.ExecuteGetPFSInfo(uuids, pools)

			for _, er := range excelResources {
				if strings.Contains(er.ResourceSubCategory, "并行文件存储") {
					storage_comparison.PFSComparison(er, pfsSet)
				} else if er.ResourceSubCategory == "文件存储" {
					storage_comparison.EFSComparison(er, efsSet)
				} else {
					logrus.Warnln("未匹配到文件存储中的资源小类：", er.ResourceSubCategory)
				}
			}

		default:
			logrus.Warnln("存储资源未匹配到Sheet:", sheet)
		}

		if len(excelResources) == 0 {
			return
		}

		utils.StorageInfoSummaryPrint(sheet)
	}
}

func initStorageExcel(root string) (*get.Excel, error) {
	storageFileName := get.GetFileName(root, ".xls", "存储资源")
	if storageFileName == "" {
		return nil, errors.New("未找到存储资源，已跳过···")
	}

	storageExcel := get.NewExcel(root, storageFileName, 2, 3)

	storageExcel.File = storageExcel.OpenFile(storageExcel.FileName)
	return storageExcel, nil
}
