package laatoocore

import (
	"laatoosdk/config"
)

//Environment hosting an application
type Application struct {
	//name of the application
	Name string
	//environments for the application
	Environments map[string]*Environment
	//url
	ParentURL string
	//config for the application
	Config config.Config
	//store containing the names of services
	Services []string
}

func NewApplication(name string) *Application {
	app := &Application{Name: name}
	return app
}
func (app *Application) AddService(name string) {

}
