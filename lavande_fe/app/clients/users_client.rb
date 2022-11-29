class UsersClient
  include HTTParty
  base_uri "http://localhost:10000/v1/users"
  format :json
  attr_reader :response

  def initialize(token)
    super()
    @token = token
  end

  # List by owners
  # GET /v1/users
  def index(query)
    self.class.get(
      "",
      query: query,
      headers: {
        "Content-Type" => "application/json",
        "Authorization" => "Bearer %s" % @token
      }
    )
  end
end