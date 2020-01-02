package usrhandler

import "context"

type VerifyUserOTP interface {
	VerifyOTP(ctx context.Context, phoneNumber string, OTP string) (string, error)
}

type verifyUserOTPrHandler struct {
	userRepo VerifyUserOTP
}

func NewVerifyUserOTPrHandler(usrRepo VerifyUserOTP) *verifyUserOTPrHandler {
	return &verifyUserOTPrHandler{usrRepo}
}

func (handler *verifyUserOTPrHandler) Response(ctx context.Context, userPhoneNumber string, otp string) (string, error) {
	return handler.userRepo.VerifyOTP(ctx, userPhoneNumber, otp)
}
