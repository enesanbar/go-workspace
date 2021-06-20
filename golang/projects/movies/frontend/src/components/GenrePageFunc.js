import React, {Fragment, useEffect, useState} from "react";
import {Link} from "react-router-dom";

function GenrePageFunc(props) {
    let [movies, setMovies] = useState([]);
    const [error, setError] = useState(null)
    let [genreName, setGenreName] = useState("")

    useEffect(() => {
        fetch(`${process.env.REACT_APP_API_URL}/v1/genres/` + props.match.params.id)
            .then((response) => {
                console.log(response.status);
                if (response.status !== 200) {
                    setError("invalid response code: " + response.status)
                } else {
                    setError(null)
                }
                return response.json()
            })
            .then((json) => {
                setMovies(json.movies)
                setGenreName(props.location.genreName)
            });

    }, [props.match.params.id, props.location.genreName])

    if (!movies) {
        movies = [];
    }

    if (error !== null) {
        return <div>Error: {error.message}</div>;
    } else {
        return (
            <Fragment>
                <h2>Genre: {genreName}</h2>
                <div className="list-group">
                    {movies.map(movie => (
                        <Link key={movie.id} to={`/movies/${movie.id}`}
                              className="list-group-item list-group-item-action">
                            {movie.title}
                        </Link>
                    ))}
                </div>
            </Fragment>
        );
    }
}

export default GenrePageFunc;
