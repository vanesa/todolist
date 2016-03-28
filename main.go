package main // package main is required for a standalone executable.

import (
    "fmt" // Implements formatted I/O. 
    "html"
    "log"
    "net/http"

    "github.com/gorilla/mux" // mux router from the Gorilla Web Toolkit.
)

func main() {
	// Here we register three routes mapping URL paths to handlers.
    router := mux.NewRouter().StrictSlash(true) // StrictSlash defines the trailing slash behavior for new routes.
    router.HandleFunc("/", Index)

    log.Fatal(http.ListenAndServe(":8080", router))
}


func Index(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

}