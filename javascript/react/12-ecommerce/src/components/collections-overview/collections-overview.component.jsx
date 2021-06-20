import React from "react";
import {connect} from 'react-redux';
import {createStructuredSelector} from "reselect";

import './collections-overview.styles.scss';
import CollectionPreview from "../collection-preview/collection-preview.component";
import {selectCollectionForPreview} from "../../redux/shop/shopSelectors";

const CollectionsOverview = ({collections}) => {
    return (
        <div className="collections-overview">
            {
                collections.map(collection => {
                    return <CollectionPreview key={collection.id} collection={collection}/>
                })
            }
        </div>
    );
};

const mapStateToProps = createStructuredSelector({
    collections: selectCollectionForPreview
});

export default connect(mapStateToProps)(CollectionsOverview);
