package authService

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {

	hashedBytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	return string(hashedBytes), err
}

func ComparePassword(
	hashed string,
	plain string,
) error {

	return bcrypt.CompareHashAndPassword(
		[]byte(hashed),
		[]byte(plain),
	)
}
