package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/marlaone/clean"
	"github.com/marlaone/clean/interfaces"
)

type ServerApp struct {
	interfaces.Registrable
	ctx interfaces.AppContext
}

var _ interfaces.App = &ServerApp{}

func NewServerApp(registry interfaces.Registry) *ServerApp {
	return &ServerApp{
		Registrable: clean.NewRegistrable(registry),
		ctx:         clean.NewAppContext(registry),
	}
}

func (server *ServerApp) SetContext(ctx interfaces.AppContext) {
	server.ctx = ctx
}

func (server *ServerApp) GetContext() interfaces.AppContext {
	return server.ctx
}

func (server *ServerApp) Setup() {

}

func (server *ServerApp) Run() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	for _, a := range server.GetRegistry().GetApps() {
		httpPresenters, err := a.GetContext().GetPresentersByType("http")
		if err != nil {
			continue
		}
		for presenterName, p := range httpPresenters {
			r.Post(fmt.Sprintf("/%s", presenterName), p.CreateAction)
			r.Get(fmt.Sprintf("/%s", presenterName), p.ReadAction)
			r.Put(fmt.Sprintf("/%s", presenterName), p.UpdateAction)
			r.Delete(fmt.Sprintf("/%s", presenterName), p.DeleteAction)
		}
	}

	http.ListenAndServe(":7292", r)
}
