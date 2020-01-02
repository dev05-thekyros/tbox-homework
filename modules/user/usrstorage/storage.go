package usrstorage

import "go.mongodb.org/mongo-driver/mongo"

type userMongoStorage struct {
	db *mongo.Client
}

func NewUserMongoStorage(db *mongo.Client) *userMongoStorage {
	return &userMongoStorage{
		db: db,
	}
}
