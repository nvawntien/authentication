package repository

import (
	"context"
	"myapp/model"
	"myapp/model/req"
)

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
	CheckLogin(context context.Context, loginReq req.ReqSignIn) (model.User, error)
}
