package papers

import (
	"net/http"

	"github.com/leoryu/leo-ryu.herokuapp.com/handler/api/render"

	"github.com/go-chi/chi"
	"github.com/leoryu/leo-ryu.herokuapp.com/core"
)

func HandleDelete(s core.PaperStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		err := s.Delete(r.Context(), id)
		if err != nil {
			render.InternalError(w, err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
