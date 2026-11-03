package internal

// RunOptions 承载命令行模式下本应来自 stdin 的输入
// 为 nil 表示当前为交互模式，各输入函数保持原有读取 stdin 的行为
type RunOptions struct {
	Root      string // 项目根目录（命令行模式必填）
	InBatches string // 是否分批，空字符串表示默认"不分批"
}

var cliOptions *RunOptions

// SetCLIOptions 设置命令行参数，调用后各输入函数不再读取 stdin
func SetCLIOptions(o *RunOptions) {
	cliOptions = o
}

// FromCLI 返回命令行参数，交互模式下返回 nil
func FromCLI() *RunOptions {
	return cliOptions
}

// FromCLIWithDefaults 返回命令行参数并补齐默认值，交互模式下返回 nil
// 未指定 -batch 时自动从根目录路径中检测"第X批"，检测不到则为"不分批"
func FromCLIWithDefaults() *RunOptions {
	if cliOptions == nil {
		return nil
	}

	o := *cliOptions
	if o.InBatches == "" {
		o.InBatches = DetectBatch(o.Root)
	}
	if o.InBatches == "" {
		o.InBatches = "不分批"
	}

	return &o
}
