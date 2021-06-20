import React, {Fragment, useEffect, useState} from "react";
import {Link} from "react-router-dom";

function GenresFunc(props) {
    const [genres, setGenres] = useState([])
    const [error, setError] = useState("")

    useEffect(() => {
        fetch(`${process.env.REACT_APP_API_URL}/v1/genres`)
            .then((response) => {
                if (response.status !== 200) {
                   setError("invalid response code: " + response.status)
                } else {
                    setError(null)
                }
                return response.json()
            })
            .then((json) => {
                setGenres(json.genres)
            });
    }, [])

    if (error) {
        return <div>Error: {error.message}</div>
    } else {
        return (
            <Fragment>
                <h2>Genres</h2>
                <div className="list-group">
                    {genres.map((genre) => (
                        <Link to={{pathname: `/genre/${genre.id}`, genreName: genre.genre_name}} key={genre.id}
                              className="list-group-item list-group-item-action">
                            {genre.genre_name}
                        </Link>
                    ))}
                </div>
            </Fragment>
        );
    }
}

export default GenresFunc;
