import React, {Component} from "react";
import {connect} from "react-redux";

import {ReactComponent as ShoppingIcon} from "../../assets/shopping-bag.svg";
import './cart-icon.styles.scss'
import {toggleCardHidden} from "../../redux/cart/cartActions";
import {selectCartItemsCount} from "../../redux/cart/cartSelectors";

class CartIcon extends Component {

    render() {
        return (
            <div className="cart-icon" onClick={this.props.toggleCardHidden}>
                <ShoppingIcon className="shopping-icon"/>
                <span className="item-count">{this.props.count}</span>
            </div>
        );
    }

}

const mapDispatchToProps = (dispatch) => ({
    toggleCardHidden: () =>  dispatch(toggleCardHidden())
});

const mapStateToProps = (state) => ({
    count: selectCartItemsCount(state)
});

export default connect(mapStateToProps, mapDispatchToProps)(CartIcon);
