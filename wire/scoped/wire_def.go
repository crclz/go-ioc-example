package scoped

import (
	"github.com/crclz/go-ioc-example/wire/singleton"
	"github.com/google/wire"
)

var ScopedSet = wire.NewSet(
	wire.Struct(new(ScopedRootComponent), "*"),

	NewAuthenticationService,

	// from root component provide ...
	FromRootComponentProvideDbConnectionService,
)

// RootComponent: put all services that you want to directly access in RootComponent
type ScopedRootComponent struct {
	authenticationService *AuthenticationService
}

func FromRootComponentProvideDbConnectionService(p *singleton.RootComponent) *singleton.DbConnectionService {
	return p.DbConnectionService
}
