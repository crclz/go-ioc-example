//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package singleton

import (
	"github.com/google/wire"
)

func InitializeRootDependencyLocator() (*RootDependencyLocator, error) {
	wire.Build(AppSet)
	return &RootDependencyLocator{}, nil
}
