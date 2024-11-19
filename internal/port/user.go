package port

import (
	"project/internal/core/application/dto"
	"project/internal/core/domain"
)

type UserRepository interface {
	FindId(id uint64) (*domain.User, error)
}

type UserService interface {
	Create(request dto.UserCreateRequest) (*dto.User, error)
}
