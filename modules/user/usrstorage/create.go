package usrstorage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"tbox-homework/modules/user/usrmodel"
	"time"
)

func (storage *userMongoStorage) CreateUser(ctx context.Context, user usrmodel.User) error {
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	collection := storage.db.Database("tbox-otp").Collection(user.CollectionName())
	_, err := collection.InsertOne(ctx, bson.M{"name": "", "phone_number": user.PhoneNumber, "verified_otp": false})
	if err != nil {
		return err
	}
	return nil
}
