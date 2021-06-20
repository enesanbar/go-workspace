import React from "react";
import {Link} from "react-router-dom";
import {connect} from 'react-redux';
import {createStructuredSelector} from 'reselect';

import './header.styles.scss';
import {ReactComponent as Logo} from "../../assets/crown.svg";
import {auth} from "../../firebase/firebase.utils";
import CartIcon from "../cart-icon/cart-icon.component";
import CartDropdown from "../cart-dropdown/cart-dropdown.component";
import {selectCurrentUser} from "../../redux/user/userSelector";
import {selectCartHidden} from "../../redux/cart/cartSelectors";

const Header = (props) => {
    return (
        <div className="header">
            <Link to="/">
                <Logo className='logo'/>
            </Link>
            <div className="options">
                <Link className='option' to='/shop'>Shop</Link>
                <Link className='option' to='/contact'>Contact</Link>
                {
                    props.currentUser ?
                        <Link className="option" to="/" onClick={() => auth.signOut()}>SIGN OUT</Link> :
                        <Link className="option" to="/signin">SIGN IN</Link>
                }
                <CartIcon/>
            </div>
            {
                props.hidden ? null : <CartDropdown/>
            }
        </div>
    );
};

const mapStateToProps = createStructuredSelector({
    currentUser: selectCurrentUser,
    hidden: selectCartHidden
});

export default connect(mapStateToProps)(Header);
