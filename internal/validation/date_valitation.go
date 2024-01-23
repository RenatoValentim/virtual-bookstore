package validation

import (
	"errors"
	"time"
)

func DateValidation(date string) error {
	if date == "" {
		return errors.New("A CreatAt field is required.")
	}
	_, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return err
	}
	return nil
}
