package environment

type Environment struct {
	ServeDir       string
	ConfigFilePath string
}

func NewEnvironment() *Environment {
	return &Environment{
		ConfigFilePath: "config.yml",
	}
}
