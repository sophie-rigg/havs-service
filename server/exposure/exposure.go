package exposure

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/sophie-rigg/havs-service/models"

	log "github.com/sirupsen/logrus"
	"github.com/sophie-rigg/havs-service/storage"
	"net/http"
)

type Handler struct {
	storage storage.Client
}

func NewHandler(storage storage.Client) http.Handler {
	return &Handler{
		storage: storage,
	}
}

func (c *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		c.Get(writer, request)
	case http.MethodPost:
		c.Post(writer, request)
	default:
		http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (c *Handler) Get(writer http.ResponseWriter, request *http.Request) {
	exposures, err := c.storage.GetExposures()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(exposures)
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

func (c *Handler) Post(writer http.ResponseWriter, request *http.Request) {
	var body bytes.Buffer
	_, err := body.ReadFrom(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	var badRequest bool
	exposure, err := models.NewExposureFromRequestBody(body.Bytes())
	if err != nil {
		if errors.Is(err, models.ErrorInvalidExposureRequest) {
			// if bad request, return 400
			// we still want to continue processing the request so we can recover from the error
			badRequest = true
		} else {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// could be a bad request, but we still want to save the exposure
	// so we can recover from the error
	// TODO: implement recovery method
	// would need to add option to provide a Exposure ID to the request
	if exposure.Equipment != nil && exposure.Equipment.ID != "" {
		equipment, err := c.storage.GetEquipmentItem(exposure.Equipment.ID)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		exposure.SetEquipment(equipment)
		err = exposure.CalculateExposure()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if exposure.User != nil && exposure.User.ID != "" {
		user, err := c.storage.GetUser(exposure.User.ID)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		exposure.SetUser(user)
	}

	err = c.storage.UpsertExposure(exposure)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if badRequest {
		http.Error(writer, models.ErrorInvalidExposureRequest.Error(), http.StatusBadRequest)
		return
	}

	responseBody, err := json.Marshal(exposure)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(responseBody)
	writer.WriteHeader(http.StatusCreated)
}
