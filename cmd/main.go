package main

import (
	"github.com/dvasyanin/http-rest-api/config"
	"github.com/dvasyanin/http-rest-api/controller/apiserver"
	"log"
)

func main() {
	cfg := config.NewConfig()
	if cfg == nil {
		log.Fatal()
	}

	if err := apiserver.Start(cfg); err != nil {
		log.Fatal(err)
	}
}
