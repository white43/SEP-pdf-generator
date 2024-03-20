package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/white43/sep401/api/handlers"
	"github.com/white43/sep401/pkg/errors"
	"github.com/white43/sep401/pkg/middleware"
	"github.com/white43/sep401/pkg/users"
)

func TopupRouter(app fiber.Router, e *errors.Service, us *users.Service, ur *users.UserRepository) {
	app.Post("/v1/user/topup", middleware.NewAuth(handlers.PostTopup(e, us), ur, e))
}
