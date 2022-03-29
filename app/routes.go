package app

// SetupRouters will register routes in router
func (app *App) setRouters() {

	// routes for the product.
	app.Post("/product", app.handleRequest(CreateProduct))
	// app.Patch("/product/{id}", app.handleRequest(UpdateProduct))
	// app.Put("/product/{id}", app.handleRequest(UpdateProduct))
	app.Get("/product/{id}", app.handleRequest(GetProduct))
	app.Get("/product", app.handleRequest(GetProducts))

	// routes for the creator.
	app.Post("/creator", app.handleRequest(CreateCreator))
	app.Get("/creator/{id}", app.handleRequest(GetCreator))
	app.Get("/creator", app.handleRequest(GetCreators))

	// routes for the brand.
	app.Post("/brand", app.handleRequest(CreateBrand))
	app.Get("/brand/{id}", app.handleRequest(GetBrand))
	app.Get("/brand", app.handleRequest(GetBrands))

	// routes for the brandcategory.
	app.Post("/brandcategory", app.handleRequest(CreateBrandCategory))
	app.Get("/brandcategory/{id}", app.handleRequest(GetBrandcategory))
	app.Get("/brandcategory", app.handleRequest(GetBrandCategories))
}
