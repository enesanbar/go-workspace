export const addItemToCart = (cartItems, cartItemToAdd) => {
    const doesCartItemExists = cartItems.find(cartItem => cartItem.id === cartItemToAdd.id);

    if (doesCartItemExists) {
        return cartItems.map(cartItem => {
            if (cartItem.id === cartItemToAdd.id) {
                return {...cartItem, quantity: cartItem.quantity + 1}
            } else {
                return cartItem
            }
        });
    }

    return [...cartItems, {...cartItemToAdd, quantity: 1}]
};


export const removeItemFromCart = (cartItems, cartItemToRemove) => {
    const doesCartItemExists = cartItems.find(cartItem => cartItem.id === cartItemToRemove.id);

    if (doesCartItemExists.quantity === 1) {
        return cartItems.filter(cartItem => cartItem.id !== cartItemToRemove.id);
    }

    return cartItems.map(cartItem => {
        if (cartItem.id === cartItemToRemove.id) {
            return {...cartItem, quantity: cartItem.quantity - 1}
        } else {
            return cartItem;
        }
    })
};


export const clearItemFromCart = (cartItems, cartItemToClear) => {
    return cartItems.filter(item => item.id !== cartItemToClear.id )
};
