package daily

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Set struct {
	ResourceCategory        []string
	ResourceSubcategory     []string
	Capacity                []string
	SystemDiskSpecification []string
	SystemDiskCapacity      []string
	Bandwidth               []string
	Number                  []string
	Comments                []string
}

// AggregatedData 用于存储聚合后的结果和规格统计
type AggregatedData struct {
	TotalCounts    int
	CapacityCounts int
	SingleCapacity string
	Comments       string
}

func (e *Excel) ComPutDailyNewspaper(set Set) map[string]*AggregatedData {
	var m = make(map[string]*AggregatedData)

	for i := 0; i < len(set.ResourceSubcategory); i++ {
		subcategory := set.ResourceSubcategory[i]

		//跳过资源小类等于NF or NFI
		if subcategory == "NF" || subcategory == "NFI" {
			continue
		}

		// 如果子类别不在映射中，添加一个新的计数器
		if _, exists := m[subcategory]; !exists {
			m[subcategory] = &AggregatedData{
				TotalCounts: 0,
				Comments:    "",
			}
		}

		// 更新子类别的总计数
		m[subcategory].TotalCounts++
		if strings.Contains(set.Comments[i], "扩容") {
			m[subcategory].Comments = "扩容"
		}
	}

	// 打印测试信息
	for k, v := range m {
		log.Println("ComputeDaily=", k, v.TotalCounts, v.Comments)
	}

	return m
}

func (e *Excel) SystemDiskDailyNewspaper(set Set) map[string]*AggregatedData {
	var m = make(map[string]*AggregatedData)

	for i := 0; i < len(set.ResourceSubcategory); i++ {
		//subResource := set.ResourceSubcategory[i]
		systemDiskSpecification := set.SystemDiskSpecification[i]

		//跳过资源小类等于NF or NFI
		if systemDiskSpecification == "NF" || systemDiskSpecification == "NFI" {
			continue
		}

		// 如果子类别不在映射中，添加一个新的计数器
		if systemDiskSpecification != "SSD" {
			if _, exists := m[systemDiskSpecification]; !exists {
				m[systemDiskSpecification] = &AggregatedData{
					TotalCounts:    0,
					CapacityCounts: 0,
					Comments:       "",
				}
			}
		}

		if systemDiskSpecification != "SSD" {
			// 更新子类别的总计数
			m[systemDiskSpecification].TotalCounts++

			//统计系统盘容量
			capaInt, err := strconv.Atoi(set.SystemDiskCapacity[i])
			if err != nil {
				log.Println("AtoI error:", err)
			}
			m[systemDiskSpecification].CapacityCounts += capaInt
		}

		//fmt.Println("systemDiskSpecification:", m[systemDiskSpecification].TotalCounts, m[systemDiskSpecification].CapacityCounts)

		if systemDiskSpecification != "SSD" {
			if strings.Contains(set.Comments[i], "扩容") {
				m[systemDiskSpecification].Comments = "扩容"
			}
		}
	}

	// 打印测试信息
	for k, v := range m {
		log.Println("SystemDiskDaily=", k, v.TotalCounts, v.CapacityCounts)
	}

	return m
}

