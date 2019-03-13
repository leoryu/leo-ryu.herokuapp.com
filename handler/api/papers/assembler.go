package papers

import (
	"errors"

	"github.com/leoryu/leo-ryu.herokuapp.com/core"
)

type (
	paperAPI struct {
		Title    *string          `json:"title"`
		Tags     *map[string]bool `json:"tags"`
		Abstract *string          `json:"abstract"`
		Content  *string          `json:"content"`
	}
)

func (in *paperAPI) assembleInput(paper *core.Paper) error {
	if in.Title == nil {
		return errors.New("Title is required")
	}
	if in.Content == nil {
		return errors.New("Content is required")
	}
	paper.Title = *in.Title
	paper.Content = *in.Content
	if in.Tags != nil {
		paper.Tags = *in.Tags
	}
	if in.Abstract != nil {
		paper.Abstract = *in.Abstract
	}
	return nil
}

