package delivery

import (
	"encoding/json"
	models2 "lib/services/models"
	"lib/services/rating/internal"
	"lib/services/rating/internal/utils"
	"net/http"
)

type RatingHandler struct {
	usecase internal.RatingUsecase
}

func NewRatingHandler(usecase internal.RatingUsecase) *RatingHandler {
	return &RatingHandler{usecase: usecase}
}

func (h *RatingHandler) GetRating(w http.ResponseWriter, r *http.Request) {
	name := r.Header.Get("X-User-Name")
	resp, st := h.usecase.GetRating(name)
	if st != models2.OK {
		utils.Response(w, models2.InternalError, "", nil)
		return
	}
	body, err := json.Marshal(resp)
	if err != nil {
		utils.Response(w, models2.InternalError, "", nil)
		return
	}
	utils.Response(w, st, "", body)
}
