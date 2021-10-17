package models

import "github.com/google/uuid"

type LibraryBookResponse struct {
	Id             uint          `json:"id"`
	BookId         uuid.UUID     `json:"book_uid"`
	Name           string        `json:"name"`
	Author         string        `json:"author"`
	Genre          string        `json:"genre"`
	Condition      BookCondition `json:"condition"`
	AvailableCount int           `json:"available_count"`
}

type BookCondition string

const (
	Excellent BookCondition = "Excelent"
	Good      BookCondition = "Good"
	Bad       BookCondition = "Bad"
)
