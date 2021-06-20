import React from "react";
import StripeCheckout from "react-stripe-checkout";

const onToken = token => {
    console.log(token);
    alert('Payment Successful');
};

const StripeCheckoutButton = ({price}) => {
    const priceForStripe = price * 100;
    const publishableKey = 'pk_test_L8pqEe3p0dAMBJrhigzSIfHi';

    return (
        <StripeCheckout
            label='Pay Now'
            name='Clothing Ltd.'
            billingAddress
            shippingAddress
            image=''
            description={`Your total is ${price}`}
            amount={priceForStripe}
            panelLabel='Pay Now'
            token={onToken}
            stripeKey={publishableKey}
        />
    );
};

export default StripeCheckoutButton;
