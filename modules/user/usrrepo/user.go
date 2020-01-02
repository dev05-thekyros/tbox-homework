package usrrepo

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"tbox-homework/common"
	"tbox-homework/modules/user/usrmodel"
	"time"
)

type UserStorage interface {
	CreateUser(ctx context.Context, user usrmodel.User) error
	UpdateUser(ctx context.Context, user usrmodel.UserUpdate, phoneNumber string) error
	GetUserByPhoneNumber(ctxt context.Context, phoneNumber string) (*usrmodel.User, error)
}

type userRepo struct {
	usrStorage    UserStorage
	redisProvider common.RedisStorageProvider
	smsService    common.SMSService
}

func NewUserRepo(storage UserStorage, redis common.RedisStorageProvider, smsService common.SMSService) *userRepo {
	return &userRepo{storage, redis, smsService}
}

func (repo *userRepo) SignUp(ctx context.Context, phoneNumber string) (string, error) {
	existedUsr, err := repo.usrStorage.GetUserByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		return "", err
	}
	if existedUsr != nil {
		// check and return token if this user is active and have token
		if existedUsr.IsVerifiedOTP == true {
			return existedUsr.AccessToken, nil
		}
	} else {
		// create  new one in db
		if err = repo.usrStorage.CreateUser(ctx, usrmodel.User{PhoneNumber: phoneNumber}); err != nil {
			return "", err
		}
	}
	// Get and verified OTP stored in Redis
	expiredTime, _, rdErr := repo.GetOTPFromRedis(ctx, phoneNumber)
	if rdErr != nil && rdErr.Error() != common.RedisNotFound {
		return "", rdErr
	}
	// init new OTP if redis value of mobile number not found or redis number is invalid
	if expiredTime == 0 { // not found OTP
		if err1 := repo.initOTP(phoneNumber); err1 != nil {
			return "", err1
		}
		return "", nil
	}

	// init new OTP, if this last 30 second expired time
	if (expiredTime - 30) <= time.Now().Unix() {
		if err1 := repo.initOTP(phoneNumber); err1 != nil {
			return "", err1
		}
		return "", nil
	}
	return "", errors.New(fmt.Sprintf("The OTP has already sent to phone number %s, you can request new OTP after 30s!", phoneNumber))
}

func (repo *userRepo) initOTP(phoneNumber string) error {
	otp, err := common.GetRandNumberToString()
	if err != nil {
		return err
	}
	//expiredAt := time.Now().Unix() + 60
	expiredAt := time.Now().Unix() + 3600
	redisKey := fmt.Sprintf("%s-%s", common.RedisPhoneOTPPrefix, phoneNumber)
	redisValue := fmt.Sprintf("%d|%s", expiredAt, otp)
	err = repo.redisProvider.SetKey(redisKey, redisValue, time.Minute)
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepo) VerifyOTP(ctx context.Context, phoneNumber string, userOTP string) (string, error) {
	existedUsr, err := repo.usrStorage.GetUserByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		return "", err
	}
	if existedUsr != nil {
		// check and return token if this user is active and have token
		if existedUsr.IsVerifiedOTP == true {
			return existedUsr.AccessToken, nil
		}
	} else {
		// create new one in db
		return "", errors.New(fmt.Sprintf("User phone number: %s is not existed!", phoneNumber))
	}

	// Verify OTP from redis
	_, storedOTP, rdErr := repo.GetOTPFromRedis(ctx, phoneNumber)
	if rdErr != nil {
		return "", rdErr
	}
	if storedOTP == "" {
		return "", errors.New("the OTP code is not expired, please request another OTP")
	}
	if storedOTP != userOTP {
		return "", errors.New("input OTP is not corrected, please insert another one")
	} else {
		accessToken := common.TokenGenerator()
		if err = repo.usrStorage.UpdateUser(ctx, usrmodel.UserUpdate{IsVerifiedOTP: true, AccessToken: accessToken}, phoneNumber); err != nil {
			return "", err
		}
		return accessToken, nil
	}
}

// return expiredTime, OTP code, error if have
func (repo *userRepo) GetOTPFromRedis(ctx context.Context, phoneNumber string) (int64, string, error) {
	rdValueStr, rdErr := repo.redisProvider.GetKey(ctx, fmt.Sprintf("%s-%s", common.RedisPhoneOTPPrefix, phoneNumber))
	if rdErr != nil && rdErr.Error() != common.RedisNotFound {
		return 0, "", rdErr
	}
	rdValues := strings.Split(rdValueStr, "|") // first time expire time, second is OTP
	if len(rdValues) != 2 {
		return 0, "", nil // Not found OTP
	}
	expiredTime, err := strconv.ParseInt(rdValues[0], 10, 64)
	if err != nil {
		return 0, "", err
	}
	return expiredTime, rdValues[1], nil //  Return expiredTime, OTP code
}
