package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/white43/sep401/pkg/errors"
	"github.com/white43/sep401/pkg/users"
)

func NewAuth(handler fiber.Handler, repository *users.UserRepository, apperr *errors.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Get("authorization")
		if token == "" {
			return apperr.Send(ctx, errors.NotAuthorized)
		}

		user, err := repository.GetOneByToken(token)
		if err != nil || user.ID == 0 {
			return apperr.Send(ctx, errors.NotAuthorized)
		}

		return handler(ctx)
	}
}
