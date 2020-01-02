package usrstorage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"tbox-homework/common"
	"tbox-homework/modules/user/usrmodel"
	"time"
)

func (storage *userMongoStorage) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*usrmodel.User, error) {

	var rs = usrmodel.User{}
	filter := bson.M{"phone_number": phoneNumber}
	collection := storage.db.Database("tbox-otp").Collection(rs.CollectionName())
	mongoCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := collection.FindOne(mongoCtx, filter).Decode(&rs)
	if err != nil {
		if err.Error() == common.MgoNotFound {
			// not found
			return nil, nil
		}
		return nil, err
	}
	return &rs, nil
}
