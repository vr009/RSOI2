package models

import "github.com/gofrs/uuid"

type Library struct {
	Id         uint      `json:"id"`
	LibraryUid uuid.UUID `json:"library_uid"`
	Name       string    `json:"name"`
	City       string    `json:"city"`
	Address    string    `json:"address"`
}
