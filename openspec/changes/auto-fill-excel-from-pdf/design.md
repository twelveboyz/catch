# Design: Auto-Fill Excel from PDF

## Context

### Current State

CMDB统计功能（选项1）的执行流程：
1. 用户输入项目目录路径
2. 调用 `internal.InputMessages()` 手动输入需求人、需求方、OA发文时间、主账号等信息
3. 调用 `cmdb/excel.go:initCountExcelFile()` 创建Excel表格，将手动输入的信息填入A1-A5单元格
4. 处理资源录入表和安全资源表，生成统计结果

### Constraints

- 必须使用现有的 `tools/pdftotext.exe` 工具进行PDF转换
- 必须保持向后兼容性：PDF提取失败时回退到手动输入
- 不能修改其他功能（功能2、3、4）的行为
- PDF文件位于用户输入目录下的固定子目录：`1、订购评审纪要`

### Stakeholders

- CMDB功能用户：减少手动输入工作
- 开发团队：维护代码清晰性和可测试性

## Goals / Non-Goals

**Goals:**
- 自动从PDF中提取申请人、需求方、OA发文时间、主账号四个字段
- 将提取的信息自动填充到Excel表格的B1、B2、B3、B5单元格
- 提供优雅的降级机制：提取失败时回退到手动输入
- 保持代码可维护性和可测试性

**Non-Goals:**
- 不修改其他功能（Daily、Config、ResourceCount）的PDF处理逻辑
- 不改变Excel表格的整体结构（仅填充前几行）
- 不实现PDF文件的上传或下载功能
- 不修改PDF内容，仅读取

## Decisions

### 1. 代码结构：新建 `cmdb/pdf_extractor.go` 模块

**决策**：创建独立的 `cmdb/pdf_extractor.go` 文件，而非在现有文件中添加代码。

**理由**：
- **单一职责**：PDF信息提取是一个独立的功能模块，应该有自己的文件
- **可测试性**：独立的文件更容易编写单元测试
- **代码复用**：虽然当前只服务于CMDB功能，但未来其他功能可能也需要类似的PDF提取能力
- **清晰性**：将PDF提取逻辑与Excel生成逻辑分离，代码更易理解

**替代方案考虑**：
- 在 `cmdb/excel.go` 中添加PDF提取函数：❌ 会使文件过长，职责不清晰
- 在 `daily/util.go` 中添加CMDB专用的提取函数：❌ 违反模块边界，CMDB不应依赖daily模块

### 2. PDF查找策略：复用 `daily/util.go` 的逻辑

**决策**：参考 `daily/util.go` 中的 `findPdf()` 和 `MostSimilarFileName()` 函数，但适配CMDB的目录结构。

**理由**：
- **代码复用**：现有的PDF查找逻辑已经经过验证
- **一致性**：保持与daily功能相同的PDF命名规则（"集中流程管理系统"或 `[A-Z]\d{3}` 模式）
- **目录适配**：CMDB需要在 `1、订购评审纪要` 子目录中查找，而daily在根目录查找

**实现**：
```go
// 在 cmdb/pdf_extractor.go 中
func FindPdfInReviewDir(root string) (string, error) {
    reviewDir := filepath.Join(root, "1、订购评审纪要")
    // 检查目录是否存在
    // 复用 daily 的 PDF 查找逻辑
}
```

### 3. 字段提取：使用正则表达式

**决策**：为每个字段编写专用的正则表达式提取函数。

**理由**：
- **简单高效**：正则表达式是Go内置功能，无需额外依赖
- **精确匹配**：可以针对PDF的具体格式编写精确的正则模式
- **可调试**：正则表达式易于测试和调试

**字段提取函数设计**：
```go
// 提取申请人：匹配 "申请人： 黄昌焕"
func ExtractApplicant(content string) string

// 提取需求方：匹配 "所属独立业务X： X_研发管理平台" → "研发管理平台"
func ExtractDemandSide(content string) string

// 提取OA发文时间：匹配 "公有云资源订购部署信息接收" 的创建时间
func ExtractOAPublishTime(content string) string

// 提取主账号：匹配 "订购帐号： cmic_jsglpt"
func ExtractPrimaryAccount(content string) string
```

### 4. Excel单元格位置：B列而非A列

**决策**：将提取的信息填充到 **B列**（B1、B2、B3、B5），而非当前的A列。

**理由**：
- **用户需求**：用户明确要求填充到B1、B2、B3、B5
- **标签分离**：A列保持为标签（"需求人："、"需求方："等），B列为值
- **可读性**：标签和值分开显示更清晰

