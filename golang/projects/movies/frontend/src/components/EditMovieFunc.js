import React, {Fragment, useEffect, useState} from "react";
import './EditMovie.css'
import Input from "./form-components/Input";
import TextArea from "./form-components/TextArea";
import Select from "./form-components/Select";
import Alert from "./ui-components/Alert";
import {Link} from "react-router-dom";
import {confirmAlert} from "react-confirm-alert";
import 'react-confirm-alert/src/react-confirm-alert.css';

function EditMovieFunc(props) {
    const [movie, setMovie] = useState({});
    const [error, setError] = useState(null)
    const [errors, setErrors] = useState([])
    const [alert, setAlert] = useState({
        type: "d-none",
        message: "",
    })
    const mpaaOptions = [
        {id: "G", value: "G"},
        {id: "PG", value: "PG"},
        {id: "PG13", value: "PG13"},
        {id: "R", value: "R"},
        {id: "NC17", value: "NC17"},
    ]

    useEffect(() => {
        if (props.jwt === "") {
            props.history.push({
                pathname: "/login"
            })
            return
        }

        const id = props.match.params.id;

        if (id > 0) {
            fetch(`${process.env.REACT_APP_API_URL}/v1/movies/` + id)
                .then((response) => {
                    if (response.status !== 200) {
                        setError("Invalid response code: " + response.status)
                    } else {
                        setError(null)
                    }
                    return response.json()
                })
                .then((json) => {
                    const releaseDate = new Date(json.movie.release_date)
                    json.movie.release_date = releaseDate.toISOString().split("T")[0]
                    setMovie(json.movie)
                })
        }
    }, [props.jwt, props.history, props.match.params.id])

    const handleSubmit = (event) => {
        event.preventDefault()

        // client side validation
        let errors = [];
        if (movie.title === "") {
            errors.push("title")
        }

        setErrors(errors)

        if (errors.length > 0) {
            return false;
        }

        const data = new FormData(event.target);
        const payload = Object.fromEntries(data.entries());
        const headers = new Headers();
        headers.append("Content-Type", "application/json")
        headers.append("Authorization", "Bearer " + props.jwt)
        const requestOptions = {
            method: 'POST',
            body: JSON.stringify(payload),
            headers: headers,
        }

        fetch(`${process.env.REACT_APP_API_URL}/v1/admin/editmovie`, requestOptions)
            .then((response) => response.json())
            .then(data => {
                if (data.error) {
                    setAlert({
                        type: 'alert-danger',
                        message: data.error.message,
                    })
                } else {
                    props.history.push({
                        pathname: "/admin"
                    })
                }
            })
    }

    const handleChange = (event) => {
        let value = event.target.value;
        let name = event.target.name
        setMovie({
            ...movie,
            [name]: value,
        })
    }

    const confirmDelete = (event) => {
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
                        headers.append("Authorization", "Bearer " + props.jwt)
                        const requestOptions = {
                            method: 'DELETE',
                            headers: headers,
                        }

                        fetch(`${process.env.REACT_APP_API_URL}/v1/movies/` + props.match.params.id, requestOptions)
                            .then((response) => response.json())
                            .then((data) => {
                                if (data.error) {
                                    setAlert({type: 'alert-danger', message: data.error.message})
                                } else {
                                    setAlert({type: 'alert-success', message: "movie deleted"})
                                    props.history.push({
                                        pathname: "/admin",
                                    })
                                }
                            })
                    }
                },
                {
                    label: 'No',
                    onClick: () => {
                    }
                }
            ]
        });
    }

    const hasError = (key) => {
        return errors.indexOf(key) !== -1;
    }


    if (error != null) {
        return <div>Error: {error.message}</div>
    } else {
        return (
            <Fragment>
                <h2>Add/Edit Movie</h2>
                <Alert alertType={alert.type} alertMessage={alert.message}/>
                <hr/>
                <form onSubmit={handleSubmit}>
                    <input type="hidden" name="id" id="id" value={movie.id} onChange={handleChange}/>
                    <Input name="title" type="text" title="Title" value={movie.title}
                           handleChange={handleChange}
                           className={hasError("title") ? "is-invalid" : ""}
                           errorDiv={hasError("title") ? "text-danger" : "d-none"}
                           errorMsg={"Please enter a title"}
                    />
                    <Input name="release_date" type="date" title="Release Date" value={movie.release_date}
                           handleChange={handleChange}/>
                    <Input name="runtime" type="text" title="Runtime" value={movie.runtime}
                           handleChange={handleChange}/>
                    <Select name="mpaa_rating" title="MPAA Rating" options={mpaaOptions}
                            value={movie.mpaa_rating} handleChange={handleChange} placeholder="Choose..."/>
                    <Input name="rating" type="text" title="Rating" value={movie.rating}
                           handleChange={handleChange}/>
                    <TextArea name="description" title="Description" value={movie.description}
                              handleChange={handleChange}/>

                    <hr/>

                    <button className="btn btn-primary">Save</button>
                    <Link to="/admin" className="btn btn-warning ms-1">Cancel</Link>
                    {
                        movie.id > 0 && (
                            <a href="#!" onClick={() => confirmDelete()} className="btn btn-danger ms-1">Delete</a>
                        )
                    }
                </form>
            </Fragment>
        );
    }

}

export default EditMovieFunc;
