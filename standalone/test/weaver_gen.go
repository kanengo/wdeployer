// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package main

import (
	"context"
	"errors"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:      "github.com/ServiceWeaver/weaver/deployer/standalone/test/A",
		Iface:     reflect.TypeOf((*A)(nil)).Elem(),
		Impl:      reflect.TypeOf(a{}),
		Listeners: []string{"aLis1", "aLis2", "aLis3"},
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return a_local_stub{impl: impl.(A), tracer: tracer, fooAMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/deployer/standalone/test/A", Method: "FooA", Remote: false, Generated: true})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return a_client_stub{stub: stub, fooAMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/deployer/standalone/test/A", Method: "FooA", Remote: true, Generated: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return a_server_stub{impl: impl.(A), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return a_reflect_stub{caller: caller}
		},
		RefData: "⟦cacf3f61:wEaVeReDgE:github.com/ServiceWeaver/weaver/deployer/standalone/test/A→github.com/ServiceWeaver/weaver/deployer/standalone/test/B⟧\n⟦c2270d3e:wEaVeReDgE:github.com/ServiceWeaver/weaver/deployer/standalone/test/A→github.com/ServiceWeaver/weaver/deployer/standalone/test/C⟧\n⟦f160865c:wEaVeRlIsTeNeRs:github.com/ServiceWeaver/weaver/deployer/standalone/test/A→aLis1,aLis2,aLis3⟧\n",
	})
	codegen.Register(codegen.Registration{
		Name:      "github.com/ServiceWeaver/weaver/deployer/standalone/test/B",
		Iface:     reflect.TypeOf((*B)(nil)).Elem(),
		Impl:      reflect.TypeOf(b{}),
		Listeners: []string{"Listener"},
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return b_local_stub{impl: impl.(B), tracer: tracer, fooBMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/deployer/standalone/test/B", Method: "FooB", Remote: false, Generated: true})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return b_client_stub{stub: stub, fooBMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/deployer/standalone/test/B", Method: "FooB", Remote: true, Generated: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return b_server_stub{impl: impl.(B), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return b_reflect_stub{caller: caller}
		},
		RefData: "⟦d0843220:wEaVeRlIsTeNeRs:github.com/ServiceWeaver/weaver/deployer/standalone/test/B→Listener⟧\n",
	})
	codegen.Register(codegen.Registration{
		Name:      "github.com/ServiceWeaver/weaver/deployer/standalone/test/C",
		Iface:     reflect.TypeOf((*C)(nil)).Elem(),
		Impl:      reflect.TypeOf(c{}),
		Listeners: []string{"cLis"},
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return c_local_stub{impl: impl.(C), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any { return c_client_stub{stub: stub} },
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return c_server_stub{impl: impl.(C), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return c_reflect_stub{caller: caller}
		},
		RefData: "⟦86ec724c:wEaVeRlIsTeNeRs:github.com/ServiceWeaver/weaver/deployer/standalone/test/C→cLis⟧\n",
	})
	codegen.Register(codegen.Registration{
		Name:      "github.com/ServiceWeaver/weaver/Main",
		Iface:     reflect.TypeOf((*weaver.Main)(nil)).Elem(),
		Impl:      reflect.TypeOf(app{}),
		Listeners: []string{"appLis"},
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return main_local_stub{impl: impl.(weaver.Main), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any { return main_client_stub{stub: stub} },
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return main_server_stub{impl: impl.(weaver.Main), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return main_reflect_stub{caller: caller}
		},
		RefData: "⟦7698fefe:wEaVeReDgE:github.com/ServiceWeaver/weaver/Main→github.com/ServiceWeaver/weaver/deployer/standalone/test/A⟧\n⟦b7bc7e7d:wEaVeRlIsTeNeRs:github.com/ServiceWeaver/weaver/Main→appLis⟧\n",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[A] = (*a)(nil)
var _ weaver.InstanceOf[B] = (*b)(nil)
var _ weaver.InstanceOf[C] = (*c)(nil)
var _ weaver.InstanceOf[weaver.Main] = (*app)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*a)(nil)
var _ weaver.Unrouted = (*b)(nil)
var _ weaver.Unrouted = (*c)(nil)
var _ weaver.Unrouted = (*app)(nil)

// Local stub implementations.

type a_local_stub struct {
	impl        A
	tracer      trace.Tracer
	fooAMetrics *codegen.MethodMetrics
}

// Check that a_local_stub implements the A interface.
var _ A = (*a_local_stub)(nil)

func (s a_local_stub) FooA(ctx context.Context) (err error) {
	// Update metrics.
	begin := s.fooAMetrics.Begin()
	defer func() { s.fooAMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.A.FooA", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.FooA(ctx)
}

type b_local_stub struct {
	impl        B
	tracer      trace.Tracer
	fooBMetrics *codegen.MethodMetrics
}

// Check that b_local_stub implements the B interface.
var _ B = (*b_local_stub)(nil)

func (s b_local_stub) FooB(ctx context.Context) (err error) {
	// Update metrics.
	begin := s.fooBMetrics.Begin()
	defer func() { s.fooBMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.B.FooB", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.FooB(ctx)
}

type c_local_stub struct {
	impl   C
	tracer trace.Tracer
}

// Check that c_local_stub implements the C interface.
var _ C = (*c_local_stub)(nil)

type main_local_stub struct {
	impl   weaver.Main
	tracer trace.Tracer
}

// Check that main_local_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_local_stub)(nil)

// Client stub implementations.

type a_client_stub struct {
	stub        codegen.Stub
	fooAMetrics *codegen.MethodMetrics
}

// Check that a_client_stub implements the A interface.
var _ A = (*a_client_stub)(nil)

func (s a_client_stub) FooA(ctx context.Context) (err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.fooAMetrics.Begin()
	defer func() { s.fooAMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.A.FooA", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	var shardKey uint64

	// Call the remote method.
	var results []byte
	results, err = s.stub.Run(ctx, 0, nil, shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

type b_client_stub struct {
	stub        codegen.Stub
	fooBMetrics *codegen.MethodMetrics
}

// Check that b_client_stub implements the B interface.
var _ B = (*b_client_stub)(nil)

func (s b_client_stub) FooB(ctx context.Context) (err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.fooBMetrics.Begin()
	defer func() { s.fooBMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.B.FooB", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	var shardKey uint64

	// Call the remote method.
	var results []byte
	results, err = s.stub.Run(ctx, 0, nil, shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

type c_client_stub struct {
	stub codegen.Stub
}

// Check that c_client_stub implements the C interface.
var _ C = (*c_client_stub)(nil)

type main_client_stub struct {
	stub codegen.Stub
}

// Check that main_client_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_client_stub)(nil)

// Note that "weaver generate" will always generate the error message below.
// Everything is okay. The error message is only relevant if you see it when
// you run "go build" or "go run".
var _ codegen.LatestVersion = codegen.Version[[0][24]struct{}](`

ERROR: You generated this file with 'weaver generate' (devel) (codegen
version v0.24.0). The generated code is incompatible with the version of the
github.com/ServiceWeaver/weaver module that you're using. The weaver module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/ServiceWeaver/weaver

We recommend updating the weaver module and the 'weaver generate' command by
running the following.

    go get github.com/ServiceWeaver/weaver@latest
    go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

Then, re-run 'weaver generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/ServiceWeaver/weaver/issues.

`)

// Server stub implementations.

type a_server_stub struct {
	impl    A
	addLoad func(key uint64, load float64)
}

// Check that a_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*a_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s a_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "FooA":
		return s.fooA
	default:
		return nil
	}
}

func (s a_server_stub) fooA(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.FooA(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

type b_server_stub struct {
	impl    B
	addLoad func(key uint64, load float64)
}

// Check that b_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*b_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s b_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "FooB":
		return s.fooB
	default:
		return nil
	}
}

func (s b_server_stub) fooB(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.FooB(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

type c_server_stub struct {
	impl    C
	addLoad func(key uint64, load float64)
}

// Check that c_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*c_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s c_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	default:
		return nil
	}
}

type main_server_stub struct {
	impl    weaver.Main
	addLoad func(key uint64, load float64)
}

// Check that main_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*main_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s main_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	default:
		return nil
	}
}

// Reflect stub implementations.

type a_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that a_reflect_stub implements the A interface.
var _ A = (*a_reflect_stub)(nil)

func (s a_reflect_stub) FooA(ctx context.Context) (err error) {
	err = s.caller("FooA", ctx, []any{}, []any{})
	return
}

type b_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that b_reflect_stub implements the B interface.
var _ B = (*b_reflect_stub)(nil)

func (s b_reflect_stub) FooB(ctx context.Context) (err error) {
	err = s.caller("FooB", ctx, []any{}, []any{})
	return
}

type c_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that c_reflect_stub implements the C interface.
var _ C = (*c_reflect_stub)(nil)

type main_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that main_reflect_stub implements the weaver.Main interface.
var _ weaver.Main = (*main_reflect_stub)(nil)
