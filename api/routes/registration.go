package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/white43/SEP401-pdf-generator/api/handlers"
	"github.com/white43/SEP401-pdf-generator/pkg/errors"
	"github.com/white43/SEP401-pdf-generator/pkg/users"
)

func RegistrationRouter(app fiber.Router, e *errors.Service, u *users.Service) {
	app.Post("/v1/user/register", handlers.PostUserRegister(e, u))
}
