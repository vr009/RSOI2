package cmd

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"library/internal/config"
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
	pool = pool

	http.Handle("/", r)
	log.Print("main running on: ", srv.Addr)
	return srv.ListenAndServe()
}
