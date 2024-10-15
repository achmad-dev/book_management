package main

import (
	"fmt"
	"net"

	grpcserver "github.com/achmad-dev/internal/author/cmd/grpc_server"
	"github.com/achmad-dev/internal/author/config"
	"github.com/achmad-dev/internal/author/internal/repository"
	"github.com/achmad-dev/internal/author/internal/service"
	pb "github.com/achmad-dev/internal/pkg/common/genproto"
	"github.com/achmad-dev/internal/pkg/logger"
	pkgUtil "github.com/achmad-dev/internal/pkg/util"
	"google.golang.org/grpc"
)

func main() {
	log := logger.InitLog()
	envPath := "./.env"
	cfg, err := config.NewConfig(envPath)
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// service implementation
	postgreUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)
	log.Infof("connecting to database: %s", postgreUrl)
	db, err := pkgUtil.InitSqlDB(postgreUrl)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	authorRepo := repository.NewAuthorRepository(db)
	authorService := service.NewAuthorService(authorRepo, log)

	serverImpl := grpcserver.NewServer(authorService)

	s := grpc.NewServer()
	pb.RegisterAuthorServiceServer(s, serverImpl)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
