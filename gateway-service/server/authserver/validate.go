package server

import (
	"fmt"
	"strings"
)

func (r *SignupRequest) Validate() error {
	if strings.TrimSpace(r.Email) == "" {
		return fmt.Errorf("email is required")
	}

	if len(strings.TrimSpace(r.Password)) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	}

	if strings.TrimSpace(r.UserType) == "" {
		return fmt.Errorf("user_type is required")
	}

	if r.UserType != "CANDIDATE" && r.UserType != "HR" {
		return fmt.Errorf("invalid user_type: must be CANDIDATE or HR")
	}

	return nil
}

func (r *LoginRequest) Validate() error {
	if strings.TrimSpace(r.Email) == "" {
		return fmt.Errorf("email is required")
	}

	if strings.TrimSpace(r.Password) == "" {
		return fmt.Errorf("password is required")
	}

	return nil
}
