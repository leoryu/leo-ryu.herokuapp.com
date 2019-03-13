package papers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/leoryu/leo-ryu.herokuapp.com/handler/api/render"

	"github.com/leoryu/leo-ryu.herokuapp.com/core"
)

func HandleList(s core.PaperStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tag := r.URL.Query().Get("tag")
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			render.BadRequest(w, err)
			return
		}
		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			render.BadRequest(w, err)
			return
		}
		papers, err := s.List(r.Context(), tag, limit, page)
		if err != nil {
			render.InternalError(w, err)
			return
		}
		if papers == nil {
			render.NotFound(w, errors.New("No record"))
			return
		}
		render.JSON(w, papers, http.StatusOK)
	}
}

