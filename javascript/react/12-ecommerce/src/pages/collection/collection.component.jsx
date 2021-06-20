import React from "react";
import './colllection.styles.scss'
import {selectCollection} from "../../redux/shop/shopSelectors";
import {connect} from "react-redux";
import CollectionItem from "../../components/collection-item/collection-item.component";

const CollectionPage = (props) => {
    const {title, items} = props.collection;

    return (
        <div className="collection-page">
            <h2 className="title">{title}</h2>
            <div className="items">
                {
                    items.map(item => <CollectionItem key={item.id} item={item} />)
                }
            </div>
        </div>
    );
};

const mapStateToProps = (state, ownProps) => ({
    collection: selectCollection(ownProps.match.params.collectionId)(state)
});

export default connect(mapStateToProps)(CollectionPage);
