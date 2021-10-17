package internal

import (
	"github.com/google/uuid"
	models2 "lib/services/models"
)

type LibraryUsecase interface {
	GetLibrariesList(page, size int, city string) ([]models2.LibraryPaginationResponse, models2.StatusCode)
	GetBooksList(page, size int, showAll bool, LibUid uuid.UUID) ([]models2.LibraryBookPaginationResponse, models2.StatusCode)
}

type LibraryRepo interface {
	GetLibraries(page, size int, city string) ([]models2.LibraryResponse, int, models2.StatusCode)
	GetBooks(page, size int, showAll bool, LibUid uuid.UUID) ([]models2.LibraryBookResponse, int, models2.StatusCode)
}
