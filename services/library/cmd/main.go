package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"library/internal/config"
	"library/internal/delivery"
	"library/internal/repo"
	"library/internal/usecase"
	"log"
	"net/http"
	"os"
)

func main2() {
	if err := run(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
func run() error {
	r := mux.NewRouter()
	srv := http.Server{Handler: r, Addr: fmt.Sprintf(":%s", "8100")}

	conn, err := config.GetConnectionString()
	if err != nil {
		return err
	}

	pool, err := pgxpool.Connect(context.Background(), conn)
	if err != nil {
		return err
	}
	repo := repo.NewLibRepo(pool)
	usecase := usecase.NewLibUsecase(repo)
	handler := delivery.NewHandler(usecase)
	api := r.PathPrefix("/api/v1").Subrouter()
	{
		api.HandleFunc("/libraries/{libraryUid}/books", handler.GetBookList)
		api.HandleFunc("/libraries", handler.GetLibraryList)
	}

	http.Handle("/", r)
	log.Print("main running on: ", srv.Addr)
	return srv.ListenAndServe()
}
