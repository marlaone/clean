package ping

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"marla.one/clean/apps/ping/repositories"
	"marla.one/clean/apps/ping/usecases"
	"marla.one/clean/pkg/apps"
	"marla.one/clean/pkg/interfaces"
)

type PingApp struct {
	router *chi.Mux
	ctx    interfaces.AppContext
}

var _ interfaces.App = &PingApp{}
var _ interfaces.HttpRoutableApp = &PingApp{}

func NewPingApp() *PingApp {
	return &PingApp{
		ctx: apps.NewAppContext("ping"),
	}
}

func (p *PingApp) OnCreate() {
	p.ctx.RegisterRepository("ping", repositories.NewPingRepository())
	p.ctx.RegisterUseCase("ping", usecases.NewPingUseCase(p.ctx))
}

func (p *PingApp) SetContext(ctx interfaces.AppContext) {
	p.ctx = ctx
}

func (p *PingApp) GetContext() interfaces.AppContext {
	return p.ctx
}

func (p *PingApp) SetHttpRouter(r *chi.Mux) {
	p.router = r
}

func (p *PingApp) GetHttpRouter() *chi.Mux {
	return p.router
}

func (p *PingApp) PresentHttp() {
	p.router.Mount("/ping", p.MountHttpRoutes())
}

func (p *PingApp) MountHttpRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		uc, err := p.ctx.GetUseCase("ping")
		if err != nil {
			panic(err)
		}

		usecase := uc.(*usecases.PingUseCase)
		w.Write([]byte(usecase.GetMessage()))
	})

	return r
}
