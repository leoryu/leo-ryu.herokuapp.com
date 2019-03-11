package user

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/leoryu/leo-ryu.herokuapp.com/handler/api/render"

	"github.com/leoryu/leo-ryu.herokuapp.com/core"
)

type (
	verifyInput struct {
		Username *string `json:"username"`
		Password *string `json:"password"`
	}
)

func HandleVrify(s core.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		in := new(verifyInput)
		if err := json.NewDecoder(r.Body).Decode(in); err != nil {
			render.BadRequest(w, err)
			return
		}
		user := new(core.User)
		if err := in.assemble(user); err != nil {
			render.BadRequest(w, err)
			return
		}
		out, err := s.Verify(r.Context(), user)
		if err != nil {
			render.InternalError(w, err)
			return
		}
		if out == nil {
			render.Unauthorized(w, errors.New("You are not owner"))
			return
		}
		render.JSON(w, out, http.StatusOK)
	}
}

func (in *verifyInput) assemble(user *core.User) error {
	if in.Username == nil || *in.Username == "" {
		return errors.New("Username is required")
	}
	if in.Password == nil || *in.Password == "" {
		return errors.New("Password is required")
	}
	user.Username = *in.Username
	user.Password = *in.Password
	return nil
}
