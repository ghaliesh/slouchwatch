package test

import (
	"testing"

	"github.com/ghaliesh/slouchwatch/loader"
)

func TestLoad(t *testing.T) {
	var fileLoaer loader.EnvFileLoader = loader.NewDotEnvFileLoader()
	actual := fileLoaer.Load("../.env")

	if actual != nil {
		t.Fatalf("FAILED: expected error to be %v but got: %v", nil, actual)
	}
}
