package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/go-chi/jwtauth"
	"github.com/leoryu/leo-ryu.herokuapp.com/config"
	"github.com/leoryu/leo-ryu.herokuapp.com/handler/api"
	"github.com/leoryu/leo-ryu.herokuapp.com/service/confimpl/user"
	"github.com/leoryu/leo-ryu.herokuapp.com/store/mongoimpl"
	"github.com/leoryu/leo-ryu.herokuapp.com/store/mongoimpl/paper"
)

var tokenAuth *jwtauth.JWTAuth

func main() {
	mongoClient := mongoimpl.NewConnect(config.GetDatabaseAddr())
	paperStore := paper.New(mongoClient, config.GetDatabaseName(), "paper")
	userService := user.New(config.GetUsername(), config.GetPassword(), config.GetSecret(), 72)
	apiServer := api.New(paperStore, userService)
	tokenAuth = jwtauth.New("HS256", []byte(config.GetSecret()), nil)
	log.Fatal(http.ListenAndServe(":7777", provideRouter(apiServer, tokenAuth)))
}

func provideRouter(api api.Server, tokenAuth *jwtauth.JWTAuth) http.Handler {
	r := chi.NewRouter()
	r.Mount("/api", api.Handler(tokenAuth))
	return r
}
