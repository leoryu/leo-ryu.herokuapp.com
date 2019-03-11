package core

import "context"

type (
	User struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	Token struct {
		TokenString string `json:"token"`
	}

	UserService interface {
		Verify(ctx context.Context, user *User) (*Token, error)
	}
)
