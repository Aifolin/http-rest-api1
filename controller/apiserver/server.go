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
		error: er,
	}
}

// err...
type err struct {
	error string
}
