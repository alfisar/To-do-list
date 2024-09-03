package service

import (
	"context"
	"todolist/domain"

	"github.com/valyala/fasthttp"
)

type AuthServiceContract interface {
	Registration(ctx context.Context, r *fasthttp.Request, data domain.User) (result domain.User, err domain.ErrorData)
	Login(ctx context.Context, data domain.Login) (token string, err domain.ErrorData)
}
