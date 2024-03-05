package validation

import (
	"errors"

	"github.com/RenatoValentim/virtual-bookstore/internal/constants/environments"
	"github.com/badoux/checkmail"
	"github.com/spf13/viper"
)

func EmailValitation(email string) error {
	if email == "" {
		return errors.New("An Email field is required.")
	}
	err := checkmail.ValidateFormat(email)
	if err != nil {
		return err
	}
	if viper.GetString(environments.Environment) == environments.Prod {
		err = checkmail.ValidateHost(email)
		if err != nil {
			return err
		}
	}
	return nil
}
