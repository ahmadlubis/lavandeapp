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
  # GET /v1/unit
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

  # Update by owners
  # PATCH /v1/unit
  def update(payload)
    self.class.patch(
      "",
      headers: {
        "Content-Type" => "application/json",
        "Authorization" => "Bearer %s" % @token
      },
      body: payload.to_json
    )
  end

  # List tenant by owners
  # GET /v1/unit/tenant
  def index_tenant(query)
    self.class.get(
      "/tenant",
      query: query,
      headers: {
        "Content-Type" => "application/json",
        "Authorization" => "Bearer %s" % @token
      }
    )
  end

  # Add tenant by owners
  # GET /v1/unit/tenant
  def add_tenant(payload)
    self.class.post(
      "/tenant",
      headers: {
        "Content-Type" => "application/json",
        "Authorization" => "Bearer %s" % @token
      },
      body: payload.to_json
    )
  end
end