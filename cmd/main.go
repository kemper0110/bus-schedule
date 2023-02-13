package main

import (
	"log"

	"github.com/kemper0110/bus-schedule/internal/api"
	db "github.com/kemper0110/bus-schedule/internal/storage"
)

func main() {
	db, err := db.NewDB("postgres://postgres:1234567890@localhost:5432/bus-schedule?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err = api.Start(db); err != nil {
		log.Fatal(err)
	}
}
