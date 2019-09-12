package cockroach_test

import (
	"context"
	"testing"

	"github.com/joshvoll/tinyuser/internal/user"
)

type createUserInput struct {
	Email    string
	Username string
}

// Ensure user service can creat and retrieved
func TestUserService(t *testing.T) {
	// open the database
	c := MustOpenClient()
	defer c.Close()
	// instance the services
	s := c.Service()

	// Test the credentials validation function
	t.Run("Test Credentials", func(t *testing.T) {

		ctx := context.Background()

		// loing service wrong email address
		if err := s.CreateUser(ctx, "joshvoll.com", "josue"); err != user.ErrInvalidEmail {
			t.Errorf("should return ErrInvalidEmail: got=%v, want=%v ", err, user.ErrInvalidEmail)
		}
		// login service wrong username
		if err := s.CreateUser(ctx, "joshvoll@yahoo.com", "josue manuel rodriguez"); err != user.ErrInvalidUsername {
			t.Errorf("shoudl return ErrInvalidUsername: got=%v, want=%v", err, user.ErrInvalidUsername)
		}
	})

	// Test if the services is saving data to the database
	t.Run("Test Create User err already exists", func(t *testing.T) {
		ctx := context.Background()

		// lets save a user
		in := createUserInput{
			Email:    "bbarahona@sanservices.hn",
			Username: "bbarahona",
		}
		if err := s.CreateUser(ctx, in.Email, in.Username); err != user.ErrEmailTaken {
			t.Errorf("should not create the user, already exists: want=%v, got=%v", user.ErrEmailTaken, err)
		}
		// send an email that already exist
	})

}
