import React from "react";
import './VideoItem.css';

class VideoItem extends React.Component {

    render() {
        return (
            <div className="video-item item" onClick={() => this.props.onVideoSelect(this.props.video)}>
                <img className="ui image" src={this.props.video.snippet.thumbnails.medium.url} alt={this.props.video.snippet.title}/>
                <div className="content">
                    <div className="header">
                        {this.props.video.snippet.title}
                    </div>
                </div>
            </div>
        );
    }

}

export default VideoItem;
