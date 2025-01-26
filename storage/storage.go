package storage

import (
	"github.com/sophie-rigg/havs-service/models"
	"time"
)

//go:generate mockgen -destination=mocks/mock_client.go -package=mocks -source=storage.go
type Client interface {
	GetEquipmentItem(id string) (*models.EquipmentItem, error)
	GetUser(id string) (*models.User, error)
	GetExposure(id string) (*models.Exposure, error)
	GetExposuresByUserID(userID string, startTime, endTime time.Time) ([]*models.Exposure, error)
	GetExposures() ([]*models.Exposure, error)

	InsertExposure(exposure *models.Exposure) error
}
