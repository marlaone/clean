package eventlogger

import (
	"context"
	"fmt"

	"github.com/marlaone/clean"
	"github.com/marlaone/clean/interfaces"
)

type EventLoggerApp struct {
	interfaces.Registrable
	ctx interfaces.AppContext
}

var _ interfaces.App = &EventLoggerApp{}

func NewEventLoggerApp(registry interfaces.Registry) *EventLoggerApp {
	return &EventLoggerApp{
		Registrable: clean.NewRegistrable(registry),
		ctx:         clean.NewAppContext(registry),
	}
}

func (ev *EventLoggerApp) SetContext(ctx interfaces.AppContext) {
	ev.ctx = ctx
}

func (ev *EventLoggerApp) GetContext() interfaces.AppContext {
	return ev.ctx
}

func (ev *EventLoggerApp) Setup() {

}

func (ev *EventLoggerApp) Run(ctx context.Context) {
	fmt.Println("registering subscriber")
	ev.GetRegistry().GetEventBus().SubscribeAsync("log:info", func(message string) {
		fmt.Printf("new event message: %s\n", message)
	}, false)
}
