package main



import (
	"github.com/ant0ine/go-json-rest/rest"
	"sync"
	"net/http"
)



type Country struct {
	Code string
	Name string
}

var store = map[string]*Country{}

var lock = sync.RWMutex{}

func GetCountry(w rest.ResponseWriter, r *rest.Request) {
	code := r.PathParam("code")

	lock.RLock()
	var country *Country
	if store[code] != nil {
		country = &Country{}
		*country = *store[code]
	}
	lock.RUnlock()

	if country == nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(country)
}

func GetAllCountries(w rest.ResponseWriter, r *rest.Request) {
	lock.RLock()
	countries := make([]Country, len(store))
	i := 0
	for _, country := range store {
		countries[i] = *country
		i++
	}
	lock.RUnlock()
	w.WriteJson(&countries)
}

func PostCountry(w rest.ResponseWriter, r *rest.Request) {
	country := Country{}
	err := r.DecodeJsonPayload(&country)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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

func DeleteCountry(w rest.ResponseWriter, r *rest.Request) {
	code := r.PathParam("code")
	lock.Lock()
	delete(store, code)
	lock.Unlock()
	w.WriteHeader(http.StatusOK)
}


// curl -i -H 'Content-Type: application/json' -d '{"Code":"FR","Name":"France"}' http://127.0.0.1:8080/countries
// curl -i -H 'Content-Type: application/json' -d '{"Code":"US","Name":"United States"}' http://127.0.0.1:8080/countries
// curl -i http://127.0.0.1:8080/countries/FR
// curl -i http://127.0.0.1:8080/countries/US
// curl -i http://127.0.0.1:8080/countries
// curl -i -X DELETE http://127.0.0.1:8080/countries/FR
// curl -i http://127.0.0.1:8080/countries
// curl -i -X DELETE http://127.0.0.1:8080/countries/US
// curl -i http://127.0.0.1:8080/countries