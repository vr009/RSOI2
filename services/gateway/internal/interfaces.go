package internal

import (
	"github.com/google/uuid"
	models2 "lib/services/models"
)

type Usecase interface {
	GetLibList(page, size int64, city string) ([]models2.LibraryPaginationResponse, models2.StatusCode)
	GetBookList(page, size int64, showAll bool, LibUid uuid.UUID) ([]models2.LibraryBookPaginationResponse, models2.StatusCode)
	GetReservationInfo() ([]models2.BookReservationResponse, models2.StatusCode)
	TakeBook() (models2.TakeBookResponse, models2.StatusCode)
	ReturnBook() models2.StatusCode
	GetRating() (models2.UserRatingResponse, models2.StatusCode)
}

type ApiClient interface {
	GetLibraries(page, size int64, city string) ([]models2.LibraryPaginationResponse, models2.StatusCode)
	GetBooks(page, size int64, showAll bool, LibUid uuid.UUID) ([]models2.LibraryBookPaginationResponse, models2.StatusCode)
	GetReservations() ([]models2.BookReservationResponse, models2.StatusCode)
	TakeBook() (models2.TakeBookResponse, models2.StatusCode)
	ReturnBook() models2.StatusCode
	GetRating() (models2.UserRatingResponse, models2.StatusCode)
}
