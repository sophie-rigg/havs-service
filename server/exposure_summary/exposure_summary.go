package exposuresummary

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sophie-rigg/havs-service/models"
	"github.com/sophie-rigg/havs-service/storage"
	"net/http"
	"time"
)

const (
	timeFormat = "2006-01-02T15:04:05Z"
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
	default:
		http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (c *Handler) Get(writer http.ResponseWriter, request *http.Request) {
	requestVars := mux.Vars(request)

	userID, ok := requestVars["userId"]
	if !ok {
		http.Error(writer, "missing userId", http.StatusBadRequest)
		return
	}

	query := request.URL.Query()

	startingAt := query.Get("starting_at")
	if startingAt == "" {
		http.Error(writer, "missing startTime", http.StatusBadRequest)
		return
	}

	endingAt := query.Get("ending_at")
	if endingAt == "" {
		http.Error(writer, "missing endTime", http.StatusBadRequest)
		return
	}

	startTime, err := time.Parse(timeFormat, startingAt)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	endTime, err := time.Parse(timeFormat, endingAt)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	exposures, err := c.storage.GetExposuresByUserID(userID, startTime, endTime)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	summary := models.NewExposureSummary(exposures)

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(summary)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
