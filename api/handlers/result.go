package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/white43/sep401/pkg/errors"
	"github.com/white43/sep401/pkg/jobs"
)

func GetAppResult(response *errors.Service, jobService *jobs.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		if err := jobService.ValidateJobResultRequest(id); err != nil {
			return response.Send(ctx, err)
		}

		result, err := jobService.GetJobResult(id)
		if err != nil {
			return response.Send(ctx, err)
		}

		return ctx.JSON(result)
	}
}
