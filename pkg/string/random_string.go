package string

import "github.com/google/uuid"

type RandomString interface {
	GenerateID() string
}

type randomString struct{}

func NewRandomString() RandomString {
	return &randomString{}
}

func (*randomString) GenerateID() string {
	uid, _ := uuid.NewV7()
	return uid.String()
}
