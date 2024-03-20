package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/white43/SEP401-pdf-generator/pkg/dto"
	"github.com/white43/SEP401-pdf-generator/pkg/errors"
	"github.com/white43/SEP401-pdf-generator/pkg/jobs"
	"github.com/white43/SEP401-pdf-generator/pkg/users"
)

func PostAppJob(response *errors.Service, jobService *jobs.Service, repository *users.UserRepository, jobType string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := dto.NewJobRequest{}
		if err := json.Unmarshal(ctx.Body(), &request); err != nil {
			return response.Send(ctx, err)
		}

		user, err := repository.GetOneByToken(ctx.Get("authorization"))
		if err != nil {
			return err
		}

		if err := jobService.ValidateNewJobRequest(user, request); err != nil {
			return response.Send(ctx, err)
		}

		result, err := jobService.AddJob(jobType, user.ID, request)
		if err != nil {
			return response.Send(ctx, err)
		}

		return ctx.JSON(result)
	}
}
