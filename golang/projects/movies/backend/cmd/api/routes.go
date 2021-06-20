package main

import (
	"context"
	"net/http"

	"github.com/justinas/alice"

	"github.com/julienschmidt/httprouter"
)

func (a *application) wrap(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), httprouter.ParamsKey, ps)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (a *application) routes() http.Handler {
	router := httprouter.New()

	secure := alice.New(a.checkToken)

	router.HandlerFunc(http.MethodGet, "/status", a.statusHandler)

	router.HandlerFunc(http.MethodPost, "/v1/signin", a.Signin)

	router.HandlerFunc(http.MethodGet, "/v1/movies", a.getAllMovies)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", a.getOneMovie)

	router.DELETE("/v1/movies/:id", a.wrap(secure.ThenFunc(a.deleteMovie)))
	router.HandlerFunc(http.MethodGet, "/v1/genres", a.getAllGenres)
	router.HandlerFunc(http.MethodGet, "/v1/genres/:genre_id", a.getAllMoviesByGenre)
	router.POST("/v1/admin/editmovie", a.wrap(secure.ThenFunc(a.editMovie)))

	router.HandlerFunc(http.MethodPost, "/v1/graphql", a.moviesListGraphQL)
	return a.enableCORS(router)
}
