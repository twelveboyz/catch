package cmdb

import (
	"testing"
)

func TestExtractApplicant(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected string
	}{
		{
			name:     "正常提取申请人",
			content:  "申请人： 黄昌焕",
			expected: "黄昌焕",
		},
		{
			name:     "申请人前后有空格",
			content:  "申请人：  张三  ",
			expected: "张三",
		},
		{
			name:     "没有申请人信息",
			content:  "没有相关内容",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractApplicant(tt.content)
			if result != tt.expected {
				t.Errorf("ExtractApplicant() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestExtractDemandSide(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected string
	}{
		{
			name:     "正常提取需求方（带X_前缀，数字编号）",
			content:  "所属独立业务1： X_研发管理平台",
			expected: "研发管理平台",
		},
		{
			name:     "不同业务编号（带X_前缀，数字编号）",
			content:  "所属独立业务2： X_测试平台",
			expected: "测试平台",
		},
		{
			name:     "实际PDF格式（字母X编号，带X_前缀）",
			content:  "所属独立业务X： X_集中运维平台",
			expected: "集中运维平台",
		},
		{
			name:     "不带X_前缀的需求方（数字编号）",
			content:  "所属独立业务1： 集中运维平台",
			expected: "集中运维平台",
		},
		{
			name:     "不带X_前缀的需求方（字母X编号）",
			content:  "所属独立业务X： 移动云平台",
			expected: "移动云平台",
		},
		{
			name:     "不带X_前缀的需求方（其他名称）",
			content:  "所属独立业务3： 移动云平台",
			expected: "移动云平台",
		},
		{
			name:     "实际PDF完整内容片段",
			content:  "所属独立业务X： X_集中运维平台 订购帐号： cmic_chinamobile",
			expected: "集中运维平台",
		},
		{
			name:     "没有需求方信息",
			content:  "没有相关内容",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractDemandSide(tt.content)
			if result != tt.expected {
				t.Errorf("ExtractDemandSide() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestExtractOAPublishTime(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected string
	}{
		{
			name:     "正常提取OA发文时间",
			content:  "公有云资源订购部署信息接收 2023-10-15",
			expected: "2023-10-15",
		},
		{
			name:     "多个空格的情况",
			content:  "公有云资源订购部署信息接收    2024-01-20",
			expected: "2024-01-20",
		},
		{
			name:     "公有云资源订购部署相关信息中的日期",
			content:  "公有云资源订购部署信息接收日期：2025-11-06",
			expected: "2025-11-06",
		},
		{
			name:     "文档中任意日期（作为后备）",
			content:  "发文时间： 2024-03-15 审批通过",
			expected: "2024-03-15",
		},
		{
			name:     "找到第一个日期作为后备",
			content:  "申请日期 2024-05-20 项目开始",
			expected: "2024-05-20",
		},
		{
			name:     "没有OA发文时间",
			content:  "没有相关内容",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractOAPublishTime(tt.content)
			if result != tt.expected {
				t.Errorf("ExtractOAPublishTime() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestExtractPrimaryAccount(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected string
	}{
		{
			name:     "正常提取订购账号",
			content:  "订购帐号： cmic_jsglpt",
			expected: "cmic_jsglpt",
		},
		{
			name:     "使用'账号'字样",
			content:  "订购账号： test_account123",
			expected: "test_account123",
		},
		{
			name:     "没有订购账号信息",
			content:  "没有相关内容",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractPrimaryAccount(tt.content)
			if result != tt.expected {
				t.Errorf("ExtractPrimaryAccount() = %v, want %v", result, tt.expected)
			}
		})
	}
}
