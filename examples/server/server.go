package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/marlaone/clean"
	"github.com/marlaone/clean/httputils"
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

func (server *ServerApp) Run(ctx context.Context) {
	r := chi.NewRouter()

	for _, a := range server.GetRegistry().GetApps() {
		httpPresenters := httputils.GetHttpPresenters(a.GetContext(), "http")
		for presenterName, p := range httpPresenters {
			r.Mount(fmt.Sprintf("/%s", presenterName), (func() http.Handler {
				r := chi.NewRouter()
				for _, mw := range p.GetMiddlewares() {
					r.Use(mw)
				}
				r.Route("/", PresenterRouter(p))
				return r
			})())

		}
	}

	http.ListenAndServe(":7292", r)
}

func PresenterRouter(p interfaces.HttpPresenter) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/", p.CreateAction)
		r.Get("/", p.ReadAction)
		r.Put("/", p.UpdateAction)
		r.Delete("/", p.DeleteAction)
	}
}
