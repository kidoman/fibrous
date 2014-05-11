require('remedial');
var http = require('http');
var url = require('url');
var jsonBody = require('body/json');
var client = require('./redis-client');

var User = function(id, name) {
  this.id = id;
  this.name = name;
};

User.prototype.key = function() {
  return 'user:{id}'.supplant({id: this.id});
};

User.prototype.save = function(cb) {
  var key = this.key();
  client.set(key, this.name, cb);
};

User.get = function(id, cb) {
  var key = User.prototype.key.call({id: id});
  client.get(key, function(error, reply) {
    if (error != null) {
      cb(error, null);
      return;
    }

    if (reply === null) {
      cb(null, null);
      return;
    }

    var user = new User(id, reply);
    cb(null, user);
  });
};

http.createServer(function(req, res) {
  var pathname = url.parse(req.url).pathname;
  var prefix = '/users/';
  if (req.method === 'GET' && pathname.substr(0, prefix.length) === prefix) {
    var id = pathname.substr(prefix.length);
    User.get(id, function(error, user) {
      if (error) {
        res.statusCode = 500;
        return res.end();
      }

      if (user === null) {
        res.statusCode = 404;
        return res.end();
      }

      res.end(JSON.stringify(user));
    });
  } else if (req.method === 'POST' && pathname === '/users') {
    jsonBody(req, res, function(err, body) {
      if (err) {
        res.statusCode = 400;
        return res.end();
      }

      var user = new User(body.id, body.name);
      user.save(function(err) {
        if (err) {
          res.statusCode = 500;
          return res.end();
        }

        res.statusCode = 201;
        res.end('OK');
      });
    })
  }
}).listen(3000, '127.0.0.1');
console.log('Started server on port 3000');
