package mapper

import (
	"project/internal/core/application/dto"
	"project/internal/core/domain"
)

type UserMapper interface {
	ToDTO(user *domain.User) *dto.User
	ToDomain(dto interface{}) *domain.User
}

type userMapper struct{}

func (r *userMapper) ToDTO(user *domain.User) *dto.User {
	return &dto.User{
		Id:        user.Id,
		Username:  user.Username,
		Status:    user.Status,
		UpdatedAt: user.UpdatedAt,
	}
}

func (r *userMapper) ToDomain(input interface{}) *domain.User {
	switch v := input.(type) {
	case *dto.User:
		return &domain.User{
			Id:        v.Id,
			Username:  v.Username,
			Status:    v.Status,
			UpdatedAt: v.UpdatedAt,
		}
	case *dto.UserCreateRequest:
		return &domain.User{
			Username: v.Username,
			Password: v.Password,
		}
	default:
		return nil
	}
}

func NewUserMapper() UserMapper {
	return &userMapper{}
}
