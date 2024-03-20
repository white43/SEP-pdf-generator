package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/white43/sep401/pkg/dto"
	apperr "github.com/white43/sep401/pkg/errors"
	"github.com/white43/sep401/pkg/users"
)

func PostTopup(response *apperr.Service, userService *users.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := dto.TopupRequest{}
		if err := json.Unmarshal(ctx.Body(), &request); err != nil {
			return response.Send(ctx, err)
		}

		if err := userService.ValidateTopupRequest(request); err != nil {
			return response.Send(ctx, err)
		}

		user, err := userService.GetUserByToken(ctx.Get("authorization"))
		if err != nil {
			return err
		}

		err = userService.AddBalance(user.ID, request)
		if err != nil {
			return response.Send(ctx, err)
		}

		return ctx.JSON(dto.DefaultResponse{Code: 200, Message: "your balance has been topped up"})
	}
}
