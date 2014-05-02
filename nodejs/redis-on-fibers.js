var Fiber = require('fibers');
var client = require('./redis-client');

exports.get = function(key) {
  var err, reply;
  var fiber = Fiber.current;
  client.get(key, function(_err, _reply) {
    err = _err;
    reply = _reply;
    fiber.run();
  });
  Fiber.yield();
  if (err != null) {
    throw err;
  }
  return reply;
}

exports.set = function(key, value) {
  var err;
  var fiber = Fiber.current;
  client.set(key, value, function(_err) {
    err = _err;
    fiber.run();
  });
  Fiber.yield();
  if (err != null) {
    throw err;
  }
}
