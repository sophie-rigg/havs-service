package main

import (
	"flag"
	"github.com/sophie-rigg/havs-service/server"
	"github.com/sophie-rigg/havs-service/storage/mongo"
	"log"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "port", 8080, "The port to listen on")
}

func main() {
	flag.Parse()

	storage := mongo.NewClient()

	s := server.NewClient(storage)

	if err := s.Listen(port); err != nil {
		log.Fatal(err)
	}
}
