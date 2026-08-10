package endpoints

import (
	"emailn/internal/domain/campaign"
	"net/http"
)

type Handler struct {
	CampaignService campaign.Service
}

func (h Handler) HandlerError(post func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)) http.HandlerFunc {
	panic("unimplemented")
}
