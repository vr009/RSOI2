package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"lib/services/library/middleware"
	"lib/services/reservation/internal/config"
	"lib/services/reservation/internal/delivery"
	repo "lib/services/reservation/internal/repo"
	usecase "lib/services/reservation/internal/usecase"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
func run() error {
	r := mux.NewRouter()
	srv := http.Server{Handler: r, Addr: fmt.Sprintf(":%s", "8000")}

	conn, err := config.GetConnectionString()
	if err != nil {
		return err
	}

	pool, err := pgxpool.Connect(context.Background(), conn)
	if err != nil {
		return err
	}
	repo := repo.NewRepo(pool)
	usecase := usecase.NewUsecase(repo)
	handler := delivery.NewResHandler(usecase)

	r.Use(middleware.CORSMiddleware)
	api := r.PathPrefix("/api/v1").Subrouter()
	{
		api.HandleFunc("/reservations", handler.ReserveBook).Methods("POST")
		api.HandleFunc("/reservations", handler.GetListReservations).Methods("GET")
		api.HandleFunc("/reservations/{reservationUid}/return", handler.GetListReservations).Methods("GET")
	}

	http.Handle("/", r)
	log.Print("main running on: ", srv.Addr)
	return srv.ListenAndServe()
}
