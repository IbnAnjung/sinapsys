package jwt

import (
	"time"

	coreerror "github.com/IbnAnjung/synapsis/pkg/error"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	GenerateAccessToken(claim UserClaim) (token string, err error)
	GenerateRefreshToken(claim UserClaim) (token string, err error)
	ValidateToken(tokenString string) (c *UserClaim, err error)
}

type UserType int8

type UserClaim struct {
	UserID int64 `json:"id"`
	jwt.RegisteredClaims
}

type JwtConfig struct {
	SecretKey            string
	AccessTokenLifeTime  time.Duration
	RefreshTokenLifeTime time.Duration
}

type jwtService struct {
	conf JwtConfig
}

func NewJwtServive(conf JwtConfig) JwtService {
	return &jwtService{
		conf,
	}
}

func (s *jwtService) GenerateAccessToken(claim UserClaim) (token string, err error) {
	claim.ExpiresAt = jwt.NewNumericDate(time.Now().Add(s.conf.AccessTokenLifeTime))
	claim.IssuedAt = jwt.NewNumericDate(time.Now())

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err = jwtToken.SignedString([]byte(s.conf.SecretKey))
	if err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
		err = e
		return
	}

	return
}

func (s *jwtService) GenerateRefreshToken(claim UserClaim) (token string, err error) {
	claim.ExpiresAt = jwt.NewNumericDate(time.Now().Add(s.conf.RefreshTokenLifeTime))
	claim.IssuedAt = jwt.NewNumericDate(time.Now())

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err = jwtToken.SignedString([]byte(s.conf.SecretKey))
	if err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
		err = e
		return
	}

	return
}

func (s *jwtService) ValidateToken(tokenString string) (c *UserClaim, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.conf.SecretKey), nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeAuthorization, err.Error())
		err = e
		return
	} else if claim, ok := token.Claims.(*UserClaim); ok {
		return claim, nil
	} else {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeAuthorization, "unknown error")
		err = e
		return
	}
}
