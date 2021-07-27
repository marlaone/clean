package httputils

import (
	"github.com/marlaone/clean/interfaces"
)

func GetHttpPresenters(ctx interfaces.AppContextable, presenterTypeName string) map[string]interfaces.HttpPresenter {
	presenters, err := ctx.GetPresentersByType(presenterTypeName)

	if err != nil {
		return map[string]interfaces.HttpPresenter{}
	}

	httpPresenters := map[string]interfaces.HttpPresenter{}

	for presenterName, presenter := range presenters {
		httpPresenter, ok := presenter.(interfaces.HttpPresenter)

		if !ok {
			continue
		}

		httpPresenters[presenterName] = httpPresenter
	}

	return httpPresenters
}
