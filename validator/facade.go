package validator

import (
	"errors"
	"fmt"
	"pruebaLDev180222/notificator"
	"time"
)

type validatorfacade struct {
	notificator notificator.INotificator
	validator   CredentialValidator
}

func NewValidatorFacade(notificator notificator.INotificator, validator CredentialValidator) *validatorfacade {
	return &validatorfacade{notificator: notificator, validator: validator}
}

func (b *validatorfacade) Validate(ID, Password string) (User, error) {
	fmt.Println("valores recibidos ID: " + ID + ", pass: " + Password)
	if b.validator == nil {
		return User{}, errors.New("validator not defined")
	}
	user, err := b.validator.Validate(ID, Password)
	if err != nil {
		//notificando error
		if errors.Is(err, errInvalidUserException) {
			fmt.Println("Notificando errores")
			b.notificator.Notify(ID, Password, time.Now().String(), time.Now().String())
		}
		return User{}, err
	}
	return user, err
}
