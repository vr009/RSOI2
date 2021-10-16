package utils

import "library/models"

func Paginate(page, PageSize, count int, items []interface{}) []models.PaginatedAnswer {
	return []models.PaginatedAnswer{
		models.PaginatedAnswer{Page: page, PageSize: PageSize, Items: items},
	}
}
