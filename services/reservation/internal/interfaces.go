package internal

import (
	"github.com/google/uuid"
	models2 "lib/services/models"
)

type ResUsecase interface {
	GetReservationsInfo(name string) ([]models2.BookReservationResponse, models2.StatusCode)
	TakeBook(name string, req models2.TakeBookRequest) (models2.TakeBookResponse, models2.StatusCode) // TODO interface ret
	ReturnBook(resUid uuid.UUID, userName string, req models2.ReturnBookRequest) models2.StatusCode
}

type ResRepo interface {
	ReservationsInfo(name string) ([]models2.BookReservationResponse, models2.StatusCode)
	ReserveBook(name string, req models2.TakeBookRequest) (models2.TakeBookResponse, models2.StatusCode)
	ReturnBook(resUid uuid.UUID, userName string, req models2.ReturnBookRequest) models2.StatusCode
}
