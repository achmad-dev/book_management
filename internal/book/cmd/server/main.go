package main

import (
	"fmt"
	"net"

	grpcserver "github.com/achmad-dev/internal/book/cmd/grpc_server"
	"github.com/achmad-dev/internal/book/config"
	"github.com/achmad-dev/internal/book/internal/repository"
	"github.com/achmad-dev/internal/book/internal/service"
	pb "github.com/achmad-dev/internal/pkg/common/genproto"
	"github.com/achmad-dev/internal/pkg/logger"
	"github.com/achmad-dev/internal/pkg/util"
	"google.golang.org/grpc"
)

func main() {
	log := logger.InitLog()
	envPath := "./.env"
	cfg, err := config.NewConfig(envPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	postgreUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)
	db, err := util.InitSqlDB(postgreUrl)

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	// Init service
	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo, log)
	server := grpcserver.NewServer(bookService)
	s := grpc.NewServer()
	pb.RegisterBookServiceServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
