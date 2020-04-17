package apiserver

import "github.com/dvasyanin/http-rest-api/service"

type server struct {
	service service.Service
}

func newServer(service service.Service) *server {
	return &server{
		service: service,
	}
}

// error response...
func (s *server) error(er string) err {
	return err{
		Error: er,
	}
}

// errResponse...
type err struct {
	Error string `json:"error"`
}
