package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/graphql-go/graphql"

	"github.com/enesanbar/workspace/golang/projects/movies/backend/models"
)

var movies []*models.Movie
var movieType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Movie",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"year": &graphql.Field{
			Type: graphql.Int,
		},
		"release_date": &graphql.Field{
			Type: graphql.DateTime,
		},
		"runtime": &graphql.Field{
			Type: graphql.Int,
		},
		"rating": &graphql.Field{
			Type: graphql.Int,
		},
		"mpaa_rating": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"updated_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"poster": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// graphql schema definition
var fields = graphql.Fields{
	"movie": &graphql.Field{
		Name: "",
		Type: movieType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type:         graphql.Int,
				DefaultValue: nil,
				Description:  "",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, ok := p.Args["id"].(int)
			if ok {
				for _, movie := range movies {
					if movie.ID == id {
						return movie, nil
					}
				}
			}
			return nil, nil
		},
		Description: "Get movie by id",
	},
	"list": &graphql.Field{
		Type: graphql.NewList(movieType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return movies, nil
		},
		Description: "Get all movies",
	},
	"search": &graphql.Field{
		Type: graphql.NewList(movieType),
		Args: graphql.FieldConfigArgument{
			"titleContains": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var result []*models.Movie
			search, ok := p.Args["titleContains"].(string)
			if ok {
				for _, movie := range movies {
					if strings.Contains(strings.ToLower(movie.Title), strings.ToLower(search)) {
						result = append(result, movie)
					}
				}
			}
			return result, nil
		},
		Description: "Search movies by title",
	},
}

func (a *application) moviesListGraphQL(w http.ResponseWriter, r *http.Request) {

	movies, _ = a.models.DB.All()
	q, _ := io.ReadAll(r.Body)
	query := string(q)

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: fields,
		}),
	})
	if err != nil {
		a.errorJSON(w, errors.New("failed to create schema"))
		a.logger.Println(err)
		return
	}

	params := graphql.Params{
		Schema:        schema,
		RequestString: query,
	}
	resp := graphql.Do(params)
	if resp.HasErrors() {
		a.errorJSON(w, fmt.Errorf("failed %+v", resp.Errors))
		return
	}

	j, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
