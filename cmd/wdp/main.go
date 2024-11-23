package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/kanengo/wdeployer/kube"
	"github.com/kanengo/wdeployer/mono"

	"os"
	"runtime"
	"runtime/debug"
)

const usage = `USAGE

  wdp standalone              // 部署单体weaver进程，所有组件在用一个进程里本地调用
  weaver version                  // show wdp version

DESCRIPTION

  使用 wdp 命令部署不同类型的weaver进程.
`

func main() {
	flag.Usage = func() { _, _ = fmt.Fprint(os.Stderr, usage) }
	flag.Parse()
	if len(flag.Args()) == 0 {
		_, _ = fmt.Fprint(os.Stderr, usage)
		os.Exit(1)
	}
	switch flag.Arg(0) {
	case "mono":
		if err := mono.Deploy(context.Background(), flag.Args()[1:]); err != nil {
			_, _ = fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
	case "kube":
		kube.Deploy()
	case "version":
		//v, err := tool.SelfVersion()
		//if err != nil {
		//	_, _ = fmt.Fprintf(os.Stderr, "%w", err)
		//	os.Exit(1)
		//}
		info, ok := debug.ReadBuildInfo()
		if !ok {
			fmt.Println("无法读取构建信息")
			return
		}

		fmt.Printf("路径: %s\n", info.Path)
		fmt.Printf("主模块版本: %s\n", info.Main.Version)
		v := info.Main.Version
		fmt.Printf("%s %s/%s\n", v, runtime.GOOS, runtime.GOARCH)
	}
}
