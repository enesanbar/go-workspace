import React from 'react';

class SearchBar extends React.Component {

    state = {term: ''};

    onFormSubmit = (event) => {
        event.preventDefault();

        // Parent component (App) decides what to do with the search
        this.props.onSubmit(this.state.term);
    };

    render() {
        return (
            <div className="ui segment">
                <form className="ui form" action="" onSubmit={this.onFormSubmit}>
                    <div className="field">
                        <label htmlFor="">Image Search</label>
                        <input type="text"
                               value={this.state.term}
                               onChange={(e) => this.setState({term: e.target.value})}
                        />
                    </div>
                </form>
            </div>
        );
    }
}

export default SearchBar;
