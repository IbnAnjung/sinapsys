package auth

import (
	"context"
	"fmt"

	"github.com/IbnAnjung/synapsis/entity/dto/auth"
	coreerror "github.com/IbnAnjung/synapsis/pkg/error"

	"github.com/IbnAnjung/synapsis/pkg/jwt"
)

type LoginInputValidation struct {
	PhoneNumber string `validate:"required,ascii,max=13"`
	Password    string `validate:"required"`
}

func (uc *AuthUsecase) Login(ctx context.Context, input auth.LoginInput) (output auth.LoginOutput, err error) {
	// validate input
	if err = uc.validator.Validate(LoginInputValidation{
		PhoneNumber: input.PhoneNumber,
		Password:    input.Password,
	}); err != nil {
		return
	}

	// find user
	user, err := uc.userRepository.FindUserByPhoneNumber(ctx, input.PhoneNumber)
	if err != nil {
		err = coreerror.NewCoreError(coreerror.CoreErrorTypeAuthorization, "invalid credentials")
		return
	}

	// create new user
	if err = uc.hasher.CompareHash(user.Password, input.Password); err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeAuthorization, "invalid credentials")
		err = e
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

	output = auth.LoginOutput{
		ID:           user.ID,
		Name:         user.Name,
		PhoneNumber:  user.PhoneNumber,
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}

	return
}
