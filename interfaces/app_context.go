package interfaces

type AppContext interface {
	Registrable
	AppContextable
	RegisterPresenter(presenterType string, name string, presenter Presenter)
	RegisterUseCase(name string, uc UseCase)
	RegisterRepository(name string, repo Repository)
}
