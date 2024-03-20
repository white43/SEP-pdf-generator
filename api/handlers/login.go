package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/white43/SEP401-pdf-generator/pkg/dto"
	apperr "github.com/white43/SEP401-pdf-generator/pkg/errors"
	"github.com/white43/SEP401-pdf-generator/pkg/users"
)

func PostUserLogin(response *apperr.Service, userService *users.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := dto.LoginRequest{}
		if err := json.Unmarshal(ctx.Body(), &request); err != nil {
			return response.Send(ctx, err)
		}

		if err := userService.ValidateLoginRequest(request); err != nil {
			return response.Send(ctx, err)
		}

		token, err := userService.Login(request)
		if err != nil {
			return response.Send(ctx, err)
		}

		return ctx.JSON(dto.LoginResponse{
			AccessToken: token,
		})
	}
}
