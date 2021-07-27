package http

import (
	"net/http"

	"github.com/marlaone/clean"
	"github.com/marlaone/clean/examples/ping/usecases"
	"github.com/marlaone/clean/interfaces"
)

type PingHttpPresenter struct {
	interfaces.Registrable
	interfaces.AppContextable
}

var _ interfaces.Presenter = &PingHttpPresenter{}

func NewPingHttpPresenter(registry interfaces.Registry, appContext interfaces.AppContextable) *PingHttpPresenter {
	return &PingHttpPresenter{
		Registrable:    clean.NewRegistrable(registry),
		AppContextable: appContext,
	}
}

func (p *PingHttpPresenter) CreateAction(w http.ResponseWriter, r *http.Request) {

}

func (p *PingHttpPresenter) ReadAction(w http.ResponseWriter, r *http.Request) {
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

func (p *PingHttpPresenter) UpdateAction(w http.ResponseWriter, r *http.Request) {

}

func (p *PingHttpPresenter) DeleteAction(w http.ResponseWriter, r *http.Request) {

}
