package services

import (
	"graphql-sample/internal/models"
	"graphql-sample/internal/repositories"
)

type UserService interface {
    GetUserByID(id string) (*models.User, error)
}

type userService struct {
    userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
    return &userService{
        userRepository: userRepository,
    }
}

func (s *userService) GetUserByID(id string) (*models.User, error) {
    return s.userRepository.GetUserByID(id)
}
