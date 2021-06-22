package usecases

import (
	"marla.one/clean/apps/ping/entities"
	"marla.one/clean/apps/ping/repositories"
	"marla.one/clean/pkg/interfaces"
)

type PingUseCase struct {
	ctx interfaces.AppContext
}

var _ interfaces.UseCase = &PingUseCase{}

func NewPingUseCase(ctx interfaces.AppContext) *PingUseCase {
	return &PingUseCase{
		ctx: ctx,
	}
}

func (uc *PingUseCase) GetMessage() string {
	r, err := uc.ctx.GetRepository("ping")

	if err != nil {
		panic(err)
	}

	repo := r.(*repositories.PingRepository)

	entity, err := repo.FetchOne(nil)

	if err != nil {
		panic(err)
	}

	message := entity.(*entities.PingMessage)

	return message.Message
}
