package user

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/leoryu/leo-ryu.herokuapp.com/core"
)

type userService struct {
	Username   string
	Password   string
	Secret     string
	ExpireHour int
}

func New(username, password, secret string, expireHour int) core.UserService {
	return &userService{
		Username:   username,
		Password:   password,
		Secret:     secret,
		ExpireHour: expireHour,
	}
}

func (s *userService) Verify(ctx context.Context, user *core.User) (*core.Token, error) {
	if s.Username != user.Username ||
		s.Password != user.Password {
		return nil, nil
	}
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = jwt.MapClaims{
		"name": s.Username,
		"exp":  time.Now().Add(time.Duration(s.ExpireHour) * time.Hour).Unix(),
	}
	var err error
	t := new(core.Token)
	t.TokenString, err = token.SignedString([]byte(s.Secret))
	return t, err
}
