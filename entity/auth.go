package entity

import (
	"context"

	"github.com/IbnAnjung/synapsis/entity/dto/auth"
)

type AuthUsecase interface {
	RegisterUser(ctx context.Context, input auth.RegisterInput) (output auth.RegisterOutput, err error)
	Login(ctx context.Context, input auth.LoginInput) (output auth.LoginOutput, err error)
}
