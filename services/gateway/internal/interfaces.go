package internal

import (
	models2 "gateway/models"
	"github.com/google/uuid"
)

type Usecase interface {
	GetLibList(page, size int64, city string) ([]models2.LibraryPaginationResponse, models2.StatusCode)
	GetBookList(page, size int64, showAll bool, LibUid uuid.UUID) ([]models2.LibraryBookPaginationResponse, models2.StatusCode)
	GetReservationInfo(name string) ([]models2.BookReservationResponse, models2.StatusCode)
	TakeBook(name string, req models2.TakeBookRequest) (models2.TakeBookResponse, models2.StatusCode)
	ReturnBook(resUid uuid.UUID, name string, req models2.ReturnBookRequest) models2.StatusCode
	GetRating(name string) (models2.UserRatingResponse, models2.StatusCode)
}

type ApiClient interface {
	GetLibraries(page, size int64, city string) ([]models2.LibraryPaginationResponse, models2.StatusCode)
	GetBooks(page, size int64, showAll bool, LibUid uuid.UUID) ([]models2.LibraryBookPaginationResponse, models2.StatusCode)
	GetReservations(name string) ([]models2.BookReservationResponse, models2.StatusCode)
	TakeBook(name string, req models2.TakeBookRequest) (models2.TakeBookResponse, models2.StatusCode)
	ReturnBook(resUid uuid.UUID, name string, req models2.ReturnBookRequest) models2.StatusCode
	GetRating(name string) (models2.UserRatingResponse, models2.StatusCode)
	GetBook(bookId uuid.UUID) (models2.BookInfo, models2.StatusCode)
	GetLibrary(libId uuid.UUID) (models2.LibraryResponse, models2.StatusCode)
}
