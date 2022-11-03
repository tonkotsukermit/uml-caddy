package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

)

const (
	purl = "http://plant-uml:8080"
	port = ":8080"
)

func main() {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("howdy"))
	})

	r.Route("/puml", func(r chi.Router) {

		r.Get("/k8s", generateK8sPUML)

	})

	r.Route("/png", func(r chi.Router){

		r.Get("/k8s", generateK8sPNG)
	})

	http.ListenAndServe(port, r)

}
