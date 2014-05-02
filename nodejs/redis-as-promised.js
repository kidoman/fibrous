var Q = require('Q');
var client = require('./redis-client');

exports.get = Q.nbind(client.get, client);
exports.set = Q.nbind(client.set, client);
