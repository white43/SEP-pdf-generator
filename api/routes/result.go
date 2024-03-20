package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/white43/SEP401-pdf-generator/api/handlers"
	"github.com/white43/SEP401-pdf-generator/pkg/errors"
	"github.com/white43/SEP401-pdf-generator/pkg/jobs"
	"github.com/white43/SEP401-pdf-generator/pkg/middleware"
	"github.com/white43/SEP401-pdf-generator/pkg/users"
)

func ResultRouter(app fiber.Router, e *errors.Service, j *jobs.Service, u *users.UserRepository) {
	app.Get("/v1/app/result/:id", middleware.NewAuth(handlers.GetAppResult(e, j), u, e))
}
