package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		name, err := os.Hostname()
		if err != nil {
			panic(err)
		}

		w.Write([]byte(name))
	})

	r.Handle("/metrics", promhttp.Handler())

	fmt.Println("Server listening on :3333")
	http.ListenAndServe(":3333", r)
}
