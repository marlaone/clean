package clean

import (
	"fmt"
	"strings"

	"github.com/marlaone/clean/interfaces"
)

type RoutableHttpPresenter struct {
	rewrites map[string]string
}

var _ interfaces.HttpPresenterRoutable = &RoutableHttpPresenter{}

var (
	validActions = []string{
		"create",
		"read",
		"update",
		"delete",
	}
)

func ActionExists(actionName string) bool {

	for _, validAction := range validActions {
		if validAction == actionName {
			return true
		}
	}

	return false
}

func ValidateRoutes(routeRewrites map[string]string) error {
	for actionName := range routeRewrites {
		if !ActionExists(actionName) {
			return fmt.Errorf("invalid route action name: %s. please use: %s", actionName, strings.Join(validActions, ","))
		}
	}
	return nil
}

func NewRoutableHttpPresenter(routeRewrites map[string]string) *RoutableHttpPresenter {

	if err := ValidateRoutes(routeRewrites); err != nil {
		panic(err)
	}
	return &RoutableHttpPresenter{
		rewrites: routeRewrites,
	}
}

func NewDefaultRoutesHttpPresenter() *RoutableHttpPresenter {
	return &RoutableHttpPresenter{
		rewrites: map[string]string{
			"create": "",
			"read":   "",
			"update": "",
			"delete": "",
		},
	}
}

func GetDefaultRoutes() map[string]string {
	return map[string]string{
		"create": "",
		"read":   "",
		"update": "",
		"delete": "",
	}
}

func (p *RoutableHttpPresenter) SetRoutes(routeRewrites map[string]string) error {
	return ValidateRoutes(routeRewrites)
}

func (p *RoutableHttpPresenter) GetRoutes() (string, string, string, string) {
	createRoute := "/"
	readRoute := "/"
	updateRoute := "/"
	deleteRoute := "/"

	if p.rewrites["create"] != "" {
		createRoute = p.rewrites["create"]
	}

	if p.rewrites["read"] != "" {
		readRoute = p.rewrites["read"]
	}

	if p.rewrites["update"] != "" {
		updateRoute = p.rewrites["update"]
	}

	if p.rewrites["delete"] != "" {
		deleteRoute = p.rewrites["delete"]
	}

	return createRoute, readRoute, updateRoute, deleteRoute
}
