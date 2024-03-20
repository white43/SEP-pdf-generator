package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/white43/sep401/pkg/dto"
	apperr "github.com/white43/sep401/pkg/errors"
	"github.com/white43/sep401/pkg/users"
)

func PostUserRegister(response *apperr.Service, userService *users.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := dto.RegistrationRequest{}
		if err := json.Unmarshal(ctx.Body(), &request); err != nil {
			return response.Send(ctx, err)
		}

		if err := userService.ValidateRegistrationRequest(request); err != nil {
			return response.Send(ctx, err)
		}

		return response.Send(ctx, userService.Register(request))
	}
}
