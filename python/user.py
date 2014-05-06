import json
import redis

redis_connection = redis.StrictRedis(host='localhost', port=6379, db=0)

class User:
    def __init__(self, id, name):
        self._id = id
        self._name = name

    @classmethod
    def fromRequest(cls, content):
        user_map = json.loads(content)
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
