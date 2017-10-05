package configuration

import (
	"io/ioutil"

	"github.com/tnoda78/posiserver/environment"
	yaml "gopkg.in/yaml.v2"
)

type Configuration struct {
	RootPath   string      `yaml:"root_path"`
	WriteFiles []WriteFile `yaml:"write_files"`
}

type WriteFile struct {
	Path     string   `yaml:"path"`
	Interval int      `yaml:"interval"`
	Data     []string `yaml:"data"`
}

func NewConfiguration() *Configuration {
	return &Configuration{}
}

func (config *Configuration) LoadConfigFile(env *environment.Environment) (string, error) {
	buffer, err := ioutil.ReadFile(env.ConfigFilePath)

	if err != nil {
		return "", err
	}

	err = yaml.Unmarshal(buffer, config)

	if err != nil {
		return "", err
	}

	d, err := yaml.Marshal(config)

	return string(d), nil
}
