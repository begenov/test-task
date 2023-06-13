package v1

import (
	"sync"

	"github.com/begenov/test-task/internal/service"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *service.Service
	stats   map[string]int
	sync.Mutex
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
		stats:   make(map[string]int),
	}
}

func (h *Handler) Init(router fiber.Router) {
	v1 := router.Group("/v1")
	h.initAvailabilityRouter(v1)
}
