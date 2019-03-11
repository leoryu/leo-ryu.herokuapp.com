package paper

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/leoryu/leo-ryu.herokuapp.com/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	paper struct {
		ID        primitive.ObjectID `bson:"_id"`
		Title     string             `bson:"title"`
		Tags      map[string]bool    `bson:"tags"`
		Abstract  string             `bson:"abstract"`
		Content   string             `bson:"content"`
		CreatedAt int64              `bson:"created_at"`
		UpdatedAt int64              `bson:"updated_at"`
	}

	paperStore struct {
		Client     *mongo.Client
		Database   string
		Collection string
	}
)

func New(client *mongo.Client, database, collection string) core.PaperStore {
	return &paperStore{
		Client:     client,
		Database:   database,
		Collection: collection,
	}
}

func (store *paperStore) Create(ctx context.Context, paper *core.Paper) error {
	collection := store.Client.Database(store.Database).Collection(store.Collection)
	_, err := collection.InsertOne(ctx, paper)
	return err
}

func (store *paperStore) Update(ctx context.Context, paper *core.Paper) error {
	_id, err := primitive.ObjectIDFromHex(paper.ID)
	if err != nil {
		return err
	}
	collection := store.Client.Database(store.Database).Collection(store.Collection)
	_, err = collection.UpdateOne(ctx, bson.M{"_id": _id}, bson.M{"$set": paper})
	return err
}

func (store *paperStore) Delete(ctx context.Context, id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	collection := store.Client.Database(store.Database).Collection(store.Collection)
	_, err = collection.DeleteOne(ctx, bson.M{"_id": _id})
	return err
}

func (store *paperStore) Find(ctx context.Context, id string) (*core.Paper, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var paper core.Paper
	collection := store.Client.Database(store.Database).Collection(store.Collection)
	err = collection.FindOne(ctx, bson.M{"_id": _id}).Decode(&paper)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &paper, err
}

func (store *paperStore) List(ctx context.Context, limit, page int) ([]*core.Paper, error) {
	return store.ListByTag(ctx, "", limit, page)
}

func (store *paperStore) ListByTag(ctx context.Context, tag string, limit, page int) ([]*core.Paper, error) {
	collection := store.Client.Database(store.Database).Collection(store.Collection)
	var filter bson.M
	if tag == "" {
		filter = nil
	} else {
		filter = bson.M{"tags." + tag: true}
	}
	options := options.Find()
	options.SetSort(bson.M{"_id": -1})
	options.SetLimit(int64(limit))
	options.SetSkip(int64(limit * (page - 1)))
	options.SetProjection(bson.M{"content": 0})
	cursor, err := collection.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var papers []*core.Paper
	for cursor.Next(ctx) {
		var p core.Paper
		var _p paper
		err := cursor.Decode(&_p)
		if err != nil {
			return nil, err
		}
		p.ID = _p.ID.Hex()
		p.Title = _p.Title
		p.Tags = _p.Tags
		p.Abstract = _p.Abstract
		p.CreatedAt = _p.CreatedAt
		p.UpdatedAt = _p.UpdatedAt
		papers = append(papers, &p)
	}
	err = cursor.Err()
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return papers, err
}
