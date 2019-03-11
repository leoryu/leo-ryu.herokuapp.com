package mongoimpl

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewConnect(dbAddr string) *mongo.Client {
	c, err := mongo.NewClient(options.Client().ApplyURI(dbAddr))
	if err != nil {
		log.Fatal(err)
	}
	return c
}
