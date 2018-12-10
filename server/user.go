package server

import (
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/leoryu/leo-ryu.herokuapp.com/model"
)

var secret = os.Getenv("SECRET")

// Login is the handler of "/login".
func Login(c *gin.Context) {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = jwt.MapClaims{
		"name": "Leo Ryu",
		"exp":  time.Now().Add(72 * time.Hour).Unix(),
	}
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
