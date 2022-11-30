class Admin::UnitClient
  include HTTParty
  base_uri "http://localhost:10000/v1/admin"
  format :json
  attr_reader :response

  def initialize(token)
    super()
    @token = token
  end

  # Units list
  # POST /v1/admin/units
  def index(query)
    p query
    self.class.get(
      "/units",
      query: query,
      headers: {
        "Authorization" => "Bearer %s" % @token
      }
    )
  end

  # Create unit
  # POST /v1/admin/units
  def create(payload)
    self.class.post(
      "/units",
      headers: {
        "Content-Type" => "application/json",
        "Authorization" => "Bearer %s" % @token
      },
      body: payload.to_json
    )
  end

  # Update unit
  # POST /v1/admin/units
  def update(payload)
    p payload
    self.class.patch(
      "/units",
      headers: {
        "Content-Type" => "application/json",
        "Authorization" => "Bearer %s" % @token
      },
      body: payload.to_json
    )
  end
end