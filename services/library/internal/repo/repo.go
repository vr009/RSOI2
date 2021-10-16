package repo

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"library/models"
)

const (
	GetBooksQuery = "WITH b AS (SELECT * FROM library_books WHERE library_id=$1 AND available_count>$2) " +
		"SELECT id, book_uid, name, author, genre, condition FROM books " +
		"INNER JOIN b ON b.book_id=id" +
		"LIMIT $3 OFFSET $4"
	GetLibsQuery = "SELECT id, library_uid, name, city, address FROM library WHERE city=$1 LIMIT $2 OFFSET $3"
)

type LibRepo struct {
	conn *pgxpool.Pool
}

func NewLibRepo(conn *pgxpool.Pool) *LibRepo {
	return &LibRepo{conn: conn}
}

func (lr *LibRepo) GetLibraries(page, size int, city string) ([]models.Library, int, models.StatusCode) {
	rows, err := lr.conn.Query(context.Background(), GetLibsQuery, city, page, size)
	if err != nil {
		return nil, 0, models.InternalError
	}
	var (
		libs  []models.Library
		count int
	)
	for rows.Next() {
		var lib models.Library
		err := rows.Scan(&lib.Id, &lib.LibraryUid, &lib.Name, &lib.City, &lib.Address, &count)
		if err != nil {
			return nil, 0, models.InternalError
		}
		libs = append(libs, lib)
	}
	return libs, count, models.OK
}

func (lr *LibRepo) GetBooks(page, size int, showAll bool, LibUid uuid.UUID) ([]models.Book, int, models.StatusCode) {
	includeZeroCountOfBooks := 0
	if showAll {
		includeZeroCountOfBooks = -1
	}

	rows, err := lr.conn.Query(context.Background(), GetBooksQuery, LibUid, includeZeroCountOfBooks, page, size)
	if err != nil {
		return nil, 0, models.InternalError
	}
	var (
		books []models.Book
		count int
	)
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.Id, &book.BookId, &book.Name, &book.Author, &book.Genre, &count)
		if err != nil {
			return nil, 0, models.InternalError
		}
		books = append(books, book)
	}
	return books, count, models.OK
}
