import React from 'react';

class CommentDetail extends React.Component {

    render() {
        return (
            <div className="comment">
                <a href="/" className="avatar">
                    <img src={this.props.avatar} alt="avatar"/>
                </a>
                <div className="content">
                    <a href="/" className="author">
                        {this.props.name}
                    </a>
                    <div className="metadata">
                            <span className="date">
                                {this.props.weekday} at 6.00 PM
                            </span>
                    </div>
                    <div className="text">
                        {this.props.comment}
                    </div>
                </div>
            </div>
        )
    }
}

export default CommentDetail;
