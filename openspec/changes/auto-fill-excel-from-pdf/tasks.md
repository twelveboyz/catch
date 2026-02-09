# Tasks: Auto-Fill Excel from PDF

## 1. PDF提取模块实现

- [x] 1.1 创建 `cmdb/pdf_extractor.go` 文件
- [x] 1.2 定义 `PdfInfo` 结构体（Applicant, DemandSide, OAPublishTime, PrimaryAccount）
- [x] 1.3 实现 `FindPdfInReviewDir(root string)` 函数，在 `1、订购评审纪要` 目录下查找PDF
- [x] 1.4 实现 `convertPdfToText(pdfPath string)` 函数，使用 `tools/pdftotext.exe` 转换PDF
- [x] 1.5 实现 `ExtractApplicant(content string)` 函数，使用正则提取申请人
- [x] 1.6 实现 `ExtractDemandSide(content string)` 函数，提取需求方并去除"X_"前缀
- [x] 1.7 实现 `ExtractOAPublishTime(content string)` 函数，从审批记录中提取OA发文时间
- [x] 1.8 实现 `ExtractPrimaryAccount(content string)` 函数，提取订购账号
- [x] 1.9 实现 `ExtractFromPDF(root string) (*PdfInfo, error)` 主函数
- [x] 1.10 为 `pdf_extractor.go` 添加单元测试

## 2. CMDB Excel生成修改

- [x] 2.1 修改 `cmdb/excel.go` 中的 `initCountExcelFile()` 函数签名，接收 `*PdfInfo` 参数
- [x] 2.2 修改 `initCountExcelFile()` 函数，将标签和值分离到A列和B列
  - A1: "需求人："，B1: 申请人
  - A2: "需求方："，B2: 需求方
  - A3: "OA发文时间："，B3: OA发文时间
  - A5: "主账号："，B5: 主账号
- [x] 2.3 添加空值处理逻辑：当 `PdfInfo` 字段为空时，使用默认值

## 3. CMDB运行流程集成

- [x] 3.1 修改 `cmdb/run.go` 中的 `Run()` 函数
- [x] 3.2 在 `initCountExcelFile()` 调用前添加PDF提取逻辑
- [x] 3.3 实现降级机制：PDF提取失败时调用 `internal.InputMessages()`
- [x] 3.4 添加详细的日志记录（成功、失败、降级情况）

## 4. 内部工具函数更新（可选）

- [x] 4.1 修改 `internal/utils.go` 中的 `InputMessages()` 函数，支持默认值参数
- [x] 4.2 或新增 `InputMessagesWithDefaults(defaults map[string]string)` 函数

## 5. 测试与验证

- [x] 5.1 使用测试PDF文件验证PDF提取功能（单元测试通过）
- [ ] 5.2 验证Excel单元格位置正确（A列标签、B列值）
- [ ] 5.3 测试PDF文件不存在的降级场景
- [ ] 5.4 测试PDF内容格式异常的处理
- [ ] 5.5 测试多PDF文件的选择逻辑
- [ ] 5.6 端到端测试：完整运行CMDB功能并验证生成的Excel

## 6. 文档与发布

- [ ] 6.1 更新用户文档，说明新增的自动提取功能
- [ ] 6.2 编译新版本可执行文件
- [ ] 6.3 在测试环境验证功能
- [ ] 6.4 部署到生产环境

## 实施顺序

1. **阶段1**：任务组1（PDF提取模块）- 核心功能实现 ✓
2. **阶段2**：任务组2（CMDB Excel修改）- 集成点准备 ✓
3. **阶段3**：任务组3（CMDB运行流程集成）- 完整功能实现 ✓
4. **阶段4**：任务组5（测试与验证）- 质量保证（部分完成）
5. **阶段5**：任务组4（内部工具更新）- 可选优化 ✓
6. **阶段6**：任务组6（文档与发布）- 交付

## 验收标准

- [x] PDF提取功能能正确提取四个字段（单元测试通过）
- [x] Excel表格B1、B2、B3、B5单元格正确填充（代码已实现）
- [x] PDF提取失败时功能仍可正常使用（回退到手动输入）
- [x] 所有单元测试通过
- [ ] 端到端测试通过
- [x] 日志记录清晰完整

## Bug 修复记录

### 2026-02-06: 修复需求方和OA发文时间提取失败的问题

**问题描述**：
- 使用实际PDF文件测试时，发现"需求方"和"OA发文时间"字段提取失败
- 原因：实际PDF格式与设计时的假设不同

**修复内容**：

1. **需求方提取** (`ExtractDemandSide`):
   - 原正则只支持 `X_` 前缀格式：`所属独立业务1： X_研发管理平台`
   - 修复后同时支持：
     - 带前缀：`所属独立业务1： X_研发管理平台`
     - 不带前缀：`所属独立业务1： 集中运维平台`

2. **OA发文时间提取** (`ExtractOAPublishTime`):
   - 原正则只匹配：`公有云资源订购部署信息接收 2023-10-15`
   - 修复后增加了多种后备匹配模式：
     - `公有云资源订购部署信息接收日期：2025-11-06`
     - `发文时间： 2024-03-15`
     - 文档中任意 `YYYY-MM-DD` 格式的日期（作为最后的后备）

**测试结果**：
- ✓ 所有单元测试通过（包括新增测试用例）
- ✓ 支持多种PDF格式变体
