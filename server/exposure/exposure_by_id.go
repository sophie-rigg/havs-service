package exposure

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/sophie-rigg/havs-service/storage"
	"net/http"
)

type IdHandler struct {
	storage storage.Client
}

func NewIDHandler(storage storage.Client) http.Handler {
	return &IdHandler{
		storage: storage,
	}
}

func (c *IdHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		c.Get(writer, request)
	default:
		http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (c *IdHandler) Get(writer http.ResponseWriter, request *http.Request) {
	requestVars := mux.Vars(request)

	exposureId, ok := requestVars["exposureId"]
	if !ok {
		http.Error(writer, "missing exposureId", http.StatusBadRequest)
		return
	}

	exposure, err := c.storage.GetExposure(exposureId)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(exposure)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(responseBody)
	if err != nil {
		log.WithField("method", "getExposures").WithError(err).Error("failed to write response")
	}
}
