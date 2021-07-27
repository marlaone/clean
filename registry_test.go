package clean_test

import (
	"testing"

	"github.com/marlaone/clean"
)

func TestRegistry(t *testing.T) {

	registry := clean.NewRegistry()

	registry.RegisterApp("mock", NewMockApp(registry))

	_, err := registry.GetApp("mock")

	if err != nil {
		t.Errorf("app registration failed: %v", err)
	}

	if len(registry.GetApps()) == 0 {
		t.Errorf("app registration failed because no app is registered")
	}

	t.Log("app registration works")
}
