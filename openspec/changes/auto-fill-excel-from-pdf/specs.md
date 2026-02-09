# Specs: Auto-Fill Excel from PDF

本变更涉及以下规范：

## 新增规范

### `pdf-info-extractor`

从集中流程管理系统的PDF文件中提取项目关键信息的能力。

**关键需求**：
- PDF-001: PDF文件定位（在 `1、订购评审纪要` 目录下查找）
- PDF-002: PDF文本提取（使用 `pdftotext.exe`）
- PDF-003: 申请人提取（正则匹配）
- PDF-004: 需求方提取（去除"X_"前缀）
- PDF-005: OA发文时间提取（提取日期部分）
- PDF-006: 主账号提取（正则匹配）
- PDF-007: 错误处理与日志
- PDF-008: 接口设计

**完整规范**：参见 `openspec/specs/pdf-info-extractor/spec.md`

## 修改规范

### `cmdb-statistics`

CMDB统计功能的增量规范，涉及PDF自动提取集成。

**变更需求**：
- CMDB-002: Excel初始化（修改单元格位置为A列标签、B列值）
- CMDB-003: PDF信息自动提取（新增）

**完整规范**：参见 `openspec/specs/cmdb-statistics/spec.md`

## 规范链接

- [pdf-info-extractor 完整规范](../../specs/pdf-info-extractor/spec.md)
- [cmdb-statistics 完整规范](../../specs/cmdb-statistics/spec.md)