**实现**：修改 `initCountExcelFile()` 函数：
```go
// 当前代码（A列，标签+值在同一行）
f.SetSheetRow(sheetName, "A1", &[]interface{}{"需求人：", messages["DemandPerson"]})

// 修改后（A列标签，B列值）
f.SetSheetRow(sheetName, "A1", &[]interface{}{"需求人："})
f.SetSheetRow(sheetName, "B1", &[]interface{}{extractedInfo.Applicant})
```

### 5. 降级机制：提取失败时回退到手动输入

**决策**：PDF提取失败时，优雅降级到原有的手动输入方式。

**理由**：
- **向后兼容**：确保在PDF不存在或格式异常时，功能仍然可用
- **灵活性**：用户可以选择手动输入或自动提取
- **安全性**：避免因PDF问题导致整个流程中断

**实现**：
```go
// 在 cmdb/run.go 中
extractedInfo, err := pdf_extractor.ExtractFromPDF(root)
if err != nil {
    logrus.Warnf("PDF提取失败: %v，回退到手动输入", err)
    messages = internal.InputMessages()  // 原有方式
} else {
    // 使用提取的信息，允许用户确认或修改
    messages = internal.InputMessagesWithDefaults(extractedInfo)
}
```

## Risks / Trade-offs

### Risk 1: PDF格式变化导致提取失败
**风险**：PDF格式变化（如字段名称调整、布局改变）会导致正则匹配失败。

**缓解措施**：
- 使用灵活的正则模式（如 `\s*` 匹配任意空白）
- 记录详细的警告日志，便于排查问题
- 降级到手动输入，确保功能可用性
- 未来可考虑添加配置文件，允许用户自定义正则模式

### Risk 2: PDF文件路径不存在
**风险**：用户输入的目录下可能没有 `1、订购评审纪要` 子目录，或该目录下没有PDF文件。

**缓解措施**：
- 检查目录是否存在，不存在则直接回退到手动输入
- 记录清晰的日志提示，告知用户为何PDF提取失败
- 不影响其他功能的正常执行

### Risk 3: 正则表达式误匹配
**风险**：PDF中可能存在多处相似文本，导致正则匹配到错误的内容。

**缓解措施**：
- 使用精确的正则模式，包含足够的上下文（如 `公有云资源订购部署信息接收`）
- 优先匹配更具体的模式（如完整字段名而非部分关键字）
- 编写单元测试覆盖各种PDF格式

### Trade-off: 代码复用 vs 模块独立性
**权衡**：参考 `daily/util.go` 的代码，但需要适配CMDB的目录结构。

**决定**：在 `cmdb/pdf_extractor.go` 中实现CMDB专用的提取逻辑，保持模块独立性。

**理由**：
- CMDB和daily的PDF处理需求略有不同（目录结构、提取字段）
- 避免模块间的耦合，便于未来独立演进
- 代码量不大，复制少量逻辑是可以接受的

## Migration Plan

### 部署步骤

1. **开发阶段**：
   - 创建 `cmdb/pdf_extractor.go` 文件
   - 实现PDF查找和字段提取函数
   - 编写单元测试
   - 修改 `cmdb/run.go` 和 `cmdb/excel.go`

2. **测试阶段**：
   - 在测试环境验证PDF提取功能
   - 测试各种PDF格式和异常场景
   - 验证降级机制是否正常工作

3. **发布阶段**：
   - 编译新版本可执行文件
   - 更新用户文档，说明新的自动填充功能
   - 部署到生产环境

### 回滚策略

- 如果出现严重问题，可快速回滚到之前的版本
- 降级机制已内置，即使PDF提取失败，功能仍可正常使用
- 代码变更集中在 `cmdb` 模块，不影响其他功能

## Open Questions

1. **Excel单元格位置确认**：
   - 用户要求填充到B1、B2、B3、B5，但现有代码是A1、A2、A3、A5
   - **需要确认**：是否同时修改A列的标签，还是仅修改B列的值？

2. **手动输入的触发时机**：
   - 是否在PDF提取成功后，仍给用户一个确认/修改的机会？
   - **建议**：始终显示提取的信息，允许用户修改

3. **PDF提取失败的日志级别**：
   - 使用 `Warn` 还是 `Info` 级别记录提取失败？
   - **建议**：使用 `Warn` 级别，因为提取失败可能需要用户关注

4. **多PDF文件的处理**：
   - 如果 `1、订购评审纪要` 目录下有多个匹配的PDF文件，如何处理？
   - **建议**：优先选择包含"集中流程管理系统"的文件，其次选择匹配项目ID的文件
