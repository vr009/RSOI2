package cmd

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"lib/services/library/internal/config"
	"lib/services/library/internal/delivery"
	"lib/services/library/internal/repo"
	"lib/services/library/internal/usecase"
	"lib/services/library/middleware"
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
	repo := repo.NewLibRepo(pool)
	usecase := usecase.NewLibUsecase(repo)
	handler := delivery.NewHandler(usecase)

	r.Use(middleware.CORSMiddleware)
	api := r.PathPrefix("/api/v1").Subrouter()
	{
		api.HandleFunc("/libraries/{libraryUid}/books", handler.GetBookList)
		api.HandleFunc("/libraries", handler.GetLibraryList)
	}

	http.Handle("/", r)
	log.Print("main running on: ", srv.Addr)
	return srv.ListenAndServe()
}
