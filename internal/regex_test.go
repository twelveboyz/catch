package internal

import "testing"

func TestDetectBatch(t *testing.T) {
	cases := []struct {
		name string
		root string
		want string
	}{
		{"中文批次", `D:\项目\C284_短信平台资源扩容-第一批`, "第一批"},
		{"数字批次", `D:\项目\C284_短信平台资源扩容-第2批`, "第2批"},
		{"下划线分隔", `C284_短信平台资源扩容_第三批`, "第三批"},
		{"中间夹目录", `D:\项目\C284_短信平台资源扩容-交维材料-第一批`, "第一批"},
		{"深层路径", `D:\项目\移动云-C284_短信平台黑名单模块资源扩容需求-交维材料-第三批`, "第三批"},
		{"无分隔符不识别", `C284_短信平台资源扩容第一批`, ""},
		{"无批次", `C284_短信平台资源扩容`, ""},
		{"只有目录名", `C284_短信平台资源扩容-交维材料`, ""},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := DetectBatch(c.root); got != c.want {
				t.Errorf("DetectBatch(%q) = %q, 期望 %q", c.root, got, c.want)
			}
		})
	}
}
