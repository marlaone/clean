package clean_test

import (
	"github.com/marlaone/clean"
	"github.com/marlaone/clean/interfaces"
)

type MockApp struct {
	interfaces.Registrable
	ctx interfaces.AppContext
}

var _ interfaces.App = &MockApp{}

func NewMockApp(registry interfaces.Registry) *MockApp {
	return &MockApp{
		Registrable: clean.NewRegistrable(registry),
		ctx:         clean.NewAppContext(registry),
	}
}

func (app *MockApp) SetContext(ctx interfaces.AppContext) {
	app.ctx = ctx
}

func (app *MockApp) GetContext() interfaces.AppContext {
	return app.ctx
}

func (app *MockApp) Setup() {

}

func (app *MockApp) Run() {

}
