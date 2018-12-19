package datastore

import (
	"context"
	"log"
	"time"

	"github.com/leoryu/leo-ryu.herokuapp.com/config"
	"github.com/leoryu/leo-ryu.herokuapp.com/model"
	"github.com/leoryu/leo-ryu.herokuapp.com/store"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type datastore struct {
	Client     *mongo.Client
	Database   string
	Collection string
}

func New() store.Store {
	return &datastore{
		Client:     new(),
		Database:   config.GetDatabaseName(),
		Collection: "paper",
	}
}

func (d *datastore) SavePaper(paper *model.Paper) error {
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	collection := d.Client.Database(d.Database).Collection(d.Collection)
	_, err := collection.InsertOne(ctx, paper)
	cancle()
	return err
}

func (d *datastore) ModifyPaper(paper *model.Paper, id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	collection := d.Client.Database(d.Database).Collection(d.Collection)
	_, err = collection.UpdateOne(ctx, bson.M{"_id": _id}, bson.M{"$set": paper})
	cancle()
	return err
}

func (d *datastore) GetPaper(id string) (*model.Paper, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var paper model.Paper
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	collection := d.Client.Database(d.Database).Collection(d.Collection)
	err = collection.FindOne(ctx, bson.M{"_id": _id}).Decode(&paper)
	cancle()
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &paper, err
}

func (d *datastore) DeletePaper(id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	collection := d.Client.Database(d.Database).Collection(d.Collection)
	_, err = collection.DeleteOne(ctx, bson.M{"_id": _id})
	cancle()
	return err
}

func new() *mongo.Client {
	c, err := mongo.Connect(context.TODO(), config.GetDatabaseAddr())
	if err != nil {
		log.Fatal(err)
	}
	return c
}
