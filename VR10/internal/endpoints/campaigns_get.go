package endpoints

import (
	"net/http"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	campaigns, err := h.CampaignService.Repository.Get()

	return map[string]interface{}{"campaigns": campaigns}, 200, err
}
