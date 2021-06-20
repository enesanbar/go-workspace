import React, {Component} from "react";
import {connect} from "react-redux";

import "./checkout-item.styles.scss";
import {addItem, removeItem, clearItemFromCart} from "../../redux/cart/cartActions";

class CheckoutItem extends Component {

    render() {
        const {name, quantity, price, imageUrl} = this.props.item;

        return (
            <div className="checkout-item">
                <div className="image-container">
                    <img src={imageUrl} alt="item"/>
                </div>
                <span className="name">{name}</span>
                <span className="quantity">
                    <div className="arrow" onClick={() => this.props.removeItem(this.props.item)}>&#10094;</div>
                    <span className="value">{quantity}</span>
                    <div className="arrow" onClick={() => this.props.addItem(this.props.item)}>&#10095;</div>
                </span>
                <span className="price">{price}</span>
                <span className="remove-button" onClick={() => this.props.clearItem(this.props.item)}>
                    &#10005;
                </span>
            </div>
        );
    }
}

const mapDispatchToProps  = (dispatch) => ({
    clearItem: item => dispatch(clearItemFromCart(item)),
    addItem: item => dispatch(addItem(item)),
    removeItem: item => dispatch(removeItem(item)),
});

export default connect(null, mapDispatchToProps)(CheckoutItem);
