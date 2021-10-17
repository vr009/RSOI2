package usecase

import (
	models2 "lib/services/models"
	"lib/services/rating/internal"
)

type RatingUsecase struct {
	repo internal.RatingRepo
}

func NewRatingUsecase(repo internal.RatingRepo) *RatingUsecase {
	return &RatingUsecase{repo: repo}
}

func (ru *RatingUsecase) GetRating(name string) (models2.UserRatingResponse, models2.StatusCode) {
	return ru.repo.FetchRating(name)
}
