package internal

import "testing"

func TestFromCLIWithDefaults(t *testing.T) {
	defer SetCLIOptions(nil)

	cases := []struct {
		name string
		opts *RunOptions
		want string
	}{
		{"显式参数优先", &RunOptions{Root: `D:\项目\C284_短信平台资源扩容-第一批`, InBatches: "第9批"}, "第9批"},
		{"自动检测批次", &RunOptions{Root: `D:\项目\C284_短信平台资源扩容-第一批`}, "第一批"},
		{"检测不到用不分批", &RunOptions{Root: `D:\项目\C284_短信平台资源扩容`}, "不分批"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			SetCLIOptions(c.opts)
			got := FromCLIWithDefaults()
			if got == nil {
				t.Fatal("FromCLIWithDefaults() 返回 nil")
			}
			if got.InBatches != c.want {
				t.Errorf("InBatches = %q, 期望 %q", got.InBatches, c.want)
			}
		})
	}
}

func TestFromCLIWithDefaultsInteractive(t *testing.T) {
	SetCLIOptions(nil)
	if got := FromCLIWithDefaults(); got != nil {
		t.Errorf("交互模式下应返回 nil，实际 %+v", got)
	}
}
