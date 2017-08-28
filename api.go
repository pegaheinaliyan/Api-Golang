package main

import (
    "fmt"
    "html"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {

    router := mux.NewRouter().StrictSlash(true)
    // router.HandleFunc("/", Index)
    router.HandleFunc("/hello", Hello)
    // log.Fatal(http.ListenAndServe(":8080", router))
    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "first, %q", html.EscapeString(r.URL.Path))

}

func Hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello my api, %q", html.EscapeString(r.URL.Path))

}