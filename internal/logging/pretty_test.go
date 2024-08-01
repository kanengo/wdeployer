package logging

import (
	"fmt"
	"github.com/ServiceWeaver/weaver/runtime/protos"
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

	p := PrettyPrinter{}
	fmt.Println(p.Format(e))
}
