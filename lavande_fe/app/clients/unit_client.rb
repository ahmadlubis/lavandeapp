class UnitClient
  include HTTParty
  base_uri "http://localhost:10000/v1/unit"
  format :json
  attr_reader :response

  def initialize(token)
    super()
    @token = token
  end

  # List by owners
  # GET /v1/user/index
  def index(query)
    self.class.get(
      "/",
      query: query,
      headers: {
        "Content-Type" => "application/json",
        "Authorization" => "Bearer %s" % @token
      }
    )
  end
end