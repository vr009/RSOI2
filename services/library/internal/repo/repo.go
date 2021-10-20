package repo

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"lib/services/models"
)

const (
	// TODO count
	GetBooksQuery = "SELECT b.id, b.book_uid, b.name, b.author, b.genre, b.condition, lb.available_count, COUNT(*) FROM books b " +
		"INNER JOIN library_books lb ON lb.book_id=b.id " + "INNER JOIN library l ON l.id=lb.library_id " +
		"WHERE l.library_uid=$1 AND lb.available_count>$2 " +
		"LIMIT $3 OFFSET $4"
	GetLibsQuery = "SELECT id, library_uid, name, city, address, COUNT(*) FROM library WHERE city=$1 LIMIT $2 OFFSET $3"
)

type LibRepo struct {
	conn *pgxpool.Pool
}

func NewLibRepo(conn *pgxpool.Pool) *LibRepo {
	return &LibRepo{conn: conn}
}

func (lr *LibRepo) GetLibraries(page, size int64, city string) ([]models.LibraryResponse, int64, models.StatusCode) {
	rows, err := lr.conn.Query(context.Background(), GetLibsQuery, city, page, size)
	if err != nil {
		return nil, 0, models.InternalError
	}
	var (
		libs  []models.LibraryResponse
		count int64
	)
	for rows.Next() {
		var lib models.LibraryResponse
		err := rows.Scan(&lib.Id, &lib.LibraryUid, &lib.Name, &lib.City, &lib.Address, &count)
		if err != nil {
			return nil, 0, models.InternalError
		}
		libs = append(libs, lib)
	}
	return libs, count, models.OK
}

func (lr *LibRepo) GetBooks(page, size int64, showAll bool, LibUid uuid.UUID) ([]models.LibraryBookResponse, int64, models.StatusCode) {
	includeZeroCountOfBooks := 0
	if showAll {
		includeZeroCountOfBooks = -1
	}

	rows, err := lr.conn.Query(context.Background(), GetBooksQuery, LibUid, includeZeroCountOfBooks, page, size)
	if err != nil {
		return nil, 0, models.InternalError
	}
	var (
		books []models.LibraryBookResponse
		count int64
	)
	for rows.Next() {
		var book models.LibraryBookResponse
		err := rows.Scan(&book.Id, &book.BookId, &book.Name, &book.Author,
			&book.Genre, &book.Condition, &book.AvailableCount, &count)
		if err != nil {
			return nil, 0, models.InternalError
		}
		books = append(books, book)
	}
	return books, count, models.OK
}
