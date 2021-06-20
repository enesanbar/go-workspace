import React, {Fragment} from "react";
import {connect} from 'react-redux';
import {deleteStream, fetchStream} from "../../actions";
import Modal from "../Modal";
import history from "../../history";
import {Link} from "react-router-dom";

class StreamDelete extends React.Component {

    componentDidMount() {
        this.props.fetchStream(this.props.match.params.id);
    }

    renderActions = () => {
        return (
            <Fragment>
                <Link to="/" className="ui button">Cancel</Link>
                <button onClick={this.onSuccess} className="ui button negative">Delete</button>
            </Fragment>
        );
    };

    renderContent = () => {
        if (!this.props.stream) {
            return 'Are you sure you want to delete this stream?'
        }

        return `Are you sure you want to delete this stream with title ${this.props.stream.title}`
    };

    render() {
        return (
            <div>
                <Modal
                    title="Delete Stream"
                    content={this.renderContent()}
                    actions={this.renderActions()}
                    onDismiss={() => history.push('/')}
                />
            </div>
        );
    }

    onSuccess = () => {
        this.props.deleteStream(this.props.match.params.id);
    };

}

const mapStateToProps = (state, ownProps) => {
    return {
        stream: state.streams[ownProps.match.params.id],
    }
};

export default connect(mapStateToProps, {fetchStream, deleteStream})(StreamDelete);
