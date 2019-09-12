package user

import "regexp"

// General errors
const (
	// ErrInvalidEmail constring the email format
	ErrInvalidEmail = Error("Invalid Email")
	// ErrInvalidUsername constrain the username format
	ErrInvalidUsername = Error("Invalid User name")
	// ErrEmailTaken constrain the error if the users already exist on the database
	ErrEmailTaken = Error("Email taken")
	// ErrUsernameTaken constrain the error if the username already exist on the db
	ErrUsernameTaken = Error("Username taken")
	// ErrrUserNotFound constrain the error if the user it is not found on the database
	ErrUserNotFound = Error("User Not Found")
)

// Regex for validat the formst of the email and user name
var (
	RXEmail    = regexp.MustCompile("^[^\\s@]+@[^\\s@]+\\.[^\\s@]+$")
	RXUsername = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_-]{0,17}$")
)

// Error is a type sting
type Error string

func (e Error) Error() string {
	return string(e)
}

// ok validate error
type ok interface {
	OK() error
}
