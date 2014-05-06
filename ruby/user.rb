require "json"
require "redis"

$redis = Redis.new(driver: :hiredis)

class User
  def self.key(id)
    "user:#{id}"
  end

  def self.load(id)
    name = $redis.get(self.key(id))
    return User.new(id, name) if name
  end

  def self.from(json)
    data = JSON.parse(json)
    User.new(data["id"], data["name"])
  end

  attr_accessor :id, :name

  def initialize(id, name)
    @id = id
    @name = name
  end

  def save
    $redis.set(key, @name)
  end

  def to_json
    JSON.generate({id: @id, name: @name})
  end

  private

  def key
    User.key(@id)
  end
end
