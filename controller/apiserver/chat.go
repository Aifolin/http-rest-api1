package apiserver

import (
	"github.com/gofiber/fiber"
	"github.com/zhs/loggr"
	"net/http"
	"strconv"
)

func (s *server) getChatByRespond() func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {
		params := c.Query(ParamsRespond)
		log := loggr.WithContext(s.ctx).
			With("method", c.Method(),
				"ip", c.IP(),
				"token", c.Get(HeaderAuthorization),
				"request", params,
			)

		log.Infow("[request]")
		respondID, err := strconv.ParseInt(params, 10, 64)
		if err != nil {
			c.Status(http.StatusBadRequest)
			_ = c.JSON(s.error("respond id is not set"))
			log.Error(err)
			return
		}

		result, err := s.service.GetChatByRespond(respondID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			_ = c.JSON(s.error(err.Error()))
			log.Error(err)
			return
		}

		if err := c.JSON(result); err != nil {
			c.Status(http.StatusInternalServerError)
			_ = c.JSON(s.error(err.Error()))
			log.Error(err)
			return
		}

		log.Infow("[response]", "response", loggr.Marshal(result))
	}
}
