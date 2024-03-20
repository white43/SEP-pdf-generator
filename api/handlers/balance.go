package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/white43/SEP401-pdf-generator/pkg/dto"
	apperr "github.com/white43/SEP401-pdf-generator/pkg/errors"
	"github.com/white43/SEP401-pdf-generator/pkg/users"
)

func GetBalance(response *apperr.Service, userService *users.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user, err := userService.GetUserByToken(ctx.Get("authorization"))
		if err != nil {
			return err
		}

		return ctx.JSON(dto.BalanceResponse{Balance: user.Balance})
	}
}
