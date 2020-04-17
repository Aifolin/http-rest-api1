package apiserver

import (
	"github.com/gofiber/fiber"
	"net/http"
	"strconv"
)

func (s *server) getChatByRespond() func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {
		params := c.Query("respond")
		respondID, err := strconv.ParseInt(params, 10, 64)
		if err != nil {
			c.Status(http.StatusBadRequest)
			_ = c.JSON(s.error("respond id is not set: " + err.Error()))
		}

		result, err := s.service.GetChatByRespond(respondID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			_ = c.JSON(s.error(err.Error()))
		}

		if err := c.JSON(result); err != nil {
			c.Status(http.StatusInternalServerError)
			_ = c.JSON(s.error(err.Error()))
		}
	}
}
