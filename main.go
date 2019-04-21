package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/go-chi/jwtauth"
	"github.com/leoryu/leo-ryu.herokuapp.com/config"
	"github.com/leoryu/leo-ryu.herokuapp.com/handler/api"
	"github.com/leoryu/leo-ryu.herokuapp.com/service/confimpl/user"
	"github.com/leoryu/leo-ryu.herokuapp.com/store/mongoimpl/paper"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(config.GetDatabaseAddr()))
	if err != nil {
		log.Fatalln("Failed to connect DB")
	}
	paperStore := paper.New(mongoClient, config.GetDatabaseName(), "paper")
	userService := user.New(config.GetUsername(), config.GetPassword(), config.GetSecret(), 72)
	apiServer := api.New(paperStore, userService)
	tokenAuth := jwtauth.New("HS256", []byte(config.GetSecret()), nil)
	router := provideRouter(apiServer, tokenAuth)
	api.FileServer(router, "/", config.GetStaticFilesPath())
	log.Fatal(http.ListenAndServe(":7777", router))
}

func provideRouter(api api.Server, tokenAuth *jwtauth.JWTAuth) chi.Router {
	r := chi.NewRouter()
	r.Mount("/api", api.Handler(tokenAuth))
	return r
}
