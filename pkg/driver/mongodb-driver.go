package driver

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Uri    string
	DbName string
	Db     *mongo.Database
}

func (mc *MongoClient) CreateClient() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mc.Uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	mc.Db = client.Database(mc.DbName)
}

type MongodbDriver struct {
	Client *MongoClient
}

func (m *MongodbDriver) Collection(collName string) *mongo.Collection {
	return m.Client.Db.Collection(collName)
}

func NewMongodbDriver(Uri string, DbName string) *MongodbDriver {
	return &MongodbDriver{
		Client: &MongoClient{Uri: Uri, DbName: DbName},
	}
}
