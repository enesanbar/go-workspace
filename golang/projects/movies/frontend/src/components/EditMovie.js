import {Component, Fragment} from "react";
import './EditMovie.css'
import Input from "./form-components/Input";
import TextArea from "./form-components/TextArea";
import Select from "./form-components/Select";
import Alert from "./ui-components/Alert";
import {Link} from "react-router-dom";
import {confirmAlert} from "react-confirm-alert";
import 'react-confirm-alert/src/react-confirm-alert.css';

export default class EditMovie extends Component {

    constructor(props) {
        super(props);

        this.state = {
            movie: {
                id: 0,
                title: "",
                release_date: "",
                runtime: "",
                mpaa_rating: "",
                rating: "",
                description: ""
            },
            mpaaOptions: [
                {id: "G", value: "G"},
                {id: "PG", value: "PG"},
                {id: "PG13", value: "PG13"},
                {id: "R", value: "R"},
                {id: "NC17", value: "NC17"},
            ],
            isLoaded: false,
            error: null,
            errors: [],
            alert: {
                type: 'd-none',
                msg: '',
            }
        }

        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleSubmit = (event) => {
        event.preventDefault()

        // client side validation
        let errors = [];
        if (this.state.movie.title === "") {
            errors.push("title")
        }

        this.setState({errors: errors})

        if (errors.length > 0) {
            return false;
        }

        const data = new FormData(event.target);
        const payload = Object.fromEntries(data.entries());
        const headers = new Headers();
        headers.append("Content-Type", "application/json")
        headers.append("Authorization", "Bearer " + this.props.jwt)
        const requestOptions = {
            method: 'POST',
            body: JSON.stringify(payload),
            headers: headers,
        }

        fetch(`${process.env.REACT_APP_API_URL}/v1/admin/editmovie`, requestOptions)
            .then((response) => response.json())
            .then(data => {
                if (data.error) {
                    this.setState({
                        alert: {
                            type: 'alert-danger',
                            message: data.error.message,
                        }
                    })
                } else {
                    this.props.history.push({
                        pathname: "/admin"
                    })
                }
            })
    }

    handleChange = (event) => {
        let value = event.target.value;
        let name = event.target.name

        this.setState((prevState) => ({
            movie: {
                ...prevState.movie,
                [name]: value,
            }
        }))
    }

    hasError = (key) => {
        return this.state.errors.indexOf(key) !== -1;
    }

    componentDidMount() {
        if (this.props.jwt === "") {
            this.props.history.push({
                pathname: "/login"
            })
            return
        }

        const id = this.props.match.params.id;

        if (id > 0) {
            fetch(`${process.env.REACT_APP_API_URL}/v1/movies/` + id)
                .then((response) => {
                    if (response.status !== "200") {
                        let err = Error
                        err.Message = "Invalid response code: " + response.status
                        this.setState({error: err})
                    }
                    return response.json()
                })
                .then((json) => {
                    const releaseDate = new Date(json.movie.release_date)

                    this.setState({
                        movie: {
                            id: id,
                            title: json.movie.title,
                            release_date: releaseDate.toISOString().split("T")[0],
                            runtime: json.movie.runtime,
                            mpaa_rating: json.movie.mpaa_rating,
                            rating: json.movie.rating,
                            description: json.movie.description,
                        },
                        isLoaded: true,
                    }, (error) => {
                        this.setState({
                            isLoaded: true,
                            error
                        })
                    })
                })
        } else {
            this.setState({
                isLoaded: true,
            })
        }
    }

    confirmDelete = (event) => {
        console.log("deleting movie", event)
        confirmAlert({
            title: 'Delete Movie',
            message: 'Are you sure?',
            buttons: [
                {
                    label: 'Yes',
                    onClick: () => {
                        const headers = new Headers();
                        headers.append("Content-Type", "application/json")
                        headers.append("Authorization", "Bearer " + this.props.jwt)
                        const requestOptions = {
                            method: 'DELETE',
                            headers: headers,
                        }

                        fetch(`${process.env.REACT_APP_API_URL}/v1/movies/` + this.props.match.params.id, requestOptions)
                            .then((response) => response.json())
                            .then((data) => {
                                if (data.error) {
                                    this.setState({
                                        alert: {
                                            type: 'alert-danger',
                                            message: data.error.message,
                                        }
                                    })
                                } else {
                                    this.props.history.push({
                                        pathname: "/admin",
                                    })
                                }
                            })
                    }
                },
                {
                    label: 'No',
                    onClick: () => {}
                }
            ]
        });
    }

    render() {
        const {movie, isLoaded, error} = this.state;
        if (error) {
            return <div>Error: {error.message}</div>
        } else if (!isLoaded) {
            return <p>Loading...</p>
        } else {
            return (
                <Fragment>
                    <h2>Add/Edit Movie</h2>
                    <Alert alertType={this.state.alert.type} alertMessage={this.state.alert.message} />
                    <hr/>
                    <form onSubmit={this.handleSubmit}>
                        <input type="hidden" name="id" id="id" value={movie.id} onChange={this.handleChange}/>
                        <Input name="title" type="text" title="Title" value={movie.title}
                               handleChange={this.handleChange}
                               className={this.hasError("title") ? "is-invalid" : ""}
                               errorDiv={this.hasError("title") ? "text-danger" : "d-none"}
                               errorMsg={"Please enter a title"}
                        />
                        <Input name="release_date" type="date" title="Release Date" value={movie.release_date}
                               handleChange={this.handleChange}/>
                        <Input name="runtime" type="text" title="Runtime" value={movie.runtime}
                               handleChange={this.handleChange}/>
                        <Select name="mpaa_rating" title="MPAA Rating" options={this.state.mpaaOptions}
                                value={movie.mpaa_rating} handleChange={this.handleChange} placeholder="Choose..."/>
                        <Input name="rating" type="text" title="Rating" value={movie.rating}
                               handleChange={this.handleChange}/>
                        <TextArea name="description" title="Description" value={movie.description}
                                  handleChange={this.handleChange}/>

                        <hr/>

                        <button className="btn btn-primary">Save</button>
                        <Link to="/admin" className="btn btn-warning ms-1">Cancel</Link>
                        {
                            movie.id > 0 && (
                                <a href="#!" onClick={() => this.confirmDelete()} className="btn btn-danger ms-1">Delete</a>
                            )
                        }
                    </form>

                    <div className="mt-3">
                        <pre>{JSON.stringify(this.state, null, 3)}</pre>
                    </div>
                </Fragment>
            );
        }
    }


};
