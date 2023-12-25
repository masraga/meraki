package system

import (
	"context"

	"github.com/masraga/meraki/pkg"
	driver "github.com/masraga/meraki/pkg/driver"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	CollName string
	Db       *driver.MongodbDriver
}

func (r *MongoRepository) InsertOne(request interface{}) (*mongo.InsertOneResult, error) {
	query, err := r.Db.Collection(r.CollName).InsertOne(context.TODO(), request)
	if err != nil {
		return nil, err
	}

	return query, nil
}

func (r *MongoRepository) FindById(id string) *mongo.SingleResult {

	filter := bson.D{
		primitive.E{Key: "_id", Value: r.Db.ObjectID(id)},
		primitive.E{Key: "isDeleted", Value: false},
	}

	cursor := r.Db.Collection(r.CollName).FindOne(context.TODO(), filter)

	return cursor
}

func (r *MongoRepository) UpdateByID(id string, fieldSet interface{}) (*mongo.UpdateResult, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: r.Db.ObjectID(id)}}
	update := bson.D{primitive.E{Key: "$set", Value: fieldSet}}
	result, err := r.Db.Collection(r.CollName).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MongoRepository) DeleteByID(id string) (*mongo.DeleteResult, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: r.Db.ObjectID(id)}}
	result, err := r.Db.Collection(r.CollName).DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *MongoRepository) FindOne(filter bson.D) (*mongo.SingleResult, error) {
	result := r.Db.Collection(r.CollName).FindOne(context.TODO(), filter)
	return result, nil
}

func (r *MongoRepository) Aggregate(filter bson.D) (*mongo.Cursor, error) {
	cursor, err := r.Db.Collection(r.CollName).Aggregate(context.TODO(), mongo.Pipeline{filter})
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {

	}

	return cursor, nil
}

func NewMongoRepository(CollName string) *MongoRepository {
	db := pkg.NewAutoload().Database()
	return &MongoRepository{
		CollName: CollName,
		Db:       db,
	}
}
