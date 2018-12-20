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
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

type datastore struct {
	Client     *mongo.Client
	Database   string
	Collection string
}

type introductionWithID struct {
	ID           primitive.ObjectID `bson:"_id"`
	Introduction model.Introduction `bson:"introduction"`
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

func (d *datastore) ModifyPaper(paper *model.Paper, id string) (int, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	collection := d.Client.Database(d.Database).Collection(d.Collection)
	result, err := collection.UpdateOne(ctx, bson.M{"_id": _id}, bson.M{"$set": paper})
	cancle()
	return int(result.MatchedCount), err
}

func (d *datastore) DeletePaper(id string) (int, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	collection := d.Client.Database(d.Database).Collection(d.Collection)
	result, err := collection.DeleteOne(ctx, bson.M{"_id": _id})
	cancle()
	return int(result.DeletedCount), err
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

func (d *datastore) GetIntroductions(limit, page int) (count int, introductions []*model.IntroductionWithID, err error) {
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	collection := d.Client.Database(d.Database).Collection(d.Collection)
	_count, err := collection.Count(ctx, nil)
	count = int(_count)
	options := options.Find()
	options.SetSort(bson.M{"_id": -1})
	options.SetLimit(int64(limit))
	options.SetSkip(int64(limit * (page - 1)))
	options.SetProjection(bson.M{"_id": 1, "introduction": 1})
	cursor, err := collection.Find(ctx, nil, options)
	if err != nil {
		return 0, nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var introduction model.IntroductionWithID
		var _introduction introductionWithID
		err := cursor.Decode(&_introduction)
		if err != nil {
			return 0, nil, err
		}
		introduction.ID = _introduction.ID.Hex()
		introduction.Introduction = _introduction.Introduction
		introductions = append(introductions, &introduction)
	}
	err = cursor.Err()
	if err == mongo.ErrNoDocuments {
		return 0, nil, nil
	}
	return count, introductions, err
}

func new() *mongo.Client {
	c, err := mongo.Connect(context.TODO(), config.GetDatabaseAddr())
	if err != nil {
		log.Fatal(err)
	}
	return c
}

