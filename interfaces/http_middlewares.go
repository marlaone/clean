package interfaces

import "net/http"

type HttpMiddlewares interface {
	AddMiddleware(func(next http.Handler) http.Handler)
	GetMiddlewares() []func(next http.Handler) http.Handler
}
