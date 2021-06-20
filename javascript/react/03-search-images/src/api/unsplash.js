import axios from 'axios';

export default axios.create({
    baseURL: 'https://api.unsplash.com',
    headers: {
        Authorization: 'Client-ID <<UNSPLASH_CLIENT_ID>>'
    }
});
