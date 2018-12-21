package store

import (
	"context"

	"github.com/leoryu/leo-ryu.herokuapp.com/model"
)

//go:generate mockgen -destination store_mock.go -package store github.com/leoryu/leo-ryu.herokuapp.com/store Store

type Store interface {
	SavePaper(paper *model.Paper) error
	ModifyPaper(paper *model.Paper, id string) (int, error)
	DeletePaper(id string) (int, error)
	GetPaper(id string) (*model.Paper, error)
	GetIntroductions(subject string, limit, page int) (count int, introductions []*model.IntroductionWithID, err error)
}

func SavePaper(c context.Context, paper *model.Paper) error {
	return FromContext(c).SavePaper(paper)

}

func ModifyPaper(c context.Context, paper *model.Paper, id string) (int, error) {
	return FromContext(c).ModifyPaper(paper, id)
}

func DeletePaper(c context.Context, id string) (int, error) {
	return FromContext(c).DeletePaper(id)
}

func GetPaper(c context.Context, id string) (*model.Paper, error) {
	return FromContext(c).GetPaper(id)
}

func GetIntroductions(c context.Context, subject string, limit, page int) (count int, introductions []*model.IntroductionWithID, err error) {
	return FromContext(c).GetIntroductions(subject, limit, page)
}

