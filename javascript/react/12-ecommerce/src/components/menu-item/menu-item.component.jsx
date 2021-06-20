import React from "react";
import history from "../../history";
import './menu-item.styles.scss'

const MenuItem = (props) => {
    const {title, imageUrl, size, linkUrl} = props.section;

    return (
        <div className={`${size} menu-item`} onClick={() => history.push(linkUrl)}>
            <div className="background-image" style={{backgroundImage: `url(${imageUrl})`}} />
            <div className="content">
                <h1 className="title">{title.toUpperCase()}</h1>
                <span className="subtitle">SHOP NOW</span>
            </div>
        </div>
    );
};

export default MenuItem;
