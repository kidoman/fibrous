require('remedial');
var express = require('express');
var bodyParser = require('body-parser')
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

    var user = new User(id, reply);
    cb(null, user);
  });
};

var app = express();
app.use(bodyParser());

app.post('/users', function(req, res) {
  var body = req.body;
  var user = new User(body.id, body.name);
  user.save(function(err) {
    res.send(200);
  });
});

app.get('/users/:id', function(req, res) {
  var id = req.params.id;
  User.get(id, function(error, user) {
    res.send(200, user);
  });
});

app.listen(3000, function() {
  console.log('Started server on port 3000');
});
