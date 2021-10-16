package internal

import (
	"github.com/google/uuid"
	"library/models"
)

type LibraryUsecase interface {
	GetLibrariesList(page, size int, city string) ([]models.PaginatedAnswer, models.StatusCode)
	GetBooksList(page, size int, showAll bool, LibUid uuid.UUID) ([]models.PaginatedAnswer, models.StatusCode)
}

type LibraryRepo interface {
	GetLibraries(page, size int, city string) ([]models.Library, int, models.StatusCode)
	GetBooks(page, size int, showAll bool, LibUid uuid.UUID) ([]models.Book, int, models.StatusCode)
}
