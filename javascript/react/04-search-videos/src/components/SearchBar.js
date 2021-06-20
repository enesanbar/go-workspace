import React from "react";

class SearchBar extends React.Component {

    state = {term: ''};

    render() {
        return (
            <div className="search-bar ui segment">
                <form action="" className="ui form"onSubmit={this.onFormSubmit}>
                    <div className="field">
                        <label htmlFor="">Video Search</label>
                        <input value={this.state.term}
                               type="text"
                               onChange={(e) => this.setState({term: e.target.value})}
                        />
                    </div>
                </form>
            </div>
        );
    }

    onFormSubmit = (event) => {
        event.preventDefault();
        this.props.onTermSubmit(this.state.term);
    }
}

export default SearchBar;
