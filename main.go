package main // package main is required for a standalone executable.

import (
    	"fmt" // Implements formatted I/O. 
    	"html"
    	"log"
    	"net/http"
	"encoding/json"

    	"github.com/gorilla/mux" // mux router from the Gorilla Web Toolkit.
)

func main() {
	// Here we register three routes mapping URL paths to handlers.
	router := mux.NewRouter().StrictSlash(true) // StrictSlash defines the trailing slash behavior for new routes.
    	router.HandleFunc("/", Index)
    	router.HandleFunc("/todos", TodoIndex)
    	router.HandleFunc("/todos/{todoId}", TodoShow)

    	log.Fatal(http.ListenAndServe(":8080", router))
}


func Index(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, %q", html.EscapeString(r.URL.Path))

}

// Simulate a real response and mock out the TodoIndex with static data
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo Show: ", todoId)
}