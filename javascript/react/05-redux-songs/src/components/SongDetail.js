import React from "react";
import {connect} from 'react-redux';

class SongDetail extends React.Component {

    render() {
        if (!this.props.song) {
            return <div>Select a song</div>
        }
        return (
            <div>
                <h3>Details for:</h3>
                <p>Title: {this.props.song.title}</p>
                <p>Duration: {this.props.song.duration}</p>
            </div>
        );
    }

}

const mapStateToProps = (state) => {
    return {
        song: state.selectedSong
    };
};

export default connect(mapStateToProps)(SongDetail);
