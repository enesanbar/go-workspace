import React, {Component, Fragment} from "react";
import {Link} from "react-router-dom";

export default class Admin extends Component {
    state = {
        movies: [],
        isLoaded: false,
        error: null,
    };

    componentDidMount() {
        console.log("Admin componentDidMount")
        if (this.props.jwt === "") {
            console.log("jwt empty")
            this.props.history.push({
                pathname: "/login"
            });
            return;
        }
        console.log("jwt not empty", this.props.jwt)

        fetch(`${process.env.REACT_APP_API_URL}/v1/movies`)
            .then((response) => {
                console.log(response.status);
                if (response.status !== 200) {
                    let err = Error
                    err.message = "invalid response code: " + response.status
                    this.setState({error: err})
                }
                return response.json()
            })
            .then((json) => {
                this.setState({
                        movies: json.movies,
                        isLoaded: true
                    }, (error) => {
                        this.setState({
                            isLoaded: true,
                            error
                        })
                    }
                )
            });
    }

    render() {
        const {movies, isLoaded, error} = this.state;

        if (error) {
            return <div>Error: {error.message}</div>
        }
        else if (!isLoaded) {
            return <p>Loading...</p>
        } else {

            return (
                <Fragment>
                    <h2>Choose a movie</h2>
                    <div className="list-group">
                        {movies.map(movie => (
                            <Link to={`/admin/movie/${movie.id}`} key={movie.id} className="list-group-item list-group-item-action">
                                {movie.title}
                            </Link>
                        ))}
                    </div>
                </Fragment>
            );
        }

    }
}
