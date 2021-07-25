package interfaces

type AppContext interface {
	Registrable
	RegisterPresenter(name string, presenterType string, presenter Presenter)
	GetPresenter(name string, presenterType string) (Presenter, error)
	RegisterUseCase(name string, uc UseCase)
	GetUseCase(name string) (UseCase, error)
	RegisterRepository(name string, repo Repository)
	GetRepository(name string) (Repository, error)
}
