package apiserver

import (
	"github.com/dvasyanin/http-rest-api/models"
	"github.com/gofiber/fiber"
	"github.com/zhs/loggr"
	"net/http"
)

func (s *server) createEmployers() func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {
		log := loggr.WithContext(s.ctx).
			With("request_id", c.Get(HeaderRequestID),
				"method", c.Method(),
				"ip", c.IP(),
				"token", c.Get(HeaderAuthorization),
				"request", loggr.Marshal(c.Body()),
			)

		log.Infow("[request]")
		var body models.Employers
		if err := s.bodyParser(c, &body); err != nil {
			log.Errorw("[response]", "response", err)
			return
		}

		if err := s.service.CreateEmployer(&body); err != nil {
			s.errorParser(c, http.StatusInternalServerError, err)
			log.Errorw("[response]", "response", err)
			return
		}

		log.Infow("[response]", "response", "success")
	}
}

func (s *server) deleteEmployer() func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {
		log := loggr.WithContext(s.ctx).
			With("request_id", c.Get(HeaderRequestID),
				"method", c.Method(),
				"ip", c.IP(),
				"token", c.Get(HeaderAuthorization),
				"request", loggr.Marshal(c.Body()),
			)

		log.Infow("[request]")

		var body models.Employers
		if err := s.bodyParser(c, &body); err != nil {
			log.Errorw("[response]", "response", err)
			return
		}

		if err := s.service.DeleteEmployer(body.Email); err != nil {
			s.errorParser(c, http.StatusInternalServerError, err)
			log.Errorw("[response]", "response", err)
			return
		}

		log.Infow("[response]", "response", "success")
	}
}

func (s *server) employerFindByEmail() func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {
		log := loggr.WithContext(s.ctx).
			With("request_id", c.Get(HeaderRequestID),
				"method", c.Method(),
				"ip", c.IP(),
				"token", c.Get(HeaderAuthorization),
				"request", loggr.Marshal(c.Body()),
			)

		log.Infow("[request]")

		var body models.Employers
		if err := s.bodyParser(c, &body); err != nil {
			log.Errorw("[response]", "response", err)
			return
		}

		if err := s.service.EmployerFindByEmail(&body); err != nil {
			s.errorParser(c, http.StatusInternalServerError, err)
			log.Errorw("[response]", "response", err)
			return
		}

		if err := c.JSON(body); err != nil {
			c.Status(http.StatusInternalServerError)
			_ = c.JSON(s.error("bad respond"))
		}

		log.Infow("[response]", "response", loggr.Marshal(body))
	}
}

func (s *server) putEmployers() func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {
		log := loggr.WithContext(s.ctx).
			With("request_id", c.Get(HeaderRequestID),
				"method", c.Method(),
				"ip", c.IP(),
				"token", c.Get(HeaderAuthorization),
				"request", loggr.Marshal(c.Body()),
			)

		log.Infow("[request]")
		var body models.Employers
		if err := s.bodyParser(c, &body); err != nil {
			log.Errorw("[response]", "response", err)
			return
		}

		if err := s.service.UpsertEmployer(&body); err != nil {
			s.errorParser(c, http.StatusInternalServerError, err)
			log.Errorw("[response]", "response", err)
			return
		}

		log.Infow("[response]", "response", "success")
	}
}
