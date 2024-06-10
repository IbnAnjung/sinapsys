package auth

type LoginInput struct {
	PhoneNumber string
	Password    string
}

type LoginOutput struct {
	ID           int64
	PhoneNumber  string
	Name         string
	AccessToken  string
	RefreshToken string
}
