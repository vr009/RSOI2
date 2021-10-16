package usecase

import (
	"github.com/google/uuid"
	"library/internal"
	"library/internal/utils"
	"library/models"
)

type LibUsecase struct {
	repo internal.LibraryRepo
}

func NewLibUsecase(repo internal.LibraryRepo) *LibUsecase {
	return &LibUsecase{repo: repo}
}

func (lu *LibUsecase) GetLibrariesList(page, size int, city string) ([]models.PaginatedAnswer, models.StatusCode) {
	libs, count, status := lu.repo.GetLibraries(page, size, city)
	if status != models.OK {
		return nil, status
	}
	answer := utils.Paginate(page, size, count, []interface{}{libs})
	return answer, status
}

func (lu *LibUsecase) GetBooksList(page, size int, showAll bool, LibUid uuid.UUID) ([]models.PaginatedAnswer, models.StatusCode) {
	books, count, status := lu.repo.GetBooks(page, size, showAll, LibUid)
	if status != models.OK {
		return nil, status
	}
	answer := utils.Paginate(page, size, count, []interface{}{books})
	return answer, status
}
