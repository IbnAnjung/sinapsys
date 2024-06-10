package auth

import (
	"context"
	"fmt"

	"github.com/IbnAnjung/synapsis/entity"
	"github.com/IbnAnjung/synapsis/entity/dto/auth"

	"github.com/IbnAnjung/synapsis/pkg/jwt"
)

type RegisterInputValidation struct {
	PhoneNumber     string `validate:"required,numeric,max=13"`
	Name            string `validate:"required,ascii,max=50,min=3"`
	Password        string `validate:"required,min=6,max=72"`
	ConfirmPassword string `validate:"required,min=6,eqfield=Password"`
}

func (uc *AuthUsecase) RegisterUser(ctx context.Context, input auth.RegisterInput) (output auth.RegisterOutput, err error) {
	// validate input
	if err = uc.validator.Validate(RegisterInputValidation{
		PhoneNumber:     input.PhoneNumber,
		Name:            input.Name,
		Password:        input.Password,
		ConfirmPassword: input.ConfirmPassword,
	}); err != nil {
		return
	}

	// create new user
	hashedPassword, err := uc.hasher.HashString(input.Password)
	if err != nil {
		return
	}

	user := entity.User{
		PhoneNumber: input.PhoneNumber,
		Name:        input.Name,
		Password:    hashedPassword,
	}

	if err = uc.userRepository.Create(ctx, &user); err != nil {
		return
	}

	// generate token
	userClaim := jwt.UserClaim{
		UserID: user.ID,
	}

	accessToken, err := uc.jwtService.GenerateAccessToken(userClaim)
	if err != nil {
		return
	}

	refreshToken, err := uc.jwtService.GenerateRefreshToken(userClaim)
	if err != nil {
		return
	}

	if e := uc.cacheService.Set(ctx, user.GetCacheKey(), user, -1); e != nil {
		fmt.Println("fail set cache user profile")
	}

	output = auth.RegisterOutput{
		ID:           user.ID,
		Name:         user.Name,
		PhoneNumber:  user.PhoneNumber,
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}

	return
}
