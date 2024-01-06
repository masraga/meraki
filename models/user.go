package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	/*
		default collection field
	*/
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	IsDeleted bool               `bson:"isDeleted"`
	DeletedId primitive.ObjectID `bson:"deletedId,omitempty"`
	DeletedAt primitive.DateTime `bson:"deletedAt,omitempty"`
	CreatedAt primitive.DateTime `bson:"createdAt"`
	CreatedId primitive.ObjectID `bson:"createdId"`
	UpdatedAt primitive.DateTime `bson:"updatedAt,omitempty"`
	UpdatedId primitive.ObjectID `bson:"updatedId,omitempty"`

	/*
		fill your collection field below
	*/
	Username    string `bson:"username"`
	Password    string `bson:"password"`
	Name        string `bson:"name"`
	AccessToken string `bson:"accessToken"`
}
