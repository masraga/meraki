package driver

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Uri    string
	DbName string
}

func (mc *MongoClient) GetDatabase() *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mc.Uri))
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		panic(err)
	}

	return client.Database(mc.DbName)
}

type MongodbDriver struct {
	Database *mongo.Database
}

func (m *MongodbDriver) Collection(collName string) *mongo.Collection {
	return m.Database.Collection(collName)
}

func (m *MongodbDriver) ObjectID(id string) primitive.ObjectID {
	objId, _ := primitive.ObjectIDFromHex(id)
	return objId
}

func NewMongodbdriver(DbUri string, DbName string) *MongodbDriver {
	client := MongoClient{DbName: DbName, Uri: DbUri}
	return &MongodbDriver{
		Database: client.GetDatabase(),
	}
}
