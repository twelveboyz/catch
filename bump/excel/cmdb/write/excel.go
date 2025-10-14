package write

import (
	"Catch/bump/pull/compute/ecs"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

type Excel struct {
	Filename  string
	file      *excelize.File
	sheet     string
	TitleLine int
}

// NewExcel creates a new Excel instance by opening the specified file，ops can set sheet.
func NewExcel(filename string, ops ...string) (*Excel, error) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, fmt.Errorf("open file %s error: %w", filename, err)
	}

	var defaultSheet string
	if len(ops) > 0 {
		defaultSheet = ops[0]
	} else {
		defaultSheet = f.GetSheetName(0)
	}

	return &Excel{
		Filename: filename,
		file:     f,
		sheet:    defaultSheet,
	}, nil
}

func (e *Excel) Write(excelEcs []*ecs.ExcelECS) {
	for i, ecs := range excelEcs {
		line := fmt.Sprintf("J%d", i+e.TitleLine+1)

		if err := e.file.SetSheetRow(e.sheet, line, &[]interface{}{ecs.Id, ecs.Name, "弹性计算", "云主机", ecs.SpecsName, ecs.VCpu, ecs.VMemory, "", ecs.ImageName, "运行中", "未交维"}); err != nil {
			logrus.Warnf("Error setting row %s: %v\n", line, err)
		}
	}
	if err := e.file.SaveAs(e.Filename); err != nil {
		fmt.Printf("Error saving file: %v\n", err)
	}

}
