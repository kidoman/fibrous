require('remedial');
var Q = require('q');
var express = require('express');
var bodyParser = require('body-parser');
var client = require('./redis-as-promised');

var User = function(id, name) {
  this.id = id;
  this.name = name;
};

User.prototype.key = function() {
  return 'user:{id}'.supplant({id: this.id});
};

User.prototype.save = function() {
  var key = this.key();
  return client.set(key, this.name);
};

User.get = function(id) {
  var key = User.prototype.key.call({id: id});
  return client.get(key)
  .then(function(name) {
    return new User(id, name);
  });
};

var app = express();
app.use(bodyParser());

app.post('/users', function(req, res) {
  var body = req.body;
  var user = new User(body.id, body.name);
  user.save()
  .then(function() {
    res.send(200);    
  });
});

app.get('/users/:id', function(req, res) {
  var id = req.params.id;
  User.get(id)
  .then(function(user) {
    res.send(200, user);
  });
});

app.listen(3000, function() {
  console.log('Started server on port 3000');
});
