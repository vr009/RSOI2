package repo

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	models2 "lib/services/models"
	"time"
)

const (
	SelectReservation = "SELECT r.reservation_uid, r.status, r.start_date, r.till_date, " +
		"r.book_uid, b.name, b.author, b.genre " +
		"r.library_uid, l.name, l.city, l.address " +
		"FROM reservation r " +
		"INNER JOIN books b ON r.book_uid=b.book_uid " +
		"INNER JOIN library l ON l.library_uid=r.library_uid " +
		"WHERE r.username=$1"

	InsertReservation = "WITH inserted AS (INSERT INTO reservation " +
		"reservation_uid, username, book_uid, library_uid, status, start_date, till_date " +
		"VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING * ) " +
		"SELECT b.name, b.author, b.genre, l.name, l.city, l.address, r.stars " +
		"FROM inserted ins INNER JOIN books b ON b.book_uid=ins.book_uid " +
		"INNER JOIN library l ON l.library_uid=ins.library_uid " +
		"INNER JOIN rating r ON r.username=ins.username"

	DeleteReservation = "DELETE FROM reservation WHERE reservation_uid=$1"
)

type Repo struct {
	conn *pgxpool.Pool
}

func NewRepo(conn *pgxpool.Pool) *Repo {
	return &Repo{conn: conn}
}

func (r *Repo) ReservationsInfo(name string) ([]models2.BookReservationResponse, models2.StatusCode) {
	var fullInfo []models2.BookReservationResponse
	res, err := r.conn.Query(context.Background(), SelectReservation, name)
	if err != nil {
		return []models2.BookReservationResponse{}, models2.InternalError
	}
	for res.Next() {
		reservation := models2.BookReservationResponse{}
		res.Scan(&reservation.ReservationUid, &reservation.Status, &reservation.StartDate, &reservation.TillDate,
			&reservation.Book.BookUid, &reservation.Book.Name, &reservation.Book.Author, &reservation.Book.Genre,
			&reservation.Lib.LibraryUid, &reservation.Lib.Name, &reservation.Lib.City, &reservation.Lib.Address)
		fullInfo = append(fullInfo, reservation)
	}
	return fullInfo, models2.OK
}
func (r *Repo) ReserveBook(name string, req models2.TakeBookRequest) (models2.TakeBookResponse, models2.StatusCode) {
	resUid := uuid.New()
	start_date := time.Now()
	res := r.conn.QueryRow(context.Background(), InsertReservation, resUid, name,
		req.BookUid, req.LibraryUid, models2.Rented, start_date, req.TillDate)

	reservation := models2.TakeBookResponse{
		ReservationUid: resUid,
		TillDate:       req.TillDate,
		StartDate:      start_date,
	}
	reservation.Book.BookUid = req.BookUid
	reservation.Library.LibraryUid = req.LibraryUid
	reservation.Status = models2.Rented

	err := res.Scan(&reservation.Book.Name, &reservation.Book.Author, &reservation.Book.Genre,
		&reservation.Library.Name, &reservation.Library.City, &reservation.Library.Address, &reservation.Rating)

	if err != nil {
		return models2.TakeBookResponse{}, models2.BadRequest
	}
	return reservation, models2.OK
}
func (r *Repo) ReturnBook(resUid uuid.UUID, name string, req models2.ReturnBookRequest) models2.StatusCode {
	_, err := r.conn.Exec(context.Background(), DeleteReservation, name)
	if err != nil {
		return models2.NotFound
	}
	return models2.OK
}
