package main

import (
	"database/sql"
	"log"

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
	// DB connection
	db, err := sql.Open("mysql", "root:anagh112@tcp(127.0.0.1:3306)/userdb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Logger, validator
	logr := logger.NewLogger()
	validate := validator.New()

	// Layers
	repo := repository.NewUserRepository(db)
	service := service.NewUserService(repo)
	handler := handlers.NewUserHandler(service, validate, logr)

	// Fiber app
	app := fiber.New()
	routes.SetupUserRoutes(app, handler)

	logr.Info("🚀 Server running", zap.String("port", "3000"))
	app.Listen(":3000")
}
