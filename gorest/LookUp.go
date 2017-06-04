package main



import (
	"github.com/ant0ine/go-json-rest/rest"
	"net"
	"net/http"
)


func LookupHostIp(w rest.ResponseWriter, req *rest.Request) {
	ip, err := net.LookupIP(req.PathParam("host"))
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteJson(&ip)
}

// curl -i http://127.0.0.1:8080/lookup/google.com
// curl -i http://127.0.0.1:8080/lookup/notadomain