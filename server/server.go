package server

import (
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/sophie-rigg/havs-service/server/exposure"
	exposuresummary "github.com/sophie-rigg/havs-service/server/exposure_summary"
	"github.com/sophie-rigg/havs-service/storage"
	"net/http"
)

type client struct {
	storage storage.Client
}

func NewClient(storage storage.Client) *client {
	return &client{
		storage: storage,
	}
}

func (c *client) Listen(port int) error {
	r := mux.NewRouter()

	c.setupExposureHandlers(r)
	c.setupUserHandlers(r)

	log.Info("Starting server on port ", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), r); err != nil {
		return err
	}
	return nil
}

func (c *client) setupExposureHandlers(mux *mux.Router) {
	exposureSubRouter := mux.PathPrefix("/exposure").Subrouter()

	exposureSubRouter.Handle("", exposure.NewHandler(c.storage))
	exposureSubRouter.Handle("/{exposureId}", exposure.NewIDHandler(c.storage))
}

func (c *client) setupUserHandlers(mux *mux.Router) {
	userSubRouter := mux.PathPrefix("/users").Subrouter()
	userByIDSubRouter := userSubRouter.PathPrefix("/{userId}").Subrouter()

	userByIDSubRouter.Handle("/exposure-summary", exposuresummary.NewHandler(c.storage))
}
