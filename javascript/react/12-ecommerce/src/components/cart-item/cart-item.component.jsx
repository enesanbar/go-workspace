import React, {Component} from "react";
import {connect} from "react-redux";

import './cart-item.styles.scss'

class CartItem extends Component {

    render() {
        const {name, price, quantity, imageUrl} = this.props.item;

        return (
            <div className="cart-item">
                <img src={imageUrl} alt="item" />
                <div className="item-details">
                    <span className="name">{name}</span>
                    <span className="price">{quantity} x {price}</span>
                </div>
            </div>
        );
    }

}

export default connect()(CartItem);
