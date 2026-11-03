package main

import (
	"Catch/cmdb"
	"Catch/config"
	"Catch/daily"
	"Catch/internal"
	"Catch/resourceCount"
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {
	//无参数时保持原有交互式菜单
	if len(os.Args) <= 1 {
		runInteractive()
		return
	}

	os.Exit(runCLI(os.Args[1:]))
}

func runInteractive() {
	//tip
	internal.Tip()

	for {
		choose := internal.InputChoose()
		switch choose {
		case "1":
			_ = cmdb.Run()
		case "2":
			_ = daily.Run()
		case "3":
			_ = config.Run()
		case "4":
			_ = resourceCount.Run()
		}
	}
}

// runCLI 解析命令行参数并执行对应任务，返回退出码
func runCLI(args []string) int {
	cmdName := args[0]

	if cmdName == "-h" || cmdName == "--help" || cmdName == "help" {
		printUsage()
		return 0
	}

	tasks := map[string]func() error{
		"cmdb":          cmdb.Run,
		"daily":         daily.Run,
		"config":        config.Run,
		"resourceCount": resourceCount.Run,
	}

	task, ok := tasks[cmdName]
	if !ok {
		fmt.Fprintf(os.Stderr, "未知的命令：%v\n\n", cmdName)
		printUsage()
		return 2
	}

	fs := flag.NewFlagSet(cmdName, flag.ContinueOnError)
	root := fs.String("root", "", "项目根目录（必填）")
	inBatches := fs.String("batch", "", "是否分批，如\"第1批\"，默认不分批")

	if err := fs.Parse(args[1:]); err != nil {
		return 2
	}

	if *root == "" {
		fmt.Fprintf(os.Stderr, "缺少必填参数 -root\n\n")
		printUsage()
		return 2
	}

	if err := checkRoot(*root); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n\n", err)
		printUsage()
		return 2
	}

	internal.SetCLIOptions(&internal.RunOptions{
		Root:      *root,
		InBatches: *inBatches,
	})

	code := 0
	if err := task(); err != nil {
		fmt.Fprintf(os.Stderr, "\n任务执行失败：%v\n", err)
		code = 1
	} else {
		fmt.Printf("\n任务执行完成\n")
	}

	waitEnter()

	return code
}

// checkRoot 校验根目录存在且为目录
func checkRoot(root string) error {
	fInfo, err := os.Stat(root)
	if err != nil {
		return fmt.Errorf("目录路径有误：%v", err)
	}

	if !fInfo.IsDir() {
		return errors.New("请选择项目根目录")
	}

	return nil
}

// waitEnter 提示并等待用户按下回车，读不到输入（如管道重定向）时不阻塞
func waitEnter() {
	fmt.Print("按回车键退出...")
	_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
}

func printUsage() {
	fmt.Print(`用法：Catch.exe <命令> -root <项目根目录> [-batch <批次>]

命令：
  cmdb           统计本批次交维信息
  daily          汇总本次建设摘要
  config         汇总本次建设信息
  resourceCount  统计多批次资源总量

参数：
  -root  项目根目录（必填）
  -batch 是否分批，如"第1批"，默认为不分批

示例：
  Catch.exe cmdb -root "D:\项目\C284_短信平台资源扩容" -batch "第1批"
  Catch.exe daily -root "D:\项目\C284_短信平台资源扩容"
  Catch.exe config -root "D:\项目\C284_短信平台资源扩容"
  Catch.exe resourceCount -root "D:\项目"

不带任何参数运行将进入交互式菜单。
`)
}
