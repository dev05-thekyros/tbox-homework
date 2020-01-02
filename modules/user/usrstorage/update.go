package usrstorage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"tbox-homework/modules/user/usrmodel"
	"time"
)

func (storage *userMongoStorage) UpdateUser(ctx context.Context, user usrmodel.UserUpdate, phoneNumber string) error {
	mgoCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"phone_number": phoneNumber}
	updateData := bson.M{"$set": user}
	collection := storage.db.Database("tbox-otp").Collection(usrmodel.User{}.CollectionName())
	if _, err := collection.UpdateOne(mgoCtx, filter, updateData); err != nil {
		return err
	}
	return nil
}
