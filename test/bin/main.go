package main

import (
	"context"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/kanengo/wdeployer/internal/logging"
	"runtime/debug"
	"time"
)

//go:generate weaver generate

type A interface {
	FooA(ctx context.Context) error
}

type B interface {
	FooB(ctx context.Context) error
}

type C interface {
}

type app struct {
	weaver.Implements[weaver.Main]
	a      weaver.Ref[A]   //lint:ignore U1000 intentionally declared but not used
	appLis weaver.Listener //lint:ignore U1000 intentionally declared but not used
}

func (*app) Main(context.Context) error { return nil }

type a struct {
	weaver.Implements[A]
	b            weaver.Ref[B]   //lint:ignore U1000 intentionally declared but not used
	c            weaver.Ref[C]   //lint:ignore U1000 intentionally declared but not used
	aLis1, aLis2 weaver.Listener //lint:ignore U1000 intentionally declared but not used
	unused       weaver.Listener `weaver:"aLis3"` //lint:ignore U1000 intentionally declared but not used
}

func (a *a) FooA(ctx context.Context) error {
	a.Logger(ctx).Info("FooA")
	a.Logger(ctx).Debug("debug FooA")
	return a.b.Get().FooB(ctx)
}

func (a *a) Shutdown(ctx context.Context) error {
	a.Logger(ctx).Info("a shutdown-1")
	time.Sleep(2 * time.Second)
	a.Logger(ctx).Info("a shutdown-2")
	//time.Sleep(100 * time.Millisecond)
	return nil
}

type b struct {
	weaver.Listener
	weaver.Implements[B]
}

func (b *b) FooB(ctx context.Context) error {
	b.Logger(ctx).Info("FooB")
	return nil
}

type c struct {
	weaver.Listener `weaver:"cLis"`
	weaver.Implements[C]
}

func main() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("无法读取构建信息")
		return
	}
	logging.NewPrettyPrinter()
	fmt.Printf("路径: %s\n", info.Path)
	fmt.Printf("主模块版本: %s\n", info.Main.Version)
	//for _, dep := range info.Deps {
	//	fmt.Printf("依赖项: %s@%s\n", dep.Path, dep.Version)
	//}
	_ = weaver.Run(context.Background(), func(ctx context.Context, t *app) error {
		_ = t.a.Get().FooA(ctx)
		<-make(chan struct{})
		return nil
	})
	//done := make(chan struct{})
	//

	//go func() {
	//	<-signalCh
	//	time.Sleep(5 * time.Second)
	//	done <- struct{}{}
	//}()
	//
	//<-done
}
