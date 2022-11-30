class Admin::UserClient
  include HTTParty
  base_uri "http://localhost:10000/v1/admin"
  format :json
  attr_reader :response

  def initialize(token)
    super()
    @token = token
  end

  # Users list
  # GET /v1/admin/users
  def index(query)
    self.class.get(
      "/users",
      query: query,
      headers: {
        "Authorization" => "Bearer %s" % @token
      },
    )
  end

  # Change user status
  # PATCH /v1/admin/users
  def status(payload)
    self.class.patch(
      "/users",
      headers: {
        "Content-Type" => "application/json",
        "Authorization" => "Bearer %s" % @token
      },
      body: payload.to_json
    )
  end
end