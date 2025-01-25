package storage

import (
	"github.com/sophie-rigg/havs-service/models"
	"time"
)

type Client interface {
	GetEquipmentItem(id string) (*models.EquipmentItem, error)
	GetUser(id string) (*models.User, error)
	GetExposure(id string) (*models.Exposure, error)
	GetExposuresByUserID(userID string, startTime, endTime time.Time) ([]*models.Exposure, error)
	GetExposures() ([]*models.Exposure, error)

	UpsertExposure(exposure *models.Exposure) error
}
