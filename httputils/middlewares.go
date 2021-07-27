package httputils

import (
	"net/http"

	"github.com/marlaone/clean/interfaces"
)

type HttpMiddlewares struct {
	middlewares []func(next http.Handler) http.Handler
}

var _ interfaces.HttpMiddlewares = &HttpMiddlewares{}

func NewHttpMiddlewares(middlewares ...func(next http.Handler) http.Handler) *HttpMiddlewares {
	return &HttpMiddlewares{
		middlewares: middlewares,
	}
}

func NewEmptyHttpMiddlewares() *HttpMiddlewares {
	return &HttpMiddlewares{
		middlewares: []func(next http.Handler) http.Handler{},
	}
}

func (hmw *HttpMiddlewares) AddMiddleware(mw func(next http.Handler) http.Handler) {
	hmw.middlewares = append(hmw.middlewares, mw)
}

func (hmw *HttpMiddlewares) GetMiddlewares() []func(next http.Handler) http.Handler {
	return hmw.middlewares
}
