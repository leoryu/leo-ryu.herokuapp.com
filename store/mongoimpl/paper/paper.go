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
	p := new(paperMongo)
	p.assembleInput(paper)
	collection := store.Client.Database(store.Database).Collection(store.Collection)
	_, err := collection.InsertOne(ctx, p)
	return err
}

func (store *paperStore) Update(ctx context.Context, paper *core.Paper) error {
	p := new(paperMongo)
	if err := p.assembleInput(paper); err != nil {
		return err
	}
	collection := store.Client.Database(store.Database).Collection(store.Collection)
	_, err := collection.UpdateOne(ctx, bson.M{"_id": p.ID}, bson.M{"$set": p})
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
	p := new(paperMongo)
	collection := store.Client.Database(store.Database).Collection(store.Collection)
	err = collection.FindOne(ctx, bson.M{"_id": _id}).Decode(p)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	paper := new(core.Paper)
	err = p.assembleOutput(paper)
	return paper, err
}

func (store *paperStore) List(ctx context.Context, tag string, limit, page int) ([]*core.Paper, error) {
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
		paper := new(core.Paper)
		p := new(paperMongo)
		if err := cursor.Decode(p); err != nil {
			return nil, err
		}
		if err := p.assembleOutput(paper); err != nil {
			return nil, err
		}
		papers = append(papers, paper)
	}
	err = cursor.Err()
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return papers, err
}

