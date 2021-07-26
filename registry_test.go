package main

import (
	"testing"
)

func TestRegistry(t *testing.T) {

	registry := NewCleanRegistry()

	registry.RegisterApp("mock", NewMockApp(registry))

	_, err := registry.GetApp("mock")

	if err != nil {
		t.Errorf("app registration failed: %v", err)
	}

	t.Log("app registration works")
}
