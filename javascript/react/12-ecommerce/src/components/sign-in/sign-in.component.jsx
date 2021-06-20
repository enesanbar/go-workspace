import React, {Component} from "react";
import './sign-in.styles.scss';
import FormInput from "../form-input/form-input.component";
import Button from "../button/button.component";
import {auth, signInWithGoogle} from '../../firebase/firebase.utils';

class SignIn extends Component {

    state = {
        email: '',
        password: ''
    };

    handleSubmit = async (e) => {
        e.preventDefault();

        const {email, password} = this.state;

        try {
            await auth.signInWithEmailAndPassword(email, password);
            this.setState({email: '', password: ''})
        } catch (err) {
            console.log(err);
        }
        this.setState({email: '', password: ''})
    };

    handleChange = (e) => {
        const {value, name} = e.target;
        this.setState({[name]: value})
    };

    render() {
        return (
            <div className="sign-in">
                <h2>I already have an account</h2>
                <span>Sign in with your email and password</span>

                <form action="" onSubmit={this.handleSubmit}>
                    <FormInput
                        name="email"
                        type="email"
                        label="email"
                        value={this.state.email}
                        handleChange={this.handleChange}
                        required />

                    <FormInput
                        name="password"
                        type="password"
                        label="password"
                        value={this.state.password}
                        handleChange={this.handleChange}
                        required/>

                    <div className="buttons">
                        <Button type="submit" value="Submit Form">
                            Sign In
                        </Button>
                        <Button value="Submit Form" onClick={signInWithGoogle} isGoogleSignIn>
                            Sign In with Google
                        </Button>
                    </div>
                </form>
            </div>
        );
    }

}

export default SignIn;
