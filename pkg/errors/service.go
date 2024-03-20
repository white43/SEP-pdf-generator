package errors

import (
	"github.com/gofiber/fiber/v2"
	"github.com/white43/SEP401-pdf-generator/pkg/dto"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s Service) Send(ctx *fiber.Ctx, err error) error {
	switch t := err.(type) {
	case *AppMessage:
		ctx.Response().SetStatusCode(t.GetCode())
		return ctx.JSON(dto.DefaultResponse{Code: t.GetCode(), Message: t.Error()})
	default:
		ctx.Response().SetStatusCode(500)
		return ctx.JSON(dto.DefaultResponse{Code: 500, Message: t.Error()})
	}
}
