package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"go-task/config"
	"go-task/db/sqlc"
	"go-task/internal/handler"
	"go-task/internal/logger"
	"go-task/internal/repository"
	"go-task/internal/routes"
	"go-task/internal/service"
)

func main() {
	// Logger init
	if err := logger.Init(); err != nil {
		log.Fatal(err)
	}
	defer logger.Log.Sync()
	
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}

	// DB Connection
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// SQLC
	queries := sqlc.New(db)

	userRepo := repository.NewUserRepository(queries)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	app := fiber.New()

	routes.SetupRoutes(app, userHandler)

	logger.Log.Info("server starting")

	// PORT from .env
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	if err := app.Listen(":" + port); err != nil {
		logger.Log.Fatal(
			"failed to start server",
			zap.Error(err),
		)
	}
}