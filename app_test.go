package main

import (
	"github.com/marlaone/clean/interfaces"
)

type MockApp struct {
	*Registrable
	ctx interfaces.AppContext
}

var _ interfaces.App = &MockApp{}

func NewMockApp(registry interfaces.Registry) *MockApp {
	return &MockApp{
		Registrable: NewRegistrable(registry),
		ctx:         NewAppContext(registry),
	}
}

func (app *MockApp) SetContext(ctx interfaces.AppContext) {
	app.ctx = ctx
}

func (app *MockApp) GetContext() interfaces.AppContext {
	return app.ctx
}
