package config

import (
	"fmt"
	"github.com/ServiceWeaver/weaver/runtime"
	"github.com/ServiceWeaver/weaver/runtime/bin"
	"github.com/ServiceWeaver/weaver/runtime/protos"
	"google.golang.org/protobuf/proto"
)

type configProtoPointer[T any, L any] interface {
	*T
	proto.Message
	GetListeners() map[string]*L
}

// GetDeployerConfig extracts and validates the deployer config from the
// specified section in the app config.
func GetDeployerConfig[T, L any, TP configProtoPointer[T, L]](key, shortKey string, app *protos.AppConfig) (*T, error) {
	// Read the config.
	config := new(T)
	if err := runtime.ParseConfigSection(key, shortKey, app.Sections, config); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	// Validate the config.
	binListeners, err := bin.ReadListeners(app.Binary)
	if err != nil {
		return nil, fmt.Errorf("cannot read listeners from binary %s: %w", app.Binary, err)
	}
	all := make(map[string]struct{})
	for _, c := range binListeners {
		for _, l := range c.Listeners {
			all[l] = struct{}{}
		}
	}
	for lis := range TP(config).GetListeners() {
		if _, ok := all[lis]; !ok {
			return nil, fmt.Errorf("listeners %s specified in the config not found in the binary", lis)
		}
	}
	return config, nil
}
