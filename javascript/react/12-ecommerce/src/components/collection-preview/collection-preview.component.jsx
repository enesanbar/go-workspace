import React, {Component} from "react";
import './collection-preview.styles.scss';
import CollectionItem from "../collection-item/collection-item.component";

class CollectionPreview extends Component {

    render() {
        return (
            <div className="collection-preview">
                <h1 className="title">{this.props.collection.title.toUpperCase()}</h1>
                <div className="preview">{this.renderPreview()}</div>
            </div>
        );
    }

    renderPreview = () => this.props.collection.items
        .filter((item, idx) => idx < 4)
        .map(item => <CollectionItem key={item.id} item={item}/>);

}

export default CollectionPreview;
