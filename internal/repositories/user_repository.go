package repositories

import "graphql-sample/internal/models"

type UserRepository interface {
    GetUserByID(id string) (*models.User, error)
}

type userRepository struct {
    users map[string]*models.User
}

func NewUserRepository() UserRepository {
    return &userRepository{
        users: map[string]*models.User{
            "1": {ID: "1", Name: "John Doe", Age: 28},
        },
    }
}

func (r *userRepository) GetUserByID(id string) (*models.User, error) {
    user, exists := r.users[id]
    if !exists {
        return nil, nil
    }
    return user, nil
}
