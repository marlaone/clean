package usecases

import (
	"github.com/marlaone/clean"
	"github.com/marlaone/clean/interfaces"
)

type PingUseCase struct {
	interfaces.Registrable
	interfaces.AppContextable
}

var _ interfaces.UseCase = &PingUseCase{}

func NewPingUseCase(registry interfaces.Registry, appContext interfaces.AppContextable) *PingUseCase {
	return &PingUseCase{
		Registrable:    clean.NewRegistrable(registry),
		AppContextable: appContext,
	}
}

func (uc *PingUseCase) GetMessage() string {
	repo, err := uc.GetRepository("ping")

	if err != nil {
		panic(err)
	}

	entities, err := repo.Read(nil)

	if err != nil {
		panic(err)
	}

	return entities[0].(string)
}
