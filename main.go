package main

import (
	"log"
	"net/http"

	"goJsonRest/app"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := app.NewRouter()
	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(*router)
	log.Fatal(http.ListenAndServe(":8000", api.MakeHandler()))
}
