package models

type LibraryBook struct {
	BookId         uint `json:"book_id"`
	LibraryId      uint `json:"library_id"`
	AvailableCount int  `json:"available_count"`
}
