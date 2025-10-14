package utils

import (
	"Catch/bump/internal/utils"
	"strconv"
	"strings"
)

func IfFieldComparison(comment string, row string, excelStr string, consoleStr string, ops ...map[string]string) {
	NoSpaceExcelStr := strings.ReplaceAll(excelStr, " ", "")
	NoSpaceConsoleStr := strings.ReplaceAll(consoleStr, " ", "")
	rowInt, _ := strconv.Atoi(row)

	if len(ops) != 0 {
		for _, o := range ops {
			for k, v := range o {
				if NoSpaceExcelStr == NoSpaceConsoleStr && k == v {
					utils.PrintDebugFFun(comment, row, excelStr, consoleStr)
				} else if k == v {
					utils.PrintWarnFFun(comment, row, excelStr, consoleStr)

					MisMatchResourceCount[rowInt]++
				}
			}
		}
		return

	} else {
		if NoSpaceExcelStr == NoSpaceConsoleStr {
			utils.PrintDebugFFun(comment, row, excelStr, consoleStr)
		} else {
			utils.PrintWarnFFun(comment, row, excelStr, consoleStr)
			MisMatchResourceCount[rowInt]++
		}
	}

}

/*func IfFieldComparisonBak(row string, excelStr string, consoleStr string, ops ...map[string]string) {
	NoSpaceExcelStr := strings.ReplaceAll(excelStr, " ", "")
	NoSpaceConsoleStr := strings.ReplaceAll(consoleStr, " ", "")
	rowInt, _ := strconv.Atoi(row)

	if len(ops) != 0 {
		for _, o := range ops {
			for k, v := range o {
				if NoSpaceExcelStr == NoSpaceConsoleStr && k == v {
					utils.PrintDebugFFun(row, excelStr, consoleStr)
				} else if k == v {
					utils.PrintWarnFFun(row, excelStr, consoleStr)

					MisMatchResourceCount[rowInt]++
				}
			}
		}
		return

	} else {
		if NoSpaceExcelStr == NoSpaceConsoleStr {
			utils.PrintDebugFFun(row, excelStr, consoleStr)
		} else {
			utils.PrintWarnFFun(row, excelStr, consoleStr)
			MisMatchResourceCount[rowInt]++
		}
	}

}*/

func IfFieldContains(comment string, row string, excelStr string, consoleStr string, ops ...map[string]string) {
	NoSpaceExcelStr := strings.ReplaceAll(excelStr, " ", "")
	NoSpaceConsoleStr := strings.ReplaceAll(consoleStr, " ", "")

	NoSpaceExcelStr = strings.ToLower(NoSpaceExcelStr)
	NoSpaceConsoleStr = strings.ToLower(NoSpaceConsoleStr)
	rowInt, _ := strconv.Atoi(row)

	if len(ops) != 0 {
		for _, o := range ops {
			for k, v := range o {
				if strings.Contains(NoSpaceConsoleStr, NoSpaceExcelStr) && k == v {
					utils.PrintDebugFFun(comment, row, excelStr, consoleStr)
				} else if k == v {
					utils.PrintWarnFFun(comment, row, excelStr, consoleStr)

					MisMatchResourceCount[rowInt]++
				}
			}
		}
		return

	} else {
		if strings.Contains(NoSpaceConsoleStr, NoSpaceExcelStr) {
			utils.PrintDebugFFun(comment, row, excelStr, consoleStr)
		} else {
			utils.PrintWarnFFun(comment, row, excelStr, consoleStr)
			MisMatchResourceCount[rowInt]++
		}
	}

}
