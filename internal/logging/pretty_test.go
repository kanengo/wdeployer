package logging

import (
	"fmt"
	"github.com/ServiceWeaver/weaver/runtime/protos"
	"os"
	"testing"
	"time"
)

func TestJsonPrinter_Format(t *testing.T) {
	e := &protos.LogEntry{
		App:        "",
		Version:    "test",
		Component:  "github.com/kanengo/akasar/component/user",
		Node:       "test",
		TimeMicros: time.Now().UnixMicro(),
		Level:      "info",
		File:       "E:\\Codes\\github\\akasar\\runtime\\logging\\pretty.go",
		Line:       42,
		Msg:        "test json printer",
		Attrs:      []string{"id", "123123213", "name", "leeka\nhello"},
	}

	p := NewPrettyPrinter(os.Stderr)
	_ = p.Format(e)
	fmt.Println("==============================")
	e = &protos.LogEntry{
		App:        "",
		Version:    "test-2",
		Component:  "github.com/kanengo/akasar/component/user",
		Node:       "test-2",
		TimeMicros: time.Now().UnixMicro(),
		Level:      "error",
		File:       "E:\\Codes\\github\\akasar\\runtime\\logging\\pretty.go",
		Line:       42,
		Msg:        "test json printe-2",
		Attrs:      []string{"id", "123123213", "name", "leeka\nhello"},
	}
	_ = p.Format(e)
}
