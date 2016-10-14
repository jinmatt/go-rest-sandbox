package app

import (
	"go-rest-sandbox/app/controllers"

	"github.com/ant0ine/go-json-rest/rest"
)

func NewRouter() (*rest.App, error) {
	router, err := rest.MakeRouter(
		rest.Get("/", controllers.SayHello),
		rest.Get("/countries", controllers.GetCountries),
		rest.Get("/country/:code", controllers.GetCountry),
		rest.Post("/country", controllers.PostCountry),
	)

	return &router, err
}
