package cmdb

import (
	"Catch/internal"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Set struct {
	ResourceCategory        []string
	ResourceSubcategory     []string
	Specification           []string
	CloudPool               []string
	Node                    []string
	AvailabilityZone        []string
	Capacity                []string
	SystemDiskSpecification []string
	SystemDiskCapacity      []string
	Bandwidth               []string
	Number                  []string
	Comments                []string
}

// AggregatedData 用于存储聚合后的结果和规格统计
type AggregatedData struct {
	totalCounts         int
	capacityCounts      int
	SingleCapacity      int
	SingleCapacitySlice []int
}

func TotalResourceSliceToMap(countSet [][]string, ResourceType string) []map[string]string {
	var countMap []map[string]string
	//将[][]string,变成[]map[string]string
	switch ResourceType {
	case "存储资源":
		for k, v := range countSet {
			countMap = append(countMap, make(map[string]string))
			countMap[k]["CloudPool"] = v[0]
			countMap[k]["Node"] = v[1]
			countMap[k]["AvailabilityZone"] = v[2]
			countMap[k]["ResourceCategory"] = v[3]
			countMap[k]["ResourceSubcategory"] = v[4]
			countMap[k]["Specification"] = v[5]
			countMap[k]["Comments"] = v[6]
			countMap[k]["Counts"] = v[7]
			countMap[k]["Capacity"] = v[8]
			countMap[k]["SingleCapacity"] = v[9]
		}
	case "计算资源":
		for k, v := range countSet {
			countMap = append(countMap, make(map[string]string))
			countMap[k]["CloudPool"] = v[0]
			countMap[k]["Node"] = v[1]
			countMap[k]["AvailabilityZone"] = v[2]
			countMap[k]["ResourceCategory"] = v[3]
			countMap[k]["ResourceSubcategory"] = v[4]
			countMap[k]["Specification"] = v[5]
			countMap[k]["Counts"] = v[6]
		}
	case "系统盘资源":
		for k, v := range countSet {
			countMap = append(countMap, make(map[string]string))
			countMap[k]["CloudPool"] = v[0]
			countMap[k]["Node"] = v[1]
			countMap[k]["AvailabilityZone"] = v[2]
			countMap[k]["ResourceCategory"] = v[3]
			countMap[k]["ResourceSubcategory"] = v[4]
			countMap[k]["SystemDiskSpecification"] = v[5]
			countMap[k]["SystemDiskCapacity"] = v[6]
			countMap[k]["Counts"] = v[7]
		}
	case "网络资源":
		for k, v := range countSet {
			countMap = append(countMap, make(map[string]string))
			countMap[k]["CloudPool"] = v[0]
			countMap[k]["Node"] = v[1]
			countMap[k]["AvailabilityZone"] = v[2]
			countMap[k]["ResourceCategory"] = v[3]
			countMap[k]["ResourceSubcategory"] = v[4]
			countMap[k]["Specification"] = v[5]
			countMap[k]["Bandwidth"] = v[6]
			countMap[k]["Counts"] = v[7]
		}
	default:
		for k, v := range countSet {
			countMap = append(countMap, make(map[string]string))
			countMap[k]["CloudPool"] = v[0]
			countMap[k]["Node"] = v[1]
			countMap[k]["AvailabilityZone"] = v[2]
			countMap[k]["ResourceCategory"] = v[3]
			countMap[k]["ResourceSubcategory"] = v[4]
			countMap[k]["Specification"] = v[5]
			countMap[k]["Counts"] = v[6]
		}
	}
	return countMap
}

func ComPutTotalResource(set Set) [][]string {
	//第一部将传入的所有数据每一行数据组成key
	// 创建一个 map 来存储聚合后的结果和规格统计
	aggregatedResources := make(map[string]*AggregatedData)

	for i := 0; i < len(set.ResourceCategory); i++ {
		key := fmt.Sprintf("%v#&%v#&%v#&%v#&%v#&%v", set.CloudPool[i], set.Node[i], set.AvailabilityZone[i], set.ResourceCategory[i], set.ResourceSubcategory[i], set.Specification[i])

		// 如果 key 不存在于 map 中，则初始化它，并使用指针 //key是全部参数组成的key，value是int,map[string]int 类型
		if _, exists := aggregatedResources[key]; !exists {
			aggregatedResources[key] = &AggregatedData{
				totalCounts: 0,
			}
		}

		// 简单说就是将Key组合在一起变成一个Set集合，value是统计同类型的个数，和统计规格和每个规格的个数
		// 增加totalCount计数
		aggregatedResources[key].totalCounts++
	}

	//第二步、拆分Key，并且把拆分的Key和Counts、spec都加入到一个[][]string切片中
	var countSlice [][]string
	//k=key ,data=aggregatedData类型
	for k, data := range aggregatedResources {
		// 每次循环开始时重新初始化 Slice
		Slice := strings.Split(k, "#&")

		// 将总数量和规格数量字符串添加到 Slice，Slice是每个key拆分后的数据，totalCounts是统计相同资源的个数，specCounts是将规格和数量组合的数据
		Slice = append(Slice, strconv.Itoa(data.totalCounts))

		//fmt.Println("Slice=", Slice)
		//跳过资源小类等于NF或者NFI的空行
		if Slice[4] == "NF" || Slice[4] == "NFI" {
			continue
		}

		// 将 Slice 添加到 countSlice
		countSlice = append(countSlice, Slice)
	}

	// 打印结果以验证
	for _, slice := range countSlice {
		log.Println("ComPutTotalResource函数打印结果：", slice)
	}
	return countSlice
}

func SystemDiskTotalResource(set Set) [][]string {
	//第一部将传入的所有数据每一行数据组成key
	// 创建一个 map 来存储聚合后的结果和规格统计
	aggregatedResources := make(map[string]*AggregatedData)

	for i := 0; i < len(set.ResourceCategory); i++ {
		key := fmt.Sprintf("%v#&%v#&%v#&%v#&%v#&%v", set.CloudPool[i], set.Node[i], set.AvailabilityZone[i], set.ResourceCategory[i], set.ResourceSubcategory[i], set.SystemDiskSpecification[i])

		// 如果 key 不存在于 map 中，则初始化它，并使用指针 //key是全部参数组成的key，value是int,map[string]int 类型
		if _, exists := aggregatedResources[key]; !exists {
			aggregatedResources[key] = &AggregatedData{
				totalCounts:    0,
				capacityCounts: 0,
			}
		}

		// 简单说就是将Key组合在一起变成一个Set集合，value是统计同类型的个数，和统计规格和每个规格的个数
		// 增加totalCount计数
		aggregatedResources[key].totalCounts++

		//统计相同规格系统盘的总容量
		capInt, err := strconv.Atoi(set.SystemDiskCapacity[i])
		if err != nil {
			log.Println("CapInt获取Atoi转换失败", err)
		}
		//统计相同小类的总容量
		aggregatedResources[key].capacityCounts = aggregatedResources[key].capacityCounts + capInt

	}

	//第二步、拆分Key，并且把拆分的Key和Counts、spec都加入到一个[][]string切片中
	var countSlice [][]string
	//k=key ,data=aggregatedData类型
	for k, data := range aggregatedResources {
		// 每次循环开始时重新初始化 Slice
		Slice := strings.Split(k, "#&")

		// 将总数量和规格数量字符串添加到 Slice，Slice是每个key拆分后的数据，totalCounts是统计相同资源的个数，specCounts是将规格和数量组合的数据
		Slice = append(Slice, strconv.Itoa(data.capacityCounts), strconv.Itoa(data.totalCounts))

		//fmt.Println("Slice=", Slice)
		//跳过资源小类等于NF或者NFI的空行
		if Slice[4] == "NF" || Slice[4] == "NFI" {
			continue
		}

		// 将 Slice 添加到 countSlice; 注意：要过滤掉裸金属的数据，不然生成excel表格会出现空行
		countSlice = append(countSlice, Slice)
	}

	// 打印结果以验证
	for _, slice := range countSlice {
		log.Println("SystemDiskTotalResource函数打印结果：", slice)
	}
	return countSlice
}

func StorageTotalResource(set Set) [][]string {
	//第一部将传入的所有数据每一行数据组成key
	// 创建一个 map 来存储聚合后的结果和规格统计
	aggregatedResources := make(map[string]*AggregatedData)

	for i := 0; i < len(set.ResourceCategory); i++ {
		var key string
		if set.ResourceSubcategory[i] == "对象存储" && set.Comments[i] != "" {
			key = fmt.Sprintf("%v#&%v#&%v#&%v#&%v#&%v#&%v", set.CloudPool[i], set.Node[i], set.AvailabilityZone[i], set.ResourceCategory[i], set.ResourceSubcategory[i], set.Specification[i], set.Comments[i])
		} else {
			key = fmt.Sprintf("%v#&%v#&%v#&%v#&%v#&%v#&%v", set.CloudPool[i], set.Node[i], set.AvailabilityZone[i], set.ResourceCategory[i], set.ResourceSubcategory[i], set.Specification[i], "")
		}

		// 如果 key 不存在于 map 中，则初始化它，并使用指针 //key是全部参数组成的key，value是int,map[string]int 类型
		if _, exists := aggregatedResources[key]; !exists {

			aggregatedResources[key] = &AggregatedData{
				totalCounts:         0, //统计相同小类总数量
				capacityCounts:      0, //统计相同小类总容量
				SingleCapacity:      0,
				SingleCapacitySlice: []int{},
			}
		}

		// 简单说就是将Key组合在一起变成一个Set集合，value是统计同类型的个数，和统计规格和每个规格的个数
		// 增加totalCount计数
		//统计相同产品的总数
		aggregatedResources[key].totalCounts++

		//统计存储相同类型的总量
		capInt, err := strconv.Atoi(set.Capacity[i])
		if err != nil {
			log.Println("CapInt获取Atoi转换失败", err)
		}
		//统计相同小类的总容量
		aggregatedResources[key].capacityCounts = aggregatedResources[key].capacityCounts + capInt

		aggregatedResources[key].SingleCapacitySlice = append(aggregatedResources[key].SingleCapacitySlice, capInt)
	}

	//第二步、拆分Key，并且把拆分的Key和Counts、spec、容量、都加入到一个[][]string切片中
	var countSlice [][]string
	//k=key ,data=aggregatedData类型
	for k, data := range aggregatedResources {
		// 每次循环开始时重新初始化 Slice
		Slice := strings.Split(k, "#&")

		//判断是否为共享桶
		ShareBucket(aggregatedResources, k, Slice[len(Slice)-1])

		// 将总数量和规格数量字符串添加到 Slice，Slice是每个key拆分后的数据，totalCounts是统计相同资源的个数，specCounts是将规格和数量组合的数据
		Slice = append(Slice, strconv.Itoa(data.totalCounts), strconv.Itoa(data.capacityCounts), strconv.Itoa(data.SingleCapacity))

		//fmt.Println("Slice=", Slice)
		//跳过资源小类等于NF或者NFI的空行
		if Slice[4] == "NF" || Slice[4] == "NFI" {
			continue
		}

		// 将 Slice 添加到 countSlice
		countSlice = append(countSlice, Slice)
	}

	// 打印结果以验证
	for _, slice := range countSlice {
		log.Println("存储资源StorageTotalResource函数打印结果：", slice)
	}
	return countSlice
}

// ShareBucket 判断是否为共享桶，判断条件是注释不为空，且所有容量相同
func ShareBucket(aggregatedResources map[string]*AggregatedData, k, comment string) {
	//如果注释为空，说明不是共享桶，直接赋值
	if comment == "" {
		aggregatedResources[k].SingleCapacity = aggregatedResources[k].capacityCounts

	} else if internal.AllSame(aggregatedResources[k].SingleCapacitySlice) {
		if len(aggregatedResources[k].SingleCapacitySlice) > 0 {
			aggregatedResources[k].SingleCapacity = aggregatedResources[k].SingleCapacitySlice[0]
		} else {
			aggregatedResources[k].SingleCapacity = 0
		}
	} else {
		aggregatedResources[k].SingleCapacity = aggregatedResources[k].capacityCounts
	}
}

func NetworkTotalResource(set Set) [][]string {
	//第一部将传入的所有数据每一行数据组成key
	// 创建一个 map 来存储聚合后的结果和规格统计
	aggregatedResources := make(map[string]*AggregatedData)

	for i := 0; i < len(set.ResourceCategory); i++ {
		key := fmt.Sprintf("%v#&%v#&%v#&%v#&%v#&%v#&%v", set.CloudPool[i], set.Node[i], set.AvailabilityZone[i], set.ResourceCategory[i], set.ResourceSubcategory[i], set.Specification[i], set.Bandwidth[i])

		// 如果 key 不存在于 map 中，则初始化它，并使用指针 //key是全部参数组成的key，value是int,map[string]int 类型
		if _, exists := aggregatedResources[key]; !exists {
			aggregatedResources[key] = &AggregatedData{
				totalCounts: 0,
			}
		}

		// 简单说就是将Key组合在一起变成一个Set集合，value是统计同类型的个数，和统计规格和每个规格的个数
		// 增加totalCount计数
		aggregatedResources[key].totalCounts++
	}

	//第二步、拆分Key，并且把拆分的Key和Counts、spec都加入到一个[][]string切片中
	var countSlice [][]string
	//k=key ,data=aggregatedData类型
	for k, data := range aggregatedResources {
		// 每次循环开始时重新初始化 Slice
		Slice := strings.Split(k, "#&")

		// 将总数量和规格数量字符串添加到 Slice，Slice是每个key拆分后的数据，totalCounts是统计相同资源的个数，specCounts是将规格和数量组合的数据
		Slice = append(Slice, strconv.Itoa(data.totalCounts))

		//fmt.Println("Slice=", Slice)
		//跳过资源小类等于NF或者NFI的空行
		if Slice[4] == "NF" || Slice[4] == "NFI" {
			continue
		}

		// 将 Slice 添加到 countSlice
		countSlice = append(countSlice, Slice)
	}

	// 打印结果以验证
	for _, slice := range countSlice {
		log.Println("NetworkTotalResource函数打印结果：", slice)
	}
	return countSlice
}

func SecurityTotalResource(set Set) [][]string {
	//第一部将传入的所有数据每一行数据组成key
	// 创建一个 map 来存储聚合后的结果和规格统计
	aggregatedResources := make(map[string]*AggregatedData)

	for i := 0; i < len(set.ResourceCategory); i++ {
		key := fmt.Sprintf("%v#&%v#&%v#&%v#&%v#&%v", set.CloudPool[i], set.Node[i], set.AvailabilityZone[i], set.ResourceCategory[i], set.ResourceSubcategory[i], set.Specification[i])

		// 如果 key 不存在于 map 中，则初始化它，并使用指针 //key是全部参数组成的key，value是int,map[string]int 类型
		if _, exists := aggregatedResources[key]; !exists {
			aggregatedResources[key] = &AggregatedData{
				totalCounts: 0,
			}
		}

		// 简单说就是将Key组合在一起变成一个Set集合，value是统计同类型的个数，和统计规格和每个规格的个数
		// 增加totalCount计数
		//统计相同产品的总数
		aggregatedResources[key].totalCounts++
	}

	//第二步、拆分Key，并且把拆分的Key和Counts、spec都加入到一个[][]string切片中
	var countSlice [][]string
	//k=key ,data=aggregatedData类型
	for k, data := range aggregatedResources {
		// 每次循环开始时重新初始化 Slice
		Slice := strings.Split(k, "#&")

		// 将总数量和规格数量字符串添加到 Slice，Slice是每个key拆分后的数据，totalCounts是统计相同资源的个数，specCounts是将规格和数量组合的数据
		Slice = append(Slice, strconv.Itoa(data.totalCounts))

		//fmt.Println("Slice=", Slice)
		//跳过资源小类等于NF或者NFI的空行
		if Slice[4] == "NF" || Slice[4] == "NFI" {
			continue
		}

		// 将 Slice 添加到 countSlice
		countSlice = append(countSlice, Slice)
	}

	// 打印结果以验证
	for _, slice := range countSlice {
		log.Println("安全资源SecurityTotalResource函数打印结果：", slice)
	}
	return countSlice
}

func PAASTotalResource(set Set) [][]string {
	//第一部将传入的所有数据每一行数据组成key
	// 创建一个 map 来存储聚合后的结果和规格统计
	aggregatedResources := make(map[string]*AggregatedData)

	for i := 0; i < len(set.ResourceCategory); i++ {
		key := fmt.Sprintf("%v#&%v#&%v#&%v#&%v#&%v", set.CloudPool[i], set.Node[i], set.AvailabilityZone[i], set.ResourceCategory[i], set.ResourceSubcategory[i], set.Specification[i])

		// 如果 key 不存在于 map 中，则初始化它，并使用指针 //key是全部参数组成的key，value是int,map[string]int 类型
		if _, exists := aggregatedResources[key]; !exists {
			aggregatedResources[key] = &AggregatedData{
				totalCounts: 0,
			}
		}

		// 简单说就是将Key组合在一起变成一个Set集合，value是统计同类型的个数，和统计规格和每个规格的个数
		// 增加totalCount计数
		aggregatedResources[key].totalCounts++

	}

	//第二步、拆分Key，并且把拆分的Key和Counts、spec都加入到一个[][]string切片中
	var countSlice [][]string
	//k=key ,data=aggregatedData类型
	for k, data := range aggregatedResources {
		// 每次循环开始时重新初始化 Slice
		Slice := strings.Split(k, "#&")

		// 将总数量和规格数量字符串添加到 Slice，Slice是每个key拆分后的数据，totalCounts是统计相同资源的个数，specCounts是将规格和数量组合的数据
		Slice = append(Slice, strconv.Itoa(data.totalCounts))

		//fmt.Println("Slice=", Slice)
		//跳过资源小类等于NF或者NFI的空行
		if Slice[4] == "NF" || Slice[4] == "NFI" {
			continue
		}

		// 将 Slice 添加到 countSlice
		countSlice = append(countSlice, Slice)
	}

	// 打印结果以验证
	for _, slice := range countSlice {
		log.Println("PAASTotalResource函数打印结果：", slice)
	}
	return countSlice
}
