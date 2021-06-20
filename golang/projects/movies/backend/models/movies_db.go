package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Repository struct {
	DB *sql.DB
}

func (d *Repository) Get(id int) (*Movie, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	query := `
		select
			id, title, description, year, release_date, 
		   	rating, runtime, mpaa_rating, created_at, updated_at, coalesce(poster, '')
		from
			movies
		where id = $1
	`

	row := d.DB.QueryRowContext(ctx, query, id)
	var movie Movie
	err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.Year,
		&movie.ReleaseDate,
		&movie.Rating,
		&movie.Runtime,
		&movie.MPAARating,
		&movie.CreatedAt,
		&movie.UpdatedAt,
		&movie.Poster,
	)
	if err != nil {
		return nil, err
	}

	// get genres
	query = `
		select
			mg.id, mg.movie_id, mg.genre_id, g.genre_name
		from
			movies_genres mg
		left join genres g on (g.id = mg.genre_id)
		where mg.movie_id = $1
	`
	rows, err := d.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			fmt.Println("unable to close rows")
		}
	}()

	genres := map[int]string{}
	for rows.Next() {
		var genre MovieGenre
		err := rows.Scan(
			&genre.ID,
			&genre.MovieID,
			&genre.GenreID,
			&genre.Genre.GenreName,
		)
		if err != nil {
			return nil, err
		}

		genres[genre.ID] = genre.Genre.GenreName
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	movie.MovieGenre = genres

	return &movie, nil
}

func (d *Repository) All(genres ...int) ([]*Movie, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	where := ""
	if len(genres) > 0 {
		where = fmt.Sprintf("where id in (select movie_id from movies_genres where genre_id = %d)", genres[0])
	}

	query := fmt.Sprintf(`
		select
			id, title, description, year, release_date, rating, 
			runtime, mpaa_rating, created_at, updated_at, coalesce(poster, '')
		from
			movies
		%s
		order by title
	`, where)

	rows, err := d.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*Movie
	for rows.Next() {
		var movie Movie
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.Year,
			&movie.ReleaseDate,
			&movie.Rating,
			&movie.Runtime,
			&movie.MPAARating,
			&movie.CreatedAt,
			&movie.UpdatedAt,
			&movie.Poster,
		)
		if err != nil {
			return nil, err
		}

		// get genres
		genreQuery := `
		select
			mg.id, mg.movie_id, mg.genre_id, g.genre_name
		from
			movies_genres mg
		left join genres g on (g.id = mg.genre_id)
		where mg.movie_id = $1
	`
		genreRows, err := d.DB.QueryContext(ctx, genreQuery, movie.ID)
		if err != nil {
			return nil, err
		}

		genres := map[int]string{}
		for genreRows.Next() {
			var genre MovieGenre
			err := genreRows.Scan(
				&genre.ID,
				&genre.MovieID,
				&genre.GenreID,
				&genre.Genre.GenreName,
			)
			if err != nil {
				return nil, err
			}

			genres[genre.ID] = genre.Genre.GenreName
		}
		if err := rows.Err(); err != nil {
			return nil, err
		}

		err = genreRows.Close()
		if err != nil {
			return nil, err
		}

		movie.MovieGenre = genres

		movies = append(movies, &movie)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}

func (d *Repository) GetGenres() ([]*Genre, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	query := `
		select
			id, genre_name, created_at, updated_at
		from
			genres
		order by genre_name
	`

	rows, err := d.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []*Genre

	for rows.Next() {
		var genre Genre
		err := rows.Scan(
			&genre.ID,
			&genre.GenreName,
			&genre.CreatedAt,
			&genre.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		genres = append(genres, &genre)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return genres, err
}

func (d *Repository) InsertMovie(movie Movie) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	query := `
		insert into movies
			(title, description, year, release_date, runtime, 
			 rating, mpaa_rating, created_at, updated_at, poster)
		values
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err := d.DB.ExecContext(ctx, query,
		movie.Title,
		movie.Description,
		movie.Year,
		movie.ReleaseDate,
		movie.Runtime,
		movie.Rating,
		movie.MPAARating,
		time.Now(),
		time.Now(),
		movie.Poster,
	)
	if err != nil {
		return err
	}

	return nil
}

func (d *Repository) UpdateMovie(movie Movie) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	query := `
		update movies set
			title = $1, 
		    description = $2, 
		    year = $3, 
		    release_date = $4, 
		    runtime = $5, 
		    rating = $6, 
		    mpaa_rating = $7, 
		    updated_at = $8,
			poster = $9
		where id = $10
	`

	_, err := d.DB.ExecContext(ctx, query,
		movie.Title,
		movie.Description,
		movie.Year,
		movie.ReleaseDate,
		movie.Runtime,
		movie.Rating,
		movie.MPAARating,
		time.Now(),
		movie.Poster,
		movie.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (d *Repository) DeleteMovie(id int) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	query := `
		delete from movies
		where id = $1
	`

	_, err := d.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
