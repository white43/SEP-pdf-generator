package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/white43/sep401/api/handlers"
	"github.com/white43/sep401/pkg/errors"
	"github.com/white43/sep401/pkg/middleware"
	"github.com/white43/sep401/pkg/users"
)

func BalanceRouter(app fiber.Router, e *errors.Service, us *users.Service, ur *users.UserRepository) {
	app.Get("/v1/user/balance", middleware.NewAuth(handlers.GetBalance(e, us), ur, e))
}
