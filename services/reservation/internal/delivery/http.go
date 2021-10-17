package delivery

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	models2 "lib/services/models"
	"lib/services/reservation/internal/usecase"
	"lib/services/reservation/internal/utils"
	"net/http"
)

type ResHandler struct {
	usecase *usecase.Usecase
}

func NewResHandler(usecase *usecase.Usecase) *ResHandler {
	return &ResHandler{usecase: usecase}
}

func (h *ResHandler) GetListReservations(w http.ResponseWriter, r *http.Request) {
	name := r.Header.Get("X-User-Name")
	list, st := h.usecase.GetReservationsInfo(name)
	if st != models2.OK {
		utils.Response(w, st, "", nil)
		return
	}
	body, err := json.Marshal(list)
	if err != nil {
		utils.Response(w, st, "", nil)
		return
	}
	utils.Response(w, st, "", body)
}

func (h *ResHandler) ReserveBook(w http.ResponseWriter, r *http.Request) {
	name := r.Header.Get("X-User-Name")
	req := models2.TakeBookRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.Response(w, models2.BadRequest, "", nil)
		return
	}

	validErr := utils.Validate(req)
	if validErr != nil {
		errBody, err := json.Marshal(*validErr)
		if err != nil {
			utils.Response(w, models2.InternalError, "", nil)
		}
		utils.Response(w, models2.BadRequest, "", errBody)
		return
	}

	list, st := h.usecase.TakeBook(name, req)
	if st != models2.OK {
		utils.Response(w, st, "", nil)
		return
	}
	body, err := json.Marshal(list)
	if err != nil {
		utils.Response(w, st, "", nil)
		return
	}
	utils.Response(w, st, "", body)
}

func (h *ResHandler) ReturnBook(w http.ResponseWriter, r *http.Request) {
	bookReq := models2.ReturnBookRequest{}
	err := json.NewDecoder(r.Body).Decode(&bookReq)

	name := r.Header.Get("X-User-Name")
	vars := mux.Vars(r)
	UidStr := vars["reservationUid"]
	Uid, err := uuid.Parse(UidStr)
	if err != nil {
		utils.Response(w, models2.InternalError, "", nil)
		return
	}
	status := h.usecase.ReturnBook(Uid, name, bookReq)
	if status == models2.NotFound {
		errBody, err := json.Marshal(models2.ErrorResponse{Message: "Not Found"})
		if err != nil {
			utils.Response(w, models2.InternalError, "", nil)
		}
		utils.Response(w, status, "", errBody)
		return
	}
	utils.Response(w, status, "", nil)
}
