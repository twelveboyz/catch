# Proposal: Auto-Fill Excel from PDF

## Why

当前CMDB统计功能（选项1）在生成统计表格时，前几行的关键信息（申请人、需求方、OA发文时间、主账号）需要用户手动输入，效率低下且容易出错。这些信息实际上已经存在于"1、订购评审纪要"目录的PDF文件中。

通过自动化从PDF提取这些信息并填充到Excel表格，可以减少用户操作步骤，提高数据准确性，提升整体工作效率。

## What Changes

在功能1（CMDB统计）中新增以下功能：

1. **PDF自动定位与提取**：
   - 自动在用户输入目录的 `1、订购评审纪要` 子目录下查找PDF文件
   - 支持两种PDF命名模式：
     - 文件名包含"集中流程管理系统"
     - 文件名匹配正则 `[A-Z]\d{3}` (如 `E020.pdf`, `E036_集中运维平台...`)

2. **PDF内容解析**：
   - 复用现有的 `tools/pdftotext.exe` 工具将PDF转换为文本
   - 使用正则表达式从PDF文本中提取以下字段：
     - **申请人**：匹配 `申请人：\s*(\p{Han}+)`，填充到单元格 **B1**
     - **需求方**：匹配 `所属独立业务X：\s*X_(.+)`，去除"X_"前缀后填充到单元格 **B2**
     - **OA发文时间**：匹配 `公有云资源订购部署信息接收.*?(\d{4}-\d{2}-\d{2})`，提取日期部分填充到单元格 **B3**
     - **主账号**：匹配 `订购帐号：\s*(\w+)`，填充到单元格 **B5**

3. **Excel表格自动填充**：
   - 修改 `cmdb/excel.go` 中的 `initCountExcelFile()` 函数
   - 在创建Excel时自动填充提取的信息到B1、B2、B3、B5单元格
   - 如果PDF提取失败或字段不存在，保持原有行为（通过用户输入）

4. **容错处理**：
   - PDF文件未找到：记录警告日志，继续执行（通过用户输入）
   - 字段提取失败：记录警告日志，对应单元格留空或使用默认值

## Capabilities

### New Capabilities
- `pdf-info-extractor`: PDF信息提取能力，包含从集中流程管理系统PDF中提取申请人、需求方、OA发文时间、主账号等字段的功能

### Modified Capabilities
- `cmdb-statistics`: CMDB统计功能的需求信息获取方式发生变化
  - **修改内容**：将原来通过 `internal.InputMessages()` 手动获取需求人、需求方、OA发文时间、主账号的方式，改为优先从PDF自动提取，提取失败时才回退到手动输入

## Impact

### 受影响的代码文件
- `cmdb/excel.go`:
  - `initCountExcelFile()` 函数：修改为接收提取的PDF信息参数
- `cmdb/run.go`:
  - `Run()` 函数：在生成Excel前调用PDF信息提取逻辑
- `cmdb/pdf_extractor.go` (新文件):
  - PDF信息提取相关函数（复用 `daily/util.go` 的PDF转换逻辑）
- `internal/utils.go` (可能修改):
  - `InputMessages()` 函数：添加可选参数支持使用PDF提取的默认值

### 依赖项
- 现有：`tools/pdftotext.exe` (Poppler工具)
- 现有：`github.com/xuri/excelize/v2` (Excel操作库)

### 兼容性
- 向后兼容：如果PDF提取失败，回退到原有的手动输入方式
- 不影响功能2、3、4（Daily、Config、ResourceCount）

### 测试需求
- 需要测试不同PDF命名格式的文件查找
- 需要测试PDF内容解析的准确性
- 需要测试PDF不存在或格式异常的容错场景
