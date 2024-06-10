package entity

import (
	"context"
	"fmt"
)

type User struct {
	ID          int64
	Name        string
	PhoneNumber string
	Password    string
}

func (u *User) GetCacheKey() string {
	return fmt.Sprintf("user_profile:%d", u.ID)
}

type UserUsecase interface {
}

type UserRepository interface {
	Create(ctx context.Context, u *User) error
	FindUserByPhoneNumber(ctx context.Context, phoneNumber string) (u User, err error)
	FindById(ctx context.Context, id string) (user User, err error)
	FindByIds(ctx context.Context, id []string) (user []User, err error)
	FindUsers(ctx context.Context, gender uint8, excludeUserIds []string) (user User, err error)
	Update(ctx context.Context, user *User) (err error)
}
