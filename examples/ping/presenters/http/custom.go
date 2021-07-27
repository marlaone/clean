package http

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/marlaone/clean"
	"github.com/marlaone/clean/examples/ping/usecases"
	"github.com/marlaone/clean/httputils"
	"github.com/marlaone/clean/interfaces"
)

type CustomPresenter struct {
	interfaces.Registrable
	interfaces.AppContextable
	interfaces.HttpMiddlewares
	interfaces.HttpPresenterRoutable
}

var _ interfaces.HttpPresenter = &CustomPresenter{}

func NewCustomHttpPresenter(registry interfaces.Registry, appContext interfaces.AppContextable) *PingHttpPresenter {

	routes := clean.GetDefaultRoutes()

	routes["read"] = "/git"

	return &PingHttpPresenter{
		Registrable:    clean.NewRegistrable(registry),
		AppContextable: appContext,
		HttpMiddlewares: httputils.NewHttpMiddlewares(
			middleware.Logger,
		),
		HttpPresenterRoutable: clean.NewRoutableHttpPresenter(routes),
	}
}

func (p *CustomPresenter) CreateAction(w http.ResponseWriter, r *http.Request) {

}

func (p *CustomPresenter) ReadAction(w http.ResponseWriter, r *http.Request) {
	uc, err := p.GetUseCase("ping")

	if err != nil {
		panic(err)
	}

	pingUseCase, ok := uc.(*usecases.PingUseCase)

	if !ok {
		panic("invalid usecase")
	}

	w.Write([]byte(pingUseCase.GetMessage()))
}

func (p *CustomPresenter) UpdateAction(w http.ResponseWriter, r *http.Request) {

}

func (p *CustomPresenter) DeleteAction(w http.ResponseWriter, r *http.Request) {

}
