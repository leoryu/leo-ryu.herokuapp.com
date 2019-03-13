package paper

import (
	"github.com/leoryu/leo-ryu.herokuapp.com/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
)

func (pm *paperMongo) assembleInput(p *core.Paper) (err error) {
	if p.ID == "" {
		pm.ID = primitive.NewObjectID()
	} else {
		pm.ID, err = primitive.ObjectIDFromHex(p.ID)
		if err != nil {
			return err
		}
	}
	pm.Abstract = p.Abstract
	pm.Content = p.Content
	pm.CreatedAt = p.CreatedAt
	pm.Tags = p.Tags
	pm.Title = p.Title
	pm.UpdatedAt = p.UpdatedAt
	return
}

func (pm *paperMongo) assembleOutput(p *core.Paper) (err error) {
	p.ID = pm.ID.Hex()
	p.Abstract = pm.Abstract
	p.Content = pm.Content
	p.CreatedAt = pm.CreatedAt
	p.Tags = pm.Tags
	p.Title = pm.Title
	p.UpdatedAt = pm.UpdatedAt
	return
}

