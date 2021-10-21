package repo

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v4/pgxpool"
	models2 "rating/models"
)

const (
	SelectRating = "SELECT stars FROM ratings.rating WHERE username=$1"
)

type Repo struct {
	conn *pgxpool.Pool
}

func NewRepo(conn *pgxpool.Pool) *Repo {
	return &Repo{conn: conn}
}

func (r *Repo) FetchRating(name string) (models2.UserRatingResponse, models2.StatusCode) {
	resp := models2.UserRatingResponse{}
	row := r.conn.QueryRow(context.Background(), SelectRating, name)
	err := row.Scan(&resp.Stars)
	if err == sql.ErrNoRows {
		return models2.UserRatingResponse{}, models2.NotFound
	}
	if err != nil {
		return models2.UserRatingResponse{}, models2.InternalError
	}
	return resp, models2.OK
}
