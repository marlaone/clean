package clean

import "github.com/marlaone/clean/interfaces"

type App struct {
	*Registrable
	ctx interfaces.AppContext
}

var _ interfaces.App = &App{}

func NewApp(registry interfaces.Registry) *App {
	return &App{
		Registrable: NewRegistrable(registry),
		ctx:         NewAppContext(registry),
	}
}
func (a *App) SetContext(ctx interfaces.AppContext) {
	a.ctx = ctx
}

func (a *App) GetContext() interfaces.AppContext {
	return a.ctx
}

func (a *App) Setup() {

}

func (a *App) Run() {

}
