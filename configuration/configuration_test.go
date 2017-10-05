package configuration

import (
	"testing"

	"github.com/tnoda78/posiserver/environment"
)

func TestNewConfiguration(t *testing.T) {
	config := NewConfiguration()

	if config == nil {
		t.Error("config should be present.")
	}
}

func TestLoadConfigFile(t *testing.T) {
	config := NewConfiguration()
	env := environment.NewEnvironment()
	env.ConfigFilePath = "config.test.yml"

	_, err := config.LoadConfigFile(env)

	if err != nil {
		t.Errorf("should not raise error.%v", err)
	}

	if config.RootPath != "static" {
		t.Errorf("config.RootPath should be static, but %v", config)
	}

	if len(config.WriteFiles) != 1 {
		t.Errorf("config.WriteFiles size should be 1, but %v", len(config.WriteFiles))
	}

	if config.WriteFiles[0].Path != "static/data/areas/1.json" {
		t.Errorf("config.WriteFiles[0].Path should be 'static/data/areas/1.json', but %v", config.WriteFiles[0].Path)
	}

	if config.WriteFiles[0].Interval != 5 {
		t.Errorf("config.WriteFiles[0].Interval should be 5, but %v", config.WriteFiles[0].Interval)
	}

	if config.WriteFiles[0].Data[0] != "{'1': '1'}" {
		t.Errorf("config.WriteFiles[0].Data[0] should be {'1': '1'}, but %v", config.WriteFiles[0].Data[0])
	}

	if config.WriteFiles[0].Data[1] != "{'2': '2'}" {
		t.Errorf("config.WriteFiles[0].Data[1] should be {'2': '2'}, but %v", config.WriteFiles[0].Data[1])
	}
}
