package main

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	"emailn/internal/infrastructure/database"
	internalerrors "emailn/internal/internal-errors"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	service := campaign.Service{
		Repository: &database.CampaignRepository{},
	}

	r.Post("/campaigns", func(w http.ResponseWriter, r *http.Request) {
		var request contract.NewCampaign
		render.DecodeJSON(r.Body, &request)

		id, err := service.Create(request)
		println(id)

		if err != nil {

			if errors.Is(err, internalerrors.ErrInternal) {
				render.Status(r, 500)

			} else {
				render.Status(r, 400)

			}

			render.JSON(w, r, map[string]string{"error": err.Error()})
			return

		}

		render.Status(r, 202)
		render.JSON(w, r, map[string]string{"id": id})

	})

	http.ListenAndServe(":8080", r)

}
