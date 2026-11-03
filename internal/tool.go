package internal

import (
	"os"
	"path/filepath"
)

// ToolPath 返回随程序分发的外部工具（tools/ 目录下的 exe）的绝对路径。
// 查找顺序：程序所在目录 -> 当前工作目录（go run 调试用）。
// 以程序目录为基准，保证从任意工作目录启动、或整体移动程序目录后仍能调用到工具。
// 不走 PATH：tools/ 是随程序分发的完整工具链（exe + 依赖 dll），
// 借用系统 PATH 里的同名程序会缺 dll 或版本不一致，失败比静默降级更好。
func ToolPath(name string) string {
	rel := filepath.Join("tools", name)

	if exe, err := os.Executable(); err == nil {
		if p := filepath.Join(filepath.Dir(exe), rel); isFile(p) {
			return p
		}
	}

	//开发调试（go run）时程序在临时目录，退回当前工作目录下的 tools/
	if isFile(rel) {
		if abs, err := filepath.Abs(rel); err == nil {
			return abs
		}
		return rel
	}

	//都找不到时返回期望位置，让报错信息能指明该把 tools/ 放在哪
	if exe, err := os.Executable(); err == nil {
		return filepath.Join(filepath.Dir(exe), rel)
	}

	return rel
}

func isFile(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}
