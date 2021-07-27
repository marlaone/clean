package ping

import (
	"context"

	"github.com/marlaone/clean"
	"github.com/marlaone/clean/examples/ping/adapters"
	"github.com/marlaone/clean/examples/ping/presenters/http"
	"github.com/marlaone/clean/examples/ping/repositories"
	"github.com/marlaone/clean/examples/ping/usecases"
	"github.com/marlaone/clean/interfaces"
)

type PingApp struct {
	interfaces.Registrable
	ctx interfaces.AppContext
}

var _ interfaces.App = &PingApp{}

func NewPingApp(registry interfaces.Registry) *PingApp {
	return &PingApp{
		Registrable: clean.NewRegistrable(registry),
		ctx:         clean.NewAppContext(registry),
	}
}

func (ping *PingApp) SetContext(ctx interfaces.AppContext) {
	ping.ctx = ctx
}

func (ping *PingApp) GetContext() interfaces.AppContext {
	return ping.ctx
}

func (ping *PingApp) Setup() {
	ping.GetRegistry().RegisterStorageAdapter("messages", adapters.NewMessagesAdapter())
	ping.GetContext().RegisterPresenter("http", "ping", http.NewPingHttpPresenter(ping.GetRegistry(), ping.GetContext()))
	ping.GetContext().RegisterPresenter("http", "custom", http.NewCustomHttpPresenter(ping.GetRegistry(), ping.GetContext()))
	ping.GetContext().RegisterRepository("ping", repositories.NewPingRepository(ping.GetRegistry(), ping.GetContext()))
	ping.GetContext().RegisterUseCase("ping", usecases.NewPingUseCase(ping.GetRegistry(), ping.GetContext()))
}

func (ping *PingApp) Run(ctx context.Context) {

}
