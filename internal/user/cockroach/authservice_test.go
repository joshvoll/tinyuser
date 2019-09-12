package cockroach_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/joshvoll/tinyuser/internal/user"
)

// test the auth services
type loginUserInput struct {
	Email string `json:"email"`
}

func TestAuthService(t *testing.T) {
	// init the database
	c := MustOpenClient()
	defer c.Close()
	// local properties
	in := &loginUserInput{
		Email: "jrodriguezsanservices.hn",
	}

	ctx := context.Background()
	// add the services
	s := c.Service()

	t.Run("login credentials", func(t *testing.T) {
		_, err := s.Login(ctx, in.Email)
		if err != user.ErrInvalidEmail {
			t.Errorf("should invalid the email for bad format: want=%v, got=%v ", user.ErrInvalidEmail, err)
		}
	})
	t.Run("Login user", func(t *testing.T) {
		var out user.LoginOutput
		out, err := s.Login(ctx, "bbarahona@sanservices.hn")
		if err == sql.ErrNoRows {
			t.Errorf("should have record for this user: %v ", err)
		}
		fmt.Println(out)

	})

}
