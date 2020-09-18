package service

import "context"

type UserService interface {
	Login(ctx context.Context, email, password string) (*UserInfoDTO, error)
	Register(ctx context.Context)
}

type UserInfoDTO struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
