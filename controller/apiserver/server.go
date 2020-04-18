package apiserver

import (
	"context"
	"github.com/dvasyanin/http-rest-api/service"
)

const (
	HeaderAuthorization = "Authorization"
	ParamsRespond       = "respond"
)

type server struct {
	ctx     context.Context
	service service.Service
}

func newServer(ctx context.Context, service service.Service) *server {
	return &server{
		ctx:     ctx,
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
