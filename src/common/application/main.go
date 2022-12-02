package application

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
)

type ApplicationOpt struct {
	Environtment string
	ConfigPath   string
	ConfigRoot   string
}

type Application struct {
	Name    string
	Context context.Context
	Enable  bool
	Opt     *ApplicationOpt
}

func (a Application) GetAppsConfig(v interface{}) {
	configRuntime := viper.New()
	configRuntime.AddConfigPath(a.Opt.ConfigPath)
	configRuntime.SetConfigName(fmt.Sprintf("config-%s-%s.yml", a.Name, a.Opt.Environtment))
	err := configRuntime.UnmarshalKey(a.Opt.ConfigRoot, &v)
	if err != nil {
		panic(fmt.Sprintf("APPLICATION: unable to decode into struct, %v", err))
	}
}

func (a Application) GetName() string {
	return a.Name
}

func (a Application) IsEnable() bool {
	return a.Enable
}

type IApplication interface {
	Init()
	TearDown(ctx context.Context)
	GetAppsConfig(v interface{})
	GetName() string
	IsEnable() bool
}

type Applications struct {
	applications []IApplication
}

func (a *Applications) RegisterApplication(app IApplication) {
	a.applications = append(a.applications, app)
}

func (a *Applications) Run() {
	for _, elm := range a.applications {
		if elm.IsEnable() {
			elm.Init()
			fmt.Printf("Application: %s%s %s- started \n", string("\033[32m"), elm.GetName(), string("\033[0m"))
		}
	}
}

func (a *Applications) Shutdown(ctx context.Context) {
	for _, elm := range a.applications {
		if elm.IsEnable() {
			elm.TearDown(ctx)
			fmt.Printf("Application: %s%s %s- stoped \n", string("\033[32m"), elm.GetName(), string("\033[0m"))

		}
	}
}
