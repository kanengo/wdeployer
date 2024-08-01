// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package standalone

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/kanengo/wdeployer/internal/config"

	"github.com/ServiceWeaver/weaver/runtime"
	"github.com/ServiceWeaver/weaver/runtime/bin"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"github.com/ServiceWeaver/weaver/runtime/tool"
	"github.com/ServiceWeaver/weaver/runtime/version"
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"time"
)

const (
	configKey      = "github.com/ServiceWeaver/weaver/standalone"
	shortConfigKey = "standalone"
)

var deployCmd = tool.Command{
	Name:        "deploy",
	Description: "Deploy a Service Weaver app",
	Help:        "Usage:\n  weaver standalone deploy <config file>",
	Flags:       flag.NewFlagSet("deploy", flag.ContinueOnError),
	Fn:          deploy,
}

// deploy deploys an application on the local machine using a multiprocess
// deployer. Note that each component is deployed as a separate OS process.
func deploy(ctx context.Context, args []string) error {
	// Validate command line arguments.
	if len(args) == 0 {
		return fmt.Errorf("no config file provided")
	}
	if len(args) > 1 {
		return fmt.Errorf("too many arguments")
	}

	// Load the config file.
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

	// Parse the multi section of the config.
	standalone, err := config.GetDeployerConfig[StandaloneConfig, StandaloneConfig_ListenerOptions](configKey, shortConfigKey, appConfig)
	if err != nil {
		return err
	}
	standalone.App = appConfig

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
		//selfVersion, err := itool.SelfVersion()
		if err != nil {
			return fmt.Errorf("read self version: %w", err)
		}
		return fmt.Errorf(`
ERROR: The binary you're trying to deploy (%q) was built with
github.com/ServiceWeaver/weaver module version %s. However, the 'wdeploy
standalone' binary you're using was built with weaver module version %s.
These versions are incompatible.

We recommend updating both the weaver module your application is built with and
updating the 'wdeploy standalone' command by running the following.

    go get github.com/ServiceWeaver/weaver@latest
    go install github.com/ServiceWeaver/weaver/cmd/weaver@latest
`,
			binary, versions.ModuleVersion, "")
	}

	// Make temporary directory.
	tmpDir, err := runtime.NewTempDir()
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)
	runtime.OnExitSignal(func() { os.RemoveAll(tmpDir) })

	// Create the deployer.
	deploymentId := uuid.New().String()
	d, err := newDeployer(ctx, deploymentId, standalone, tmpDir)
	if err != nil {
		return fmt.Errorf("create deployer: %w", err)
	}

	// Run a status server.
	//lis, err := net.Listen("tcp", "localhost:0")
	//if err != nil {
	//	return fmt.Errorf("listen: %w", err)
	//}
	//mux := http.NewServeMux()
	//status.RegisterServer(mux, d, d.logger)
	//go func() {
	//	if err := serveHTTP(ctx, lis, mux); err != nil {
	//		fmt.Fprintf(os.Stderr, "status server: %v\n", err)
	//	}
	//}()

	// Deploy main.
	if err := d.startMain(); err != nil {
		return fmt.Errorf("start main process: %w", err)
	}

	// Wait for the status server to become active.
	//client := status.NewClient(lis.Addr().String())
	//for r := retry.Begin(); r.Continue(ctx); {
	//	_, err := client.Status(ctx)
	//	if err == nil {
	//		break
	//	}
	//	_, _ = fmt.Fprintf(os.Stderr, "status server %q unavailable: %#v\n", lis.Addr(), err)
	//}

	//errCh := make(chan error)
	//runtime.OnExitSignal(func() {
	//	errCh <- d.wait()
	//})
	//return <-errCh

	errCh := make(chan error, 1)
	go func() {
		errCh <- d.wait()
	}()
	done := make(chan struct{}, 1)
	runtime.OnExitSignal(func() {
		err = <-errCh
		time.Sleep(3 * time.Second)
		done <- struct{}{}
	})
	<-done
	return err
}

// defaultRegistry returns a registry in defaultRegistryDir().
//func defaultRegistry(ctx context.Context) (*status.Registry, error) {
//	return status.NewRegistry(ctx, registryDir)
//}
