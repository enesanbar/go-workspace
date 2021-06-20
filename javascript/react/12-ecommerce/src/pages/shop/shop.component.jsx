import React, {Component} from "react";

import CollectionsOverview from "../../components/collections-overview/collections-overview.component";
import {Route} from "react-router-dom";
import CollectionPage from "../collection/collection.component";

class ShopPage extends Component {

    render() {
        return (
            <div className="shop-page">
                <Route exact path={this.props.match.path} component={CollectionsOverview} />
                <Route path={`${this.props.match.path}/:collectionId`} component={CollectionPage} />
            </div>
        );

    }

}

export default ShopPage;
