package main

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"github.com/GoWebProd/uuid7"
	"github.com/white43/sep401/pkg/database"
	"github.com/white43/sep401/pkg/generator"
	"github.com/white43/sep401/pkg/jobs"
	"github.com/white43/sep401/pkg/mail"
	"github.com/white43/sep401/pkg/users"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	db := database.NewDatabase(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	defer db.Close()

	mailService := mail.NewMail("", "")

	jobRepository := jobs.NewJobRepository(db)
	userRepository := users.NewUserRepository(db)
	uuidService := uuid7.New()
	jobService := jobs.NewService(jobRepository, uuidService)
	userService := users.NewService(userRepository, mailService)
	processorFactory := generator.NewFactory(os.Getenv("CHROME_HEADLESS_URL"))
	generatorService := generator.NewService(jobService, userService, processorFactory)

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

Loop:
	for {
		select {
		case <-quit:
			log.Println("Daemon has been stopped gracefully")
			break Loop
		case <-ticker.C:
			job, err := generatorService.GetNextJob()
			if err == sql.ErrNoRows {
				break
			} else if err != nil {
				fmt.Printf("GetNextJob: %#v\n", err)
				break
			}

			result, err := generatorService.Process(job)
			if err == nil {
				err = generatorService.MarkJobSuccessful(job, base64.StdEncoding.EncodeToString(result))
			}
			if err == nil {
				err = generatorService.UpdateUserBalance(job, 1)
			}
			if err != nil {
				fmt.Printf("Process %#v\n", err)
				err = generatorService.MarkJobFailed(job)
			}
		}
	}
}
