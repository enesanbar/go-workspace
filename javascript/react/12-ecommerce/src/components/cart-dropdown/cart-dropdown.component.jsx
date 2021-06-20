import React, {Component} from "react";
import {connect} from "react-redux";
import {createStructuredSelector} from 'reselect';

import Button from "../button/button.component";
import './cart-dropdown.styles.scss'
import CartItem from "../cart-item/cart-item.component";
import {selectCartItems} from "../../redux/cart/cartSelectors";
import history from "../../history";
import {toggleCardHidden} from "../../redux/cart/cartActions";

class CartDropdown extends Component {

    render() {
        return (
            <div className="cart-dropdown">
                <div className="cart-items">{this.renderCartItems()}</div>
                <Button onClick={() => {
                    history.push('/checkout');
                    this.props.dispatch(toggleCardHidden())
                }}>
                    GO TO CHECKOUT
                </Button>
            </div>
        );
    }

    renderCartItems = () => {
        if (this.props.cartItems.length === 0) {
            return <span className="empty-message">Your cart is empty</span>
        } else {
            return this.props.cartItems.map(cartItem => {
                return <CartItem key={cartItem.id} item={cartItem}/>
            });
        }
    }
}

const mapStateToProps = createStructuredSelector({
    cartItems: selectCartItems
});

export default connect(mapStateToProps)(CartDropdown);
