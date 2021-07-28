package main

import (
	"context"
	"sync"

	"github.com/marlaone/clean"
	"github.com/marlaone/clean/examples/eventlogger"
	"github.com/marlaone/clean/examples/ping"
	"github.com/marlaone/clean/examples/server"
	"github.com/marlaone/clean/interfaces"
)

func main() {
	var wg sync.WaitGroup
	ctx := context.Background()
	registry := clean.NewRegistry()

	registry.RegisterApp("ping", ping.NewPingApp(registry))
	registry.RegisterApp("server", server.NewServerApp(registry))
	registry.RegisterApp("eventlogger", eventlogger.NewEventLoggerApp(registry))

	for _, a := range registry.GetApps() {
		a.Setup()
	}

	for _, a := range registry.GetApps() {
		wg.Add(1)
		go (func(a interfaces.App, ctx context.Context) {
			defer wg.Done()
			a.Run(ctx)
		})(a, ctx)
	}

	wg.Wait()
}
