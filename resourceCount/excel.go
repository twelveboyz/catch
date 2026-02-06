package resourceCount

import (
	"Catch/internal"
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type resource struct {
	CloudPool           string
	NodeName            string
	ResourceCategory    string
	ResourceSubcategory string
	Specification       string
	Unit                string
	TotalCounts         string
	CapacityCounts      string
	Comment             string
	ProjectID           string //项目ID
	Batch               string //批次
	ResourceProperties  string //资源属性
}

func Run() {
	projectBatchFiles, err := loopFindProjectBatchResources()
	if err != nil {
		logrus.Fatalf("LoopFindProjectBatchResources:%v", err)
	}

	resourceCount, err := ReadExcelFile(projectBatchFiles)
	if err != nil {
		logrus.Fatalf("ReadExcelFile:%v", err)
	}

	err = WriteExcelFile(resourceCount)
	if err != nil {
		logrus.Fatalf("WriteExcelFile:%v", err)
	}

}

// WriteExcelFile 将resource结构体切片写入Excel文件
func WriteExcelFile(resources []resource) error {
	excel := internal.CreateExcel()
	defer func() {
		if err := excel.Close(); err != nil {
			logrus.Fatalf("关闭文件失败:%v", err)
		}
	}()

	sheet := excel.GetSheetName(0)

	for i, res := range resources {
		nodeName := strings.ReplaceAll(res.NodeName, "_", "-")
		var resourceCategory string
		if strings.Contains(strings.ToUpper(res.ResourceCategory), "PAAS") {
			resourceCategory = "云PaaS-SaaS"
		} else {
			resourceCategory = res.ResourceCategory
		}

		cell := fmt.Sprintf("A%d", i+1)

		row := &[]string{
			res.CloudPool,
			nodeName,
			resourceCategory,
			res.ResourceSubcategory,
			res.Specification,
			res.Unit,
			res.TotalCounts,
			res.CapacityCounts,
			res.Comment,
			res.ProjectID,
			res.Batch,
			res.ResourceProperties,
		}

		if err := excel.SetSheetRow(sheet, cell, row); err != nil {
			return fmt.Errorf("SetSheetRow:%w", err)
		}

	}

	fileName := fmt.Sprintf("资源统计结果-%v.xlsx", time.Now().Format("20060102"))
	if err := excel.SaveAs(fileName); err != nil {
		return fmt.Errorf("SaveAs:%w", err)
	}

	logrus.Infoln("已生成文件:", fileName)

	return nil
}

// ReadExcelFile 读取多个Excel文件并转换为resource结构体切片
func ReadExcelFile(files []string) ([]resource, error) {

	var allResources []resource
	for _, file := range files {
		resources, err := readExcelRowConvertResource(file)
		if err != nil {
			return nil, fmt.Errorf("readExcelRow:%w", err)
		}
		allResources = append(allResources, resources...)
	}

	return allResources, nil
}

// readExcelRowConvertResource 读取Excel文件的行内容并转换为resource结构体
func readExcelRowConvertResource(file string) ([]resource, error) {
	excel, err := internal.NewExcel(file, "Sheet1", 0, 0)
	if err != nil {

		return nil, fmt.Errorf("打开文件失败:%w", err)
	}

	rows, err := excel.File.GetRows(excel.Sheet)
	if err != nil {
		return nil, fmt.Errorf("获取行内容失败:%w", err)
	}

	titleLine, err := getTitleLine(rows)
	if err != nil {
		return nil, fmt.Errorf("获取标题行失败:%w", err)
	}

	excel.SetTitleLine(titleLine)
	excel.SetCopyLine(titleLine + 1)

	titleRow := rows[titleLine]
	logrus.Debugln("titleRow:", titleRow)

	rows = rows[excel.CopyLine:]
	logrus.Debugln("rows before fill:", len(rows), rows)
	for _, row := range rows {
		logrus.Debugln("len row:", len(row), row)
	}

	rows = fillRows(rows)
	logrus.Debugln("rows after fill:", len(rows), rows)

	for _, row := range rows {
		logrus.Debugln("len row:", len(row), row)
	}

	cpIndex := excel.CaptureColNumber(titleRow, internal.CloudPool, false)
	nodeIndex := excel.CaptureColNumber(titleRow, internal.Node, true)
	rcIndex := excel.CaptureColNumber(titleRow, internal.ResourceCategory, true)
	rsIndex := excel.CaptureColNumber(titleRow, internal.ResourceSubcategory, true)
	specIndex := excel.CaptureColNumber(titleRow, internal.Specification, true)
	unitIndex := excel.CaptureColNumber(titleRow, internal.Unit, true)
	resourceCountIndex := excel.CaptureColNumber(titleRow, internal.ResourceCount, true)
	vbIndex := excel.CaptureColNumber(titleRow, internal.VolumeOrBucketNumber, true)
	commentIndex := excel.CaptureColNumber(titleRow, internal.Comment, true)

	var resources []resource

	for _, row := range rows {
		var resource resource
		resource.CloudPool = row[cpIndex]
		resource.NodeName = row[nodeIndex]
		resource.ResourceCategory = row[rcIndex]
		resource.ResourceSubcategory = row[rsIndex]
		resource.Specification = row[specIndex]
		resource.Unit = row[unitIndex]
		resource.TotalCounts = row[resourceCountIndex]
		resource.CapacityCounts = row[vbIndex]
		resource.Comment = row[commentIndex]

		resource.ProjectID = getProjectID(file)
		resource.Batch = getBatch(file)
		resource.ResourceProperties = getResourceProperties(resource.Comment)
		resources = append(resources, resource)
	}

	return resources, nil
}

func fillRows(rows [][]string) [][]string {
	newRows := make([][]string, 0, len(rows))
	for _, row := range rows {
		if len(row) < 9 {
			row = append(row, "", "")
			newRows = append(newRows, row)
		} else if len(row) < 10 {
			row = append(row, "")
			newRows = append(newRows, row)
		} else {
			newRows = append(newRows, row)
		}
	}

	return newRows
}

// getProjectID 从文件路径中提取项目ID
func getProjectID(filePath string) string {
	dir := filepath.Dir(filePath)
	file := filepath.Base(filePath)
	re := regexp.MustCompile(`[A-Z]\d{2,3}`)

	if s := re.FindString(dir); s != "" {
		return s
	}

	if s := re.FindString(file); s != "" {
		return s
	}

	return ""
}

// getBatch 从文件路径中提取批次
func getBatch(filePath string) string {
	dir := filepath.Dir(filePath)
	re := regexp.MustCompile(`[-_](第.{1,2}批)`)

	subStr := re.FindStringSubmatch(dir)
	if len(subStr) == 2 {
		return subStr[1]
	}

	return "不分批"
}

// getResourceProperties 从注释中提取资源属性
func getResourceProperties(comment string) string {
	if strings.Contains(comment, "扩容") {
		return "扩容"
	} else {
		return "新订购"
	}
}

func getTitleLine(rows [][]string) (int, error) {
	for i, row := range rows {
		for _, col := range row {
			if col == internal.CloudPool {
				return i, nil
			}
		}
	}
	return 0, fmt.Errorf("未找到标题行")
}
