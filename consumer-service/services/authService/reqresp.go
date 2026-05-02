package authService

type SignupRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	UserType    string `json:"user_type"`
	PhoneNumber string `json:"phone_number"`
}

type SignupResponse struct {
	Message string `json:"message"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
