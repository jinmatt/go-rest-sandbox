package controllers

import (
	"go-rest-sandbox/app/types"
	"sync"

	"github.com/ant0ine/go-json-rest/rest"
)

var store = map[string]*types.Country{}

var lock = sync.RWMutex{}

func GetCountries(w rest.ResponseWriter, req *rest.Request) {
	lock.RLock()
	countries := make([]types.Country, len(store))
	i := 0
	for _, country := range store {
		countries[i] = *country
		i++
	}
	lock.RUnlock()

	w.WriteJson(&countries)
}

func GetCountry(w rest.ResponseWriter, req *rest.Request) {
	lock.RLock()
	code := req.PathParam("code")

	var country *types.Country
	if store[code] != nil {
		country = &types.Country{}
		*country = *store[code]
	}
	lock.RUnlock()

	if country == nil {
		rest.NotFound(w, req)
		return
	}

	w.WriteJson(country)
}

func PostCountry(w rest.ResponseWriter, req *rest.Request) {
	country := types.Country{}
	req.DecodeJsonPayload(&country)
	if country.Code == "" {
		rest.Error(w, "country code required", 400)
		return
	}
	if country.Name == "" {
		rest.Error(w, "country name required", 400)
		return
	}

	lock.Lock()
	store[country.Code] = &country
	lock.Unlock()
	w.WriteJson(&country)
}
