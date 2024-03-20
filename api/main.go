package main

import (
	"github.com/GoWebProd/uuid7"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/white43/SEP401-pdf-generator/api/routes"
	"github.com/white43/SEP401-pdf-generator/pkg/database"
	"github.com/white43/SEP401-pdf-generator/pkg/errors"
	"github.com/white43/SEP401-pdf-generator/pkg/jobs"
	"github.com/white43/SEP401-pdf-generator/pkg/mail"
	"github.com/white43/SEP401-pdf-generator/pkg/users"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := fiber.New()

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		MaxAge:       3600,
	}))

	errorService := errors.NewService()

	db := database.NewDatabase(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	defer db.Close()

	userRepository := users.NewUserRepository(db)
	jobRepository := jobs.NewJobRepository(db)

	uuidService := uuid7.New()

	mailService := mail.NewMail(os.Getenv("SMTP_HOST"), os.Getenv("SMTP_PORT"))
	userService := users.NewService(userRepository, mailService)
	jobService := jobs.NewService(jobRepository, uuidService)

	routes.RegistrationRouter(app, errorService, userService)
	routes.LoginRouter(app, errorService, userService)
	routes.HtmlRouter(app, errorService, jobService, userRepository)
	routes.ResultRouter(app, errorService, jobService, userRepository)
	routes.TopupRouter(app, errorService, userService, userRepository)
	routes.BalanceRouter(app, errorService, userService, userRepository)

	port := os.Getenv("PORT")

	if port == "" {
		port = "80"
	}

	go func() {
		if err := app.Listen(":" + port); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := app.Shutdown(); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server has been stopped gracefully")
}
