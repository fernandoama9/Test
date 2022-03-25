package validator

import (
	"errors"
)
var errInvalidUserException = errors.New("InvalidUserException")
type CredentialValidator interface {
	Validate(ID, Password string) (User, error)
}

type User struct {
	ID, pass, email string
}

type validatorMock struct {
}

func GetInstance() *validatorMock {
	return &validatorMock{}
}
func (v validatorMock) Validate(ID, Password string) (User, error) {
	if ID == "fallo"{
		return User{}, errInvalidUserException
	}
	return User{
		ID:    ID,
		pass:  Password,
		email: ID + "@" + Password,
	}, nil
}