var Fiber = require('fibers');

module.exports = function() {
  return function(req, res, next) {
    Fiber(function() {
      next();
    }).run();
  };
};
