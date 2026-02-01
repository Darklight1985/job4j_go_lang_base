package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

type DeleteItemRequest struct {
	ID uuid.UUID `json:"id"`
}

func (s *Server) DeleteItem(c *fiber.Ctx) error {
	var req DeleteItemRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
	}

	if req.ID.String() == "" {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	err := s.Repository.Delete(c.Context(), req.ID)

	if err != nil {
		log.Errorw("Error deleting item", "id", req.ID, "err", err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}
