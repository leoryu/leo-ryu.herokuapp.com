package server

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/leoryu/leo-ryu.herokuapp.com/config"
	"github.com/leoryu/leo-ryu.herokuapp.com/model"
)

// Login is the handler of "/login".
func Login(c *gin.Context) {
	user := new(model.User)
	if err := c.ShouldBind(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.Email != config.GetEmail() || user.Password != config.GetPassword() {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong email or password"})
		return
	}
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = jwt.MapClaims{
		"name": "Leo Ryu",
		"exp":  time.Now().Add(72 * time.Hour).Unix(),
	}
	tokenString, err := token.SignedString([]byte(config.GetSecret()))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

