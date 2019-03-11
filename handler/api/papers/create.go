package papers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/leoryu/leo-ryu.herokuapp.com/core"
	"github.com/leoryu/leo-ryu.herokuapp.com/handler/api/render"
)

type (
	createInput struct {
		Title    *string          `json:"title"`
		Tags     *map[string]bool `json:"tags"`
		Abstract *string          `json:"abstract"`
		Content  *string          `json:"content"`
	}
)

func HandleCreate(s core.PaperStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		in := new(updateInput)
		if err := json.NewDecoder(r.Body).Decode(in); err != nil {
			render.BadRequest(w, err)
			return
		}
		paper := new(core.Paper)
		if err := in.assemble(paper); err != nil {
			render.BadRequest(w, err)
			return
		}
		paper.CreatedAt = time.Now().Unix()
		paper.UpdatedAt = paper.CreatedAt
		if err := s.Create(r.Context(), paper); err != nil {
			render.InternalError(w, err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func (in *createInput) assemble(paper *core.Paper) error {
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
