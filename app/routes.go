package app

// SetupRouters will register routes in router
func (app *App) setRouters() {

	// routes for the person(users).
	app.Post("/product", app.handleRequest(CreateProduct))
	// app.Patch("/product/{id}", app.handleRequest(UpdateProduct))
	// app.Put("/product/{id}", app.handleRequest(UpdateProduct))
	app.Get("/product/{id}", app.handleRequest(GetProduct))
	app.Get("/product", app.handleRequest(GetProducts))
}
