package main

import "github.com/marlaone/clean/interfaces"

type CleanAppContext struct {
	*Registrable
}

var _ interfaces.AppContext = &CleanAppContext{}

func NewAppContext(registry interfaces.Registry) *CleanAppContext {
	return &CleanAppContext{
		Registrable: NewRegistrable(registry),
	}
}
