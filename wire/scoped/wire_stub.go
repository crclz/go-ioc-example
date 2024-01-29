//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package scoped

import (
	"github.com/crclz/go-ioc-example/wire/singleton"
	"github.com/google/wire"
)

func initializeScopedRootDependencyLocator(RootDependencyLocator *singleton.RootDependencyLocator) (*ScopedRootDependencyLocator, error) {
	wire.Build(ScopedSet)
	return &ScopedRootDependencyLocator{}, nil
}
