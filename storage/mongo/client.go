package mongo

import (
	"github.com/sophie-rigg/havs-service/models"
	"github.com/sophie-rigg/havs-service/storage"
	"time"
)

type client struct {
}

func NewClient() storage.Client {
	return &client{}
}

func (c *client) GetEquipmentItem(id string) (*models.EquipmentItem, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) GetUser(id string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) GetExposure(id string) (*models.Exposure, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) GetExposuresByUserID(userID string, startTime, endTime time.Time) ([]*models.Exposure, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) GetExposures() ([]*models.Exposure, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) UpsertExposure(exposure *models.Exposure) error {
	//TODO implement me
	panic("implement me")
}
