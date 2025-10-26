package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Anagh3/go-backend/internal/handlers"
	"github.com/Anagh3/go-backend/internal/logger"
	"github.com/Anagh3/go-backend/internal/repository"
	"github.com/Anagh3/go-backend/internal/routes"
	"github.com/Anagh3/go-backend/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Read DB credentials from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	if dbUser == "" || dbPass == "" || dbHost == "" || dbName == "" {
		log.Fatal("Database environment variables not set")
	}

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ")/" + dbName + "?parseTime=true"

	// DB connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Logger and validator
	logr := logger.NewLogger()
	validate := validator.New()

	// Layers
	repo := repository.NewUserRepository(db)
	svc := service.NewUserService(repo)
	handler := handlers.NewUserHandler(svc, validate, logr)

	// Fiber app
	app := fiber.New()
	routes.SetupUserRoutes(app, handler)

	logr.Info("🚀 Server running", zap.String("port", "3000"))
	if err := app.Listen(":3000"); err != nil {
		logr.Fatal("Failed to start server", zap.Error(err))
	}
}
