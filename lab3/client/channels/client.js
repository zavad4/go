const http = require('../common/http');

const Client = (baseUrl) => {

    const client = http.Client(baseUrl);

    return {
        listChannels: () => client.get('/users'),
        createChannel: (name) => client.post('/users', { name })
    }

};

module.exports = { Client };