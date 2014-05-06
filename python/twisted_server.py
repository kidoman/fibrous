from twisted.web.resource import Resource
from twisted.web.server import Site
from twisted.internet import reactor

from user import User

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
        user = User.fromRequest(request.content.getvalue())
        user.save()
        return "OK"

resource = Site(UserResource())
reactor.listenTCP(3000, resource)
reactor.run()
