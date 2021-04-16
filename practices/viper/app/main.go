package app

type App struct {
	AppName string `mapstructure:"app-name"`
}

var Setting = &App{}
