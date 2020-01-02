package usrmodel

import (
	"errors"
	"tbox-homework/common"
)

type UserUpdate struct {
	FullName      string `bson:"full_name,omitempty" json:"full_name"`
	PhoneNumber   string `bson:"phone_number,omitempty" json:"phone_number"`
	AccessToken   string `bson:"access_token,omitempty"`
	IsVerifiedOTP bool   `bson:"is_verified_otp,omitempty" json:"is_verified_otp"`
}

type VerifyOTPRequest struct {
	PhoneNumber string `json:"phone_number"`
	OTP         string `json:"otp"`
}
type VerifyOTPResponse struct {
	Message string `json:"message, omitempty"`
	IsError bool   `json:"is_error"`
	Data    User   `json:"data"`
}

func (s *VerifyOTPRequest) Validate() error {
	if !common.ValidatePhone(s.PhoneNumber) {
		return errors.New("Invalid phone number.")
	}
	return nil
}
