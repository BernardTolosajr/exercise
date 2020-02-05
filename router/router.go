package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
}

func Handler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello")
}

func New() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/hello", Handler())

	return r
}
