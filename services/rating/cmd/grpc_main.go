package main

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"lib/services/proto/rating "
	"lib/services/rating/internal/config"
	delivery2 "lib/services/rating/internal/delivery"
	repo2 "lib/services/rating/internal/repo"
	usecase2 "lib/services/rating/internal/usecase"
	"log"
	"net"
	"os"
)

const (
	grpcPort = "50051"
)

func main() {

	if err := runGRPC(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}

func runGRPC() error {
	conn, err := config.GetConnectionString()
	if err != nil {
		return err
	}

	pool, err := pgxpool.Connect(context.Background(), conn)
	if err != nil {
		return err
	}
	repo := repo2.NewRepo(pool)
	usecase := usecase2.NewRatingUsecase(repo)

	grpcServer := grpc.NewServer()
	ratingService := delivery2.NewGRPCHandler(usecase)
	rating.RegisterRatingServiceServer(grpcServer, ratingService)

	lis, err := net.Listen("tcp", ":" + grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
		return err
	}
	return nil
}