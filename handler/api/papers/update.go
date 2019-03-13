package papers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/leoryu/leo-ryu.herokuapp.com/handler/api/render"

	"github.com/go-chi/chi"

	"github.com/leoryu/leo-ryu.herokuapp.com/core"
)

func HandleUpdate(s core.PaperStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := chi.URLParam(r, "id")
		in := new(paperAPI)
		if err := json.NewDecoder(r.Body).Decode(in); err != nil {
			render.BadRequest(w, err)
			return
		}
		paper := new(core.Paper)
		if err := in.assembleInput(paper); err != nil {
			render.BadRequest(w, err)
			return
		}
		paper.ID = id
		paper.UpdatedAt = time.Now().Unix()
		err := s.Update(ctx, paper)
		if err != nil {
			render.InternalError(w, err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