func (e *Excel) StorageDailyNewspaper(set Set) map[string]*AggregatedData {
	var m = make(map[string]*AggregatedData)

	for i := 0; i < len(set.ResourceSubcategory); i++ {
		var subcategory string
		if set.ResourceSubcategory[i] == "对象存储" {
			subcategory = fmt.Sprintf("%v#&%v", set.ResourceSubcategory[i], set.Comments[i])
		} else {
			subcategory = set.ResourceSubcategory[i]
		}

		if subcategory == "NF" || subcategory == "NFI" {
			continue
		}

		// 如果子类别不在映射中，添加一个新的计数器
		if _, exists := m[subcategory]; !exists {
			m[subcategory] = &AggregatedData{
				TotalCounts:    0,
				CapacityCounts: 0,
				Comments:       "",
			}
		}

		// 更新子类别的总计数
		m[subcategory].TotalCounts++

		if strings.Contains(subcategory, "对象存储") {
			capaInt, err := strconv.Atoi(set.Capacity[i])
			if err != nil {
				log.Println("DailyNewspaper函数Atoi转换失败:", err)
			}
			m[subcategory].CapacityCounts = capaInt

		} else {
			capaInt, err := strconv.Atoi(set.Capacity[i])
			if err != nil {
				log.Println("DailyNewspaper函数Atoi转换失败:", err)
			}
			m[subcategory].CapacityCounts += capaInt
		}

		if strings.Contains(set.Comments[i], "扩容") {
			m[subcategory].Comments = "扩容"
		}
	}

	//将多个对象存储的数据合并成一个k,v
	eos := "对象存储"
	for k, v := range m {
		if strings.Contains(k, "对象存储#&") {
			if _, exists := m[eos]; !exists {
				m[eos] = &AggregatedData{
					TotalCounts:    0,
					CapacityCounts: 0,
				}
			}
			m[eos].TotalCounts += v.TotalCounts
			m[eos].CapacityCounts += v.CapacityCounts
			delete(m, k)
		}
	}

	// 打印测试信息
	for k, v := range m {
		log.Println("StorageDaily=", k, v.TotalCounts, v.CapacityCounts)
	}

	return m
}

func (e *Excel) NetworkDailyNewspaper(set Set) map[string]*AggregatedData {
	var m = make(map[string]*AggregatedData)

	for i := 0; i < len(set.ResourceSubcategory); i++ {
		subcategory := set.ResourceSubcategory[i]

		//跳过资源小类等于NF or NFI
		if subcategory == "NF" || subcategory == "NFI" {
			continue
		}

		// 如果子类别不在映射中，添加一个新的计数器
		if _, exists := m[subcategory]; !exists {
			m[subcategory] = &AggregatedData{
				TotalCounts: 0,
				Comments:    "",
			}
		}

		// 更新子类别的总计数
		m[subcategory].TotalCounts++

		if strings.Contains(set.Comments[i], "扩容") {
			m[subcategory].Comments = "扩容"
		}
	}

	// 打印测试信息
	for k, v := range m {
		log.Println("NetworkDaily=", k, v.TotalCounts)
	}
	return m
}

func (e *Excel) SecurityDailyNewspaper(set Set) map[string]*AggregatedData {
	var m = make(map[string]*AggregatedData)

	for i := 0; i < len(set.ResourceSubcategory); i++ {
		subcategory := set.ResourceSubcategory[i]

		//跳过资源小类等于NF or NFI
		if subcategory == "NF" || subcategory == "NFI" {
			continue
		}

		// 如果子类别不在映射中，添加一个新的计数器
		if _, exists := m[subcategory]; !exists {
			m[subcategory] = &AggregatedData{
				TotalCounts: 0,
				Comments:    "",
			}
		}

		// 更新子类别的总计数
		number, err := strconv.Atoi(set.Number[i])
		if err != nil {
			log.Println("AtoI error:", err)
		}
		m[subcategory].TotalCounts += number

		if strings.Contains(set.Comments[i], "扩容") {
			m[subcategory].Comments = "扩容"
		}
	}

	// 打印测试信息
	for k, v := range m {
		log.Println("SecurityDaily=", k, v.TotalCounts)
	}
	return m
}

func (e *Excel) PAASDailyNewspaper(set Set) map[string]*AggregatedData {
	var m = make(map[string]*AggregatedData)

	for i := 0; i < len(set.ResourceSubcategory); i++ {
		subcategory := set.ResourceSubcategory[i]

		//跳过资源小类等于NF or NFI
		if subcategory == "NF" || subcategory == "NFI" {
			continue
		}

		// 如果子类别不在映射中，添加一个新的计数器
		if _, exists := m[subcategory]; !exists {
			m[subcategory] = &AggregatedData{
				TotalCounts: 0,
				Comments:    "",
			}
		}

		// 更新子类别的总计数
		m[subcategory].TotalCounts++

		if strings.Contains(set.Comments[i], "扩容") {
			m[subcategory].Comments = "扩容"
		}
	}

	// 打印测试信息
	for k, v := range m {
		log.Println("PAASDaily=", k, v.TotalCounts)
	}

	return m

}
