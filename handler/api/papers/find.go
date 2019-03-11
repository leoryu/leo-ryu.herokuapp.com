package papers

import (
	"errors"
	"net/http"

	"github.com/leoryu/leo-ryu.herokuapp.com/handler/api/render"

	"github.com/go-chi/chi"

	"github.com/leoryu/leo-ryu.herokuapp.com/core"
)

func HandleFind(s core.PaperStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		paper, err := s.Find(r.Context(), id)
		if err != nil {
			render.InternalError(w, err)
			return
		}
		if paper == nil {
			render.NotFound(w, errors.New(id+" dosen't exists"))
			return
		}
		render.JSON(w, paper, http.StatusOK)
	}
}
