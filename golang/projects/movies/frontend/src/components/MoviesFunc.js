import React, {useEffect, useState, Fragment} from "react";
import {Link} from "react-router-dom";

function MoviesFunc(props) {

    const [movies, setMovies] = useState([])
    const [error, setError] = useState("")

    useEffect(() => {
        fetch(`${process.env.REACT_APP_API_URL}/v1/movies`)
            .then((response) => {
                console.log(response.status);
                if (response.status !== 200) {
                    setError("Invalid response code: ", response.status)
                } else {
                    setError(null)
                }
                return response.json()
            })
            .then((json) => {
                setMovies(json.movies)
            });
    }, [])

    if (error != null) {
        return <div>Error: {error}</div>
    } else {
        return (
            <Fragment>
                <h2>Choose a movie</h2>
                <div className="list-group">
                    {movies.map(movie => (
                        <Link to={`/movies/${movie.id}`} key={movie.id}
                              className="list-group-item list-group-item-action">
                            {movie.title}
                        </Link>
                    ))}
                </div>
            </Fragment>
        );
    }

}

export default MoviesFunc;
