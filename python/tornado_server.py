import tornado.escape
from tornado.ioloop import IOLoop
from tornado.web import RequestHandler
from tornado.web import Application

from user import User

class GetUserByIdHandler(RequestHandler):
    def get(self, id):
        user = User.load(id)
        if user is None:
            self.set_status(404)
            return

        self.write(user.to_json())

class SaveUserHandler(RequestHandler):
    def post(self):
        user = User.fromRequest(self.request.body)
        user.save()
        self.write("OK")

application = tornado.web.Application([
    (r"/users/([0-9]+)", GetUserByIdHandler),
    (r"/users", SaveUserHandler)
])

application.listen(3000)
IOLoop.instance().start()
