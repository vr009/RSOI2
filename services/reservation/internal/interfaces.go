package internal

import (
	models2 "lib/services/models"
)

type ResUsecase interface {
	GetReservationsInfo() models2.StatusCode
	TakeBook() models2.StatusCode
	ReturnBook() models2.StatusCode
}

type ResRepo interface {
	ReservationsInfo() models2.StatusCode
	ReserveBook() models2.StatusCode
	ReturnBook() models2.StatusCode
}
