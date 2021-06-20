package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/enesanbar/workspace/golang/projects/movies/backend/models"

	"github.com/julienschmidt/httprouter"
)

type jsonResp struct {
	OK      bool   `json:"ok,omitempty"`
	Message string `json:"message"`
}

func (a *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		a.logger.Println(errors.New("invalid id parameter"))
		a.errorJSON(w, err)
		return
	}

	movie, err := a.models.DB.Get(id)
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	a.writeJSON(w, http.StatusOK, movie, "movie")
}

func (a *application) getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := a.models.DB.All()
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	a.writeJSON(w, http.StatusOK, movies, "movies")
}

func (a *application) deleteMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		a.logger.Println(errors.New("invalid id parameter"))
		a.errorJSON(w, err)
		return
	}

	err = a.models.DB.DeleteMovie(id)
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	a.writeJSON(w, http.StatusOK, jsonResp{
		OK:      true,
		Message: "Deleted movie",
	}, "response")
}

type MoviePayload struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        string `json:"year"`
	ReleaseDate string `json:"release_date"`
	Runtime     string `json:"runtime"`
	Rating      string `json:"rating"`
	MPAARating  string `json:"mpaa_rating"`
}

func (a *application) editMovie(w http.ResponseWriter, r *http.Request) {
	var payload MoviePayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		a.errorJSON(w, err)
		return
	}

	movieID, _ := strconv.Atoi(payload.ID)
	releaseDate, err := time.Parse("2006-01-02", payload.ReleaseDate)
	if err != nil {
		a.logger.Println(err)
		a.errorJSON(w, errors.New("unable to parse release date"))
		return
	}
	runtime, _ := strconv.Atoi(payload.Runtime)
	rating, _ := strconv.Atoi(payload.Rating)
	movie := models.Movie{
		ID:          movieID,
		Title:       payload.Title,
		Description: payload.Description,
		Year:        releaseDate.Year(),
		ReleaseDate: releaseDate,
		Runtime:     runtime,
		Rating:      rating,
		MPAARating:  payload.MPAARating,
	}

	if movie.Poster == "" {
		movie = getPoster(movie)
	}

	if movie.ID == 0 {
		err = a.models.DB.InsertMovie(movie)
		if err != nil {
			a.errorJSON(w, err)
			return
		}
	} else {
		err = a.models.DB.UpdateMovie(movie)
		if err != nil {
			a.errorJSON(w, err)
			return
		}
	}

	a.writeJSON(w, http.StatusOK, jsonResp{
		OK:      true,
		Message: "",
	}, "response")
}

func getPoster(movie models.Movie) models.Movie {
	type TheMovieDB struct {
		Page    int `json:"page"`
		Results []struct {
			Adult            bool    `json:"adult"`
			BackdropPath     string  `json:"backdrop_path"`
			GenreIds         []int   `json:"genre_ids"`
			ID               int     `json:"id"`
			OriginalLanguage string  `json:"original_language"`
			OriginalTitle    string  `json:"original_title"`
			Overview         string  `json:"overview"`
			Popularity       float64 `json:"popularity"`
			PosterPath       string  `json:"poster_path"`
			ReleaseDate      string  `json:"release_date"`
			Title            string  `json:"title"`
			Video            bool    `json:"video"`
			VoteAverage      float64 `json:"vote_average"`
			VoteCount        int     `json:"vote_count"`
		} `json:"results"`
		TotalPages   int `json:"total_pages"`
		TotalResults int `json:"total_results"`
	}

	client := http.Client{}
	key := "ffdd68f00915b2590f5abda40d96d0e2"
	posterUrl := fmt.Sprintf(
		"https://api.themoviedb.org/3/search/movie?api_key=%s&query=%s",
		key, url.QueryEscape(movie.Title))
	fmt.Println(posterUrl)

	request, err := http.NewRequest("GET", posterUrl, nil)
	if err != nil {
		log.Println(err)
		return models.Movie{}
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		return models.Movie{}
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return models.Movie{}
	}

	var responseObject TheMovieDB
	err = json.Unmarshal(bodyBytes, &responseObject)
	if err != nil {
		log.Println(err)
		return models.Movie{}
	}

	if len(responseObject.Results) > 0 {
		movie.Poster = responseObject.Results[0].PosterPath
	}

	return movie
}
