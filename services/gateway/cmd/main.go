package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"lib/services/gateway/internal/config"
	"lib/services/gateway/internal/delivery"
	"lib/services/gateway/internal/usecase"
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
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	clients, err := config.GetConnectionString()
	if err != nil {
		return err
	}
	libConn, err := grpc.Dial(clients.LibraryURL, opts...)
	if err != nil {
		return err
	}
	defer libConn.Close()
	reservationConn, err := grpc.Dial(clients.ReservationURL)
	if err != nil {
		return err
	}
	defer reservationConn.Close()
	ratingConn, err := grpc.Dial(clients.RatingURL)
	if err != nil {
		return err
	}
	defer ratingConn.Close()
	apiClient := delivery.NewGRPCClient(ratingConn, reservationConn, libConn)
	usecase := usecase.NewGatewayUsecase(apiClient)
	gatewayHandler := delivery.NewGatewayHandler(usecase)

	r.Use(middleware.CORSMiddleware)
	api := r.PathPrefix("/api/v1").Subrouter()
	{
		api.HandleFunc("/libraries", gatewayHandler.GetLibraries).Methods("GET")
		api.HandleFunc("/libraries/{libraryUid}/books", gatewayHandler.GetLibraries).Methods("GET")
		api.HandleFunc("/reservations", gatewayHandler.GetReservations).Methods("GET")
		api.HandleFunc("/reservations", gatewayHandler.GetBook).Methods("POST")
		api.HandleFunc("/reservations/{reservationUid}/return", gatewayHandler.ReturnBook).Methods("POST")
		api.HandleFunc("/rating", gatewayHandler.GetRating).Methods("GET")
	}

	http.Handle("/", r)
	log.Print("main running on: ", srv.Addr)
	return srv.ListenAndServe()
}
