package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mycandys-order-analytics/internal/database"
	"mycandys-order-analytics/internal/models"
)

type AnalyticsService struct {
	db *mongo.Collection
}

func NewAnalyticsService() *AnalyticsService {
	return &AnalyticsService{
		db: database.Db.Collection("analytics"),
	}
}

func (h *AnalyticsService) GetLastCalledEndpoint() (*models.Endpoint, error) {
	var endpoint models.Endpoint

	opts := options.FindOne()
	opts.SetSort(bson.D{{"timestamp", -1}}) // Sort by timestamp in descending order

	err := h.db.FindOne(nil, bson.D{}, opts).Decode(&endpoint)
	if err != nil {
		return nil, err
	}

	return &endpoint, nil
}

func (h *AnalyticsService) AddCalledEndpoint(endpoint *models.Endpoint) error {
	endpoint.SetTimestamp()
	_, err := h.db.InsertOne(nil, endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (h *AnalyticsService) GetNumberOfCallForEachEndpoint() (map[string]int, error) {
	pipeline := mongo.Pipeline{
		{{"$group", bson.D{
			{"_id", "$url"},                // Group by URL field
			{"count", bson.D{{"$sum", 1}}}, // Count occurrences
		}}},
	}

	cursor, err := h.db.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	result := make(map[string]int)
	for cursor.Next(context.Background()) {
		var data struct {
			URL   string `bson:"_id"`
			Count int    `bson:"count"`
		}
		if err := cursor.Decode(&data); err != nil {
			return nil, err
		}
		result[data.URL] = data.Count
	}

	return result, nil
}

func (h *AnalyticsService) GetMostCalledEndpoint() (interface{}, error) {
	pipeline := mongo.Pipeline{
		{{"$group", bson.D{
			{"_id", "$url"},                // Group by URL field
			{"count", bson.D{{"$sum", 1}}}, // Count occurrences
		}}},
		{{"$sort", bson.D{{"count", -1}}}}, // Sort by count in descending order
		{{"$limit", 1}},                    // Limit to the top-most endpoint
	}

	cursor, err := h.db.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var mostCalledEndpoint struct {
		URL   string `bson:"_id" json:"url"`
		Count int    `bson:"count" json:"count"`
	}

	if cursor.Next(context.Background()) {
		if err := cursor.Decode(&mostCalledEndpoint); err != nil {
			return nil, err
		}
	}

	return mostCalledEndpoint, nil
}
