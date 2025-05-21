package repository

import (
	"context"
	"myapp/model"
)

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
}
