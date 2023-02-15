package app

import (
	"testing"

	"go.uber.org/fx"
)

func TestGetAllOptions(t *testing.T) {
	options := getAllOptions()

	if err := fx.ValidateApp(options...); err != nil {
		t.Fatal(err)
	}
}
