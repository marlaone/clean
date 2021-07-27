package interfaces

import "net/http"

type HttpPresenter interface {
	Presenter
	HttpMiddlewares
	CreateAction(w http.ResponseWriter, r *http.Request)
	ReadAction(w http.ResponseWriter, r *http.Request)
	UpdateAction(w http.ResponseWriter, r *http.Request)
	DeleteAction(w http.ResponseWriter, r *http.Request)
}
