package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Product struct {
	ID      int    `json:"id"`
	Product string `json:"produto"`
}

type mHandler struct{}

func (m mHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world handler"))
}

func main() {

	r := chi.NewRouter()

	m := mHandler{}
	r.Handle("/hanlder", m)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		params := chi.URLParam(r, "id")
		w.Write([]byte("hello world " + params))
	})

	r.Get("/product", func(w http.ResponseWriter, r *http.Request) { // /product?id=2311231
		params := r.URL.Query().Get("id")
		w.Write([]byte("hello world " + params))
	})

	r.Get("/productTotal", func(w http.ResponseWriter, r *http.Request) { // /product/2311231
		w.Header().Set("Content-Type", "application/json")
		params := map[string]string{"message": "sucess1"}
		b, _ := json.Marshal(params)
		w.Write(b)
	})

	r.Get("/productTotal2", func(w http.ResponseWriter, r *http.Request) { // /product/2311231

		obj := map[string]string{"message": "sucess2"}
		render.JSON(w, r, obj)
	})

	r.Post("/produto", func(w http.ResponseWriter, r *http.Request) {
		var product Product
		render.DecodeJSON(r.Body, &product)
		render.JSON(w, r, product)
	})

	http.ListenAndServe(":8080", r)

}
