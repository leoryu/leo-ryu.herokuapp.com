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
	paperMongo struct {
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
	if err := store.Client.Connect(ctx); err != nil {
		return err
	}
	p := new(paperMongo)
	p.assembleWriter(paper)
	collection := store.Client.Database(store.Database).Collection(store.Collection)
	_, err := collection.InsertOne(ctx, p)
	return err
}

func (store *paperStore) Update(ctx context.Context, paper *core.Paper) error {
	if err := store.Client.Connect(ctx); err != nil {
		return err
	}
	p := new(paperMongo)
	if err := p.assembleWriter(paper); err != nil {
		return err
	}
	collection := store.Client.Database(store.Database).Collection(store.Collection)
	_, err := collection.UpdateOne(ctx, bson.M{"_id": p.ID}, bson.M{"$set": p})
	return err
}

func (store *paperStore) Delete(ctx context.Context, id string) error {
	if err := store.Client.Connect(ctx); err != nil {
		return err
	}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	collection := store.Client.Database(store.Database).Collection(store.Collection)
	_, err = collection.DeleteOne(ctx, bson.M{"_id": _id})
	return err
}

func (store *paperStore) Find(ctx context.Context, id string) (*core.Paper, error) {
	if err := store.Client.Connect(ctx); err != nil {
		return nil, err
	}
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
	err = p.assembleReader(paper)
	return paper, err
}

func (store *paperStore) List(ctx context.Context, limit, page int) ([]*core.Paper, error) {
	return store.ListByTag(ctx, "", limit, page)
}

func (store *paperStore) ListByTag(ctx context.Context, tag string, limit, page int) ([]*core.Paper, error) {
	if err := store.Client.Connect(ctx); err != nil {
		return nil, err
	}
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
		if err := p.assembleReader(paper); err != nil {
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

func (pm *paperMongo) assembleWriter(p *core.Paper) (err error) {
	if p.ID == "" {
		pm.ID = primitive.NewObjectID()
	}
	pm.ID, err = primitive.ObjectIDFromHex(p.ID)
	if err != nil {
		return
	}
	pm.Abstract = p.Abstract
	pm.Content = p.Content
	pm.CreatedAt = p.CreatedAt
	pm.Tags = p.Tags
	pm.Title = p.Title
	pm.UpdatedAt = p.UpdatedAt
	return
}

func (pm *paperMongo) assembleReader(p *core.Paper) (err error) {
	p.ID = pm.ID.Hex()
	p.Abstract = pm.Abstract
	p.Content = pm.Content
	p.CreatedAt = pm.CreatedAt
	p.Tags = pm.Tags
	p.Title = pm.Title
	p.UpdatedAt = pm.UpdatedAt
	return
}

