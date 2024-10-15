package v1

import (
	"fmt"
	"strconv"

	"github.com/achmad-dev/internal/pkg/logger"
	pkgUtil "github.com/achmad-dev/internal/pkg/util"
	grpcclient "github.com/achmad-dev/internal/user/api/grpc_client"
	"github.com/achmad-dev/internal/user/api/handler"
	"github.com/achmad-dev/internal/user/config"
	"github.com/achmad-dev/internal/user/internal/repository"
	"github.com/achmad-dev/internal/user/internal/service"
	"github.com/achmad-dev/internal/user/internal/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLog "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ServeRoutes(envPath string) {
	log := logger.InitLog()
	cfg, err := config.NewConfig(envPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Initialize the sqlx database
	postgreUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)
	db, err := pkgUtil.InitSqlDB(postgreUrl)
	if err != nil {
		log.Fatalf("failed to connect to database: %v url:%s", err, postgreUrl)
	}

	// Initialize the user service
	cost, err := strconv.Atoi(cfg.Cost)
	if err != nil {
		log.Fatalf("failed to convert cost to int: %v", err)
	}
	bcryptUtil := util.NewBcryptUtil(cost)
	// user service
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, bcryptUtil, log, cfg.JwtSecret)
	// user borrow book service
	userBrRepo := repository.NewUserBorrowedBookRepository(db)
	userBrService := service.NewUserBorrowedBookService(userBrRepo, log)

	// Initialize the user handler
	userHandler := handler.NewAuthHandler(userService)

	// Initialize the fiber app
	app := fiber.New(
		fiber.Config{
			StrictRouting: true,
			Prefork:       true,
			AppName:       "User Service",
		},
	)
	app.Use(
		cors.New(cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "Origin, Content-Type, Accept",
		}),
		fiberLog.New(
			fiberLog.Config{
				Format:     "${time} ${status} - ${latency} ${method} ${path}\n",
				TimeFormat: "02-Jan-2006",
			},
		),
	)

	// Define the routes
	api := app.Group("/api/v1")
	api.Post("/signup", userHandler.SignUp)
	api.Post("/signin", userHandler.SignIn)

	// Microservices
	// Author
	authorGrpcConn, err := grpc.NewClient(cfg.AuthorServiceHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to author service: %v", err)
	}
	grpcclient.RegisterAuthorRoutes(api, authorGrpcConn, cfg.JwtSecret, userService)

	// Category
	categoryGrpcConn, err := grpc.NewClient(cfg.CategoryServiceHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to category service: %v", err)
	}
	grpcclient.RegisterCategoryRoutes(api, categoryGrpcConn, cfg.JwtSecret, userService)

	// Book
	bookGrpcConn, err := grpc.NewClient(cfg.BookServiceHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to book service: %v", err)
	}

	redisAddr := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)
	redisClient := pkgUtil.InitRedisDB(redisAddr)

	grpcclient.RegisterBookRoutes(api, bookGrpcConn, authorGrpcConn, categoryGrpcConn, userBrService, cfg.JwtSecret, userService, redisClient)

	// Metrics
	app.Get("/metrics", monitor.New(
		monitor.Config{
			Title: "Book Management Metrics",
		},
	))
	log.Fatal(app.Listen(fmt.Sprintf(":%s", cfg.Port)))

}
