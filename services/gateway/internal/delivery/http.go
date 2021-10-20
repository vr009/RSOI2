package delivery

import (
	"lib/services/gateway/internal/usecase"
	"net/http"
)

type GatewayHandler struct {
	usecase *usecase.GatewayUsecase
}

func NewGatewayHandler(usecase *usecase.GatewayUsecase) *GatewayHandler {
	return &GatewayHandler{usecase: usecase}
}

func (u *GatewayHandler) GetLibraries(w http.ResponseWriter, r *http.Request) {

}

func (u *GatewayHandler) GetBooks(w http.ResponseWriter, r *http.Request) {

}

func (u *GatewayHandler) GetReservations(w http.ResponseWriter, r *http.Request) {

}

func (u *GatewayHandler) GetBook(w http.ResponseWriter, r *http.Request) {

}

func (u *GatewayHandler) ReturnBook(w http.ResponseWriter, r *http.Request)  {

}

func (u *GatewayHandler) GetRating(w http.ResponseWriter, r *http.Request) {

}