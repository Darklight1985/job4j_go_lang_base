package api

import (
	"github.com/gofiber/fiber/v2"
	"job4j.ru/go-lang-base/internal/tracker"
)

type SearchItemsRequest struct {
	Name string `json:"name"`
}

type SearchItemsResponse struct {
	tracker.Item `json:"item"`
}

func (s *Server) SearchItems(c *fiber.Ctx) error {
	var req SearchItemsRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON body")
	}

	if req.Name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid name")
	}

	item, err := s.Repository.SearchByName(c.Context(), req.Name)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(SearchItemsResponse{Item: item})
}
