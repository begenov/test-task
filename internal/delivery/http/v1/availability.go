package v1

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initAvailabilityRouter(v1 fiber.Router) {
	v1.Get("/site/:url", h.getSiteAvailability)
	v1.Get("/min-availability", h.getMinAvailability)
	v1.Get("/max-availability", h.getMaxAvailability)
	v1.Get("/stats", h.getStats)
}

func (h *Handler) getSiteAvailability(c *fiber.Ctx) error {
	h.Mutex.Lock()
	h.stats["/site/:url"]++
	h.Unlock()

	url := c.Params("url")

	site, err := h.service.Availability.GetSite(url)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusNotFound).JSON(Response{err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(site)
}

func (h *Handler) getMinAvailability(c *fiber.Ctx) error {
	h.Mutex.Lock()
	h.stats["/min-availability"]++
	h.Unlock()

	siteName, err := h.service.Availability.GetSiteWithMinAvailability()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{err.Error()})
	}

	return c.JSON(Response{"Сайт с минимальной доступностью: " + siteName})
}

func (h *Handler) getMaxAvailability(c *fiber.Ctx) error {
	h.Mutex.Lock()
	h.stats["/max-availability"]++
	h.Unlock()

	siteName, err := h.service.Availability.GetSiteWithMaxAvailability()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{err.Error()})
	}

	return c.JSON(Response{"Сайт с максимальной доступностью: " + siteName})
}

func (h *Handler) getStats(c *fiber.Ctx) error {
	h.Lock()
	defer h.Unlock()

	stats := h.stats

	return c.JSON(Response{stats})
}
