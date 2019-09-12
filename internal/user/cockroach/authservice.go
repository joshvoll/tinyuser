package cockroach

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/joshvoll/tinyuser/internal/user"
)

// userInput model definition
type authUserInput struct {
	Email string `json:"email"`
}

const (
	// TokenLifespan represent until the token are valid
	TokenLifespan = time.Hour * 24 * 14
)

func (u *authUserInput) OK() error {
	// trim the email
	u.Email = strings.TrimSpace(u.Email)
	if !user.RXEmail.MatchString(u.Email) {
		return user.ErrInvalidEmail
	}
	return nil
}

// Login define the login service
func (s *UserService) Login(ctx context.Context, email string) (user.LoginOutput, error) {
	// local properties
	var out user.LoginOutput

	u := &authUserInput{
		Email: email,
	}

	// validate the email
	if err := decode(u); err != nil {
		return out, err
	}

	// make the query to check the database
	query := "SELECT id, username FROM users WHERE email = $1"
	err := s.client.db.QueryRowContext(ctx, query, u.Email).Scan(&out.AuthUser.ID, &out.AuthUser.Username)

	// check the errors
	if err == sql.ErrNoRows {
		return out, user.ErrUserNotFound
	}
	if err != nil {
		return out, fmt.Errorf("Could not query the user: %v ", err)
	}

	// make the token using branca token
	out.Token, err = s.client.codec.EncodeToString(strconv.FormatInt(out.AuthUser.ID, 10))
	if err != nil {
		return out, fmt.Errorf("could not encode the id of this user: %v ", err)
	}

	// add the expiration of the token
	out.ExpireAt = time.Now().Add(TokenLifespan)

	return out, nil

}
