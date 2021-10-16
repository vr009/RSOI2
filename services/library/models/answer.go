package models

type PaginatedAnswer struct {
	Page          int
	PageSize      int
	TotalElements int
	Items         interface{}
}
