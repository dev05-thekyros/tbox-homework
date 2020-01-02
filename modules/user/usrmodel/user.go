package usrmodel

import (
	"errors"
	"tbox-homework/common"
)

type User struct {
	FullName      string `bson:"full_name" json:"full_name,omitempty"`
	PhoneNumber   string `bson:"phone_number" json:"phone_number,omitempty"`
	IsVerifiedOTP bool   `bson:"is_verified_otp" json:"is_verified_otp,omitempty"`
	AccessToken   string `bson:"access_token" json:"access_token,omitempty"`
}

func (s User) CollectionName() string {
	return "users"
}

type SignUpUser struct {
	PhoneNumber string `json:"phone_number"`
}

func (s *SignUpUser) Validate() error {
	if !common.ValidatePhone(s.PhoneNumber) {
		return errors.New("Invalid phone number.")
	}
	return nil
}

type SignUpUsrResponse struct {
	Message string `json:"message, omitempty"`
	IsError bool   `json:"is_error"`
	Data    User   `json:"data"`
}
