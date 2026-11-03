package cmdb

import (
	"strings"
	"testing"
)

// TestCombinationKeyOther 其他资源的key末段取备注，云PAAS仍留在空串分支
func TestCombinationKeyOther(t *testing.T) {
	key := CombinationKey("移动云-华北-呼和浩特", "其他资源", "公有云-其他资源", "APISIX", "容器服务 KCS")
	if !strings.HasSuffix(key, "#&容器服务 KCS") {
		t.Errorf("其他资源key末段应取备注，实际=%v", key)
	}

	key = CombinationKey("移动云-华北-呼和浩特", "云PAAS", "移动云-云数据库Redis", "2C4G", "备注内容")
	if !strings.HasSuffix(key, "#&") {
		t.Errorf("云PAASkey末段应保持为空，实际=%v", key)
	}
}

// TestSlicesToStructOtherUsesComment 其他资源有备注时覆盖资源小类，无备注时保留原小类
func TestSlicesToStructOtherUsesComment(t *testing.T) {
	slices := [][]string{
		{"移动云-华北-呼和浩特", "其他资源", "公有云-其他资源", "规格A", "容器服务 KCS", "1", "0", "0"},
		{"移动云-华北-呼和浩特", "其他资源", "公有云-其他资源", "规格B", "", "1", "0", "0"},
	}

	result := slicesToStruct(slices)
	if len(result) != 2 {
		t.Fatalf("期望2条数据，实际=%d", len(result))
	}

	if result[0].CloudPool != "公有云" || result[0].resourceSubcategory != "容器服务 KCS" {
		t.Errorf("云池=%v 资源小类=%v，期望 公有云/容器服务 KCS", result[0].CloudPool, result[0].resourceSubcategory)
	}

	if result[1].CloudPool != "公有云" || result[1].resourceSubcategory != "其他资源" {
		t.Errorf("云池=%v 资源小类=%v，期望 公有云/其他资源", result[1].CloudPool, result[1].resourceSubcategory)
	}
}

// TestTotalResourceOtherGroups 其他资源按备注和规格聚合，备注相同规格相同才合并
func TestTotalResourceOtherGroups(t *testing.T) {
	set := Set{
		ResourceCategory:    []string{"其他资源", "其他资源", "其他资源"},
		ResourceSubcategory: []string{"公有云-其他资源", "公有云-其他资源", "公有云-其他资源"},
		Specification:       []string{"规格A", "规格B", "规格A"},
		Node:                []string{"_移动云_华北_呼和浩特", "_移动云_华北_呼和浩特", "_移动云_华北_呼和浩特"},
		Comments:            []string{"容器服务 KCS", "容器服务 CNP版", "容器服务 KCS"},
	}

	result := TotalResource(set)
	if len(result) != 2 {
		t.Fatalf("期望2组数据，实际=%d，内容=%v", len(result), result)
	}

	for _, slice := range result {
		if slice[4] == "容器服务 KCS" && slice[5] != "2" {
			t.Errorf("容器服务 KCS 统计数期望2，实际=%v", slice[5])
		}
	}
}
