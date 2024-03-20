package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/white43/sep401/api/handlers"
	"github.com/white43/sep401/pkg/errors"
	"github.com/white43/sep401/pkg/jobs"
	"github.com/white43/sep401/pkg/middleware"
	"github.com/white43/sep401/pkg/users"
)

func ResultRouter(app fiber.Router, e *errors.Service, j *jobs.Service, u *users.UserRepository) {
	app.Get("/v1/app/result/:id", middleware.NewAuth(handlers.GetAppResult(e, j), u, e))
}
