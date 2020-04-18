package apiserver

import (
	"github.com/dvasyanin/http-rest-api/models"
	"github.com/gofiber/fiber"
	"github.com/zhs/loggr"
	"net/http"
	"strconv"
)

func (s *server) getChatByRespond() func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {
		params := c.Query(ParamsRespond)
		log := loggr.WithContext(s.ctx).
			With("request_id", c.Get(HeaderRequestID),
				"method", c.Method(),
				"ip", c.IP(),
				"token", c.Get(HeaderAuthorization),
				"request", params,
			)

		log.Infow("[request]")
		respondID, err := strconv.ParseInt(params, 10, 64)
		if err != nil {
			c.Status(http.StatusBadRequest)
			_ = c.JSON(s.error("respond id is not set"))
			log.Errorw("[response]", "response", err)
			return
		}

		result, err := s.service.GetChatByRespond(respondID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			_ = c.JSON(s.error(err.Error()))
			log.Errorw("[response]", "response", err)
			return
		}

		if err := c.JSON(result); err != nil {
			c.Status(http.StatusInternalServerError)
			_ = c.JSON(s.error(err.Error()))
			log.Errorw("[response]", "response", err)
			return
		}

		log.Infow("[response]", "response", loggr.Marshal(result))
	}
}

func (s *server) createMessage() func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {
		log := loggr.WithContext(s.ctx).
			With("request_id", c.Get(HeaderRequestID),
				"method", c.Method(),
				"ip", c.IP(),
				"token", c.Get(HeaderAuthorization),
				"request", loggr.Marshal(c.Body()),
			)

		log.Infow("[request]")

		var body models.Chat
		if err := c.BodyParser(&body); err != nil {
			c.Status(http.StatusBadRequest)
			_ = c.JSON(s.error(err.Error()))
			log.Errorw("[response]", "response", err)
			return
		}

		if err := s.service.Create(&body); err != nil {
			c.Status(http.StatusInternalServerError)
			_ = c.JSON(s.error(err.Error()))
			log.Errorw("[response]", "response", err)
			return
		}

		log.Infow("[response]", "response", "success")
	}
}
