package usecase

import (
	"github.com/google/uuid"
	"lib/services/gateway/internal"
	models2 "lib/services/models"
)

type GatewayUsecase struct {
	client internal.ApiClient
}

func NewGatewayUsecase(client internal.ApiClient) *GatewayUsecase {
	return &GatewayUsecase{
		client: client,
	}
}

func (u *GatewayUsecase) GetLibList(page, size int64, city string) ([]models2.LibraryPaginationResponse, models2.StatusCode) {
	return u.client.GetLibraries(page, size, city)
}
func (u *GatewayUsecase) GetBookList(page, size int64, showAll bool, LibUid uuid.UUID) ([]models2.LibraryBookPaginationResponse, models2.StatusCode) {
	return u.client.GetBooks(page, size, showAll, LibUid)
}
func (u *GatewayUsecase) GetReservationInfo() ([]models2.BookReservationResponse, models2.StatusCode) {
	return nil, 0
}
func (u *GatewayUsecase) TakeBook() (models2.TakeBookResponse, models2.StatusCode) {
	return models2.TakeBookResponse{}, 0
}
func (u *GatewayUsecase) ReturnBook() models2.StatusCode {
	return 0
}
func (u *GatewayUsecase) GetRating() (models2.UserRatingResponse, models2.StatusCode) {
	return models2.UserRatingResponse{}, 0
}
