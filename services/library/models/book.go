package models

import "github.com/gofrs/uuid"

type Book struct {
	Id        uint          `json:"id"`
	BookId    uuid.UUID     `json:"book_uid"`
	Name      string        `json:"name"`
	Author    string        `json:"author"`
	Genre     string        `json:"genre"`
	Condition BookCondition `json:"condition"`
}

type BookCondition string

const (
	Excellent BookCondition = "Excelent"
	Good      BookCondition = "Good"
	Bad       BookCondition = "Bad"
)
