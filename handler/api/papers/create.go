package papers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/leoryu/leo-ryu.herokuapp.com/core"
	"github.com/leoryu/leo-ryu.herokuapp.com/handler/api/render"
)

func HandleCreate(s core.PaperStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		paper.CreatedAt = time.Now().Unix()
		paper.UpdatedAt = paper.CreatedAt
		if err := s.Create(r.Context(), paper); err != nil {
			render.InternalError(w, err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

