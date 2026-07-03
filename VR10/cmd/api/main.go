package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Product struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type myHandler struct{}

func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("myHandler"))
}

func main() {
	r := chi.NewRouter()

	// All middlewares must come FIRST
	r.Use(myMiddleware)

	// Now you can register routes
	m := myHandler{}
	r.Handle("/handler", m)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("product")
		id := r.URL.Query().Get("id")

		if param != "" || id != "" {
			w.Write([]byte(param + " " + id))
			return // optional: avoid writing "Hello world" after query params
		}
		w.Write([]byte("Hello world"))
	})

	r.Get("/{productName}", func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "productName")
		w.Write([]byte("Hello " + param))
	})

	r.Get("/json", func(w http.ResponseWriter, r *http.Request) {
		obj := map[string]string{"product": "teste", "id": "1"}
		render.JSON(w, r, obj)
	})

	r.Post("/product", func(w http.ResponseWriter, r *http.Request) {
		var product Product
		if err := render.DecodeJSON(r.Body, &product); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		render.JSON(w, r, product)
	})

	http.ListenAndServe(":8080", r)
}

func myMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("myMiddleware before")
		next.ServeHTTP(w, r)
		println("myMiddleware after")
	})
}
