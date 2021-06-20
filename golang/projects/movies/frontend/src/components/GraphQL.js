import React, {Component, Fragment} from "react";
import Input from "./form-components/Input";
import {Link} from "react-router-dom";

export default class GraphQL extends Component {
    constructor(props) {
        super(props);

        this.state = {
            movies: [],
            isLoaded: false,
            error: null,
            alert: {
                type: "d-none",
                msg: ""
            },
            searchTerm: "",
        }

        this.handleChange = this.handleChange.bind(this)
    }

    handleChange = (event) => {
        let value = event.target.value;
        this.setState((prevState) => ({
            searchTerm: value,
        }))

        if (value.length > 2) {
            this.performSearch();
        } else {
            this.setState({movies: []})
        }
    }

    performSearch = () => {
        const payload = `
            {
                search(titleContains: "${this.state.searchTerm}") {
                    id
                    title
                    description
                    runtime
                    year
                }
            }
        `

        const headers = new Headers();
        headers.append("Content-Type", "application/json");

        const requestOptions = {
            method: "POST",
            body: payload,
            headers: headers,
        }

        fetch(`${process.env.REACT_APP_API_URL}/v1/graphql`, requestOptions)
            .then((resp) => resp.json())
            .then((data) => {
                console.log(data)
                return data
            })
            .then((data) => Object.values(data.data.search))
            .then((list) => {
                if (list.length > 0) {
                    this.setState({
                        movies: list,
                    })
                } else {
                    this.setState({
                        movies: [],
                    })
                }
            })
    }

    componentDidMount() {
        const payload = `
            {
                list {
                    id
                    title
                    description
                    runtime
                    year
                }
            }
        `

        const headers = new Headers();
        headers.append("Content-Type", "application/json");

        const requestOptions = {
            method: "POST",
            body: payload,
            headers: headers,
        }

        fetch(`${process.env.REACT_APP_API_URL}/v1/graphql`, requestOptions)
            .then((resp) => resp.json())
            .then((data) => Object.values(data.data.list))
            .then((list) => {
                this.setState({
                    movies: list,
                    isLoaded: true,
                })
            })
    }


    render() {
        let {movies} = this.state;
        return (
            <Fragment>
                <h2>GraphQL</h2>
                <hr/>
                <Input title={"Search"} type={"text"} name={"search"} value={this.state.searchTerm} handleChange={this.handleChange}/>
                <div className="list-group">
                    {movies.map((movie) => (
                        <Link to={`/moviesgraphql/${movie.id}`} key={movie.id} className="list-group-item list-group-item-action">
                            <strong>{movie.title}</strong><br/>
                            <small className="text-muted">{movie.year} - {movie.runtime} minutes</small><br/>
                            {movie.description.slice(0, 100)}...
                        </Link>
                    ))}
                </div>

            </Fragment>
        );
    }

}
