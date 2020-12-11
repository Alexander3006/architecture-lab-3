const request = require('request');

const Client = (baseUrl) => {

    const respHandler = (resp) => {
        if (resp.ok) {
            return resp.json();
        }
        throw new Error(`Unexpected response from the server ${resp.status} ${resp.statusText}`)
    };

    const method = async (path, req, data) => {
        return new Promise((resolve, reject) => {
            request(`${baseUrl}${path}`, {json: true, method: req, body: data}, (err, res, body) => {
                if (err) {
                    reject(err);
                    return;
                }
                resolve(body);
            });
        });
    }

    return {
        get: (path) => method(path, 'GET', null),
        post: (path, data) => method(path, 'POST', data),
        delete: (path, data) => method(path, 'DELETE', data)
    };
};

module.exports = { Client };