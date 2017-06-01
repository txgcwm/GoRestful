package main



import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net"
	"net/http"
	"sync"
)

func main() {
	users := Users{
		Store: map[string]*User{},
	}

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/lookup/#host", LookupHostIp),

		rest.Get("/countries", GetAllCountries),
		rest.Post("/countries", PostCountry),
		rest.Get("/countries/:code", GetCountry),
		rest.Delete("/countries/:code", DeleteCountry),

		rest.Get("/users", users.GetAllUsers),
		rest.Post("/users", users.PostUser),
		rest.Get("/users/:id", users.GetUser),
		rest.Put("/users/:id", users.PutUser),
		rest.Delete("/users/:id", users.DeleteUser),
	)
	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

// curl -i http://127.0.0.1:8080/lookup/google.com
// curl -i http://127.0.0.1:8080/lookup/notadomain
func LookupHostIp(w rest.ResponseWriter, req *rest.Request) {
	ip, err := net.LookupIP(req.PathParam("host"))
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteJson(&ip)
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


// https://github.com/txgcwm/go-json-rest