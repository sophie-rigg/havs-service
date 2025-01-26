package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/sophie-rigg/havs-service/models"
	"github.com/sophie-rigg/havs-service/storage"
	"time"
)

type client struct {
	client *mongo.Client
	db     string
}

func NewClient(uri, db string) (storage.Client, error) {
	clientOption := options.Client()
	clientOption.ApplyURI(uri)
	c, err := mongo.Connect(clientOption)
	if err != nil {
		return nil, err
	}

	err = c.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return &client{
		client: c,
		db:     db,
	}, err
}

func (c *client) GetEquipmentItem(id string) (*models.EquipmentItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := c.client.Database(c.db).Collection("equipment")

	res := collection.FindOne(ctx, bson.D{{Key: "_id", Value: id}})
	var result models.EquipmentItem
	if err := res.Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *client) GetUser(id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := c.client.Database(c.db).Collection("users")

	res := collection.FindOne(ctx, bson.D{{Key: "_id", Value: id}})
	var result models.User
	if err := res.Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *client) GetExposure(id string) (*models.Exposure, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := c.client.Database(c.db).Collection("exposures")

	res := collection.FindOne(ctx, bson.D{{Key: "_id", Value: id}})
	var result models.Exposure
	if err := res.Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *client) GetExposuresByUserID(userID string, startTime, endTime time.Time) ([]*models.Exposure, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := c.client.Database(c.db).Collection("exposures")

	filter := bson.D{
		{Key: "user._id", Value: userID},
		{Key: "created_time", Value: bson.D{
			{Key: "$gte", Value: startTime},
			{Key: "$lte", Value: endTime},
		}},
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var results []*models.Exposure
	for cursor.Next(ctx) {
		var result models.Exposure
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, &result)
	}

	return results, nil
}

func (c *client) GetExposures() ([]*models.Exposure, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := c.client.Database(c.db).Collection("exposures")

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var results []*models.Exposure
	for cursor.Next(ctx) {
		var result models.Exposure
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, &result)
	}

	return results, nil
}

func (c *client) InsertExposure(exposure *models.Exposure) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := c.client.Database(c.db).Collection("exposures")

	_, err := collection.InsertOne(ctx, exposure)
	return err
}
