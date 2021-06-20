import {combineReducers} from "redux";

const songsReducer = () => {
    return [
        {
            title: "Black",
            duration: '3.54'
        },
        {
            title: "Sweet Virginia",
            duration: '4.53'
        },
        {
            title: "How do you sleep",
            duration: '5.34'
        }
    ];
};

const selectedSongReducer = (selectedSong = null, action) => {
    if (action.type === 'SONG_SELECTED') {
        return action.payload;
    }

    return selectedSong;
};

export default combineReducers({
    songs: songsReducer,
    selectedSong: selectedSongReducer
});
