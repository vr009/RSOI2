package utils

import models2 "lib/services/models"

func Validate(req models2.TakeBookRequest) *models2.ValidationErrorResponse {
	return &models2.ValidationErrorResponse{}
}
