# Spec: CMDB Statistics (Delta)

CMDB统计功能的增量规范。

## REQUIREMENTS (Delta)

### CMDB-002: Excel初始化 (修改)
- **变更**：A列包含标签，B列包含值（而非A列同时包含标签和值）
- 新的单元格结构：
  - A1: "需求人："，B1: 申请人
  - A2: "需求方："，B2: 需求方
  - A3: "OA发文时间："，B3: OA发文时间
  - A5: "主账号："，B5: 主账号

### CMDB-003: PDF信息自动提取 (新增)
- 在Excel初始化前调用 `pdf_extractor.ExtractFromPDF(root)`
- 提取成功时使用提取的信息
- 提取失败时回退到手动输入
- 记录详细日志
