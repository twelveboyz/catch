package cmdb

import (
	"Catch/internal"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"regexp"
	"sort"
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

type ResourceCountInfo struct {
	CloudPool           string
	NodeName            string
	resourceCategory    string
	resourceSubcategory string
	Specification       string
	comment             string
	Capacity            string
	SingleCapacity      string
	Counts              string
}

func TotalResource(set Set) [][]string {
	//第一部将传入的所有数据每一行数据组成key
	// 创建一个 map 来存储聚合后的结果和规格统计
	aggregatedResources := make(map[string]*AggregatedData)

	for i := 0; i < len(set.ResourceCategory); i++ {
		var key string
		node := strings.Replace(set.Node[i], "_", "", 1)
		resourceCategory := set.ResourceCategory[i]
		cloudPoolAndResourceSubcategory := set.ResourceSubcategory[i]
		specification := set.Specification[i]
		comment := set.Comments[i]

		switch resourceCategory {
		case "弹性计算", "KCS":
			systemDiskSpec := regexp.MustCompile(`\p{Han}+`).FindString(comment)
			key = fmt.Sprintf("%v#&%v#&%v#&%v#&%v", node, resourceCategory, cloudPoolAndResourceSubcategory, specification, systemDiskSpec)
		case "云存储":
			if strings.Contains(cloudPoolAndResourceSubcategory, "对象存储") {
				sharedBucket := regexp.MustCompile(`共享桶`).FindString(comment)
				key = fmt.Sprintf("%v#&%v#&%v#&%v#&%v", node, resourceCategory, cloudPoolAndResourceSubcategory, specification, sharedBucket)
			} else {
				key = fmt.Sprintf("%v#&%v#&%v#&%v#&%v", node, resourceCategory, cloudPoolAndResourceSubcategory, specification, "")
			}
		case "云网络":
			key = fmt.Sprintf("%v#&%v#&%v#&%v#&%v", node, resourceCategory, cloudPoolAndResourceSubcategory, specification, comment)
		default:
			key = fmt.Sprintf("%v#&%v#&%v#&%v#&%v", node, resourceCategory, cloudPoolAndResourceSubcategory, specification, "")
		}

		// 如果 key 不存在于 map 中，则初始化它，并使用指针 //key是全部参数组成的key，value是int,map[string]int 类型
		if _, exists := aggregatedResources[key]; !exists {
			aggregatedResources[key] = &AggregatedData{
				capacityCounts:      0,
				SingleCapacity:      0,
				SingleCapacitySlice: []int{},
				totalCounts:         0,
			}
		}

		// 简单说就是将Key组合在一起变成一个Set集合，value是统计同类型的个数，和统计规格和每个规格的个数
		// 增加totalCount计数
		classifiedStatistics(resourceCategory, comment, aggregatedResources, key)
	}

	return splitKey(aggregatedResources)
}

// classifiedStatistics 根据资源大类进行不同的统计处理
func classifiedStatistics(resourceCategory, comment string, aggregatedResources map[string]*AggregatedData, key string) {
	//获取注释中的数字
	comment = regexp.MustCompile("[0-9]*").FindString(comment)

	//根据资源大类进行不同的统计处理
	switch resourceCategory {
	case "弹性计算", "KCS":
		aggregatedResources[key].totalCounts++

		capInt, err := strconv.Atoi(comment)
		if err != nil {
			log.Println("CapInt获取Atoi转换失败", err)
		}
		//统计相同小类的总容量
		aggregatedResources[key].capacityCounts = aggregatedResources[key].capacityCounts + capInt

	case "云存储":
		aggregatedResources[key].totalCounts++

		//统计存储相同类型的总量
		capInt, err := strconv.Atoi(comment)
		if err != nil {
			log.Println("CapInt获取Atoi转换失败", err)
		}

		//统计相同小类的总容量
		aggregatedResources[key].capacityCounts = aggregatedResources[key].capacityCounts + capInt

		aggregatedResources[key].SingleCapacitySlice = append(aggregatedResources[key].SingleCapacitySlice, capInt)

	default:
		aggregatedResources[key].totalCounts++
	}
}

func splitKey(aggregatedResources map[string]*AggregatedData) [][]string {
	//第二步、拆分Key，并且把拆分的Key和Counts、spec都加入到一个[][]string切片中
	countSlice := make([][]string, 0, len(aggregatedResources))

	fmt.Println(len(aggregatedResources), "aggregatedResources=", aggregatedResources)

	//将map中的key排序，排序规则是根据资源大类的权重
	keys := sortKey(aggregatedResources)

	//遍历排序后的key
	for _, k := range keys {
		data, ok := aggregatedResources[k]
		if !ok {
			logrus.Errorf("splitKey函数获取aggregatedResources[%v]失败", k)
			continue
		}

		// 每次循环开始时重新初始化 Slice
		slice := strings.Split(k, "#&")

		//跳过资源小类空行
		if slice[2] == "NF" || slice[2] == "NFI" {
			continue
		}

		cloudPoolAndResourceSubcategory := slice[2]
		if strings.Contains(cloudPoolAndResourceSubcategory, "对象存储") {
			//判断是否为共享桶
			ShareBucket(aggregatedResources, k, slice[len(slice)-1])
		}

		// 将总数量和规格数量字符串添加到 Slice，Slice是每个key拆分后的数据，totalCounts是统计相同资源的个数，specCounts是将规格和数量组合的数据
		slice = append(slice, strconv.Itoa(data.totalCounts), strconv.Itoa(data.capacityCounts), strconv.Itoa(data.SingleCapacity))

		// 将 Slice 添加到 countSlice
		countSlice = append(countSlice, slice)
	}

	// 打印结果以验证
	for _, slice := range countSlice {
		log.Println("totalResource函数打印结果：", slice)
	}
	return countSlice
}

func sortKey(aggregatedResources map[string]*AggregatedData) []string {
	items := make([]string, 0, len(aggregatedResources))
	for k := range aggregatedResources {
		items = append(items, k)
	}

	keywords := []string{"弹性计算", "云存储", "云网络", "云PAAS"}

	sort.Slice(items, func(i, j int) bool {
		scoreI := calculateScore(items[i], keywords)
		scoreJ := calculateScore(items[j], keywords)

		if scoreI == scoreJ {
			return items[i] < items[j]
		}

		return scoreI < scoreJ
	})

	/*	fmt.Println("按关键词权重排序:")
		for i, item := range items {
			fmt.Printf("%d: %s (权重: %d)\n", i+1, item, calculateScore(item, keywords))
		}*/

	return items
}

// 计算字符串的权重分数（越小优先级越高）
func calculateScore(s string, keywords []string) int {
	for i, keyword := range keywords {
		if strings.Contains(s, keyword) {
			return i
		}
	}
	return len(keywords)
}

func splitCloudPortAndResourceSubcategory(cloudPoolAndResourceSubcategory string) (string, string) {
	//拆分云池，资源小类
	var cloudpool, resourceSubcategory string
	split := strings.Split(cloudPoolAndResourceSubcategory, "-")
	if len(split) == 2 {
		cloudpool = split[0]
		resourceSubcategory = split[1]
	} else {
		logrus.Errorf("格式错误无法拆分:%v", cloudPoolAndResourceSubcategory)
	}

	return cloudpool, resourceSubcategory
}

func slicesToStruct(slices [][]string) []ResourceCountInfo {
	var resourceCountInfos []ResourceCountInfo
	for _, slice := range slices {
		nodeName := slice[0]
		resourceCategory := slice[1]
		cloudPoolAndResourceSubcategory := slice[2]
		cloudPool, resourceSubcategory := splitCloudPortAndResourceSubcategory(cloudPoolAndResourceSubcategory)
		specification := slice[3]
		comment := slice[4]
		counts := slice[5]
		capacity := slice[6]
		singleCapacity := slice[7]

		rci := ResourceCountInfo{
			CloudPool:           cloudPool,
			NodeName:            nodeName,
			resourceCategory:    resourceCategory,
			resourceSubcategory: resourceSubcategory,
			Specification:       specification,
			comment:             comment,
			Counts:              counts,
			Capacity:            capacity,
			SingleCapacity:      singleCapacity,
		}

		resourceCountInfos = append(resourceCountInfos, rci)
	}

	return resourceCountInfos
}

// ShareBucket 判断是否为共享桶，判断条件是注释不为空，且所有容量相同
func ShareBucket(aggregatedResources map[string]*AggregatedData, k, comment string) {
	//如果注释为空，说明不是共享桶，直接赋值
	if !strings.Contains(comment, "共享桶") {
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
