package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client         *mongo.Client
	databaseName   string
	collectionName string
}

// NewMongoDB initializes a MongoDB connection.
func NewMongoDB(uri, databaseName, collectionName string) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	return &MongoDB{
		client:         client,
		databaseName:   databaseName,
		collectionName: collectionName,
	}, nil
}

func (m *MongoDB) GetConfigs(pkg string, country string, randomPercentile int) ([]*Config, error) {
	collection := m.client.Database(m.databaseName).Collection(m.collectionName)
	filter := bson.M{
		"package":        pkg,
		"country_code":   country,
		"percentile_min": bson.M{"$lt": randomPercentile},
		"percentile_max": bson.M{"$gte": randomPercentile},
	}

	var configs []*Config
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &configs); err != nil {
		return nil, err
	}
	return configs, nil
}

func (m *MongoDB) Close() error {
	return m.client.Disconnect(context.TODO())
}
