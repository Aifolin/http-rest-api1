package main

import (
	"context"
	"fmt"
	"github.com/dvasyanin/http-rest-api/config"
	"github.com/dvasyanin/http-rest-api/models"
	"github.com/dvasyanin/http-rest-api/repository/store/pgstore"
	"github.com/dvasyanin/http-rest-api/service"
	"github.com/jackc/pgx/v4"
	"log"
)

func main() {
	cfg := config.NewConfig()
	if cfg == nil {
		log.Fatal()
	}

	chat := &models.Chat{
		Message:   "Пидарас что ли?",
		RespondID: 1,
		ClientID:  1,
	}

	db, err := newDB(cfg.DatabaseURL())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())

	// create server
	store := pgstore.New(db)
	src := service.NewService(store)

	if err := src.Create(chat); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Success")
}

func newDB(databaseURL string) (*pgx.Conn, error) {
	db, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(context.Background()); err != nil {
		return nil, err
	}

	return db, nil
}
