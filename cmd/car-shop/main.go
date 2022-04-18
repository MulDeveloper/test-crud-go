package main

import (
	"os"
	"os/signal"

	"log"

	"github.com/MulDeveloper/go-test-crud/internal/data"
	"github.com/MulDeveloper/go-test-crud/internal/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	port := os.Getenv("PORT")

	serv, err := server.New(port)

	if err != nil {
		log.Fatal(err)
	}

	// db
	d := data.New()
	if err := d.DB.Ping(); err != nil {
		log.Fatal(err)
	}

	go serv.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	serv.Close()
	data.Close()
}
