package models

type LibraryPaginationResponse struct {
	Page          int               `json:"page"`
	PageSize      int               `json:"page_size"`
	TotalElements int               `json:"total_elements"`
	Items         []LibraryResponse `json:"items"`
}

type LibraryBookPaginationResponse struct {
	Page          int                   `json:"page"`
	PageSize      int                   `json:"page_size"`
	TotalElements int                   `json:"total_elements"`
	Items         []LibraryBookResponse `json:"items"`
}
