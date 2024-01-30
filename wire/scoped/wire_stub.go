//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package scoped

import (
	"github.com/crclz/go-ioc-example/wire/singleton"
	"github.com/google/wire"
)

func initializeScopedRootServiceLocator(RootServiceLocator *singleton.RootServiceLocator) (*ScopedRootServiceLocator, error) {
	wire.Build(ScopedSet)
	return &ScopedRootServiceLocator{}, nil
}
