package tests

import (
	"os"
	"testing"

	"github.com/ghaliesh/slouchwatch/config"
)

const (
	key           = "AnyKeyName"
	expectedValue = "AnyValueReally"
)

type MockEnvLoaer struct{}

func (l *MockEnvLoaer) Load(filenames ...string) error {
	os.Setenv(key, expectedValue)
	return nil
}

func NewMockLoader() config.EnvLoader {
	return &MockEnvLoaer{}
}

func TestGetEnvVar(t *testing.T) {
	value, err := config.GetEnvVar(NewMockLoader(), key)

	if err != nil {
		t.Fatalf("Failed, Expected err to be nil but got %v", err)
	}

	if value != expectedValue {
		t.Fatalf("Expected result of GetEnvVar to be %v but got %v", expectedValue, value)
	}

}
