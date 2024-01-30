//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package singleton

import (
	"github.com/google/wire"
)

func InitializeRootServiceLocator() (*RootServiceLocator, error) {
	wire.Build(AppSet)
	return &RootServiceLocator{}, nil
}
