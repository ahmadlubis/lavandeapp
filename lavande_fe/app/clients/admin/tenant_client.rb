class Admin::TenantClient
  include HTTParty
  base_uri "http://localhost:10000/v1/admin"
  format :json
  attr_reader :response

  def initialize(token)
    super()
    @token = token
  end

  # Tenants list
  # POST /v1/admin/tenants
  def index(query)
    self.class.get(
      "/tenants",
      query: query,
      headers: {
        "Authorization" => "Bearer %s" % @token
      }
    )
  end

  # Create tenant
  # POST /v1/admin/tenants
  def create(payload)
    self.class.post(
      "/tenants",
      headers: {
        "Content-Type" => "application/json",
        "Authorization" => "Bearer %s" % @token
      },
      body: payload.to_json
    )
  end

  # Delete tenant
  # POST /v1/admin/tenants
  def delete(payload)
    self.class.delete(
      "/tenants",
      headers: {
        "Content-Type" => "application/json",
        "Authorization" => "Bearer %s" % @token
      },
      body: payload.to_json
    )
  end
end