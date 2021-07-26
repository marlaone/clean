package main

import "github.com/marlaone/clean/interfaces"

type CleanApp struct {
	*Registrable
	ctx interfaces.AppContext
}

var _ interfaces.App = &CleanApp{}

func NewCleanApp(registry interfaces.Registry) *CleanApp {
	return &CleanApp{
		Registrable: NewRegistrable(registry),
		ctx:         NewAppContext(registry),
	}
}
func (a *CleanApp) SetContext(ctx interfaces.AppContext) {
	a.ctx = ctx
}

func (a *CleanApp) GetContext() interfaces.AppContext {
	return a.ctx
}
