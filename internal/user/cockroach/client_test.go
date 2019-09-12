package cockroach_test

import (
	_ "github.com/jackc/pgx/stdlib" // just another lib
	"github.com/joshvoll/branca"
	"github.com/joshvoll/tinyuser/internal/user/cockroach"
)

// Client is the testing wrrape of the database client
type Client struct {
	*cockroach.Client
}

func NewClient() *Client {

	codec := branca.NewBranca("supersecretkeyyoushouldnotcommit")
	c := &Client{
		Client: cockroach.New("postgresql://root@127.0.0.1:26257/tinyuser?sslmode=disable", codec),
	}

	return c
}

func MustOpenClient() *Client {
	c := NewClient()
	if err := c.Open(); err != nil {
		panic(err)
	}

	return c
}

func (c *Client) Close() error {
	return c.Client.Close()
}
