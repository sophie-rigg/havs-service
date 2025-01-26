package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/sophie-rigg/havs-service/server"
	"github.com/sophie-rigg/havs-service/storage/mongo"
)

var (
	port        int
	mongoDBHost string
	mongoUser   string
	mongoPass   string
	db          string // The name of the database to use in MongoDB (could be swapped for staging/production etc.)
)

const (
	mongoDBURI = "mongodb://%s:%s@%s"
)

func init() {
	flag.IntVar(&port, "port", 8080, "The port to listen on")
	flag.StringVar(&mongoDBHost, "mongo-url", "mongo:27017", "The URL for the MongoDB instance")
	flag.StringVar(&mongoUser, "mongo-user", "user", "The username for the MongoDB instance")
	flag.StringVar(&mongoPass, "mongo-pass", "password", "The password for the MongoDB instance")
	flag.StringVar(&db, "db", "havs", "The name of the database to use in MongoDB")
}

func main() {
	flag.Parse()
	log.SetFormatter(&log.JSONFormatter{})

	storage, err := mongo.NewClient(fmt.Sprintf(mongoDBURI, mongoUser, mongoPass, mongoDBHost), db)
	if err != nil {
		log.WithError(err).Fatal("failed to create storage client")
	}

	s := server.NewClient(storage)

	if err := s.Listen(port); err != nil {
		log.WithError(err).Fatal("failed to start server")
	}
}
