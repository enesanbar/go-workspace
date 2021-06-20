package main

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (a *application) getAllGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := a.models.DB.GetGenres()
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	a.writeJSON(w, http.StatusOK, genres, "genres")
}

func (a *application) getAllMoviesByGenre(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	genreID, err := strconv.Atoi(params.ByName("genre_id"))
	if err != nil {
		return
	}

	movies, err := a.models.DB.All(genreID)
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	a.writeJSON(w, http.StatusOK, movies, "movies")
}
