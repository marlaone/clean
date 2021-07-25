package main

import "github.com/marlaone/clean/interfaces"

type CleanApp struct {
	registry interfaces.Registry
	ctx      interfaces.AppContext
}

var _ interfaces.App = &CleanApp{}

func NewCleanApp() *CleanApp {
	return &CleanApp{
		registry: NewCleanRegistry(),
	}
}

func (a *CleanApp) SetRegistry(r interfaces.Registry) {
	a.registry = r
}

func (a *CleanApp) GetRegistry() interfaces.Registry {
	return a.registry
}

func (a *CleanApp) SetContext(ctx interfaces.AppContext) {
	a.ctx = ctx
}

func (a *CleanApp) GetContext() interfaces.AppContext {
	return a.ctx
}
