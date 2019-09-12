package cockroach

import (
	"database/sql"

	"github.com/joshvoll/branca"
	"github.com/joshvoll/tinyuser/internal/user"
	"github.com/pkg/errors"

	_ "github.com/jackc/pgx/stdlib" // pgx postgres library
)

// Client reprenset the client of the cockroach database
type Client struct {
	// databaseURL usually is a path to the db
	databaseURL string
	// point to the user service
	service UserService
	// database
	db *sql.DB
	// token going to use branca
	codec *branca.Branca
}

// UserService manager the user
type UserService struct {
	client *Client
}

// New contructor function to get the database
func New(databaseURL string, codec *branca.Branca) *Client {
	c := &Client{
		databaseURL: databaseURL,
		codec:       codec,
	}

	c.service.client = c

	return c
}

// Open just open the db connection base on the configuration
func (c *Client) Open() error {
	// open the database
	db, err := sql.Open("pgx", c.databaseURL)
	if err != nil {
		return errors.Wrap(err, "OPEN DATABASE")
	}
	// ping the database
	if err := db.Ping(); err != nil {
		return errors.Wrap(err, "PING DATABASE")
	}

	c.db = db

	return nil
}

// Close just close the db
func (c *Client) Close() error {
	return c.db.Close()
}

// Service return the user service asociated with the client
func (c *Client) Service() user.Service {
	return &c.service
}
