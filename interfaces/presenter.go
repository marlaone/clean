package interfaces

import "net/http"

type Presenter interface {
	Registrable
	AppContextable
	CreateAction(w http.ResponseWriter, r *http.Request)
	ReadAction(w http.ResponseWriter, r *http.Request)
	UpdateAction(w http.ResponseWriter, r *http.Request)
	DeleteAction(w http.ResponseWriter, r *http.Request)
}
