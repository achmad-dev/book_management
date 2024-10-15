package main

import (
	"fmt"
	"net"

	grpcserver "github.com/achmad-dev/internal/category/cmd/grpc_server"
	"github.com/achmad-dev/internal/category/config"
	"github.com/achmad-dev/internal/category/internal/repository"
	"github.com/achmad-dev/internal/category/internal/service"
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

	db, err := pkgUtil.InitSqlDB(postgreUrl)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo, log)

	serverImpl := grpcserver.NewServer(categoryService)

	s := grpc.NewServer()
	pb.RegisterCategoryServiceServer(s, serverImpl)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
