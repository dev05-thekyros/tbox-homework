package usrhandler

import "context"

type SignUpUserRepo interface {
	SignUp(ctx context.Context, phoneNumber string) (string, error)
}

type signUpUserHandler struct {
	userRepo SignUpUserRepo
}

func NewSignUpUserHandler(usrRepo SignUpUserRepo) *signUpUserHandler {
	return &signUpUserHandler{usrRepo}
}

func (handler *signUpUserHandler) Response(ctx context.Context, userPhoneNumber string) (string, error) {
	return handler.userRepo.SignUp(ctx, userPhoneNumber)
}
