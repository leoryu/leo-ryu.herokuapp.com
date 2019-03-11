package papers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/leoryu/leo-ryu.herokuapp.com/handler/api/render"

	"github.com/go-chi/chi"

	"github.com/leoryu/leo-ryu.herokuapp.com/core"
)

type (
	updateInput struct {
		Title    *string          `json:"title"`
		Tags     *map[string]bool `json:"tags"`
		Abstract *string          `json:"abstract"`
		Content  *string          `json:"content"`
	}
)

func HandleUpdate(s core.PaperStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := chi.URLParam(r, "id")
		in := new(updateInput)
		if err := json.NewDecoder(r.Body).Decode(in); err != nil {
			render.BadRequest(w, err)
			return
		}
		paper, err := s.Find(ctx, id)
		if err != nil {
			render.InternalError(w, err)
			return
		}
		if paper == nil {
			render.NotFound(w, errors.New(id+" doesn't exists"))
			return
		}
		if err = in.assemble(paper); err != nil {
			render.BadRequest(w, err)
			return
		}
		paper.UpdatedAt = time.Now().Unix()
		err = s.Update(ctx, paper)
		if err != nil {
			render.InternalError(w, err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func (in *updateInput) assemble(paper *core.Paper) error {
	if in.Title != nil {
		if *in.Title == "" {
			return errors.New("Title is required")
		}
		paper.Title = *in.Title
	}
	if in.Tags != nil {
		paper.Tags = *in.Tags
	}
	if in.Abstract != nil {
		paper.Abstract = *in.Abstract
	}
	if in.Content != nil {
		if *in.Content == "" {
			return errors.New("Content is required")
		}
		paper.Content = *in.Content
	}
	return nil
}
