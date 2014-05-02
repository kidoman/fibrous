var redis = require('redis');

var options = {
  no_ready_check: true,
  enable_offline_queue: false
};

var port = process.env.REDIS_PORT;
var host = process.env.REDIS_HOST;

module.exports = redis.createClient(port, host, options);
