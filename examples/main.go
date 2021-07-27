package main

import (
	"github.com/marlaone/clean"
	"github.com/marlaone/clean/examples/ping"
	"github.com/marlaone/clean/examples/server"
)

func main() {
	registry := clean.NewRegistry()

	registry.RegisterApp("ping", ping.NewPingApp(registry))
	registry.RegisterApp("server", server.NewServerApp(registry))

	for _, a := range registry.GetApps() {
		a.Setup()
	}

	for _, a := range registry.GetApps() {
		a.Run()
	}
}
