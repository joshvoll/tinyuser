package cockroach

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx"
	"github.com/joshvoll/tinyuser/internal/user"
)

type userInput struct {
	Email    string
	Username string
}

// OK implemente de validation for a user credential
func (u *userInput) OK() error {
	// check the white space
	u.Email = strings.TrimSpace(u.Email)
	u.Username = strings.TrimSpace(u.Username)

	// validate the regex
	if !user.RXEmail.MatchString(u.Email) {
		return user.ErrInvalidEmail
	}
	if !user.RXUsername.MatchString(u.Username) {
		return user.ErrInvalidUsername
	}
	return nil
}

// Decode acces the interface to return each of the error
func decode(v user.Validator) error {
	return v.OK()
}

// User model definition
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

// CreateUser defin the login services of the user
func (s *UserService) CreateUser(ctx context.Context, email, username string) error {
	// first check the email and usernmae
	u := &userInput{
		Email:    email,
		Username: username,
	}
	if err := decode(u); err != nil {
		return err
	}

	// make the query and save the users
	query := "INSERT INTO users (email, username) VALUES ($1, $2)"
	_, err := s.client.db.ExecContext(ctx, query, u.Email, u.Username)

	unique := isUniqueViolation(err)

	// check if the unique constrina have a problem
	if unique && strings.Contains(err.Error(), "email") {
		return user.ErrEmailTaken
	}

	if unique && strings.Contains(err.Error(), "username") {
		return user.ErrUsernameTaken
	}
	if err != nil {
		return fmt.Errorf("could not save the user to the database: %v ", err)
	}

	return nil
}

func isUniqueViolation(err error) bool {
	pgerr, ok := err.(pgx.PgError)
	return ok && pgerr.Code == "23505"

}
