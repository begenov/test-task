package http

import (
	v1handler "github.com/begenov/test-task/internal/delivery/http/v1"
	"github.com/begenov/test-task/internal/service"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Init(app *fiber.App) {
	router := app.Group("/api")
	h.init(router)
}

func (h *Handler) init(router fiber.Router) {
	v1 := v1handler.NewHandler(h.service)
	v1.Init(router)
}
