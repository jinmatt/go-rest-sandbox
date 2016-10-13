package controllers

import "github.com/ant0ine/go-json-rest/rest"

func SayHello(w rest.ResponseWriter, req *rest.Request) {
	w.WriteJson(map[string]string{"Body": "Hello World!"})
}
