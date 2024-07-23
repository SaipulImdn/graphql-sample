package graphql

import (
	"context"
	"graphql-sample/internal/models"
	"graphql-sample/internal/services"
)

type Resolver struct {
    UserService services.UserService
}

func (r *Resolver) Query_user(ctx context.Context, id string) (*models.User, error) {
    return r.UserService.GetUserByID(id)
}
