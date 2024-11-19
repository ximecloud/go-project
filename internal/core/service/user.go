package service

import (
	"project/internal/core/application/dto"
	"project/internal/port"
)

type userService struct {
	user port.UserRepository
}

func (r *userService) Create(req dto.UserCreateRequest) (*dto.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserService(user port.UserRepository) port.UserService {
	return &userService{user: user}
}
