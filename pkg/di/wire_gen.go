// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"github.com/int128/kauthproxy/pkg/adaptors/cmd"
	"github.com/int128/kauthproxy/pkg/adaptors/logger"
	"github.com/int128/kauthproxy/pkg/adaptors/network"
	"github.com/int128/kauthproxy/pkg/adaptors/portforwarder"
	"github.com/int128/kauthproxy/pkg/adaptors/resolver"
	"github.com/int128/kauthproxy/pkg/adaptors/reverseproxy"
	"github.com/int128/kauthproxy/pkg/usecases/authproxy"
)

// Injectors from di.go:

func NewCmd() cmd.Interface {
	reverseProxy := &reverseproxy.ReverseProxy{}
	portForwarder := &portforwarder.PortForwarder{}
	loggerLogger := &logger.Logger{}
	factory := &resolver.Factory{
		Logger: loggerLogger,
	}
	networkNetwork := &network.Network{}
	authProxy := &authproxy.AuthProxy{
		ReverseProxy:    reverseProxy,
		PortForwarder:   portForwarder,
		ResolverFactory: factory,
		Network:         networkNetwork,
		Logger:          loggerLogger,
	}
	cmdCmd := &cmd.Cmd{
		AuthProxy: authProxy,
		Logger:    loggerLogger,
	}
	return cmdCmd
}
