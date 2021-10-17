package internal

import models2 "lib/services/models"

type RatingUsecase interface {
	GetRating(name string) (models2.UserRatingResponse, models2.StatusCode)
}

type RatingRepo interface {
	FetchRating(name string) (models2.UserRatingResponse, models2.StatusCode)
}
