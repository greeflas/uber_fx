package app

import (
	"go.uber.org/fx"
	"testing"
)

func TestGetAllOptions(t *testing.T) {
	options := getAllOptions()

	if err := fx.ValidateApp(options...); err != nil {
		t.Fatal(err)
	}
}
