package validate

import "net/mail"

// Email cheks whether given email string is invalid format

func Email(email string) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return err
	}
	return nil
}
