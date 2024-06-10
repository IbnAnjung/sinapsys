package auth

import (
	"github.com/IbnAnjung/synapsis/entity"
	"github.com/IbnAnjung/synapsis/pkg/cache"
	"github.com/IbnAnjung/synapsis/pkg/crypt"
	"github.com/IbnAnjung/synapsis/pkg/structvalidator"

	"github.com/IbnAnjung/synapsis/pkg/jwt"
)

type AuthUsecase struct {
	hasher         crypt.Hash
	jwtService     jwt.JwtService
	cacheService   cache.CacheService
	validator      structvalidator.Validator
	userRepository entity.UserRepository
}

func NewUsecase(
	hasher crypt.Hash,
	jwtService jwt.JwtService,
	cacheService cache.CacheService,
	validator structvalidator.Validator,
	userRepository entity.UserRepository,
) entity.AuthUsecase {
	return &AuthUsecase{
		hasher, jwtService, cacheService, validator, userRepository,
	}
}
