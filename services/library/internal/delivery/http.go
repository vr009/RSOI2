package delivery

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"library/internal"
	"library/internal/utils"
	"library/models"
	"net/http"
	"strconv"
)

type Handler struct {
	usecase internal.LibraryUsecase
}

func (h *Handler) GetLibraryList(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		utils.Response(w, models.BadRequest, "", nil)
	}
	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil {
		utils.Response(w, models.BadRequest, "", nil)
	}
	city := r.URL.Query().Get("city")
	if city == "" {
		utils.Response(w, models.BadRequest, "", nil)
	}

	libs, status := h.usecase.GetLibrariesList(page, size, city)
	body, err := json.Marshal(libs)
	if err != nil {
		utils.Response(w, models.InternalError, "", nil)
	}
	utils.Response(w, status, "", body)
}

func (h *Handler) GetBookList(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		utils.Response(w, models.BadRequest, "", nil)
	}
	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil {
		utils.Response(w, models.BadRequest, "", nil)
	}

	vars := mux.Vars(r)
	uidStr := vars["libraryUid"]
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		utils.Response(w, models.BadRequest, "", nil)
	}

	var showAll bool
	showAllStr := r.URL.Query().Get("showAll")
	if showAllStr == "true" {
		showAll = true
	} else {
		showAll = false
	}

	libs, status := h.usecase.GetBooksList(page, size, showAll, uid)
	body, err := json.Marshal(libs)
	if err != nil {
		utils.Response(w, models.InternalError, "", nil)
	}
	utils.Response(w, status, "", body)
}
