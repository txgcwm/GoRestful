package main


import (
    "fmt"
    "log"
    "net/http"
    "github.com/issue9/mux"
)

func main() {
    router := mux.New(true, true, nil, nil)
    router.HandleFunc("/", Index)
    router.HandleFunc("/todos", TodoIndex)
    router.HandleFunc("/todos/{todoId}", TodoShow)

    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    // vars := mux.Vars(r)
    // todoId := vars["todoId"]
    // fmt.Fprintln(w, "Todo show:", todoId)
}


// https://git.oschina.net/caixw/mux
// https://my.oschina.net/zijingshanke/blog/907955