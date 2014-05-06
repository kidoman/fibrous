from twisted.web.resource import Resource
from twisted.web.server import Site
from twisted.internet import reactor

import json
import redis

redis_connection = redis.StrictRedis(host='localhost', port=6379, db=0)

class User:
    def __init__(self, id, name):
        self._id = id
        self._name = name

    @classmethod
    def fromRequest(cls, content):
        user_map = json.loads(content.getvalue())
        return User(user_map["id"], user_map["name"])

    @classmethod
    def load(cls, id):
        name = redis_connection.get(User._key(id))
        if name is None:
            return None
        return User(id, name)

    @classmethod
    def _key(cls, id):
        return "user:" + str(id)

    def name(self):
        return self._name

    def save(self):
        redis_connection.set(User._key(self._id), self._name)

    def to_json(self):
        return json.dumps({"id": self._id, "name": self._name})

class UserResource(Resource):
    isLeaf = True

    def user_id(self, path):
        path_variables = path.split("/")
        if len(path_variables) < 3 or path_variables[1] != "users": return None
        return path_variables[2]

    def render_GET(self, request):
        user_id = self.user_id(request.path)
        if user_id is None:
            request.setResponseCode(400)
            return ""

        user = User.load(user_id)
        if user is None:
            request.setResponseCode(404)
            return ""

        return user.to_json()

    def render_POST(self, request):
        user = User.fromRequest(request.content)
        user.save()
        return "OK"

resource = Site(UserResource())
reactor.listenTCP(3000, resource)
reactor.run()
