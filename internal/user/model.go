package user

import "time"

// User model definition
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

// LoginOutput model definition
type LoginOutput struct {
	Token    string    `json:"token"`
	ExpireAt time.Time `json:"expireat"`
	AuthUser User      `json:"authuser"`
}
