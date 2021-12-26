package test

import (
	"os"
	"testing"

	"github.com/ghaliesh/slouchwatch/config"
	"github.com/ghaliesh/slouchwatch/loader"
)

const (
	testEnvKeyName = "Some_Key_Name"
	testEnvValue   = "Some_Value"
)

type mockFileLoader struct {
}

func (m *mockFileLoader) Load(filenames ...string) error {
	return os.Setenv(testEnvKeyName, testEnvValue)
}

// Mock subject
var mockFileLoaderInstance loader.EnvFileLoader = &mockFileLoader{}

// Test subject
var dotEnvConfig config.EnvConfig = config.NewDotEnvConfigInstance(mockFileLoaderInstance)

func TestGetEnvVar(t *testing.T) {
	expectedValue := testEnvValue
	actualValue, err := dotEnvConfig.GetEnvVar(testEnvKeyName)

	if err != nil {
		t.Fatalf("FAILED: Expected error to be %v \n got: %v", nil, err)
	}

	if expectedValue != actualValue {
		t.Fatalf("FAILED: Expected result to be %v \n got: %v", expectedValue, actualValue)
	}
}
