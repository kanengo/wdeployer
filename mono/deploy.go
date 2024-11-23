package mono

import (
	"context"
	"errors"
	"fmt"
	"github.com/ServiceWeaver/weaver/runtime"
	"github.com/ServiceWeaver/weaver/runtime/bin"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"github.com/ServiceWeaver/weaver/runtime/envelope"
	"github.com/ServiceWeaver/weaver/runtime/logging"
	"github.com/ServiceWeaver/weaver/runtime/protos"
	"github.com/ServiceWeaver/weaver/runtime/version"
	"github.com/google/uuid"
	logging2 "github.com/kanengo/wdeployer/internal/logging"
	"golang.org/x/sync/errgroup"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

func Deploy(ctx context.Context, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no config file provided")
	}
	if len(args) > 1 {
		return fmt.Errorf("too many arguments")
	}

	configFile := args[0]
	bytes, err := os.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("load config file %q: %w\n", configFile, err)
	}

	// Parse and sanity-check the application section of the config.
	appConfig, err := runtime.ParseConfig(configFile, string(bytes), codegen.ComponentConfigValidator)
	if err != nil {
		return fmt.Errorf("load config file %q: %w\n", configFile, err)
	}
	if _, err := os.Stat(appConfig.Binary); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("binary %q doesn't exist", appConfig.Binary)
	}

	// Check version compatibility.
	versions, err := bin.ReadVersions(appConfig.Binary)
	if err != nil {
		return fmt.Errorf("read versions: %w", err)
	}
	if versions.DeployerVersion != version.DeployerVersion {
		// Try to relativize the binary, defaulting to the absolute path if
		// there are any errors..
		binary := appConfig.Binary
		if cwd, err := os.Getwd(); err == nil {
			if rel, err := filepath.Rel(cwd, appConfig.Binary); err == nil {
				binary = rel
			}
		}
		return fmt.Errorf(`
ERROR: The binary you're trying to deploy (%q) was built with
github.com/ServiceWeaver/weaver module version %s. However, the 'wdp
standalone' binary you're using was built with weaver module version %s.
These versions are incompatible.

We recommend updating both the weaver module your application is built with and
updating the 'wdp standalone' command by running the following.

    go get github.com/ServiceWeaver/weaver@latest
    go install github.com/ServiceWeaver/weaver/cmd/weaver@latest
`,
			binary, versions.ModuleVersion, "")
	}

	return deploy(ctx, appConfig)
}

func deploy(ctx context.Context, app *protos.AppConfig) error {
	deploymentId := uuid.Must(uuid.NewV7()).String()
	wLet := &protos.WeaveletArgs{
		App:             app.Name,
		DeploymentId:    deploymentId,
		Id:              deploymentId,
		RunMain:         true,
		InternalAddress: ":0", // 内部rpc 监听地址
		LogLevel:        app.LogLevel,
	}
	logger := logging.StderrLogger(logging.Options{
		App:        app.Name,
		Deployment: deploymentId,
		Component:  "mono",
		Weavelet:   wLet.Id,
		LogLevel:   slog.Level(app.LogLevel),
		Attrs:      []string{"serviceweaver/system", ""},
	})

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	e, err := envelope.NewEnvelope(ctx, wLet, app, envelope.Options{
		Logger: logger,
	})
	if err != nil {
		return fmt.Errorf("new mono: create envelope: %w", err)
	}

	var running errgroup.Group

	h := &handler{
		e:       e,
		printer: logging2.NewPrettyPrinter(os.Stdout),
	}
	running.Go(func() error {
		err := e.Serve(h)
		logger.Info("envelope serve end", "err", err)
		cancel()
		return err
	})

	//running.Go(func() error {
	//	<-ctx.Done()
	//	return nil
	//})

	done := make(chan struct{})
	runtime.OnExitSignal(func() {
		timer := time.NewTimer(10 * time.Second)
		select {
		case <-timer.C:
			cancel()
		case <-ctx.Done():
		}
		done <- struct{}{}
	})

	<-done
	return nil

}

type handler struct {
	e       *envelope.Envelope
	printer *logging2.PrettyPrinter
}

func (h *handler) ActivateComponent(ctx context.Context, request *protos.ActivateComponentRequest) (*protos.ActivateComponentReply, error) {
	_ = h.e.UpdateRoutingInfo(&protos.RoutingInfo{
		Component: request.Component,
		Local:     true,
	})
	return &protos.ActivateComponentReply{}, nil
}

func (h *handler) GetListenerAddress(ctx context.Context, request *protos.GetListenerAddressRequest) (*protos.GetListenerAddressReply, error) {
	return &protos.GetListenerAddressReply{Address: "localhost:0"}, nil
}

func (h *handler) ExportListener(ctx context.Context, request *protos.ExportListenerRequest) (*protos.ExportListenerReply, error) {
	return &protos.ExportListenerReply{}, nil
}

func (h *handler) GetSelfCertificate(ctx context.Context, request *protos.GetSelfCertificateRequest) (*protos.GetSelfCertificateReply, error) {
	return &protos.GetSelfCertificateReply{}, nil
}

func (h *handler) VerifyClientCertificate(ctx context.Context, request *protos.VerifyClientCertificateRequest) (*protos.VerifyClientCertificateReply, error) {
	return &protos.VerifyClientCertificateReply{}, nil
}

func (h *handler) VerifyServerCertificate(ctx context.Context, request *protos.VerifyServerCertificateRequest) (*protos.VerifyServerCertificateReply, error) {
	return &protos.VerifyServerCertificateReply{}, nil
}

func (h *handler) LogBatch(ctx context.Context, batch *protos.LogEntryBatch) error {
	for _, entry := range batch.Entries {
		log(h.printer, entry)
	}

	return nil
}

func (h *handler) HandleTraceSpans(ctx context.Context, spans *protos.TraceSpans) error {
	return nil
}

var _ envelope.EnvelopeHandler = (*handler)(nil)

func log(printer *logging2.PrettyPrinter, e *protos.LogEntry) {
	if !logging.IsSystemGenerated(e) {
		_, _ = fmt.Fprintln(os.Stderr, printer.Format(e))
	}
	//db.Add(e)
}
