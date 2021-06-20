import React from 'react';
import './SeasonDisplay.css'

const seasonConfig = {
    summer: {
        text: "Let's hit the beach",
        iconName: 'sun'
    },
    winter: {
        text: "It's cold",
        iconName: 'snowflake'
    }
};
const getSeason = (latitude, month) => {
    if (month > 2 && month < 9) {
        return latitude > 0 ? 'summer' : 'winter';
    } else {
        return latitude > 0 ? 'winter' : 'summer';
    }
};

class SeasonDisplay extends React.Component{

    render() {
        const season = getSeason(this.props.latitude, new Date().getMonth());
        const { text, iconName } = seasonConfig[season];

        return (
            <div className={`season-display ${season}`}>
                <i className={`icon-left massive icon ${iconName}`}/>
                <h1>{text}</h1>
                <i className={`icon-right massive icon ${iconName}`}/>
            </div>
        )
    }

}

export default SeasonDisplay;
