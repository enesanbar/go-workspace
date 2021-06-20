import React, {Component} from "react";
import {connect} from "react-redux";
import {createStructuredSelector} from 'reselect';

import {selectCartItems, selectCartTotal} from "../../redux/cart/cartSelectors";
import "./checkout.styles.scss";
import CheckoutItem from "../checkout-item/checkout-item.component";
import StripeCheckoutButton from "../stripe-button/stripe-button.component";

class Checkout extends Component {

    render() {
        return (
            <div className="checkout-page">
                <div className="checkout-header">
                    <div className="header-block">Product</div>
                    <div className="header-block">Description</div>
                    <div className="header-block">Quantity</div>
                    <div className="header-block">Price</div>
                    <div className="header-block">Remove</div>
                </div>
                {this.renderCartItems()}
                <div className="total">TOTAL: {this.props.total}</div>
                <div className="test-warning">
                    Please use the following test card
                    <br/>
                    4242 4242 4242 4242 - 01/20 - 123
                </div>
                <StripeCheckoutButton price={this.props.total}/>
            </div>
        );
    }

    renderCartItems = () => {
        return this.props.cartItems.map(cartItem => {
            return <CheckoutItem key={cartItem.id} item={cartItem}/>
        });
    };

}

const mapStateToProps = createStructuredSelector({
    cartItems: selectCartItems,
    total: selectCartTotal
});

export default connect(mapStateToProps)(Checkout);
