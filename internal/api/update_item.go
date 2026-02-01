package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"job4j.ru/go-lang-base/internal/tracker"
)

type UpdateItemRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (s *Server) UpdateItem(c *fiber.Ctx) error {
	var req UpdateItemRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if req.ID == "" {
		return fiber.NewError(fiber.StatusBadRequest, "empty id")
	}
	if req.Name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "empty name")
	}

	err := s.Repository.Update(c.Context(), tracker.Item{
		ID:   req.ID,
		Name: req.Name,
	})
	if err != nil {
		log.Errorw("UpdateItem error", err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}
