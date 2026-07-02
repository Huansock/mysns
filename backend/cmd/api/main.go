package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main () {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func (w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello World"))
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

	})
	if err := http.ListenAndServe(":3000",r); err != nil {
		fmt.Printf("err: %v\n", err)
	}

}
