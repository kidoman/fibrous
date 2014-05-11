require('remedial');
var express = require('express');
var bodyParser = require('body-parser');
var expressOnFibers = require('./express-on-fibers');
var client = require('./redis-on-fibers');

var User = function(id, name) {
  this.id = id;
  this.name = name;
};

User.prototype.key = function() {
  return 'user:{id}'.supplant({id: this.id});
};

User.prototype.save = function() {
  var key = this.key();
  client.set(key, this.name);
};

User.get = function(id) {
  var key = User.prototype.key.call({id: id});
  var name = client.get(key);
  if (name === null) {
    return null;
  }

  return new User(id, name);
};

var app = express();
app.use(bodyParser());
app.use(expressOnFibers());

app.post('/users', function(req, res) {
  var body = req.body;
  var user = new User(body.id, body.name);
  user.save();
  res.send(201, 'OK');
});

app.get('/users/:id', function(req, res) {
  var id = req.params.id;
  var user = User.get(id);
  if (user === null) {
    return res.send(404);
  }

  res.send(200, user);
});

app.listen(3000, function() {
  console.log('Started server on port 3000');
});
