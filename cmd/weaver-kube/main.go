package main

import (
	"context"
	"github.com/ServiceWeaver/weaver-kube/tool"
	"github.com/ServiceWeaver/weaver/runtime/protos"
	"github.com/kanengo/wdeployer/internal/logging"
	"os"
)

func main() {
	loggingPrettyPrinter := logging.NewPrettyPrinter(os.Stdout)
	handleLogEntry := func(ctx context.Context, entry *protos.LogEntry) error {
		return loggingPrettyPrinter.Format(entry)
	}
	tool.Run("weaver-kube", tool.Plugins{
		HandleLogEntry:   handleLogEntry,
		HandleTraceSpans: nil,
		HandleMetrics:    nil,
	})
}
