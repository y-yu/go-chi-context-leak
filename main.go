package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const timeout = 30 * time.Second

const (
	nameParam    = "name"
	privateParam = "private"
)

var database = make(map[string]string)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(func(h http.Handler) http.Handler {
		return http.TimeoutHandler(h, timeout, "Timeout!")
	})

	r.Get(fmt.Sprintf("/{%s}/{%s}/update", nameParam, privateParam), update)

	r.Get(fmt.Sprintf("/{%s}/show", nameParam), show)

	fmt.Println("Listening on localhost:3000")
	http.ListenAndServe("localhost:3000", r)
}

func update(w http.ResponseWriter, r *http.Request) {
	routeCtx := chi.RouteContext(r.Context())
	name := routeCtx.URLParam(nameParam)

	// very slow logic
	time.Sleep(3 * time.Second)

	private := routeCtx.URLParam(privateParam)

	fmt.Printf("name %s,  private %s\n", name, private)

	database[name] = private

	w.Write([]byte(fmt.Sprintf("%s private was updated!", name)))
}

func show(w http.ResponseWriter, r *http.Request) {
	routeCtx := chi.RouteContext(r.Context())

	name := routeCtx.URLParam(nameParam)

	w.Write([]byte(fmt.Sprintf("%s private is %s", name, database[name])))
}
