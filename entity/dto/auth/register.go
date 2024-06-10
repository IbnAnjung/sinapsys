package auth

type RegisterInput struct {
	PhoneNumber     string
	Password        string
	ConfirmPassword string
	Name            string
}

type RegisterOutput struct {
	ID           int64
	PhoneNumber  string
	Name         string
	AccessToken  string
	RefreshToken string
}
