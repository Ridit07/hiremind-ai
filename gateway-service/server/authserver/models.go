package server

type UserType string

const (
	UserTypeCandidate UserType = "CANDIDATE"
	UserTypeHR        UserType = "HR"
)

type SignupRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	UserType    string `json:"user_type"`
	PhoneNumber string `json:"phone_number"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupResponse struct {
	Message string `json:"message"`
	Error   *Error `json:"error,omitempty"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Error        *Error `json:"error,omitempty"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Error        *Error `json:"error,omitempty"`
}

type LogoutResponse struct {
	Message string `json:"message"`
	Error   *Error `json:"error,omitempty"`
}

type Error struct {
	Code    int32  `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}
