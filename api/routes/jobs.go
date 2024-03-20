package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/white43/SEP401-pdf-generator/api/handlers"
	"github.com/white43/SEP401-pdf-generator/pkg/errors"
	"github.com/white43/SEP401-pdf-generator/pkg/jobs"
	"github.com/white43/SEP401-pdf-generator/pkg/middleware"
	"github.com/white43/SEP401-pdf-generator/pkg/users"
)

func HtmlRouter(app fiber.Router, e *errors.Service, j *jobs.Service, u *users.UserRepository) {
	app.Post("/v1/app/html", middleware.NewAuth(handlers.PostAppJob(e, j, u, "html"), u, e))
	app.Post("/v1/app/url", middleware.NewAuth(handlers.PostAppJob(e, j, u, "url"), u, e))
}
