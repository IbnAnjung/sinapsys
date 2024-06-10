package presenter

type RegisterRequest struct {
	Name            string `json:"name"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	PhoneNumber     string `json:"phone_number"`
}

type RegisterResponse struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	PhoneNumber  string `json:"phone_number"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	PhoneNumber  string `json:"phone_number"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
