package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
)

var httpRequests = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "todo_http_requests_total",
	Help: "Number of http requests.",
})

func init() {
	prometheus.Register(httpRequests)
}

func Index(w http.ResponseWriter, r *http.Request) {
	httpRequests.Inc()
	fmt.Fprintln(w, "Welcome!")
}

// Simulate a real response and mock out the TodoIndex with static data
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	httpRequests.Inc()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	httpRequests.Inc()
	vars := mux.Vars(r)
	var todoId int
	var err error
	if todoId, err = strconv.Atoi(vars["todoId"]); err != nil {
		panic(err)
	}
	todo := RepoFindTodo(todoId)
	if todo.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(err); err != nil {
		panic(err)
	}

}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	httpRequests.Inc()
	var todo Todo
	// open up the body of the request with a size limit
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	// Unmarshal it to our Todo struct
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity and send back the error in a json string
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	// send back the json representation of entity we created
	t := RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
