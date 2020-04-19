package apiserver

import (
	"github.com/gofiber/fiber"
	"net/http"
)

func (s *server) bodyParser(c *fiber.Ctx, body interface{}) error {
	if err := c.BodyParser(&body); err != nil {
		c.Status(http.StatusBadRequest)
		_ = c.JSON(s.error(err.Error()))
		return err
	}

	return nil
}

func (s *server) errorParser(c *fiber.Ctx, status int, err error) {
	c.Status(status)
	_ = c.JSON(s.error(err.Error()))
	return
}
