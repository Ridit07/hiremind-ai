package authService

type UserType string

const (
	UserTypeCandidate UserType = "candidate"
	UserTypeHR        UserType = "hr"
)

type contextKey string

const userIDKey contextKey = "user_id"

type SignupRequest struct {
	Email       string   `json:"email"`
	Password    string   `json:"password"`
	UserType    UserType `json:"user_type"`
	PhoneNumber string   `json:"phone_number"`
}

type SignupResponse struct {
	Message string `json:"message"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
