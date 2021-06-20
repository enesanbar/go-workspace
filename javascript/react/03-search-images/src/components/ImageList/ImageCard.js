import React from "react";

class ImageCard extends React.Component {

    state = {span: 0};

    constructor(props) {
        super(props);

        this.imageRef = React.createRef();
    }

    setSpans = () => {
        const height = this.imageRef.current.clientHeight;
        const span = Math.ceil(height / 10) + 1;
        this.setState({span: span})
    };

    componentDidMount() {
        this.imageRef.current.addEventListener('load', this.setSpans);
        console.log(this.imageRef.current.clientHeight);
    }

    render() {
        const {alt_description, urls} = this.props.image;

        return (
            <div style={{gridRowEnd: `span ${this.state.span}`}}>
                <img ref={this.imageRef} src={urls.regular} alt={alt_description}/>
            </div>
        );
    }
}

export default ImageCard;
