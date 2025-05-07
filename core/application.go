package core

type ServiceProvider interface {
	Register(Container)
	Boot(Container)
}

type Application struct {
	container Container
	providers []ServiceProvider
	booted    bool
}

func NewApplication() *Application {
	app := &Application{
		container: NewContainer(),
	}
	app.container.Bind("app", app)
	return app
}

func (app *Application) Register(provider ServiceProvider) {
	provider.Register(app.container)
	app.providers = append(app.providers, provider)

	if app.booted {
		provider.Boot(app.container)
	}
}

func (app *Application) Boot() {
	if app.booted {
		return
	}

	for _, provider := range app.providers {
		provider.Boot(app.container)
	}

	app.booted = true
}

func (app *Application) GetContainer() Container {
	return app.container
}
