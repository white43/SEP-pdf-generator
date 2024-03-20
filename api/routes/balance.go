package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/white43/SEP401-pdf-generator/api/handlers"
	"github.com/white43/SEP401-pdf-generator/pkg/errors"
	"github.com/white43/SEP401-pdf-generator/pkg/middleware"
	"github.com/white43/SEP401-pdf-generator/pkg/users"
)

func BalanceRouter(app fiber.Router, e *errors.Service, us *users.Service, ur *users.UserRepository) {
	app.Get("/v1/user/balance", middleware.NewAuth(handlers.GetBalance(e, us), ur, e))
}
