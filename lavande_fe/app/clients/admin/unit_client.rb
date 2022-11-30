class Admin::UnitClient
  include HTTParty
  base_uri "http://localhost:10000/v1/admin"
  format :json
  attr_reader :response

  def initialize(token)
    super()
    @token = token
  end

  # Users list
  # POST /v1/admin/users
  def get_users(query) 
    self.class.get(
      "/users",
      headers: {
        "Authorization" => "Bearer %s" % @token
      },
      query: query
    )
  end

  # Change user status
  # PATCH /v1/admin/users
  def change_status(payload) 
    self.class.patch(
      "/users",
      headers: {
        "Content-Type" => "application/json",
        "Authorization" => "Bearer %s" % @token
      },
      body: payload.to_json
    )
  end

  # Units list
  # POST /v1/admin/units
  def get_units(token, unit_data, page) 
    query = {
      "limit" => 5,
      "offset" => 5 * page
    }.merge(unit_data)
    self.class.get(
      "/units",
      headers: {
        "Authorization" => "Bearer %s" % token
      },
      query: query
    )
  end

  # Create unit
  # POST /v1/admin/units
  def create_unit(token, unit_data)
    self.class.post(
      "/units",
      headers: {
        "Content-Type" => "application/json",
        "Authorization" => "Bearer %s" % token
      },
      body: unit_data.to_json
    )
  end

  # Create unit
  # POST /v1/admin/units
  def create_unit(token, unit_data)
    self.class.post(
      "/units",
      headers: {
        "Content-Type" => "application/json",
        "Authorization" => "Bearer %s" % token
      },
      body: unit_data.to_json
    )
  end
end