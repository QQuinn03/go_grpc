package main

import (
	"log"

	"github.com/QQuinn03/go_grpc/internal/db"
	service "github.com/QQuinn03/go_grpc/internal/rocket"
)

//responsible for initializing and starting grpc server
func Run() error {
	database, err := db.New()
	if err != nil {
		return err
	}

	err = database.Migrate()
	if err != nil {
		log.Println("failures to migrate databases")
		return err

	}
	_ = service.New(database)

	return nil
}
func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
