package environment

import (
	"testing"
)

func TestEnvironment(t *testing.T) {
	environment := NewEnvironment()

	if environment.ConfigFilePath == "" {
		t.Error("ConfigFilePath should be present.")
	}
}
