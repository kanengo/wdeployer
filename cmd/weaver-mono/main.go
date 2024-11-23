package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/kanengo/wdeployer/mono"
	"os"
)

func main() {
	if err := mono.Deploy(context.Background(), flag.Args()[1:]); err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
