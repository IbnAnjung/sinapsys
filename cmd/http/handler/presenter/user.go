package presenter

type GetDatingUserProfileResponse struct {
	ID       string `json:"id"`
	Fullname string `json:"fullname"`
	Age      uint8  `json:"age"`
	Gender   uint8  `json:"gender"`
}

type SwipeUserRequest struct {
	SwipeUserID string `json:"user_id"`
	Type        *uint8 `json:"type"`
}
