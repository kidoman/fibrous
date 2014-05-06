require "sinatra"
require "./user"

disable :logging
set :port, 3000

get "/users/:id" do |id|
  user = User.load(id)
  return [404] if !user

  user.to_json
end

post "/users" do
  user = User.from(request.body.read)
  user.save

  [201, "OK"]
end
