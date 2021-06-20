import React, {Component} from "react";
import MenuItem from "../menu-item/menu-item.component";
import './directory.styles.scss';
import {connect} from "react-redux";
import {selectDirectorySections} from "../../redux/directory/directorySelectors";
import {createStructuredSelector} from "reselect";

class Directory extends Component{

    state = {
        sections: [
            {
                title: 'hats',
                imageUrl: 'https://i.ibb.co/cvpntL1/hats.png',
                id: 1,
                linkUrl: '/shop/hats'
            },
            {
                title: 'jackets',
                imageUrl: 'https://i.ibb.co/px2tCc3/jackets.png',
                id: 2,
                linkUrl: '/shop/jackets'
            },
            {
                title: 'sneakers',
                imageUrl: 'https://i.ibb.co/0jqHpnp/sneakers.png',
                id: 3,
                linkUrl: '/shop/sneakers'
            },
            {
                title: 'women',
                imageUrl: 'https://i.ibb.co/GCCdy8t/womens.png',
                size: 'large',
                id: 4,
                linkUrl: '/shop/women'
            },
            {
                title: 'mens',
                imageUrl: 'https://i.ibb.co/R70vBrQ/men.png',
                size: 'large',
                id: 5,
                linkUrl: '/shop/mens'
            }
        ]
    };

    renderMenuItems = () => {
        return this.props.sections.map(section => {
            return <MenuItem key={section.id} section={section} />
        });
    };

    render() {
        return (
            <div className="directory-menu">
                {this.renderMenuItems()}
            </div>
        );
    }

}

const mapStateToProps = createStructuredSelector({
    sections: selectDirectorySections
});

export default connect(mapStateToProps)(Directory);
