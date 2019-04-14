package api

import (
	"net/http"

	"github.com/go-chi/jwtauth"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"github.com/leoryu/leo-ryu.herokuapp.com/core"
	"github.com/leoryu/leo-ryu.herokuapp.com/handler/api/papers"
	"github.com/leoryu/leo-ryu.herokuapp.com/handler/api/user"
)

type (
	Server struct {
		PaperStore  core.PaperStore
		UserService core.UserService
	}
)

func New(paper core.PaperStore, user core.UserService) Server {
	return Server{
		PaperStore:  paper,
		UserService: user,
	}
}

func (s Server) Handler(ja *jwtauth.JWTAuth) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Post("/verify", user.HandleVrify(s.UserService))

	r.Route("/papers", func(r chi.Router) {
		r.Get("/", papers.HandleList(s.PaperStore))
		r.Get("/{id}", papers.HandleFind(s.PaperStore))

		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(ja))
			r.Post("/", papers.HandleCreate(s.PaperStore))
			r.Put("/{id}", papers.HandleUpdate(s.PaperStore))
			r.Delete("/{id}", papers.HandleDelete(s.PaperStore))
		})
	})

	return r
}
