package server

import (
	"net/http"
	"os"
	"time"

	"github.com/leoryu/leo-ryu.herokuapp.com/model"

	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo"
)

// Login is the handler of "/login".
func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == os.Getenv("USERNAME") && password == os.Getenv("PASSWORD") {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Leo Ryu"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString(model.Secret)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}